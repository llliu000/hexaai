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
import { Activity, BarChart3, Wallet, Eye, EyeOff } from 'lucide-react'
import { useState } from 'react'
import { useTranslation } from 'react-i18next'

import { PageIcon, type PageIconTone } from '@/components/ui/page-icon'
import { Skeleton } from '@/components/ui/skeleton'
import { formatQuota } from '@/lib/format'

import type { UserWalletData } from '../types'

interface WalletStatsCardProps {
  user: UserWalletData | null
  loading?: boolean
}

export function WalletStatsCard(props: WalletStatsCardProps) {
  const { t } = useTranslation()
  const [hidden, setHidden] = useState<Record<string, boolean>>({})
  if (props.loading) {
    return (
      <div className='grid grid-cols-3 divide-x rounded-lg border'>
        {['balance', 'usage', 'requests'].map((key) => (
          <div key={key} className='min-w-0 px-2.5 py-2.5 sm:px-5 sm:py-4'>
            <Skeleton className='h-3.5 w-full' />
            <Skeleton className='mt-2 h-6 w-full sm:h-7' />
            <Skeleton className='mt-1.5 hidden h-3.5 w-24 md:block' />
          </div>
        ))}
      </div>
    )
  }

  const stats: {
    label: string
    value: string
    description: string
    icon: typeof Wallet
    tone: PageIconTone
  }[] = [
    {
      label: t('Current Balance'),
      value: formatQuota(props.user?.quota ?? 0),
      description: t('Remaining quota'),
      icon: Wallet,
      tone: 'green',
    },
    {
      label: t('Total Usage'),
      value: formatQuota(props.user?.used_quota ?? 0),
      description: t('Total consumed quota'),
      icon: BarChart3,
      tone: 'blue',
    },
    {
      label: t('API Requests'),
      value: (props.user?.request_count ?? 0).toLocaleString(),
      description: t('Total requests made'),
      icon: Activity,
      tone: 'purple',
    },
  ]

  return (
    <div className='wallet-design-stats'>
      {stats.map((item, index) => (
        <div
          key={item.label}
          className={`wallet-design-stat wallet-stat-${index}`}
        >
          <div className='flex items-center gap-1.5 sm:gap-2.5'>
            <PageIcon icon={item.icon} tone={item.tone} />
            <div className='text-muted-foreground truncate text-[12px] font-medium tracking-wider uppercase sm:text-xs'>
              {item.label}
            </div>
            <button
              type='button'
              className='wallet-visibility'
              aria-label={t('Toggle visibility')}
              aria-pressed={Boolean(hidden[item.label])}
              onClick={() =>
                setHidden((previous) => ({
                  ...previous,
                  [item.label]: !previous[item.label],
                }))
              }
            >
              {hidden[item.label] ? <EyeOff size={14} /> : <Eye size={14} />}
            </button>
          </div>

          <div className='text-foreground mt-1.5 font-mono text-sm font-bold tracking-tight break-all tabular-nums sm:mt-2.5 sm:text-2xl'>
            {hidden[item.label] ? '••••' : item.value}
          </div>
          <span
            className={`wallet-art wallet-stat-art wallet-stat-art-${index}`}
            aria-hidden='true'
          />
          <div className='text-muted-foreground/60 mt-1 hidden text-xs md:block'>
            {item.description}
          </div>
        </div>
      ))}
    </div>
  )
}
