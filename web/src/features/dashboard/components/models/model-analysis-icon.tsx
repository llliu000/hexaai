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
  ArrowUpRight,
  Circle,
  Hash,
  Infinity as InfinityIcon,
  Menu,
  SquareSquare,
  Zap,
  type LucideIcon,
} from 'lucide-react'
import type { ReactElement } from 'react'

// Retain the symbols from the existing model-page artwork, rendered with the
// same SVG dimensions and line treatment as the overview API route icon.
const icons = new Map<string, LucideIcon>([
  ['count', Hash],
  ['quota', InfinityIcon],
  ['tokens', Circle],
  ['avgRpm', ArrowUpRight],
  ['avgTpm', Zap],
  ['consumption', Menu],
  ['trend', SquareSquare],
])

export function ModelAnalysisIcon(props: { name: string }): ReactElement {
  const Icon = icons.get(props.name) ?? Hash
  return (
    <span className='model-analysis-icon' aria-hidden='true'>
      <Icon
        className='model-analysis-icon-svg'
        size={20}
        strokeWidth={2}
        fill='none'
      />
    </span>
  )
}
