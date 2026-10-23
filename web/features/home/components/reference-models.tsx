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
import { Link } from '@tanstack/react-router'
import { ArrowLeft, ArrowRight } from 'lucide-react'
import { useEffect, useMemo, useRef, useState } from 'react'
import { useTranslation } from 'react-i18next'

import { formatPrice, formatRequestPrice } from '@/features/pricing/lib/price'
import { getLobeIcon } from '@/lib/lobe-icon'

import { useHomeModels } from '../hooks/use-home-models'
import {
  getDomesticHomeModels,
  getHomeModelCategory,
  HOME_MODEL_CATEGORIES,
  type HomeModelCategory,
} from '../lib/home-model-catalog'
import { getHomepageModelIcon } from '../lib/model-brand'
import {
  getLoopedModelScrollPosition,
  getModelScrollEdges,
} from '../lib/model-scroll'
import { HomeParallaxDecorations } from './home-parallax'

export function ReferenceModels() {
  const { t } = useTranslation()
  const data = useHomeModels()
  const [category, setCategory] = useState<HomeModelCategory>('featured')
  const scroller = useRef<HTMLDivElement>(null)
  const loopWidth = useRef(0)
  const measuredModelsKey = useRef('')
  const [edges, setEdges] = useState({ previous: false, next: false })
  const [loopEnabled, setLoopEnabled] = useState(false)
  const [autoScrollPaused, setAutoScrollPaused] = useState(false)
  const domesticModels = useMemo(
    () => getDomesticHomeModels(data.models),
    [data.models]
  )
  const models = useMemo(
    () =>
      domesticModels
        .filter(
          (model) =>
            category === 'featured' || getHomeModelCategory(model) === category
        )
        .slice(0, 12),
    [category, domesticModels]
  )
  const modelsKey = useMemo(
    () => models.map((model) => model.model_name).join('\n'),
    [models]
  )
  const cardItems = useMemo(() => {
    const originals = models.map((model) => ({
      model,
      copy: 'original' as const,
    }))
    if (!loopEnabled) return originals

    return [
      ...models.map((model) => ({ model, copy: 'leading' as const })),
      ...originals,
      ...models.map((model) => ({ model, copy: 'trailing' as const })),
    ]
  }, [loopEnabled, models])
  useEffect(() => {
    const viewport = scroller.current
    if (!viewport) return

    const modelsChanged = measuredModelsKey.current !== modelsKey
    if (modelsChanged) {
      measuredModelsKey.current = modelsKey
      loopWidth.current = 0
      viewport.scrollTo({ left: 0, behavior: 'instant' })
    }
    let preservePosition = !modelsChanged
    const updateEdges = () => {
      if (loopEnabled && loopWidth.current > 1) {
        setEdges((current) =>
          current.previous && current.next
            ? current
            : { previous: true, next: true }
        )
        return
      }
      const next = getModelScrollEdges(
        viewport.scrollLeft,
        viewport.clientWidth,
        viewport.scrollWidth
      )
      setEdges((current) =>
        current.previous === next.previous && current.next === next.next
          ? current
          : next
      )
    }
    const measure = () => {
      const originalCards = viewport.querySelectorAll<HTMLElement>(
        "[data-model-copy='original']"
      )
      const firstOriginal = originalCards.item(0)
      const lastOriginal = originalCards.item(originalCards.length - 1)
      if (!firstOriginal || !lastOriginal) {
        loopWidth.current = 0
        setLoopEnabled(false)
        setEdges({ previous: false, next: false })
        return
      }

      const contentWidth =
        lastOriginal.offsetLeft +
        lastOriginal.offsetWidth -
        firstOriginal.offsetLeft
      const shouldLoop =
        originalCards.length > 1 && contentWidth - viewport.clientWidth > 1
      if (!shouldLoop) {
        loopWidth.current = 0
        setLoopEnabled(false)
        viewport.scrollTo({ left: 0, behavior: 'instant' })
        setEdges({ previous: false, next: false })
        return
      }
      if (!loopEnabled) {
        setLoopEnabled(true)
        return
      }

      const firstTrailing = viewport.querySelector<HTMLElement>(
        "[data-model-copy='trailing']"
      )
      if (!firstTrailing) return
      const nextLoopWidth = firstTrailing.offsetLeft - firstOriginal.offsetLeft
      if (nextLoopWidth <= 1) return

      const previousLoopWidth = loopWidth.current
      const offset =
        preservePosition && previousLoopWidth > 1
          ? (((viewport.scrollLeft - previousLoopWidth) % previousLoopWidth) +
              previousLoopWidth) %
            previousLoopWidth
          : 0
      loopWidth.current = nextLoopWidth
      preservePosition = true
      viewport.scrollTo({
        left: nextLoopWidth + Math.min(offset, nextLoopWidth - 1),
        behavior: 'instant',
      })
      setEdges({ previous: true, next: true })
    }
    const normalizeLoopPosition = () => {
      if (!loopEnabled || loopWidth.current <= 1) return
      const position = getLoopedModelScrollPosition(
        viewport.scrollLeft,
        loopWidth.current,
        0
      )
      if (Math.abs(position.start - viewport.scrollLeft) > 1) {
        viewport.scrollTo({ left: position.start, behavior: 'instant' })
      }
    }

    measure()
    updateEdges()
    const observer = new ResizeObserver(measure)
    observer.observe(viewport)
    viewport.addEventListener('scroll', updateEdges, { passive: true })
    viewport.addEventListener('scrollend', normalizeLoopPosition, {
      passive: true,
    })
    return () => {
      observer.disconnect()
      viewport.removeEventListener('scroll', updateEdges)
      viewport.removeEventListener('scrollend', normalizeLoopPosition)
    }
  }, [data.isLoading, loopEnabled, modelsKey])
  useEffect(() => {
    if (
      !loopEnabled ||
      autoScrollPaused ||
      data.isLoading ||
      data.error ||
      window.matchMedia('(prefers-reduced-motion: reduce)').matches
    ) {
      return
    }

    const timer = window.setInterval(() => {
      const viewport = scroller.current
      if (!viewport) return
      const card = viewport.querySelector<HTMLElement>(
        "[data-model-copy='original']"
      )
      const distance = (card?.getBoundingClientRect().width ?? 420) + 20
      const position = getLoopedModelScrollPosition(
        viewport.scrollLeft,
        loopWidth.current,
        distance
      )
      if (Math.abs(position.start - viewport.scrollLeft) > 1) {
        viewport.scrollTo({ left: position.start, behavior: 'instant' })
      }
      viewport.scrollTo({
        left: position.target,
        behavior: 'smooth',
      })
    }, 4000)

    return () => window.clearInterval(timer)
  }, [autoScrollPaused, data.error, data.isLoading, loopEnabled, models])
  const scroll = (direction: number) => {
    const viewport = scroller.current
    if (!viewport) return
    const card = viewport.querySelector<HTMLElement>(
      "[data-model-copy='original']"
    )
    const distance =
      direction * ((card?.getBoundingClientRect().width ?? 420) + 20)
    const behavior = window.matchMedia('(prefers-reduced-motion: reduce)')
      .matches
      ? 'instant'
      : 'smooth'
    if (loopEnabled) {
      const position = getLoopedModelScrollPosition(
        viewport.scrollLeft,
        loopWidth.current,
        distance
      )
      if (Math.abs(position.start - viewport.scrollLeft) > 1) {
        viewport.scrollTo({ left: position.start, behavior: 'instant' })
      }
      viewport.scrollTo({ left: position.target, behavior })
      return
    }
    viewport.scrollBy({
      left: distance,
      behavior,
    })
  }

  if (data.access !== 'public') {
    return (
      <section className='wm-container wm-models' id='models'>
        <div className='wm-section-heading'>
          <h2>{t('homepage.models.title')}</h2>
          <p>
            {t(
              data.access === 'login'
                ? 'homepage.models.login'
                : 'homepage.models.unavailable'
            )}
          </p>
        </div>
        {data.access === 'login' && (
          <Link
            to='/sign-in'
            search={{ redirect: '/pricing' }}
            className='wm-section-link'
          >
            {t('Sign in')}
            <ArrowRight size={16} />
          </Link>
        )}
      </section>
    )
  }
  return (
    <section
      className='wm-container wm-models wm-home-parallax-section'
      id='models'
    >
      <HomeParallaxDecorations />
      <div className='wm-model-content wm-home-parallax-panel'>
        <div className='wm-section-heading'>
          <h2>{t('homepage.models.title')}</h2>
          <p>{t('homepage.models.subtitle')}</p>
        </div>
        <div className='wm-model-toolbar'>
          <div className='wm-vendor-tabs' role='group' aria-label={t('Vendor')}>
            {HOME_MODEL_CATEGORIES.map((item) => (
              <button
                key={item.id}
                type='button'
                className={category === item.id ? 'is-active' : ''}
                aria-pressed={category === item.id}
                onClick={() => setCategory(item.id)}
              >
                {t(item.labelKey)}
              </button>
            ))}
          </div>
          <div className='wm-scroll-buttons'>
            <button
              type='button'
              onClick={() => scroll(-1)}
              aria-label={t('homepage.models.previous')}
              title={t('homepage.models.previous')}
              disabled={data.isLoading || data.error || !edges.previous}
            >
              <ArrowLeft size={20} />
            </button>
            <button
              type='button'
              onClick={() => scroll(1)}
              aria-label={t('homepage.models.next')}
              title={t('homepage.models.next')}
              disabled={data.isLoading || data.error || !edges.next}
            >
              <ArrowRight size={20} />
            </button>
          </div>
        </div>
        {data.error ? (
          <div className='wm-data-state' role='alert'>
            <p>{t('Failed to load models')}</p>
            <button type='button' onClick={() => void data.refetch()}>
              {t('Retry')}
            </button>
          </div>
        ) : (
          <div
            key={category}
            className='wm-model-track'
            ref={scroller}
            aria-busy={data.isLoading}
            aria-live='off'
            onMouseEnter={() => setAutoScrollPaused(true)}
            onMouseLeave={() => setAutoScrollPaused(false)}
            onFocusCapture={() => setAutoScrollPaused(true)}
            onBlurCapture={() => setAutoScrollPaused(false)}
            onTouchStart={() => setAutoScrollPaused(true)}
            onTouchEnd={() => setAutoScrollPaused(false)}
          >
            {data.isLoading
              ? ['a', 'b', 'c', 'd'].map((id) => (
                  <div
                    key={id}
                    className='wm-model-card wm-skeleton'
                    aria-hidden='true'
                  />
                ))
              : cardItems.map((item) => {
                  const model = item.model
                  const tokens = model.quota_type === 0
                  const price = tokens
                    ? formatPrice(model, 'input', 'M')
                    : formatRequestPrice(model)
                  const icon = getHomepageModelIcon(model)
                  return (
                    <Link
                      key={`${item.copy}-${model.model_name}`}
                      to='/pricing/$modelId'
                      params={{ modelId: model.model_name }}
                      className='wm-model-card'
                      data-model-copy={item.copy}
                      aria-hidden={item.copy !== 'original' || undefined}
                      tabIndex={item.copy === 'original' ? undefined : -1}
                    >
                      <div className='wm-model-card-header'>
                        <div className='wm-model-identity'>
                          <h3 title={model.model_name}>{model.model_name}</h3>
                          <p>@ {model.vendor_name || t('Unknown')}</p>
                        </div>
                        <div className='wm-model-logo'>
                          {icon ? (
                            getLobeIcon(icon, 36)
                          ) : (
                            <span>
                              {(model.vendor_name || model.model_name)
                                .slice(0, 1)
                                .toUpperCase()}
                            </span>
                          )}
                        </div>
                      </div>
                      <div className='wm-model-pricing'>
                        <p>
                          {t(tokens ? 'Input' : 'Price')}
                          <strong>{price}</strong>
                          <span>/ {tokens ? 'M tokens' : t('Request')}</span>
                        </p>
                        {tokens && (
                          <p>
                            {t('Output')}
                            <strong>{formatPrice(model, 'output', 'M')}</strong>
                            <span>/ M tokens</span>
                          </p>
                        )}
                      </div>
                    </Link>
                  )
                })}
            {!data.isLoading && models.length === 0 && (
              <p className='wm-data-state'>{t('No models found')}</p>
            )}
          </div>
        )}
        <Link to='/pricing' className='wm-section-link'>
          {t('homepage.models.explore')}
          <ArrowRight size={16} />
        </Link>
      </div>
    </section>
  )
}
