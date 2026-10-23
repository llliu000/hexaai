/*
Copyright (C) 2023-2026 QuantumNous

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.

For commercial licensing, please contact support@quantumnous.com
*/
const POSITION_TOLERANCE = 2

export type SectionScrollRange = { start: number; end: number }

export function getSectionScrollRanges(
  sections: { top: number; height: number }[],
  viewportHeight: number,
  headerHeight: number,
  maximum: number
): SectionScrollRange[] {
  const visibleHeight = Math.max(1, viewportHeight - headerHeight)
  const clamp = (value: number): number =>
    Math.max(0, Math.min(Math.max(0, maximum), value))

  return sections
    .filter((section) => section.height > 0)
    .map((section) => {
      if (section.height > visibleHeight) {
        return {
          start: clamp(section.top - headerHeight),
          end: clamp(section.top + section.height - viewportHeight),
        }
      }
      const center = clamp(
        section.top + section.height / 2 - headerHeight - visibleHeight / 2
      )
      return { start: center, end: center }
    })
}

function settleDistanceLimit(viewportHeight: number): number {
  return Math.min(220, viewportHeight * 0.28)
}

export function getSectionSettleTarget(
  position: number,
  ranges: SectionScrollRange[],
  viewportHeight: number,
  origin: number
): number | null {
  const travel = position - origin
  if (Math.abs(travel) < 24 || !ranges.length) return null
  const last = ranges.at(-1)
  if (!last) return null
  // Past the last module: ease back to its center while the overshoot is still small.
  if (position > last.end + POSITION_TOLERANCE) {
    const overshoot = position - last.end
    if (overshoot > settleDistanceLimit(viewportHeight)) return null
    return last.end
  }
  if (
    ranges.some(
      (range) =>
        range.end - range.start > POSITION_TOLERANCE &&
        position >= range.start - POSITION_TOLERANCE &&
        position <= range.end + POSITION_TOLERANCE
    )
  ) {
    return null
  }

  let nearest: number | null = null
  let distance = Infinity
  for (const anchor of ranges.flatMap((range) => [range.start, range.end])) {
    const difference = Math.abs(anchor - position)
    if (
      difference < distance ||
      (difference === distance && (anchor - position) * travel > 0)
    ) {
      nearest = anchor
      distance = difference
    }
  }
  if (nearest === null || distance <= POSITION_TOLERANCE) return null
  if (distance > settleDistanceLimit(viewportHeight)) return null
  // Do not undo a small departure from a module or pull past the gesture's origin.
  if (
    (nearest - position) * travel < 0 &&
    distance >= Math.abs(travel) - POSITION_TOLERANCE
  ) {
    return null
  }
  return nearest
}

export function getSectionNudgeTarget(
  position: number,
  ranges: SectionScrollRange[],
  viewportHeight: number,
  gesture: { origin: number; distance: number; durationMs: number }
): number | null {
  const delta = position - gesture.origin
  const distance = Math.abs(delta)
  // Cover medium up/down slides: upward often travels past the old 140px window
  // because the previous module already peeks into view and reveal will not fire.
  const maxTravel = Math.min(420, viewportHeight * 0.48)
  if (
    distance < 48 ||
    distance > maxTravel ||
    gesture.distance > maxTravel + 48 ||
    gesture.distance - distance > 48 ||
    gesture.durationMs >= 1200
  ) {
    return null
  }

  const index = ranges.findIndex(
    (range) =>
      range.end - range.start <= POSITION_TOLERANCE &&
      Math.abs(range.start - gesture.origin) <= 12
  )
  if (index < 0) return null
  const direction = Math.sign(delta)
  const adjacent = ranges[index + direction]
  if (!adjacent) return null
  const target = direction > 0 ? adjacent.start : adjacent.end
  if (
    (target - position) * direction <= POSITION_TOLERANCE ||
    Math.abs(target - gesture.origin) > viewportHeight * 1.15
  ) {
    return null
  }
  return target
}
