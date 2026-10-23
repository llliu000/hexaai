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
import type { SectionScrollRange } from './section-scroll'

export type SectionScrollVisibility = {
  top: number
  height: number
  range: SectionScrollRange
}

export function getLastVisibleSectionTarget(
  position: number,
  sections: SectionScrollVisibility[],
  viewport: { height: number; headerHeight: number },
  gesture: { origin: number; direction: number }
): number | null {
  if (
    !gesture.direction ||
    (position - gesture.origin) * gesture.direction < 120
  ) {
    return null
  }
  const last = sections.at(-1)
  if (!last || position > last.range.end + 2) return null
  if (
    sections.some(
      (section) =>
        section.range.end - section.range.start > 2 &&
        position >= section.range.start - 2 &&
        position <= section.range.end + 2
    )
  ) {
    return null
  }

  const visibleHeight = viewport.height - viewport.headerHeight
  if (visibleHeight <= 0) return null
  const visiblePixels = (
    section: SectionScrollVisibility,
    scrollTop: number
  ): number =>
    Math.max(
      0,
      Math.min(section.top + section.height, scrollTop + viewport.height) -
        Math.max(section.top, scrollTop + viewport.headerHeight)
    )
  // Walk bottom-to-top: downward lands on the last revealed module, upward on the
  // nearest module above instead of skipping several pages.
  const candidates = [...sections].reverse()
  for (const section of candidates) {
    if (section.height <= 0) continue
    const minimumVisible = Math.min(
      180,
      section.height * 0.35,
      visibleHeight * 0.35
    )
    const nowVisible = visiblePixels(section, position)
    const wasVisible = visiblePixels(section, gesture.origin)
    if (nowVisible < minimumVisible) continue
    if (gesture.direction < 0) {
      // Previous modules often already peek while centered; require a clear gain.
      if (nowVisible - wasVisible < minimumVisible * 0.55) continue
    } else if (wasVisible >= minimumVisible) {
      continue
    }
    const target =
      gesture.direction > 0 ? section.range.start : section.range.end
    const maxCorrection = Math.min(
      gesture.direction < 0 ? 560 : 480,
      visibleHeight * (gesture.direction < 0 ? 0.75 : 0.65)
    )
    if (
      (target - position) * gesture.direction <= 2 ||
      Math.abs(target - position) > maxCorrection
    ) {
      continue
    }
    return target
  }
  return null
}
