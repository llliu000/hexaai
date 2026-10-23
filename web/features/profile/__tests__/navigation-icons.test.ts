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

const navigationSource = readFileSync(
  new URL('../components/profile-navigation.tsx', import.meta.url),
  'utf8'
)
const profileStyles = readFileSync(
  new URL('../profile-design.css', import.meta.url),
  'utf8'
)
const pageStylePaths = [
  '../../../styles/api-keys-lanhu.css',
  '../../../styles/flow-lanhu.css',
  '../../../styles/models-lanhu.css',
  '../../../styles/overview-lanhu.css',
  '../../../styles/task-logs-lanhu.css',
  '../../../styles/wallet-lanhu.css',
] as const

describe('profile navigation icons', () => {
  test('preserves configured branding when the supplied LuJiyun wordmark does not apply', () => {
    assert.match(
      navigationSource,
      /<SystemLogo src=\{logo\}[^>]*\/>[\s\S]*?<span>\{systemName\}<\/span>/
    )
    assert.match(
      profileStyles,
      /\.profile-brand\s*\{[^}]*background: none !important;/s
    )
    assert.match(
      profileStyles,
      /\.profile-brand > \.system-logo\s*\{[^}]*width: 32px;[^}]*height: 32px;/s
    )
    assert.doesNotMatch(navigationSource, /home-header-logo\.png/)
  })

  test('resolves every supplied profile SVG locally without external resources', () => {
    const iconModule = new URL('../components/design-icon.tsx', import.meta.url)
    const iconSource = readFileSync(iconModule, 'utf8')
    const assetImports = [
      ...iconSource.matchAll(
        /from ['"](\.\.\/assets\/icons\/[^'"]+\.svg)['"]/g
      ),
    ]
    assert.ok(assetImports.length > 0, 'The profile has no local design assets')
    for (const [, relativePath] of assetImports) {
      const svg = readFileSync(new URL(relativePath, iconModule), 'utf8')
      assert.match(
        svg,
        /<svg\b[^>]*xmlns=["']http:\/\/www\.w3\.org\/2000\/svg["']/
      )
      assert.match(svg, /viewBox=["'][^"']+["']/)
      assert.match(svg, /<(?:path|circle|rect|line|polyline|polygon|ellipse)\b/)
      assert.doesNotMatch(svg, /(?:href|src)=["']https?:\/\//)
      assert.match(svg, /<\/svg>\s*$/)
    }
  })

  test('does not layer page artwork behind sidebar navigation icons', () => {
    for (const path of pageStylePaths) {
      const styles = readFileSync(new URL(path, import.meta.url), 'utf8')
      assert.doesNotMatch(
        styles,
        /\.profile-nav-item[^,{]*\.icon-[^{]*\{[^}]*background-(?:image|position|size)\s*:/s
      )
      assert.doesNotMatch(
        styles,
        /\.profile-nav-item[^,{]*\.icon-[^{]*\{[^}]*\bmask\s*:/s
      )
      assert.doesNotMatch(
        styles,
        /\.profile-design-icon(?!:not\(\.profile-nav-icon\))\s*,/
      )
      assert.doesNotMatch(
        styles,
        /\.profile-notifications button::before\s*\{[^}]*background-(?:image|position|size)\s*:/s
      )
    }
  })

  test('keeps the notification bell visible without page artwork', () => {
    assert.match(
      profileStyles,
      /\.profile-notifications button > svg\s*\{[^}]*opacity:\s*1;/s
    )
    assert.match(
      profileStyles,
      /\.profile-notifications button::before\s*\{[^}]*display:\s*none;/s
    )
  })

  test('keeps the profile icon neutral until its navigation item is current', () => {
    const userIcon = readFileSync(
      new URL('../assets/icons/user.svg', import.meta.url),
      'utf8'
    )
    assert.match(userIcon, /stroke="#344054"/)
    assert.doesNotMatch(userIcon, /stroke="#1C64F2"/)
    assert.match(
      profileStyles,
      /\.profile-nav-item\.is-current \.profile-nav-icon\s*\{[^}]*filter:/s
    )
    assert.doesNotMatch(profileStyles, /\.icon-user\s*\{[^}]*filter:\s*none/s)
  })
})
