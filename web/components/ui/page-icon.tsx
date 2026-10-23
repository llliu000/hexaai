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
import type { LucideIcon } from 'lucide-react'
import type { ReactElement } from 'react'

import { cn } from '@/lib/utils'

import '@/styles/page-icon.css'

export type PageIconTone =
  | 'blue'
  | 'green'
  | 'purple'
  | 'cyan'
  | 'amber'
  | 'red'
  | 'slate'

interface PageIconProps {
  icon: LucideIcon
  tone?: PageIconTone
  className?: string
}

export function PageIcon(props: PageIconProps): ReactElement {
  const Icon = props.icon

  return (
    <span
      className={cn(
        'page-icon',
        `page-icon-${props.tone ?? 'blue'}`,
        props.className
      )}
      aria-hidden='true'
    >
      <Icon size={20} strokeWidth={2} fill='none' />
    </span>
  )
}
