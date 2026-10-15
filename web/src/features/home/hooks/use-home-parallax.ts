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
import { useReducedMotion } from 'motion/react'
import { useEffect, type RefObject } from 'react'

import { attachGroupedHomeParallax } from '../lib/grouped-home-parallax'

export function useHomeParallax(ref: RefObject<HTMLElement | null>): void {
  const reducedMotion = useReducedMotion()

  useEffect(() => {
    const root = ref.current
    if (!root || reducedMotion) return

    return attachGroupedHomeParallax(root, [
      { selector: '.wm-home-parallax-backdrop-slow', speed: -4 },
      { selector: '.wm-home-parallax-backdrop-drift', speed: -2 },
      { selector: '.wm-home-parallax-heading', speed: 1 },
      { selector: '.wm-home-parallax-panel', speed: 2 },
    ])
  }, [reducedMotion, ref])
}
