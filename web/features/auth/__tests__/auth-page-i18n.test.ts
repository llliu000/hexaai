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

const locales = ['en', 'zh', 'zh-TW', 'ja', 'fr', 'ru', 'vi'] as const
const pageKeys = [
  'auth.page.welcome',
  'auth.page.subtitle',
  'auth.page.hint',
  'auth.page.navigation',
  'auth.page.heroEyebrow',
  'auth.page.heroTitle',
  'auth.page.prompt',
  'auth.page.promptExplore',
  'auth.page.promptCreate',
  'auth.page.promptAgent',
  'auth.page.footnote',
] as const

for (const locale of locales) {
  const { translation } = JSON.parse(
    readFileSync(
      new URL(`../../../i18n/locales/${locale}.json`, import.meta.url),
      'utf8'
    )
  ) as { translation: Record<string, string> }

  describe(`authentication page translations: ${locale}`, () => {
    test('selecting this language provides all login and registration page copy', () => {
      for (const key of pageKeys) {
        assert.equal(typeof translation[key], 'string', `${locale}: ${key}`)
        assert.ok(translation[key].trim(), `${locale}: ${key} is empty`)
        assert.notEqual(
          translation[key],
          key,
          `${locale}: ${key} is untranslated`
        )
      }
    })

    test('legal footer copy is translated for both authentication modes', () => {
      for (const key of [
        'By clicking sign in, you agree to our',
        'By creating an account, you agree to our',
      ]) {
        assert.equal(typeof translation[key], 'string')
        assert.ok(translation[key].trim())
        if (locale !== 'en') assert.notEqual(translation[key], key)
      }
    })

    test('configuring a brand keeps the welcome and creative prompt interpolatable', () => {
      for (const key of [
        'auth.page.welcome',
        'auth.page.prompt',
        'auth.page.promptExplore',
        'auth.page.promptCreate',
        'auth.page.promptAgent',
      ]) {
        assert.deepEqual(
          translation[key].match(/{{\s*brand\s*}}/g),
          ['{{brand}}'],
          `${locale}: ${key} must display the configured brand once`
        )
      }
    })
  })
}
