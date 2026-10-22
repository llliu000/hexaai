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
import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'
import { test } from 'node:test'

const source = readFileSync(
  new URL('../components/profile-navigation.tsx', import.meta.url),
  'utf8'
)
const authenticatedLayout = readFileSync(
  new URL(
    '../../../components/layout/components/authenticated-layout.tsx',
    import.meta.url
  ),
  'utf8'
)
const styles = readFileSync(
  new URL('../profile-design.css', import.meta.url),
  'utf8'
)

test('profile navigation shows admin routes to administrator accounts', () => {
  assert.match(source, /const userRole = useAuthStore\([\s\S]*?ROLE\.GUEST\)/)
  assert.match(source, /const isAdmin = userRole >= ROLE\.ADMIN/)
  assert.match(source, /\{isAdmin && \([\s\S]*?\{t\('Admin'\)\}/)

  for (const route of [
    '/channels',
    '/models',
    '/users',
    '/redemption-codes',
    '/subscriptions',
  ]) {
    assert.match(source, new RegExp(`to=['"]${route}['"]`))
  }
})

test('profile navigation restricts root-only routes to super administrators', () => {
  assert.match(source, /const isSuperAdmin = userRole >= ROLE\.SUPER_ADMIN/)
  assert.match(
    source,
    /\{isSuperAdmin && \([\s\S]*?to='\/system-info'[\s\S]*?to='\/system-settings'/
  )
})

test('profile admin entries keep the shared sidebar module visibility gate', () => {
  assert.match(
    source,
    /function ProfileNavItem[\s\S]*?useIsSidebarModuleVisible\(props\.to\)/
  )
})

test('admin pages keep the same profile header and sidebar shell', () => {
  for (const route of [
    '/channels',
    '/models',
    '/users',
    '/redemption-codes',
    '/subscriptions',
    '/system-info',
    '/system-settings',
  ]) {
    assert.match(authenticatedLayout, new RegExp(`['"]${route}['"]`))
  }

  assert.match(
    authenticatedLayout,
    /usesProfileNavigation \? <ProfileNavigation \/> : <AppHeader \/>/
  )
  assert.match(
    authenticatedLayout,
    /usesSectionPagePadding && 'profile-section-page-layout'/
  )
  assert.match(
    styles,
    /\.profile-design-layout\.profile-section-page-layout[\s\S]*?> \[data-slot='sidebar-inset'\]\s*\{[^}]*padding:\s*0;/
  )
})

test('nested admin routes keep their parent menu active', () => {
  assert.match(
    source,
    /props\.to === '\/models'[\s\S]*?props\.to === '\/system-settings'[\s\S]*?pathname\.startsWith\(`\$\{props\.to\}\/`\)/
  )
})

test('system settings keep their nested navigation inside the unified sidebar', () => {
  assert.match(source, /function ProfileSystemSettingsNavigation\(\)/)
  assert.match(source, /SYSTEM_SETTINGS_VIEW\.getNavGroups\(t\)/)
  assert.match(
    source,
    /\{isSystemSettingsPage &&(?: \()?\s*<ProfileSystemSettingsNavigation \/>/
  )
  assert.match(
    styles,
    /\.profile-system-settings-subnav a\.is-current\s*\{[^}]*color:\s*var\(--design-link, #1966ff\);[^}]*background:\s*var\(--design-blue-surface, #eaf1ff\);/s
  )
})

test('user analytics keeps the unified shell, padding, and dashboard menu active', () => {
  assert.match(authenticatedLayout, /'\/dashboard\/users'/)
  assert.match(
    authenticatedLayout,
    /const usesSectionPagePadding =\s*isAdminSection \|\| pathname === '\/dashboard\/users'/
  )
  assert.match(
    authenticatedLayout,
    /usesSectionPagePadding && 'profile-section-page-layout'/
  )
  assert.match(
    source,
    /props\.to === '\/dashboard\/models'[\s\S]*?pathname === '\/dashboard\/flow'[\s\S]*?pathname === '\/dashboard\/users'/
  )
})
