import { useTranslation } from 'react-i18next'

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
import { Main } from '@/components/layout'
import { useStatus } from '@/hooks/use-status'
import { useAuthStore } from '@/stores/auth-store'

import { CheckinCalendarCard } from './components/checkin-calendar-card'
import { LanguagePreferencesCard } from './components/language-preferences-card'
import { LoginSessionsCard } from './components/login-sessions-card'
import { ProfileHeader } from './components/profile-header'
import { ProfileSecurityCard } from './components/profile-security-card'
import { ProfileSettingsCard } from './components/profile-settings-card'
import { SidebarModulesCard } from './components/sidebar-modules-card'
import { useProfile } from './hooks'

import './profile-design.css'

export function Profile() {
  const { t } = useTranslation()
  const { profile, loading, refreshProfile } = useProfile()
  const { status } = useStatus()
  const permissions = useAuthStore((s) => s.auth.user?.permissions)

  const checkinEnabled = status?.checkin_enabled === true
  const turnstileEnabled = !!(
    status?.turnstile_check && status?.turnstile_site_key
  )
  const turnstileSiteKey = status?.turnstile_site_key || ''
  const canConfigureSidebar = permissions?.sidebar_settings !== false
  const hasSideContent = checkinEnabled || canConfigureSidebar

  return (
    <Main>
      <div className='profile-design'>
        <div className='profile-layout-header'>
          <div className='profile-page-title'>
            <h1>{t('Profile')}</h1>
            <p>
              {t(
                'Manage your account information, security settings and personal preferences',
                { defaultValue: '管理您的账户信息、安全设置与个人偏好' }
              )}
            </p>
          </div>
        </div>
        <div className='profile-content-scroll'>
          <div
            className={`profile-columns${hasSideContent ? '' : ' profile-columns-full'}`}
          >
            <div className='profile-left'>
              <div className='profile-summary'>
              <ProfileHeader profile={profile} loading={loading} />
            </div>
            <div className='profile-main'>
              <ProfileSettingsCard
                profile={profile}
                loading={loading}
                onProfileUpdate={refreshProfile}
              />
              <LanguagePreferencesCard
                profile={profile}
                onProfileUpdate={refreshProfile}
              />
              <ProfileSecurityCard profile={profile} loading={loading} />
              <LoginSessionsCard />
            </div>
          </div>
            {hasSideContent && (
              <div className='profile-side'>
                {checkinEnabled && (
                  <CheckinCalendarCard
                    checkinEnabled={checkinEnabled}
                    turnstileEnabled={turnstileEnabled}
                    turnstileSiteKey={turnstileSiteKey}
                  />
                )}
                {canConfigureSidebar && <SidebarModulesCard />}
                {/* <PasskeyCard loading={loading} /> */}
                {/* <TwoFACard loading={loading} /> */}
              </div>
            )}
          </div>
        </div>
      </div>
    </Main>
  )
}
