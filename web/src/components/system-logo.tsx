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
import { useState } from 'react'

import { Logo } from '@/assets/logo'
import { resolveSystemLogo, SYSTEM_LOGO_URL } from '@/lib/system-logo'
import { cn } from '@/lib/utils'

interface SystemLogoProps {
  src?: string | null
  alt?: string
  width?: number
  height?: number
  className?: string
}

export function SystemLogo(props: SystemLogoProps) {
  const source = resolveSystemLogo(props.src)
  const [failedSource, setFailedSource] = useState<string | null>(null)
  const className = cn('system-logo shrink-0 object-contain', props.className)
  if (source !== SYSTEM_LOGO_URL && source !== failedSource) {
    return (
      <img
        src={source}
        alt={props.alt ?? ''}
        width={props.width ?? 32}
        height={props.height ?? 32}
        className={className}
        onError={() => setFailedSource(source)}
      />
    )
  }
  return (
    <Logo
      width={props.width ?? 32}
      height={props.height ?? 32}
      className={className}
      aria-hidden={props.alt ? undefined : true}
      aria-label={props.alt}
    />
  )
}
