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
import { CircleHelp, Clock3, Info, UsersRound, Wallet } from 'lucide-react'
import { useTranslation } from 'react-i18next'

import { CopyButton } from '@/components/copy-button'
import { Button } from '@/components/ui/button'
import { Card, CardContent } from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import { PageIcon, type PageIconTone } from '@/components/ui/page-icon'
import { Skeleton } from '@/components/ui/skeleton'
import { formatQuota } from '@/lib/format'

import type { UserWalletData } from '../types'

interface AffiliateRewardsCardProps {
  user: UserWalletData | null
  affiliateLink: string
  onTransfer: () => void
  complianceConfirmed?: boolean
  loading?: boolean
}

export function AffiliateRewardsCard({
  user,
  affiliateLink,
  onTransfer,
  complianceConfirmed = true,
  loading,
}: AffiliateRewardsCardProps) {
  const { t } = useTranslation()
  if (loading) {
    return (
      <Card data-card-hover='false' className='bg-muted/20 py-0'>
        <CardContent className='grid gap-4 p-3 sm:p-4 lg:grid-cols-[minmax(220px,1fr)_minmax(220px,0.72fr)_minmax(320px,1.15fr)] lg:items-center'>
          <div>
            <Skeleton className='h-5 w-32' />
            <Skeleton className='mt-2 h-4 w-48' />
          </div>
          <Skeleton className='h-14 rounded-lg' />
          <Skeleton className='h-10 rounded-lg' />
        </CardContent>
      </Card>
    )
  }

  const hasRewards = (user?.aff_quota ?? 0) > 0
  const stats: {
    label: string
    value: string
    icon: typeof Wallet
    tone: PageIconTone
  }[] = [
    {
      label: t('Pending'),
      value: formatQuota(user?.aff_quota ?? 0),
      icon: Clock3,
      tone: 'green',
    },
    {
      label: t('Total Earned'),
      value: formatQuota(user?.aff_history_quota ?? 0),
      icon: Wallet,
      tone: 'blue',
    },
    {
      label: t('Invites'),
      value: String(user?.aff_count ?? 0),
      icon: UsersRound,
      tone: 'amber',
    },
  ]

  return (
    <Card data-card-hover='false' className='wallet-affiliate-card'>
      <CardContent className='wallet-affiliate-content'>
        <div className='wallet-affiliate-intro'>
          <PageIcon icon={UsersRound} tone='purple' />
          <div className='min-w-0'>
            <h3 className='truncate text-sm font-semibold'>
              {t('Referral Program')}
            </h3>
            <p className='text-muted-foreground text-xs'>
              {t(
                'Earn rewards when users join through your referral link. Transfer accumulated rewards to your balance anytime.'
              )}
            </p>
          </div>
          {!complianceConfirmed && (
            <p className='wallet-affiliate-notice'>
              <Info size={17} />
              {t(
                'Referral reward transfer is disabled until the administrator confirms compliance terms.'
              )}
            </p>
          )}
        </div>

        <div className='wallet-affiliate-stats'>
          {stats.map((item) => (
            <div key={item.label}>
              <PageIcon
                icon={item.icon}
                tone={item.tone}
                className='wallet-referral-stat-icon'
              />
              <div className='text-muted-foreground truncate text-[12px] font-medium tracking-wider uppercase'>
                {item.label}
                <CircleHelp size={12} aria-hidden='true' />
              </div>
              <div className='mt-0.5 truncate text-sm font-semibold tabular-nums'>
                {item.value}
              </div>
            </div>
          ))}
        </div>

        <div className='wallet-affiliate-link'>
          <label htmlFor='wallet-referral-link'>
            {t('Referral Link', { defaultValue: '推荐链接' })}
          </label>
          <div className='wallet-affiliate-link-control'>
            <Input
              id='wallet-referral-link'
              value={affiliateLink}
              readOnly
              className='border-muted bg-background/70 h-9 min-w-0 flex-1 font-mono text-xs'
            />
            <CopyButton
              value={affiliateLink}
              variant='outline'
              className='bg-background size-9 shrink-0'
              iconClassName='size-4'
              tooltip={t('Copy referral link')}
              aria-label={t('Copy referral link')}
            />
          </div>
          <p>
            {t('Share your link to invite friends and earn more rewards.', {
              defaultValue: '分享链接邀请好友，获得更多奖励',
            })}
          </p>
          {hasRewards && (
            <Button
              onClick={onTransfer}
              disabled={!complianceConfirmed}
              className='h-9 shrink-0 px-3'
              size='sm'
            >
              {t('Transfer to Balance')}
            </Button>
          )}
        </div>
      </CardContent>
    </Card>
  )
}
