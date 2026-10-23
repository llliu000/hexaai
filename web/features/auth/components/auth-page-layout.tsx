import { Link } from '@tanstack/react-router'
import { Check, Globe2, ShieldCheck } from 'lucide-react'
import type { ReactNode } from 'react'
import { useTranslation } from 'react-i18next'

import { SystemLogo } from '@/components/system-logo'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import { useStatus } from '@/hooks/use-status'
import { useSystemConfig } from '@/hooks/use-system-config'
import {
  INTERFACE_LANGUAGE_OPTIONS,
  normalizeInterfaceLanguage,
  toIntlLocale,
} from '@/i18n/languages'

import { AuthPageShowcase } from './auth-page-showcase'

import '@/styles/auth-pages.css'

function AuthLanguageMenu() {
  const { t, i18n } = useTranslation()
  const language = normalizeInterfaceLanguage(
    i18n.resolvedLanguage || i18n.language
  )
  return (
    <DropdownMenu modal={false}>
      <DropdownMenuTrigger
        render={
          <button
            type='button'
            className='auth-page-language'
            aria-label={t('Change language')}
          />
        }
      >
        <Globe2 size={20} aria-hidden='true' />
        <span>{(toIntlLocale(language) || 'en').toUpperCase()}</span>
      </DropdownMenuTrigger>
      <DropdownMenuContent align='end' className='auth-page-language-menu'>
        {INTERFACE_LANGUAGE_OPTIONS.map((option) => (
          <DropdownMenuItem
            key={option.code}
            onClick={() => {
              void i18n.changeLanguage(option.code)
            }}
          >
            <span>{option.label}</span>
            {language === option.code && <Check size={16} aria-hidden='true' />}
          </DropdownMenuItem>
        ))}
      </DropdownMenuContent>
    </DropdownMenu>
  )
}

export function AuthPageLayout(props: {
  mode: 'sign-in' | 'sign-up'
  redirectTo?: string
  children: ReactNode
}) {
  const { t } = useTranslation()
  const { systemName, logo } = useSystemConfig()
  const { status } = useStatus()
  const showSignUp =
    props.mode === 'sign-up' ||
    (!status?.self_use_mode_enabled && status?.register_enabled !== false)

  return (
    <div className={`auth-page auth-page-${props.mode}`}>
      <header className='auth-page-header'>
        <Link to='/' className='auth-page-brand' aria-label={systemName}>
          <SystemLogo src={logo} width={36} height={36} />
          <span>{systemName}</span>
        </Link>
        <AuthLanguageMenu />
      </header>
      <main className='auth-page-main'>
        <section className='auth-page-panel' aria-labelledby='auth-page-title'>
          <div className='auth-page-heading'>
            <h1 id='auth-page-title'>
              {t('auth.page.welcome', { brand: systemName })}
            </h1>
            <p>{t('auth.page.subtitle')}</p>
          </div>
          <div className='auth-page-note'>
            <ShieldCheck size={17} aria-hidden='true' />
            <span>{t('auth.page.hint')}</span>
          </div>
          <nav
            className='auth-page-modes'
            aria-label={t('auth.page.navigation')}
          >
            <Link
              to='/sign-in'
              search={{ redirect: props.redirectTo }}
              aria-current={props.mode === 'sign-in' ? 'page' : undefined}
            >
              {t('Sign in')}
            </Link>
            {showSignUp && (
              <Link
                to='/sign-up'
                aria-current={props.mode === 'sign-up' ? 'page' : undefined}
              >
                {t('Sign up')}
              </Link>
            )}
          </nav>
          <div className='auth-page-fields'>{props.children}</div>
        </section>
        <AuthPageShowcase brand={systemName} />
      </main>
    </div>
  )
}
