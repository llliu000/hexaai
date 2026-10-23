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
import Autoplay from 'embla-carousel-autoplay'
import Fade from 'embla-carousel-fade'
import useEmblaCarousel from 'embla-carousel-react'
import { ArrowRight } from 'lucide-react'
import { useReducedMotion } from 'motion/react'
import { useEffect, useMemo, useRef, useState } from 'react'
import { useTranslation } from 'react-i18next'

import { SystemLogo } from '@/components/system-logo'
import { useSystemConfig } from '@/hooks/use-system-config'
import { getLobeIcon } from '@/lib/lobe-icon'
import { resolveSystemLogo, SYSTEM_LOGO_URL } from '@/lib/system-logo'

import { attachGroupedHomeParallax } from '../lib/grouped-home-parallax'
import {
  getHeroText,
  getVisibleHeroSlides,
  type HeroConfig,
  type HeroSlide,
} from '../lib/hero-config'
import { HeroHalo } from './hero-halo'

const BRAND_ICONS = {
  openai: 'OpenAI',
  deepseek: 'DeepSeek.Color',
  qwen: 'Qwen.Color',
  zhipu: 'ZAI',
  moonshot: 'Kimi.Color',
}

function HeroLogo(props: { slide: HeroSlide; logo: string }) {
  const [failed, setFailed] = useState(false)
  const image = props.slide.logoUrl
  if (image && resolveSystemLogo(image) === SYSTEM_LOGO_URL) {
    return <SystemLogo src={image} width={36} height={36} />
  }
  if (image && !failed) {
    return (
      <img
        src={image}
        alt=''
        width={36}
        height={36}
        onError={() => setFailed(true)}
      />
    )
  }
  if (props.slide.icon === 'platform' || props.slide.icon === 'custom') {
    return <SystemLogo src={props.logo} width={36} height={36} />
  }
  if (props.slide.icon in BRAND_ICONS) {
    return getLobeIcon(
      BRAND_ICONS[props.slide.icon as keyof typeof BRAND_ICONS],
      36
    )
  }
  return (
    <span className='wm-logo-fallback' aria-hidden='true'>
      AI
    </span>
  )
}

function HeroAction(props: {
  url: string
  newTab: boolean
  label: string
  primary?: boolean
  preview?: boolean
}) {
  const className = `wm-button ${props.primary ? 'wm-button-primary' : 'wm-button-secondary'}`
  const contents = (
    <>
      {props.label}
      {props.primary && <ArrowRight size={18} />}
    </>
  )
  if (props.preview) {
    return (
      <button type='button' className={className} disabled>
        {contents}
      </button>
    )
  }
  if (props.url.startsWith('/')) {
    return (
      <Link
        className={className}
        to={props.url}
        target={props.newTab ? '_blank' : undefined}
        rel={props.newTab ? 'noopener noreferrer' : undefined}
      >
        {contents}
      </Link>
    )
  }
  return (
    <a
      className={className}
      href={props.url}
      target={props.newTab ? '_blank' : undefined}
      rel='noopener noreferrer'
    >
      {contents}
    </a>
  )
}

