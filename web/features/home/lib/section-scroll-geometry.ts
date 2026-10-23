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
import {
  getSectionScrollRanges,
  type SectionScrollRange,
} from './section-scroll'

type SectionGeometry = {
  element: HTMLElement
  top: number
  height: number
  range: SectionScrollRange
}

export function createSectionScrollGeometry(root: HTMLElement): {
  read: () => SectionGeometry[]
  destroy: () => void
} {
  let dirty = true
  let entries: SectionGeometry[] = []
  const observed = new Set<Element>()
  const invalidate = (): void => {
    dirty = true
  }
  const resize = new ResizeObserver(invalidate)
  const mutations = new MutationObserver(invalidate)
  mutations.observe(root, { childList: true, subtree: true })
  window.addEventListener('resize', invalidate, { passive: true })

  return {
    read() {
      if (!dirty) return entries
      const sections = [...root.querySelectorAll<HTMLElement>('main > section')]
      const header = root.querySelector('header')
      const current = new Set<Element>([root, ...sections])
      if (header) current.add(header)
      for (const element of observed) {
        if (!current.has(element)) {
          resize.unobserve(element)
          observed.delete(element)
        }
      }
      for (const element of current) {
        if (!observed.has(element)) {
          resize.observe(element)
          observed.add(element)
        }
      }

      // Read geometry together only after layout changes, never between writes.
      const position = window.scrollY
      const viewportHeight = window.innerHeight
      const headerHeight = header?.getBoundingClientRect().height ?? 0
      const maximum = document.documentElement.scrollHeight - viewportHeight
      const visible = sections
        .map((element) => {
          const rect = element.getBoundingClientRect()
          return { element, top: position + rect.top, height: rect.height }
        })
        .filter((section) => section.height > 0)
      const ranges = getSectionScrollRanges(
        visible,
        viewportHeight,
        headerHeight,
        maximum
      )
      entries = visible.map((section, index) => ({
        element: section.element,
        top: section.top,
        height: section.height,
        range: ranges[index],
      }))
      dirty = false
      return entries
    },
    destroy() {
      resize.disconnect()
      mutations.disconnect()
      window.removeEventListener('resize', invalidate)
      observed.clear()
      entries = []
    },
  }
}
