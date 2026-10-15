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
import type { SectionScrollVisibility } from './section-scroll-reveal'

function visiblePixels(
  section: SectionScrollVisibility,
  scrollTop: number,
  viewport: { height: number; headerHeight: number }
): number {
  return Math.max(
    0,
    Math.min(section.top + section.height, scrollTop + viewport.height) -
      Math.max(section.top, scrollTop + viewport.headerHeight)
  )
}

function originSectionIndex(
  sections: SectionScrollVisibility[],
  origin: number,
  viewport: { height: number; headerHeight: number }
): number {
  let index = 0
  let best = -1
  for (let i = 0; i < sections.length; i += 1) {
    const pixels = visiblePixels(sections[i], origin, viewport)
    if (
      pixels > best ||
      (pixels === best &&
        origin >= sections[i].range.start - 2 &&
        origin <= sections[i].range.end + 2)
    ) {
      best = pixels
      index = i
    }
  }
  return index
}

/**
 * One-page turn when the adjacent module starts appearing.
 * Long flings are left to settle/reveal so large-range scrolling stays native.
 */
export function getAppearingSectionTarget(
  position: number,
  sections: SectionScrollVisibility[],
  viewport: { height: number; headerHeight: number },
  gesture: { origin: number; direction: number }
): number | null {
  if (!gesture.direction || sections.length < 2) return null
  const travel = (position - gesture.origin) * gesture.direction
  if (travel < 28) return null

  const visibleHeight = viewport.height - viewport.headerHeight
  if (visibleHeight <= 0) return null
  // Do not steal multi-module flings.
  if (travel > Math.min(viewport.height * 0.72, 580)) return null

  const originIndex = originSectionIndex(sections, gesture.origin, viewport)
  const adjacentIndex = originIndex + gesture.direction
  if (adjacentIndex < 0 || adjacentIndex >= sections.length) return null

  const adjacent = sections[adjacentIndex]
  const appearThreshold = Math.min(48, Math.max(28, visibleHeight * 0.05))
  const nowVisible = visiblePixels(adjacent, position, viewport)
  const wasVisible = visiblePixels(adjacent, gesture.origin, viewport)
  const newlyAppeared =
    wasVisible < appearThreshold && nowVisible >= appearThreshold
  const grew = nowVisible - wasVisible >= appearThreshold * 0.3
  // Downward often already peeks the next module; modest travel still counts.
  const clearTravel =
    travel >= Math.min(56, visibleHeight * 0.07) && nowVisible >= appearThreshold
  if (!newlyAppeared && !grew && !clearTravel) return null

  const target =
    gesture.direction > 0 ? adjacent.range.start : adjacent.range.end
  if ((target - gesture.origin) * gesture.direction <= 2) return null
  if (Math.abs(target - position) > Math.min(viewport.height * 1.15, 980)) {
    return null
  }
  return target
}

/**
 * Complete a downward page turn that overshot into the adjacent tall module's
 * reading range, without stealing deep in-module reading or long flings.
 */
export function getEnteredSectionTarget(
  position: number,
  sections: SectionScrollVisibility[],
  viewport: { height: number; headerHeight: number },
  gesture: { origin: number; direction: number }
): number | null {
  if (gesture.direction <= 0 || sections.length < 2) return null
  const travel = position - gesture.origin
  if (travel < 28) return null

  const originIndex = originSectionIndex(sections, gesture.origin, viewport)
  const adjacent = sections[originIndex + 1]
  if (!adjacent) return null
  const span = adjacent.range.end - adjacent.range.start
  if (span <= 2) return null
  if (gesture.origin >= adjacent.range.start - 2) return null
  if (
    position < adjacent.range.start - 2 ||
    position > adjacent.range.end + 2
  ) {
    return null
  }
  // Only finish the entry while still near the top of that tall module.
  // Travel can be large because short→tall snap points are far apart.
  if (position - adjacent.range.start > Math.min(260, span * 0.45 + 60)) {
    return null
  }
  return adjacent.range.start
}
