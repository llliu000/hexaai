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
import { ArrowRight, Copy } from 'lucide-react'
import { memo, type KeyboardEvent, type ReactNode } from 'react'
import { useTranslation } from 'react-i18next'

import { useCopyToClipboard } from '@/hooks/use-copy-to-clipboard'
import { getLobeIcon } from '@/lib/lobe-icon'
import { cn } from '@/lib/utils'

import { DEFAULT_TOKEN_UNIT } from '../constants'
import {
  getDynamicDisplayGroupRatio,
  getDynamicPricingSummary,
} from '../lib/dynamic-price'
import { parseTags } from '../lib/filters'
import { isTokenBasedModel } from '../lib/model-helpers'
import { getPricingVendorIcon } from '../lib/model-vendor-icon'
import { formatPrice, formatRequestPrice } from '../lib/price'
import type { PricingModel, TokenUnit } from '../types'
import type { ModelPerfBadgeData } from './model-perf-badge'

export interface ModelCardProps {
  model: PricingModel
  onClick: () => void
  priceRate?: number
  usdExchangeRate?: number
  tokenUnit?: TokenUnit
  showRechargePrice?: boolean
  selectedGroup?: string
  perf?: ModelPerfBadgeData
}

export const ModelCard = memo(function ModelCard(props: ModelCardProps) {
  const { t } = useTranslation()
  const { copyToClipboard } = useCopyToClipboard()
  const tokenUnit = props.tokenUnit ?? DEFAULT_TOKEN_UNIT
  const priceRate = props.priceRate ?? 1
  const usdExchangeRate = props.usdExchangeRate ?? 1
  const showRechargePrice = props.showRechargePrice ?? false
  const isTokenBased = isTokenBasedModel(props.model)
  const tokenUnitLabel = tokenUnit === 'K' ? '1K' : '1M'
  const tags = parseTags(props.model.tags)
  const groups = props.model.enable_groups || []
  const modelIconKey = getPricingVendorIcon(props.model)
  const modelIcon = modelIconKey ? getLobeIcon(modelIconKey, 28) : null
  const initial =
    (props.model.vendor_name?.trim() || props.model.model_name)
      ?.charAt(0)
      .toUpperCase() || '?'
  const isDynamicPricing =
    props.model.billing_mode === 'tiered_expr' &&
    Boolean(props.model.billing_expr)
  const hasCachedPrice = isTokenBased && props.model.cache_ratio != null
  const dynamicSummary = isDynamicPricing
    ? getDynamicPricingSummary(props.model, {
        tokenUnit,
        showRechargePrice,
        priceRate,
        usdExchangeRate,
        groupRatioMultiplier: getDynamicDisplayGroupRatio(
          props.model,
          props.selectedGroup
        ),
      })
    : null

  const primaryGroup = groups[0]

  const handleCopy = (e: React.MouseEvent) => {
    e.stopPropagation()
    copyToClipboard(props.model.model_name || '')
  }

  const handleCardKeyDown = (e: KeyboardEvent<HTMLDivElement>) => {
    if (e.target !== e.currentTarget) return

    if (e.key === 'Enter' || e.key === ' ') {
      e.preventDefault()
      props.onClick()
    }
  }

  let priceSummary: ReactNode
  if (dynamicSummary) {
    if (dynamicSummary.isSpecialExpression) {
      priceSummary = (
        <span className='min-w-0'>
          <span className='text-amber-700 dark:text-amber-300'>
            {t('Special billing expression')}
          </span>
          <code className='text-muted-foreground/70 mt-0.5 line-clamp-1 block font-mono text-[11px] break-all'>
            {dynamicSummary.rawExpression}
          </code>
        </span>
      )
    } else if (dynamicSummary.primaryEntries.length > 0) {
      priceSummary = (
        <>
          {dynamicSummary.primaryEntries.map((entry) => (
            <span
              key={entry.key}
              className='text-muted-foreground whitespace-nowrap'
            >
              {t(entry.shortLabel)}{' '}
              <span className='text-foreground font-mono font-semibold'>
                {entry.formatted}
              </span>
            </span>
          ))}
        </>
      )
    } else {
      priceSummary = (
        <span className='text-muted-foreground text-sm'>
          {t('Dynamic Pricing')}
        </span>
      )
    }
  } else if (isTokenBased) {
    priceSummary = (
      <>
        <span className='text-muted-foreground whitespace-nowrap'>
          {t('Input')}{' '}
          <span className='text-foreground font-mono font-semibold'>
            {formatPrice(
              props.model,
              'input',
              tokenUnit,
              showRechargePrice,
              priceRate,
              usdExchangeRate,
              props.selectedGroup
            )}
          </span>
        </span>
        <span className='text-muted-foreground whitespace-nowrap'>
          {t('Output')}{' '}
          <span className='text-foreground font-mono font-semibold'>
            {formatPrice(
              props.model,
              'output',
              tokenUnit,
              showRechargePrice,
              priceRate,
              usdExchangeRate,
              props.selectedGroup
            )}
          </span>
        </span>
        {hasCachedPrice && (
          <span className='text-muted-foreground whitespace-nowrap'>
            {t('Cached')}{' '}
            <span className='text-foreground font-mono font-semibold'>
              {formatPrice(
                props.model,
                'cache',
                tokenUnit,
                showRechargePrice,
                priceRate,
                usdExchangeRate,
                props.selectedGroup
              )}
            </span>
          </span>
        )}
      </>
    )
  } else {
    priceSummary = (
      <span className='text-muted-foreground whitespace-nowrap'>
        <span className='text-foreground font-mono font-semibold'>
          {formatRequestPrice(
            props.model,
            showRechargePrice,
            priceRate,
            usdExchangeRate,
            props.selectedGroup
          )}
        </span>{' '}
        / {t('request')}
      </span>
    )
  }

  return (
    <div
      className={cn(
        'lanhu-model-card group relative flex cursor-pointer flex-col rounded-xl border p-3 transition-colors sm:p-5',
        'hover:bg-muted/20'
      )}
      role='button'
      tabIndex={0}
      onClick={props.onClick}
      onKeyDown={handleCardKeyDown}
    >
      {/* Header: model identity and actions */}
      <div className='lanhu-model-card-header flex items-start justify-between gap-2.5 sm:gap-3'>
        <div className='lanhu-model-card-identity flex min-w-0 items-start gap-2.5 sm:gap-3'>
          <div
            className='lanhu-model-card-logo bg-muted/40 text-foreground flex size-9 shrink-0 items-center justify-center rounded-lg sm:size-10 sm:rounded-xl'
            aria-hidden='true'
          >
            {modelIcon || (
              <span className='text-muted-foreground text-sm font-bold'>
                {initial}
              </span>
            )}
          </div>
          <div className='min-w-0'>
            <h3
              className='lanhu-model-card-name text-foreground truncate font-mono text-[15px] leading-tight font-bold'
              title={props.model.model_name}
            >
              {props.model.model_name}
            </h3>
          </div>
        </div>

        <div className='lanhu-model-card-actions flex shrink-0 items-center gap-1.5'>
          <button
            type='button'
            onClick={handleCopy}
            className='lanhu-model-card-copy text-muted-foreground hover:text-foreground hover:bg-muted rounded-md border p-1.5 transition-colors'
            aria-label={t('Copy')}
            title={t('Copy')}
          >
            <Copy className='size-3.5' />
          </button>
        </div>
      </div>

      <div className='lanhu-model-tags'>
        {(props.model.vendor_name || primaryGroup) && (
          <span>{props.model.vendor_name || primaryGroup}</span>
        )}
        <span className='lanhu-model-billing-tag'>
          {isDynamicPricing
            ? t('Dynamic Pricing')
            : t(isTokenBased ? 'Token-based' : 'Per Request')}
        </span>
        {tags.slice(0, 2).map((tag) => (
          <span key={tag}>{tag}</span>
        ))}
      </div>
      <div
        className='lanhu-model-prices'
        data-unit={isTokenBased ? `/${tokenUnitLabel}` : undefined}
      >
        {priceSummary}
      </div>

      {/* Description */}
      <p className='lanhu-model-card-description text-muted-foreground mt-2 line-clamp-1 flex-1 text-[13px] leading-relaxed sm:mt-4 sm:line-clamp-2 sm:min-h-[2.5rem]'>
        {props.model.description || t('No description available.')}
      </p>

      <div className='lanhu-model-footer'>
        <span>
          {t('Latency short')}{' '}
          {props.perf && Number.isFinite(props.perf.avg_latency_ms)
            ? `${(props.perf.avg_latency_ms / 1000).toFixed(1)}s`
            : '—'}
        </span>
        <span>
          {t('Context')}{' '}
          {props.model.context_length
            ? new Intl.NumberFormat('en', {
                notation: 'compact',
                maximumFractionDigits: 1,
              }).format(props.model.context_length)
            : '—'}
        </span>
        <button
          type='button'
          onClick={(e) => {
            e.stopPropagation()
            props.onClick()
          }}
        >
          {t('Details')}
          <ArrowRight className='size-3' aria-hidden />
        </button>
      </div>
    </div>
  )
})
