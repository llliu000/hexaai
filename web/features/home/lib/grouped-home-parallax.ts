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
import Rellax from 'rellax'

type ParallaxLayer = { selector: string; speed: number }
let nextGroup = 0

export function attachGroupedHomeParallax(
  root: HTMLElement,
  layers: ParallaxLayer[]
): () => void {
  const group = `home-parallax-${nextGroup++}`
  const originals = new Map<
    HTMLElement,
    {
      owner: string | null
      speed: string | null
      min: string | null
      max: string | null
    }
  >()
  for (const layer of layers) {
    for (const element of root.querySelectorAll<HTMLElement>(layer.selector)) {
      if (!originals.has(element)) {
        originals.set(element, {
          owner: element.getAttribute('data-home-parallax-owner'),
          speed: element.getAttribute('data-rellax-speed'),
          min: element.getAttribute('data-rellax-min'),
          max: element.getAttribute('data-rellax-max'),
        })
      }
      element.setAttribute('data-home-parallax-owner', group)
      element.setAttribute('data-rellax-speed', String(layer.speed))
      // Visible layers stay within ±50 × speed; stop changing distant layers.
      const limit = Math.abs(layer.speed) * 60
      if (!element.hasAttribute('data-rellax-min')) {
        element.setAttribute('data-rellax-min', String(-limit))
      }
      if (!element.hasAttribute('data-rellax-max')) {
        element.setAttribute('data-rellax-max', String(limit))
      }
    }
  }
  if (!originals.size) return () => {}

  // One scheduler batches all layer writes instead of one loop per element.
  const instance = new Rellax(`[data-home-parallax-owner="${group}"]`, {
    center: true,
  })
  return () => {
    instance.destroy()
    for (const [element, original] of originals) {
      if (original.owner === null) {
        element.removeAttribute('data-home-parallax-owner')
      } else {
        element.setAttribute('data-home-parallax-owner', original.owner)
      }
      if (original.speed === null) element.removeAttribute('data-rellax-speed')
      else element.setAttribute('data-rellax-speed', original.speed)
      if (original.min === null) element.removeAttribute('data-rellax-min')
      else element.setAttribute('data-rellax-min', original.min)
      if (original.max === null) element.removeAttribute('data-rellax-max')
      else element.setAttribute('data-rellax-max', original.max)
    }
  }
}
