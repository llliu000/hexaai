import { useLocation } from '@tanstack/react-router'

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
import { AnimatedOutlet } from '@/components/page-transition'
import { SkipToMain } from '@/components/skip-to-main'
import { SidebarInset, SidebarProvider } from '@/components/ui/sidebar'
import { LayoutProvider } from '@/context/layout-provider'
import { SearchProvider } from '@/context/search-provider'
import { ProfileNavigation } from '@/features/profile/components/profile-navigation'

import '@/features/profile/profile-design.css'
import { getCookie } from '@/lib/cookies'
import { cn } from '@/lib/utils'

import { AppHeader } from './app-header'
import { AppSidebar } from './app-sidebar'

type AuthenticatedLayoutProps = {
  children?: React.ReactNode
}

const PROFILE_NAVIGATION_EXACT_PATHS = new Set([
  '/profile',
  '/keys',
  '/usage-logs/common',
  '/usage-logs/task',
  '/wallet',
  '/dashboard/overview',
  '/dashboard/models',
  '/dashboard/flow',
  '/dashboard/users',
])

const PROFILE_NAVIGATION_ADMIN_PREFIXES = [
  '/channels',
  '/models',
  '/users',
  '/redemption-codes',
  '/subscriptions',
  '/system-info',
  '/system-settings',
] as const

function matchesRoutePrefix(pathname: string, prefix: string): boolean {
  return pathname === prefix || pathname.startsWith(`${prefix}/`)
}

export function AuthenticatedLayout(props: AuthenticatedLayoutProps) {
  const pathname = useLocation({
    select: (location) => location.pathname.replace(/\/$/, ''),
  })
  const isAdminSection = PROFILE_NAVIGATION_ADMIN_PREFIXES.some((prefix) =>
    matchesRoutePrefix(pathname, prefix)
  )
  const usesProfileNavigation =
    PROFILE_NAVIGATION_EXACT_PATHS.has(pathname) || isAdminSection
  const usesSectionPagePadding =
    isAdminSection || pathname === '/dashboard/users'
  const defaultOpen = getCookie('sidebar_state') !== 'false'

  return (
    <LayoutProvider>
      <SearchProvider>
        <SidebarProvider defaultOpen={defaultOpen} className='flex-col'>
          <SkipToMain />
          {usesProfileNavigation ? <ProfileNavigation /> : <AppHeader />}
          <div
            className={cn(
              usesProfileNavigation
                ? 'profile-design-layout'
                : 'flex min-h-0 w-full flex-1',
              usesSectionPagePadding && 'profile-section-page-layout'
            )}
          >
            {!usesProfileNavigation && <AppSidebar />}
            <SidebarInset
              className={cn(
                '@container/content',
                usesProfileNavigation && 'min-w-0 max-w-none',
                'h-[calc(100svh-var(--app-header-height,0px))]',
                'min-h-0 overflow-hidden',
                'peer-data-[variant=inset]:h-[calc(100svh-var(--app-header-height,0px)-(var(--spacing)*4))]'
              )}
            >
              {props.children ?? <AnimatedOutlet />}
            </SidebarInset>
          </div>
        </SidebarProvider>
      </SearchProvider>
    </LayoutProvider>
  )
}
