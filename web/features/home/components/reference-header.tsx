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
import { Menu } from 'lucide-react'
import { useEffect, useState } from 'react'
import { useTranslation } from 'react-i18next'

import { LanguageSwitcher } from '@/components/language-switcher'
import { ProfileDropdown } from '@/components/profile-dropdown'
import { SystemLogo } from '@/components/system-logo'
import { Button } from '@/components/ui/button'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import { useSystemConfig } from '@/hooks/use-system-config'
import { useAuthStore } from '@/stores/auth-store'

import { watchHeaderScroll } from '../lib/lanhu-home-interactions'

export function ReferenceHeader(props: { docsUrl: string }) {
  const { t } = useTranslation()
  const { logo, systemName } = useSystemConfig()
  const user = useAuthStore((state) => state.auth.user)
  const [scrolled, setScrolled] = useState(false)
  useEffect(() => watchHeaderScroll(window, setScrolled), [])
  return (
    <header className={`wm-header${scrolled ? ' is-scrolled' : ''}`}>
      <div className='wm-container wm-header-inner'>
        <Link to='/' className='wm-brand'>
          <SystemLogo src={logo} width={36} height={36} />
          <span>{systemName}</span>
        </Link>
        <nav className='wm-desktop-nav' aria-label={t('Navigation')}>
          <Link to='/'>{t('Home')}</Link>
          <Link to='/dashboard'>{t('Console')}</Link>
          <Link to='/pricing'>{t('Model Square')}</Link>
          <a href={props.docsUrl} target='_blank' rel='noreferrer'>
            {t('Docs')}
          </a>
        </nav>
        <div className='wm-header-actions'>
          <LanguageSwitcher />
          {user ? (
            <ProfileDropdown />
          ) : (
            <Link to='/sign-in' className='wm-login'>
              {t('Sign in')}
            </Link>
          )}
          <DropdownMenu>
            <DropdownMenuTrigger
              render={
                <Button
                  variant='ghost'
                  size='icon'
                  className='wm-mobile-menu'
                  aria-label={t('Toggle navigation menu')}
                />
              }
            >
              <Menu size={22} />
            </DropdownMenuTrigger>
            <DropdownMenuContent align='end'>
              <DropdownMenuItem render={<Link to='/' />}>
                {t('Home')}
              </DropdownMenuItem>
              <DropdownMenuItem render={<Link to='/dashboard' />}>
                {t('Console')}
              </DropdownMenuItem>
              <DropdownMenuItem render={<Link to='/pricing' />}>
                {t('Model Square')}
              </DropdownMenuItem>
              <DropdownMenuItem
                render={
                  <a href={props.docsUrl} target='_blank' rel='noreferrer' />
                }
              >
                {t('Docs')}
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        </div>
      </div>
    </header>
  )
}
