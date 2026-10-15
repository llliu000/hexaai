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
import { useId, type SVGProps } from 'react'

import { DEFAULT_SYSTEM_NAME } from '@/lib/constants'
import {
  SYSTEM_LOGO_PATH,
  SYSTEM_LOGO_STOPS,
  SYSTEM_LOGO_VIEW_BOX,
} from '@/lib/system-logo'
import { cn } from '@/lib/utils'

export function Logo(props: SVGProps<SVGSVGElement>) {
  const gradientId = `system-logo-${useId()}`
  return (
    <svg
      {...props}
      data-system-logo='true'
      viewBox={SYSTEM_LOGO_VIEW_BOX}
      xmlns='http://www.w3.org/2000/svg'
      height={props.height ?? 24}
      width={props.width ?? 24}
      fill='none'
      focusable='false'
      role={props.role ?? 'img'}
      aria-label={props['aria-label'] ?? DEFAULT_SYSTEM_NAME}
      className={cn('shrink-0', props.className)}
    >
      <defs>
        <linearGradient
          id={gradientId}
          x1='16'
          y1='16'
          x2='216'
          y2='168'
          gradientUnits='userSpaceOnUse'
        >
          {SYSTEM_LOGO_STOPS.map((stop) => (
            <stop
              key={stop.offset}
              offset={stop.offset}
              stopColor={stop.color}
            />
          ))}
        </linearGradient>
      </defs>
      <path d={SYSTEM_LOGO_PATH} fill={`url(#${gradientId})`} />
    </svg>
  )
}
