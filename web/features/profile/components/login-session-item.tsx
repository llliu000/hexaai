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
import { useTranslation } from 'react-i18next'

import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import dayjs from '@/lib/dayjs'
import type { LoginSession } from '@/stores/auth-store'

import { DesignIcon } from './design-icon'
import { loginMethodLabel, sessionDevice } from './login-session-utils'

interface LoginSessionItemProps {
  session: LoginSession
  onRevoke: (session: LoginSession) => void
}

export function LoginSessionItem({ session, onRevoke }: LoginSessionItemProps) {
  const { t } = useTranslation()
  const maxTouchPoints =
    session.current && typeof navigator !== 'undefined'
      ? navigator.maxTouchPoints
      : 0

  return (
    <div className='profile-session-row'>
      <div className='profile-session-device' aria-hidden='true'>
        <DesignIcon name='laptop' />
      </div>
      <div className='profile-session-details'>
        <div className='profile-session-identity'>
          <p className='profile-session-name'>
            {sessionDevice(
              session.user_agent,
              t('Unknown device'),
              t('Browser'),
              maxTouchPoints
            )}
          </p>
          {session.current && <Badge variant='secondary'>{t('Current')}</Badge>}
        </div>
        <p className='profile-session-meta'>
          {t('IP: {{ip}} · Method: {{method}}', {
            ip: session.ip || t('Unknown'),
            method: loginMethodLabel(session.login_method, t),
          })}
        </p>
        <p className='profile-session-time'>
          {t('Last active {{time}} · Expires {{expires}}', {
            time: dayjs.unix(session.last_active_at).fromNow(),
            expires: dayjs.unix(session.expires_at).format('YYYY-MM-DD HH:mm'),
          })}
        </p>
      </div>
      <Button
        type='button'
        variant={session.current ? 'outline' : 'ghost'}
        size='sm'
        className='profile-session-action'
        onClick={() => onRevoke(session)}
      >
        <DesignIcon name='logout' />
        {session.current ? t('Sign out') : t('Revoke')}
      </Button>
    </div>
  )
}
