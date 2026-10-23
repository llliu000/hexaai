import { Link, useLocation } from '@tanstack/react-router'
import { ChevronDown, Menu, X } from 'lucide-react'
import { useEffect, useRef, useState } from 'react'
import { useTranslation } from 'react-i18next'

import { LanguageSwitcher } from '@/components/language-switcher'
import { SYSTEM_SETTINGS_VIEW } from '@/components/layout/config/system-settings.config'
import { checkIsActive } from '@/components/layout/lib/url-utils'
import { NotificationPopover } from '@/components/notification-popover'
import { ProfileDropdown } from '@/components/profile-dropdown'
import { Search } from '@/components/search'
import { SystemLogo } from '@/components/system-logo'
import { useNotifications } from '@/hooks/use-notifications'
import { useIsSidebarModuleVisible } from '@/hooks/use-sidebar-config'
import { useStatus } from '@/hooks/use-status'
import { useSystemConfig } from '@/hooks/use-system-config'
import { ROLE } from '@/lib/roles'
import { useAuthStore } from '@/stores/auth-store'

import profileWordmark from '../assets/profile-wordmark.svg'
import { DesignIcon, type DesignIconName } from './design-icon'

type ProfileNavPath =
  | '/dashboard/overview'
  | '/dashboard/models'
  | '/keys'
  | '/usage-logs/common'
  | '/usage-logs/task'
  | '/wallet'
  | '/profile'
  | '/channels'
  | '/models'
  | '/users'
  | '/redemption-codes'
  | '/subscriptions'
  | '/system-info'
  | '/system-settings'

function ProfileNavItem(props: {
  to: ProfileNavPath
  label: string
  icon: DesignIconName
}) {
  const visible = useIsSidebarModuleVisible(props.to)
  const pathname = useLocation({
    select: (location) => location.pathname.replace(/\/$/, ''),
  })
  const matchesNestedAdminRoute =
    (props.to === '/models' || props.to === '/system-settings') &&
    pathname.startsWith(`${props.to}/`)
  const current =
    pathname === props.to ||
    matchesNestedAdminRoute ||
    (props.to === '/dashboard/models' &&
      (pathname === '/dashboard/flow' || pathname === '/dashboard/users'))
  if (!visible) return null
  const content = (
    <>
      <DesignIcon name={props.icon} className='profile-nav-icon' />
      <span>{props.label}</span>
    </>
  )
  if (props.to === '/dashboard/overview' || props.to === '/dashboard/models') {
    return (
      <Link
        to='/dashboard/$section'
        params={{
          section: props.to === '/dashboard/overview' ? 'overview' : 'models',
        }}
        className={`profile-nav-item ${current ? 'is-current' : ''}`}
        aria-current={current ? 'page' : undefined}
      >
        {content}
      </Link>
    )
  }
  if (props.to === '/usage-logs/common' || props.to === '/usage-logs/task') {
    return (
      <Link
        to='/usage-logs/$section'
        params={{
          section: props.to === '/usage-logs/common' ? 'common' : 'task',
        }}
        className={`profile-nav-item ${current ? 'is-current' : ''}`}
        aria-current={current ? 'page' : undefined}
      >
        {content}
      </Link>
    )
  }
  return (
    <Link
      to={props.to}
      className={`profile-nav-item ${current ? 'is-current' : ''}`}
      aria-current={current ? 'page' : undefined}
    >
      {content}
    </Link>
  )
}

function ProfileSystemSettingsNavigation() {
  const { t } = useTranslation()
  const href = useLocation({ select: (location) => location.href })
  const groups = SYSTEM_SETTINGS_VIEW.getNavGroups(t)

  return (
    <div className='profile-system-settings-nav'>
      {groups
        .flatMap((group) => group.items)
        .map((item) => {
          if (!item.items) return null
          const Icon = item.icon
          const active = checkIsActive(href, item)

          return (
            <details key={item.title} open={active}>
              <summary className={active ? 'is-active' : undefined}>
                {Icon && <Icon aria-hidden='true' />}
                <span>{item.title}</span>
                <ChevronDown
                  aria-hidden='true'
                  className='profile-settings-arrow'
                />
              </summary>
              <div className='profile-system-settings-subnav'>
                {item.items.map((subItem) => {
                  const subItemActive = checkIsActive(href, subItem)
                  return (
                    <Link
                      key={subItem.title}
                      to={subItem.url}
                      className={subItemActive ? 'is-current' : undefined}
                      aria-current={subItemActive ? 'page' : undefined}
                    >
                      {subItem.title}
                    </Link>
                  )
                })}
              </div>
            </details>
          )
        })}
    </div>
  )
}

