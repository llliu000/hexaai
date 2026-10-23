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
import {
  ArrowRight,
  Globe2,
  Layers3,
  Quote,
  ShieldCheck,
  Wallet,
} from 'lucide-react'
import { useRef } from 'react'
import { useTranslation } from 'react-i18next'

import { Footer } from '@/components/layout/components/footer'
import { useStatus } from '@/hooks/use-status'
import { useSystemConfig } from '@/hooks/use-system-config'
import { getLobeIcon } from '@/lib/lobe-icon'

import { useHomeHero } from '../hooks/use-home-hero'
import { useHomeParallax } from '../hooks/use-home-parallax'
import { createDefaultHeroConfig } from '../lib/hero-config'
import { HOME_MODEL_PROVIDERS } from '../lib/home-model-catalog'
import { HeroCarousel } from './hero-carousel'
import { HomeParallaxDecorations } from './home-parallax'
import { ReferenceHeader } from './reference-header'
import { ReferenceIntegration } from './reference-integration'
import { ReferenceModels } from './reference-models'

import '@/styles/reference-home.css'

const BENEFITS = [
  { id: 'routing', Icon: Globe2 },
  { id: 'models', Icon: Layers3 },
  { id: 'security', Icon: ShieldCheck },
  { id: 'billing', Icon: Wallet },
]
const TESTIMONIALS = ['lin', 'chen', 'zhou', 'wang', 'zhao', 'liu'] as const
const DEFAULT_HERO = createDefaultHeroConfig()

