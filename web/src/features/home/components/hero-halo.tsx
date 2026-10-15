import { useEffect, useRef, useState } from 'react'

import type { HeroSlide } from '../lib/hero-config'
import { createHeroHaloRenderer } from '../lib/hero-halo-renderer'

import '@/styles/hero-halo.css'

type HeroHaloProps = {
  theme: HeroSlide['theme']
  slideId: string
  reducedMotion: boolean
}

export function HeroHalo(props: HeroHaloProps) {
  const hostRef = useRef<HTMLDivElement>(null)
  const canvasRef = useRef<HTMLCanvasElement>(null)
  const rendererRef = useRef<ReturnType<typeof createHeroHaloRenderer>>(null)
  const [generation, setGeneration] = useState(0)

  useEffect(() => {
    const canvas = canvasRef.current
    const host = hostRef.current
    if (!canvas || !host) return
    const renderer = createHeroHaloRenderer(
      canvas,
      props.reducedMotion,
      (color) => {
        const channels = color.map((value) => Math.round(value * 255)).join(' ')
        host.style.setProperty(
          '--wm-halo-background-accent',
          `rgb(${channels})`
        )
      }
    )
    rendererRef.current = renderer
    canvas.dataset.renderer = renderer ? 'webgl' : 'fallback'
    if (!renderer) {
      host.style.removeProperty('--wm-halo-background-accent')
      return
    }
    let intersecting = true
    const updateVisibility = () =>
      renderer.setVisible(intersecting && !document.hidden)
    const updateSize = () =>
      renderer.resize(host.clientWidth, host.clientHeight)
    const updateTheme = () =>
      renderer.setDark(document.documentElement.classList.contains('dark'))
    const resize = new ResizeObserver(updateSize)
    resize.observe(host)
    const visibility = new IntersectionObserver((entries) => {
      intersecting = entries[0]?.isIntersecting ?? true
      updateVisibility()
    })
    visibility.observe(host)
    const theme = new MutationObserver(updateTheme)
    theme.observe(document.documentElement, {
      attributes: true,
      attributeFilter: ['class'],
    })
    const onLost = (event: Event) => {
      event.preventDefault()
      renderer.destroy()
      canvas.dataset.renderer = 'fallback'
      host.style.removeProperty('--wm-halo-background-accent')
    }
    const onRestored = () => setGeneration((value) => value + 1)
    canvas.addEventListener('webglcontextlost', onLost)
    canvas.addEventListener('webglcontextrestored', onRestored)
    document.addEventListener('visibilitychange', updateVisibility)
    updateSize()
    updateVisibility()
    updateTheme()
    return () => {
      resize.disconnect()
      visibility.disconnect()
      theme.disconnect()
      canvas.removeEventListener('webglcontextlost', onLost)
      canvas.removeEventListener('webglcontextrestored', onRestored)
      document.removeEventListener('visibilitychange', updateVisibility)
      renderer.destroy()
      rendererRef.current = null
    }
  }, [props.reducedMotion, generation])

  useEffect(() => {
    rendererRef.current?.setTheme(props.theme, props.slideId)
  }, [props.theme, props.slideId, props.reducedMotion, generation])

  return (
    <div
      ref={hostRef}
      className={`wm-hero-halo wm-theme-${props.theme}`}
      aria-hidden='true'
    >
      <div className='wm-hero-halo-fallback' />
      <canvas ref={canvasRef} className='wm-hero-halo-canvas' />
    </div>
  )
}
