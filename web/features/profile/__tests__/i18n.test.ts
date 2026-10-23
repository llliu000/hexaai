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
import fs from 'node:fs'
import { describe, test } from 'node:test'

import { createInstance } from 'i18next'

const zh = JSON.parse(
  fs.readFileSync(
    new URL('../../../i18n/locales/zh.json', import.meta.url),
    'utf8'
  )
) as { translation: Record<string, string> }

const profileTranslationKeys = [
  'Account',
  'Billing Center',
  'Contact Support',
  'Documentation',
  'Earn rewards',
  'Help',
  'Invite Friends',
  'Main navigation',
  'Manage your account information, security settings and personal preferences',
  'profile.design.deleteDescription',
  'profile.design.passwordDescription',
  'profile.design.remainingQuota',
  'profile.design.settingsDescription',
  'profile.design.tokenDescription',
  'Search models, docs, APIs...',
  'Share AI capabilities',
  'Sidebar Settings',
  'Simpler and safer sign-in',
  'Support',
  'Toggle Sidebar',
  'View Details',
  'Workspace',
] as const

describe('profile Chinese translations', () => {
  test('renders the redesigned navigation and controls without English fallback', async () => {
    const i18n = createInstance()
    await i18n.init({
      lng: 'zhCN',
      fallbackLng: 'en',
      resources: { zhCN: zh },
    })

    for (const key of profileTranslationKeys) {
      const translated = i18n.t(key)
      assert.notEqual(
        translated,
        key,
        `Missing Chinese translation for: ${key}`
      )
      assert.ok(translated.trim().length > 0)
    }
  })
})