export function ReferenceHome() {
  const homeRef = useRef<HTMLDivElement>(null)
  useHomeParallax(homeRef)
  const { t } = useTranslation()
  const { status } = useStatus()
  const { footerHtml, systemName } = useSystemConfig()
  const hero = useHomeHero()
  const docsUrl =
    typeof status?.docs_link === 'string' && status.docs_link
      ? status.docs_link
      : 'https://docs.newapi.pro'
  const steps = [
    { id: 'account', href: '/sign-up' },
    { id: 'key', href: '/keys' },
    { id: 'connect', href: docsUrl },
    { id: 'monitor', href: '/dashboard' },
  ]
  const footerColumns = [
    {
      title: 'homepage.nav.products',
      links: [
        { text: 'Model Square', href: '/pricing' },
        { text: 'Playground', href: '/playground' },
        { text: 'API Keys', href: '/keys' },
        { text: 'Wallet', href: '/wallet' },
      ],
    },
    {
      title: 'homepage.footer.support',
      links: [
        { text: 'Docs', href: docsUrl },
        { text: 'Console', href: '/dashboard' },
        { text: 'Profile', href: '/profile' },
      ],
    },
    {
      title: 'homepage.footer.legal',
      links: [
        ...(status?.user_agreement_enabled
          ? [{ text: 'User Agreement', href: '/user-agreement' }]
          : []),
        ...(status?.privacy_policy_enabled
          ? [{ text: 'Privacy Policy', href: '/privacy-policy' }]
          : []),
      ],
    },
  ]
  return (
    <div className='wm-home' ref={homeRef}>
      <ReferenceHeader docsUrl={docsUrl} />
      <main>
        <HeroCarousel config={hero.data?.config ?? DEFAULT_HERO} />
        <ReferenceModels />
        <ReferenceIntegration docsUrl={docsUrl} />
        <section
          className='wm-container wm-benefits wm-home-parallax-section'
          id='advantages'
        >
          <HomeParallaxDecorations />
          <div className='wm-section-heading wm-home-parallax-heading'>
            <h2>{t('homepage.benefits.title', { brand: systemName })}</h2>
            <p>{t('homepage.benefits.subtitle')}</p>
          </div>
          <div className='wm-four-grid wm-home-parallax-panel'>
            {BENEFITS.map(({ id, Icon }) => (
              <article className={`wm-benefit wm-benefit-${id}`} key={id}>
                <span className='wm-feature-icon'>
                  <Icon size={28} />
                </span>
                <h3>{t(`homepage.benefits.${id}.title`)}</h3>
                <p>{t(`homepage.benefits.${id}.description`)}</p>
              </article>
            ))}
          </div>
        </section>
        <section
          className='wm-testimonials wm-home-parallax-section'
          id='capabilities'
        >
          <HomeParallaxDecorations />
          <div className='wm-container'>
            <div className='wm-section-heading wm-home-parallax-heading'>
              <h2>{t('homepage.testimonials.title')}</h2>
              <p>{t('homepage.testimonials.subtitle')}</p>
            </div>
            <div className='wm-testimonial-grid wm-home-parallax-panel'>
              {TESTIMONIALS.map((id, index) => (
                <article className='wm-testimonial' key={id}>
                  <blockquote>
                    <Quote
                      className='wm-testimonial-quote wm-testimonial-quote-open'
                      aria-hidden='true'
                    />
                    <p>
                      {t(`homepage.testimonials.${id}.quote`, {
                        brand: systemName,
                      })}
                    </p>
                    <Quote
                      className='wm-testimonial-quote wm-testimonial-quote-close'
                      aria-hidden='true'
                    />
                  </blockquote>
                  <div className='wm-testimonial-person'>
                    <img
                      src={`/lanhu/testimonial-${index + 1}.png`}
                      alt=''
                      loading='lazy'
                    />
                    <div>
                      <strong>{t(`homepage.testimonials.${id}.name`)}</strong>
                      <span>{t(`homepage.testimonials.${id}.role`)}</span>
                    </div>
                  </div>
                </article>
              ))}
            </div>
          </div>
        </section>
        <section
          className='wm-container wm-steps wm-home-parallax-section'
          id='get-started'
        >
          <HomeParallaxDecorations />
          <div className='wm-section-heading wm-home-parallax-heading'>
            <h2>{t('homepage.steps.title')}</h2>
            <p>{t('homepage.steps.subtitle')}</p>
          </div>
          <ol className='wm-four-grid wm-steps-track wm-home-parallax-panel'>
            {steps.map(({ id, href }, index) => (
              <li key={id}>
                <article>
                  <div className='wm-step-number'>0{index + 1}</div>
                  <h3>{t(`homepage.steps.${id}.title`)}</h3>
                  <p>{t(`homepage.steps.${id}.description`)}</p>
                  {href.startsWith('http') ? (
                    <a href={href} target='_blank' rel='noreferrer'>
                      {t(`homepage.steps.${id}.action`)}
                      <ArrowRight size={16} />
                    </a>
                  ) : (
                    <Link to={href}>
                      {t(`homepage.steps.${id}.action`)}
                      <ArrowRight size={16} />
                    </Link>
                  )}
                </article>
              </li>
            ))}
          </ol>
        </section>
        <section className='wm-providers wm-home-parallax-section'>
          <HomeParallaxDecorations />
          <div className='wm-container'>
            <div className='wm-section-heading wm-home-parallax-heading'>
              <h2>{t('homepage.providers.title')}</h2>
              <p>{t('homepage.providers.subtitle')}</p>
            </div>
            <div className='wm-provider-row wm-home-parallax-panel'>
              {HOME_MODEL_PROVIDERS.map((provider) => (
                <div key={provider.id}>
                  {getLobeIcon(provider.icon, 36)}
                  <span>{provider.name}</span>
                </div>
              ))}
            </div>
            <Link to='/pricing' className='wm-button wm-button-primary'>
              {t('homepage.models.explore')}
              <ArrowRight size={18} />
            </Link>
          </div>
        </section>
      </main>
      <Footer
        columns={footerColumns}
        description={t('homepage.footer.description')}
        showColumns
        className='wm-footer'
        extraContent={
          <div className='wm-footer-wechat'>
            <p>{t('homepage.footer.wechatSupport')}</p>
            <img
              src='/lanhu/home-service-qr.png'
              alt={t('homepage.footer.wechatSupport')}
              width={120}
              height={120}
              loading='lazy'
            />
          </div>
        }
      />
      {!footerHtml && (
        <div className='wm-footer-records'>
          <a href='https://beian.miit.gov.cn/' target='_blank' rel='noreferrer'>
            粤ICP备2023020253号-1
          </a>
          <span>电信业务经营许可证粤B1.B2-20263198</span>
        </div>
      )}
    </div>
  )
}