export function ProfileNavigation() {
  const { t } = useTranslation()
  const { logo, systemName } = useSystemConfig()
  const userRole = useAuthStore((state) => state.auth.user?.role ?? ROLE.GUEST)
  const isAdmin = userRole >= ROLE.ADMIN
  const isSuperAdmin = userRole >= ROLE.SUPER_ADMIN
  const pathname = useLocation({
    select: (location) => location.pathname.replace(/\/$/, ''),
  })
  const isProfilePage = pathname === '/profile'
  const isSystemSettingsPage =
    pathname === '/system-settings' || pathname.startsWith('/system-settings/')
  const { status } = useStatus()
  const notifications = useNotifications()
  const [mobileOpen, setMobileOpen] = useState(false)
  const sidebarRef = useRef<HTMLElement>(null)
  useEffect(() => {
    if (!mobileOpen) return
    const previousFocus = document.activeElement
    sidebarRef.current?.querySelector<HTMLButtonElement>('button')?.focus()
    const handleKey = (event: KeyboardEvent) => {
      if (event.key === 'Escape') setMobileOpen(false)
      if (event.key !== 'Tab') return
      const controls = [
        ...(sidebarRef.current?.querySelectorAll<HTMLElement>(
          'a[href], button:not([disabled])'
        ) ?? []),
      ]
      const first = controls[0]
      const last = controls.at(-1)
      if (event.shiftKey && document.activeElement === first) {
        event.preventDefault()
        last?.focus()
      } else if (!event.shiftKey && document.activeElement === last) {
        event.preventDefault()
        first?.focus()
      }
    }
    document.addEventListener('keydown', handleKey)
    return () => {
      document.removeEventListener('keydown', handleKey)
      if (previousFocus instanceof HTMLElement) previousFocus.focus()
    }
  }, [mobileOpen])
  const docs =
    typeof status?.docs_link === 'string'
      ? status.docs_link
      : 'https://docs.666api.ai/'

  return (
    <>
      <header className='profile-design-header'>
        <button
          type='button'
          className='profile-mobile-menu'
          aria-label={t('Toggle Sidebar')}
          aria-expanded={mobileOpen}
          onClick={() => setMobileOpen(!mobileOpen)}
        >
          <Menu size={20} />
        </button>
        <nav aria-label={t('Main navigation')}>
          <Link to='/'>{t('Home')}</Link>
          <Link to='/dashboard'>{t('Console')}</Link>
          <Link to='/pricing'>{t('Model Square')}</Link>
          <a href={docs} target='_blank' rel='noreferrer'>
            {t('Documentation', { defaultValue: '文档中心' })}
          </a>
        </nav>
        <div className='profile-header-tools'>
          <Search
            className='profile-header-search'
            placeholder={t('Search models, docs, APIs...', {
              defaultValue: '搜索模型、文档、接口...',
            })}
          />
          <a
            href={docs}
            target='_blank'
            rel='noreferrer'
            className='profile-help'
            title={t('Help')}
            aria-label={t('Help')}
          >
            <DesignIcon name='help' />
          </a>
          <div className='profile-notifications'>
            <NotificationPopover
              open={notifications.popoverOpen}
              onOpenChange={notifications.setPopoverOpen}
              unreadCount={notifications.unreadCount}
              activeTab={notifications.activeTab}
              onTabChange={notifications.setActiveTab}
              notice={notifications.notice}
              announcements={notifications.announcements}
              loading={notifications.loading}
            />
          </div>
          {!isProfilePage && <LanguageSwitcher />}
          <ProfileDropdown />
        </div>
      </header>
      {mobileOpen && (
        <button
          type='button'
          className='profile-nav-backdrop'
          aria-label={t('Close')}
          onClick={() => setMobileOpen(false)}
        />
      )}
      <aside
        ref={sidebarRef}
        role={mobileOpen ? 'dialog' : undefined}
        aria-modal={mobileOpen || undefined}
        aria-label={t('Sidebar')}
        className={`profile-design-sidebar ${mobileOpen ? 'is-open' : ''}`}
      >
        <Link className='profile-brand' to='/' aria-label={systemName}>
          {systemName.toUpperCase() === 'LUJIYUN' ? (
            <img
              className='profile-brand-wordmark'
              src={profileWordmark}
              alt={systemName}
            />
          ) : (
            <>
              <SystemLogo src={logo} width={32} height={32} />
              <span>{systemName}</span>
            </>
          )}
        </Link>
        <button
          type='button'
          className='profile-mobile-close'
          aria-label={t('Close')}
          onClick={() => setMobileOpen(false)}
        >
          <X size={18} />
        </button>
        <nav aria-label={t('Sidebar')}>
          <ProfileNavItem
            to='/dashboard/overview'
            icon='overview'
            label={t('Overview')}
          />
          <p className='profile-nav-label'>
            {t('Workspace', { defaultValue: '工作台' })}
          </p>
          <ProfileNavItem
            to='/dashboard/models'
            icon='dashboard'
            label={t('Dashboard')}
          />
          <ProfileNavItem to='/keys' icon='key' label={t('API Keys')} />
          <ProfileNavItem
            to='/usage-logs/common'
            icon='file'
            label={t('Usage Logs')}
          />
          <ProfileNavItem
            to='/usage-logs/task'
            icon='tasks'
            label={t('Task Logs')}
          />
          <p className='profile-nav-label profile-account-label'>
            {t('Account', { defaultValue: '账户' })}
          </p>
          <ProfileNavItem
            to='/wallet'
            icon='wallet'
            label={t('Billing Center')}
          />
          <ProfileNavItem to='/profile' icon='user' label={t('Profile')} />
          {isAdmin && (
            <>
              <p className='profile-nav-label'>{t('Admin')}</p>
              <ProfileNavItem
                to='/channels'
                icon='channel'
                label={t('Channels')}
              />
              <ProfileNavItem to='/models' icon='models' label={t('Models')} />
              <ProfileNavItem to='/users' icon='users' label={t('Users')} />
              <ProfileNavItem
                to='/redemption-codes'
                icon='redemption'
                label={t('Redemption Codes')}
              />
              <ProfileNavItem
                to='/subscriptions'
                icon='subscriptions'
                label={t('Subscriptions')}
              />
              {isSuperAdmin && (
                <>
                  <ProfileNavItem
                    to='/system-info'
                    icon='system-info'
                    label={t('System Info')}
                  />
                  <ProfileNavItem
                    to='/system-settings'
                    icon='system-settings'
                    label={t('System Settings')}
                  />
                  {isSystemSettingsPage && <ProfileSystemSettingsNavigation />}
                </>
              )}
            </>
          )}
          <p className='profile-nav-label profile-support-label'>
            {t('Support', { defaultValue: '支持' })}
          </p>
          <a
            className='profile-nav-item'
            href={docs}
            target='_blank'
            rel='noreferrer'
          >
            <DesignIcon name='docs' className='profile-nav-icon' />
            <span>{t('Documentation', { defaultValue: '文档中心' })}</span>
          </a>
          {/* Temporarily hidden: the contact support menu item was not part of the original menu. */}
          {/* <a
            className='profile-nav-item'
            href={docs}
            target='_blank'
            rel='noreferrer'
          >
            <DesignIcon name='help' className='profile-nav-icon' />
            <span>{t('Contact Support', { defaultValue: '联系支持' })}</span>
          </a> */}
        </nav>
        <div className='profile-sidebar-footer'>
          <Link to='/wallet' className='profile-invite'>
            <strong>{t('Invite Friends', { defaultValue: '邀请好友' })}</strong>
            <span>
              {t('Share AI capabilities', { defaultValue: '共享 AI 能力' })}
              <br />
              {t('Earn rewards', { defaultValue: '获取奖励' })}
            </span>
            <b>
              {t('View Details', { defaultValue: '查看详情' })}
              <DesignIcon name='arrow-right' />
            </b>
            <img
              className='profile-gift'
              src='/lanhu/profile-gift.png'
              alt=''
              width={175}
              height={175}
              aria-hidden='true'
            />
          </Link>
          <span className='profile-footer-brand' aria-hidden='true' />
          <p>© 2026 陆极云AI</p>
          <small>v2.1.0</small>
        </div>
      </aside>
    </>
  )
}
