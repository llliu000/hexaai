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
import { describe, test } from 'node:test'

const authenticatedLayout = readFileSync(
  new URL(
    '../../../components/layout/components/authenticated-layout.tsx',
    import.meta.url
  ),
  'utf8'
)
const profileNavigation = readFileSync(
  new URL('../components/profile-navigation.tsx', import.meta.url),
  'utf8'
)
const apiKeysStyles = readFileSync(
  new URL('../../../styles/api-keys-lanhu.css', import.meta.url),
  'utf8'
)
const overviewStyles = readFileSync(
  new URL('../../../styles/overview-lanhu.css', import.meta.url),
  'utf8'
)
const pageStyleUrls = [
  '../profile-design.css',
  '../../../styles/api-keys-lanhu.css',
  '../../../styles/flow-lanhu.css',
  '../../../styles/models-lanhu.css',
  '../../../styles/overview-lanhu.css',
  '../../../styles/task-logs-lanhu.css',
  '../../../styles/usage-logs-lanhu.css',
  '../../../styles/wallet-lanhu.css',
] as const

describe('profile sidebar footer', () => {
  test('common usage logs use the navigation that owns the footer', () => {
    assert.match(authenticatedLayout, /'\/usage-logs\/common'/)
    assert.match(profileNavigation, /className='profile-sidebar-footer'/)
  })

  test('menu-specific styles do not hide footer content', () => {
    for (const styleUrl of pageStyleUrls) {
      const styles = readFileSync(new URL(styleUrl, import.meta.url), 'utf8')
      assert.doesNotMatch(
        styles,
        /profile-sidebar-footer[^{}]*\{[^}]*display:\s*none/
      )
    }
  })

  test('API keys use the shared footer geometry', () => {
    assert.doesNotMatch(
      apiKeysStyles,
      /body:has\(\.lujiyun-keys-page\) \.profile-(?:sidebar-footer|invite|footer-brand)/
    )
  })

  test('overview uses the shared footer geometry', () => {
    assert.doesNotMatch(
      overviewStyles,
      /body:has\(\.lujiyun-overview-page\) \.profile-(?:sidebar-footer|invite|footer-brand)/
    )
  })
})
