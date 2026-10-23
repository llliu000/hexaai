export const HALO_TRANSITION_DURATION = 2
// Bounds shared with the shader keep even the ragged edge outside at rest.
export const HALO_WAVE_FEATHER = 0.1
export const HALO_INK_DISTORTION = 0.26

export type HaloPlacement = {
  origin: [number, number]
  size: number
  startRadius: number
  exitRadius: number
}

function smoothProgress(value: number): number {
  const progress = Math.max(0, Math.min(1, value))
  return progress * progress * (3 - 2 * progress)
}

export function getHaloPlacement(width: number, height: number): HaloPlacement {
  const safeWidth = Math.max(width, 1)
  const safeHeight = Math.max(height, 1)
  const mobile = width <= 520
  const size = mobile ? 0.82 : 1.05
  const offset = ((size / 5.2 + 0.045) * safeHeight) / safeWidth
  const origin: [number, number] = [1 + offset, mobile ? 0.73 : 0.68]
  return {
    origin,
    size,
    startRadius: size / 5.2 - HALO_INK_DISTORTION - HALO_WAVE_FEATHER,
    // Include maximum paper distortion and feather to finish without a snap.
    exitRadius:
      Math.hypot(
        (origin[0] * safeWidth) / safeHeight,
        Math.max(origin[1], 1 - origin[1])
      ) +
      HALO_INK_DISTORTION +
      HALO_WAVE_FEATHER,
  }
}

export function getHaloWaveRadius(
  elapsed: number,
  placement: HaloPlacement
): number {
  const progress = smoothProgress(elapsed / HALO_TRANSITION_DURATION)
  return (
    placement.startRadius +
    (placement.exitRadius - placement.startRadius) * progress
  )
}

/** The paper-warped distance is fixed while the reveal radius advances. */
export function getHaloReveal(distance: number, radius: number): number {
  return (
    1 -
    smoothProgress(
      (distance - radius + HALO_WAVE_FEATHER) / (HALO_WAVE_FEATHER * 2)
    )
  )
}

export function getHaloWaveOpacity(elapsed: number): number {
  if (elapsed <= 0 || elapsed >= HALO_TRANSITION_DURATION) return 0
  return (
    smoothProgress(elapsed / 0.18) *
    (1 - smoothProgress((elapsed - HALO_TRANSITION_DURATION + 0.25) / 0.25))
  )
}
