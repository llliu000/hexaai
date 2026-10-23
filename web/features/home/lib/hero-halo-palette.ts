import type { HeroSlide } from './hero-config'

export type HaloColor = [number, number, number]

const colors: Record<HeroSlide['theme'], HaloColor> = {
  violet: [0.49, 0.27, 0.96],
  blue: [0.12, 0.46, 0.96],
  teal: [0.02, 0.72, 0.61],
  rose: [0.95, 0.25, 0.55],
}

export function getHaloColor(theme: HeroSlide['theme']): HaloColor {
  return [...colors[theme]]
}

export function interpolateHaloColor(
  from: HaloColor,
  to: HaloColor,
  progress: number
): HaloColor {
  const amount = Math.max(0, Math.min(progress, 1))
  const eased = amount * amount * (3 - 2 * amount)
  return [
    from[0] + (to[0] - from[0]) * eased,
    from[1] + (to[1] - from[1]) * eased,
    from[2] + (to[2] - from[2]) * eased,
  ]
}

// Match the pale gradients used by the original carousel themes.
const backgroundColors: Record<HeroSlide['theme'], [number, number]> = {
  violet: [0x7650d8, 0xb6a0f0],
  blue: [0x2c73ce, 0x7bb5fa],
  teal: [0x138176, 0x69cdbf],
  rose: [0xbd5184, 0xe99aba],
}

export function getHaloBackgroundColor(
  theme: HeroSlide['theme'],
  dark: boolean
): HaloColor {
  const color = backgroundColors[theme][dark ? 1 : 0]
  return [
    ((color >> 16) & 255) / 255,
    ((color >> 8) & 255) / 255,
    (color & 255) / 255,
  ]
}
