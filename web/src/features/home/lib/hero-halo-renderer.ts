import type { HeroSlide } from './hero-config'
import {
  createHaloProgram,
  createHaloSurface,
  resizeHaloSurface,
  type HaloSurface,
} from './hero-halo-gl'
import {
  getHaloBackgroundColor,
  getHaloColor,
  type HaloColor,
} from './hero-halo-palette'
import { HALO_REVEAL_SHADER, HALO_VERTEX_SHADER } from './hero-halo-shaders'
import {
  getHaloPlacement,
  getHaloWaveOpacity,
  getHaloWaveRadius,
  HALO_TRANSITION_DURATION,
} from './hero-halo-transition'

export function createHeroHaloRenderer(
  canvas: HTMLCanvasElement,
  reducedMotion: boolean,
  onBackgroundColor?: (color: HaloColor) => void
) {
  const gl = canvas.getContext('webgl', {
    alpha: true,
    antialias: false,
    depth: false,
    stencil: false,
    premultipliedAlpha: false,
    powerPreference: 'low-power',
  })
  if (!gl) return null

  let program: WebGLProgram | null = null
  const surfaces: HaloSurface[] = []
  const quad = gl.createBuffer()
  let frame = 0
  let lastFrame = 0
  let destroyed = false
  let visible = true
  let allocated = false
  let dark = false
  let sourceIndex = 0
  let elapsed = HALO_TRANSITION_DURATION
  let currentSlide = ''
  let inkVariation = 0
  let currentTheme: HeroSlide['theme'] = 'violet'
  let targetColor = getHaloBackgroundColor(currentTheme, dark)
  let ringColor = getHaloColor(currentTheme)
  let placement = getHaloPlacement(canvas.clientWidth, canvas.clientHeight)

  const destroy = () => {
    if (destroyed) return
    destroyed = true
    cancelAnimationFrame(frame)
    for (const surface of surfaces) {
      gl.deleteTexture(surface.texture)
      gl.deleteFramebuffer(surface.framebuffer)
    }
    gl.deleteProgram(program)
    gl.deleteBuffer(quad)
  }

  try {
    if (!quad) throw new Error('Unable to create ripple vertices')
    program = createHaloProgram(gl, HALO_VERTEX_SHADER, HALO_REVEAL_SHADER)
    surfaces.push(createHaloSurface(gl))
    surfaces.push(createHaloSurface(gl))
    gl.bindBuffer(gl.ARRAY_BUFFER, quad)
    gl.bufferData(
      gl.ARRAY_BUFFER,
      new Float32Array([-1, -1, 3, -1, -1, 3]),
      gl.STATIC_DRAW
    )
  } catch {
    destroy()
    return null
  }

  const uniforms = Object.fromEntries(
    [
      'uFrom',
      'uResolution',
      'uOrigin',
      'uTargetColor',
      'uSurfaceColor',
      'uRingColor',
      'uRadius',
      'uComplete',
      'uRingStrength',
      'uInkSeed',
    ].map((name) => [name, gl.getUniformLocation(program, name)])
  )
  const attribute = gl.getAttribLocation(program, 'aPosition')

  const render = (destination: WebGLFramebuffer | null, highlight: boolean) => {
    if (destroyed || !allocated || gl.isContextLost()) return
    const complete = elapsed >= HALO_TRANSITION_DURATION
    gl.useProgram(program)
    gl.bindBuffer(gl.ARRAY_BUFFER, quad)
    gl.enableVertexAttribArray(attribute)
    gl.vertexAttribPointer(attribute, 2, gl.FLOAT, false, 0, 0)
    gl.bindFramebuffer(gl.FRAMEBUFFER, destination)
    gl.viewport(0, 0, canvas.width, canvas.height)
    gl.activeTexture(gl.TEXTURE0)
    gl.bindTexture(gl.TEXTURE_2D, surfaces[sourceIndex].texture)
    gl.uniform1i(uniforms.uFrom, 0)
    gl.uniform2f(uniforms.uResolution, canvas.width, canvas.height)
    gl.uniform2fv(uniforms.uOrigin, placement.origin)
    gl.uniform3fv(uniforms.uTargetColor, targetColor)
    gl.uniform3fv(
      uniforms.uSurfaceColor,
      dark ? [21 / 255, 29 / 255, 42 / 255] : [1, 1, 1]
    )
    gl.uniform3fv(uniforms.uRingColor, ringColor)
    gl.uniform2f(
      uniforms.uInkSeed,
      (inkVariation * 5.37) % 37,
      (inkVariation * 3.19) % 29
    )
    gl.uniform1f(uniforms.uRadius, getHaloWaveRadius(elapsed, placement))
    gl.uniform1f(uniforms.uComplete, complete ? 1 : 0)
    const strength = dark ? 0.085 : 0.055
    gl.uniform1f(
      uniforms.uRingStrength,
      highlight ? getHaloWaveOpacity(elapsed) * strength : 0
    )
    gl.drawArrays(gl.TRIANGLES, 0, 3)
  }

  const commitBackground = () => onBackgroundColor?.(targetColor)

  const animate = (timestamp: number) => {
    frame = 0
    if (destroyed || !visible || reducedMotion || gl.isContextLost()) return
    if (!lastFrame || timestamp - lastFrame >= 1000 / 30) {
      elapsed = Math.min(
        HALO_TRANSITION_DURATION,
        elapsed +
          (lastFrame ? Math.min((timestamp - lastFrame) / 1000, 0.06) : 1 / 30)
      )
      lastFrame = timestamp
      render(null, true)
    }
    if (elapsed < HALO_TRANSITION_DURATION) {
      frame = requestAnimationFrame(animate)
    } else commitBackground()
  }

  const schedule = () => {
    cancelAnimationFrame(frame)
    frame = 0
    lastFrame = 0
    if (destroyed || !allocated || !visible || gl.isContextLost()) return
    render(null, true)
    if (!reducedMotion && elapsed < HALO_TRANSITION_DURATION) {
      frame = requestAnimationFrame(animate)
    }
  }

  return {
    resize(width: number, height: number) {
      if (destroyed || gl.isContextLost()) return
      placement = getHaloPlacement(width, height)
      const scale = Math.min(
        window.devicePixelRatio || 1,
        1.25,
        1500 / Math.max(width, 1),
        850 / Math.max(height, 1)
      )
      const nextWidth = Math.max(1, Math.round(width * scale))
      const nextHeight = Math.max(1, Math.round(height * scale))
      if (
        allocated &&
        canvas.width === nextWidth &&
        canvas.height === nextHeight
      ) {
        return
      }
      canvas.width = nextWidth
      canvas.height = nextHeight
      try {
        for (const surface of surfaces) {
          resizeHaloSurface(gl, surface, nextWidth, nextHeight)
        }
      } catch {
        destroy()
        canvas.dataset.renderer = 'fallback'
        commitBackground()
        return
      }
      // The new-sized surface starts with the selected gradient, never black.
      allocated = true
      elapsed = HALO_TRANSITION_DURATION
      render(surfaces[1 - sourceIndex].framebuffer, false)
      sourceIndex = 1 - sourceIndex
      commitBackground()
      schedule()
    },
    setTheme(theme: HeroSlide['theme'], slideId: string) {
      if (destroyed || gl.isContextLost()) return
      const key = `${slideId}:${theme}`
      if (currentSlide === key) return
      const animateChange = Boolean(currentSlide) && allocated && !reducedMotion
      // Freeze the visible colors without the old ring; even rapid clicks keep
      // one wavefront and preserve the already-revealed area without flashing.
      if (animateChange) {
        render(surfaces[1 - sourceIndex].framebuffer, false)
        sourceIndex = 1 - sourceIndex
      }
      // Change the paper field only after freezing the previous background.
      inkVariation += 1
      currentTheme = theme
      currentSlide = key
      targetColor = getHaloBackgroundColor(theme, dark)
      ringColor = getHaloColor(theme)
      elapsed = animateChange ? 0 : HALO_TRANSITION_DURATION
      if (!animateChange) commitBackground()
      schedule()
    },
    setVisible(value: boolean) {
      if (visible === value) return
      visible = value
      schedule()
    },
    setDark(value: boolean) {
      if (destroyed || dark === value) return
      dark = value
      targetColor = getHaloBackgroundColor(currentTheme, dark)
      elapsed = HALO_TRANSITION_DURATION
      commitBackground()
      schedule()
    },
    destroy,
  }
}