export function HeroCarousel(props: {
  config: HeroConfig
  preview?: boolean
  locale?: string
}) {
  const { t, i18n } = useTranslation()
  const { systemName, logo } = useSystemConfig()
  const reducedMotion = useReducedMotion()
  const slides = useMemo(
    () => getVisibleHeroSlides(props.config),
    [props.config]
  )
  const sectionRef = useRef<HTMLElement>(null)
  const plugins = useMemo(
    () => [
      Fade(),
      Autoplay({
        delay: props.config.intervalMs,
        playOnInit: false,
        stopOnInteraction: true,
        stopOnMouseEnter: false,
        stopOnFocusIn: false,
      }),
    ],
    [props.config.intervalMs]
  )
  const [viewportRef, api] = useEmblaCarousel(
    {
      loop: props.config.loop && slides.length > 1,
      watchDrag: !props.preview,
      duration: reducedMotion ? 0 : 35,
    },
    plugins
  )
  const [selected, setSelected] = useState(0)
  const [focused, setFocused] = useState(false)
  const [visible, setVisible] = useState(() => !document.hidden)
  useEffect(() => {
    const change = () => setVisible(!document.hidden)
    document.addEventListener('visibilitychange', change)
    return () => document.removeEventListener('visibilitychange', change)
  }, [])
  useEffect(() => {
    const hero = sectionRef.current
    if (!hero || props.preview || reducedMotion) return

    return attachGroupedHomeParallax(hero, [
      { selector: '.wm-hero-parallax-art', speed: -3 },
      { selector: '.wm-hero-parallax-copy', speed: -1 },
      { selector: '.wm-hero-parallax-decoration', speed: 5 },
    ])
  }, [props.preview, reducedMotion, slides.length])
  useEffect(() => {
    if (!api) return
    const onSelect = () => setSelected(api.selectedScrollSnap())
    onSelect()
    api.on('select', onSelect)
    api.on('reInit', onSelect)
    return () => {
      api.off('select', onSelect)
      api.off('reInit', onSelect)
    }
  }, [api])
  useEffect(() => {
    if (!api) return
    const autoplay = api.plugins().autoplay
    const sync = () => {
      if (
        props.config.enabled &&
        props.config.autoplay &&
        slides.length > 1 &&
        !props.preview &&
        !reducedMotion &&
        visible &&
        !focused
      ) {
        autoplay?.play()
      } else autoplay?.stop()
    }
    sync()
    api.on('pointerUp', sync)
    return () => {
      api.off('pointerUp', sync)
      autoplay?.stop()
    }
  }, [
    api,
    props.config.enabled,
    props.config.autoplay,
    props.preview,
    slides.length,
    reducedMotion,
    visible,
    focused,
  ])
  return (
    <section
      ref={sectionRef}
      className={`wm-hero${props.preview ? ' wm-hero-preview' : ''}`}
      aria-roledescription='carousel'
      aria-label={t('homepage.carousel')}
      onFocusCapture={() => setFocused(true)}
      onBlurCapture={(event) => {
        if (!event.currentTarget.contains(event.relatedTarget)) {
          setFocused(false)
        }
      }}
      onKeyDown={(event) => {
        if (event.key === 'ArrowLeft') {
          event.preventDefault()
          api?.scrollPrev()
        }
        if (event.key === 'ArrowRight') {
          event.preventDefault()
          api?.scrollNext()
        }
      }}
    >
      <HeroHalo
        theme={slides[Math.min(selected, slides.length - 1)].theme}
        slideId={slides[Math.min(selected, slides.length - 1)].id}
        reducedMotion={!!reducedMotion}
      />
      <div className='wm-hero-viewport' ref={viewportRef}>
        <div className='wm-hero-track'>
          {slides.map((slide, index) => {
            const text = getHeroText(
              slide,
              props.locale || i18n.resolvedLanguage || i18n.language,
              props.config.defaultLocale,
              systemName
            )
            const Heading = index === selected && !props.preview ? 'h1' : 'h2'
            return (
              <article
                key={slide.id}
                className={`wm-hero-slide wm-theme-${slide.theme}`}
                aria-hidden={index !== selected}
                inert={index !== selected}
              >
                <div className='wm-hero-parallax-art' aria-hidden='true'>
                  {slide.desktopImage ? (
                    <picture className='wm-hero-art'>
                      <source
                        media='(max-width: 767px)'
                        srcSet={slide.mobileImage || slide.desktopImage}
                      />
                      <img
                        src={slide.desktopImage}
                        alt=''
                        style={{ objectPosition: slide.imagePosition }}
                        onError={(event) => {
                          event.currentTarget.style.visibility = 'hidden'
                        }}
                      />
                    </picture>
                  ) : (
                    <img
                      className='wm-hero-art wm-hero-default-art'
                      src='/lanhu/home-hero-background.png'
                      alt=''
                    />
                  )}
                </div>
                <span
                  className='wm-hero-decoration wm-hero-decoration-left wm-hero-parallax-decoration'
                  aria-hidden='true'
                >
                  <svg
                    className='wm-hero-decoration-shape'
                    viewBox='0 0 160 144'
                    width={160}
                    height={144}
                    fill='currentColor'
                    focusable='false'
                  >
                    <path d='M27 40 109 12Q125 7 133 22L155 76Q162 94 147 105L128 119Q121 124 113 123L28 106Q11 102 10 85L8 64Q6 48 27 40Z' />
                  </svg>
                </span>
                <span
                  className='wm-hero-decoration wm-hero-decoration-right wm-hero-parallax-decoration'
                  aria-hidden='true'
                >
                  <svg
                    className='wm-hero-decoration-shape'
                    viewBox='0 0 160 160'
                    width={160}
                    height={160}
                    fill='currentColor'
                    focusable='false'
                  >
                    <path d='M75 16Q88 4 101 16L148 63Q162 77 149 91L102 140Q88 154 74 141L15 96Q2 86 14 72Z' />
                  </svg>
                </span>
                <div className='wm-hero-copy wm-hero-parallax-copy'>
                  <div className='wm-hero-badge'>
                    <HeroLogo
                      key={`${slide.id}-${slide.logoUrl}`}
                      slide={slide}
                      logo={logo}
                    />
                    <span>{text.badge}</span>
                  </div>
                  <Heading>
                    {text.titlePrefix}
                    <strong>{text.highlight}</strong>
                    {text.titleSuffix}
                  </Heading>
                  <p>{text.description}</p>
                  <div className='wm-hero-actions'>
                    <HeroAction
                      {...slide.primaryAction}
                      label={text.primaryLabel}
                      primary
                      preview={props.preview}
                    />
                    {slide.secondaryAction && (
                      <HeroAction
                        {...slide.secondaryAction}
                        label={text.secondaryLabel}
                        preview={props.preview}
                      />
                    )}
                  </div>
                </div>
              </article>
            )
          })}
        </div>
      </div>
      {slides.length > 1 && (
        <div className='wm-hero-controls'>
          <div className='wm-hero-dots'>
            {slides.map((slide, index) => (
              <button
                type='button'
                key={slide.id}
                aria-label={t('homepage.slide', { number: index + 1 })}
                aria-current={selected === index ? 'true' : undefined}
                onClick={() => api?.scrollTo(index)}
              >
                <span />
              </button>
            ))}
          </div>
        </div>
      )}
    </section>
  )
}
