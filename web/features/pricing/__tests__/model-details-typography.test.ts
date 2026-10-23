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

const detailsSource = readFileSync(
  new URL('../components/model-details.tsx', import.meta.url),
  'utf8'
)
const styles = readFileSync(
  new URL('../../../styles/pricing-lanhu.css', import.meta.url),
  'utf8'
)

describe('model details drawer typography', () => {
  test('scopes the model-card typography scale to the details drawer', () => {
    assert.match(detailsSource, /model-details-drawer sm:max-w-2xl/)
    assert.match(
      styles,
      /\.model-details-drawer\s*\{[^}]*font-family: var\(--font-body\);[^}]*font-size: 13px;[^}]*line-height: 20px;[^}]*letter-spacing: 0;/s
    )
  })

  test('keeps the model identity and section hierarchy compact', () => {
    assert.match(
      styles,
      /\.model-details-drawer h1\s*\{[^}]*font-family: inherit;[^}]*font-size: 18px;[^}]*line-height: 24px;[^}]*font-weight: 650;[^}]*letter-spacing: 0;/s
    )
    assert.match(
      styles,
      /\.model-details-drawer h2,\s*\.model-details-drawer h3\s*\{[^}]*font-family: inherit;[^}]*font-size: 14px;[^}]*line-height: 20px;[^}]*font-weight: 600;[^}]*letter-spacing: 0;/s
    )
  })

  test('uses the same compact body and value sizes as the model cards', () => {
    assert.match(
      styles,
      /\.model-details-drawer \.text-lg\s*\{[^}]*font-size: 16px;[^}]*line-height: 24px;/s
    )
    assert.match(
      styles,
      /\.model-details-drawer \.text-base\s*\{[^}]*font-size: 15px;[^}]*line-height: 22px;/s
    )
    assert.match(
      styles,
      /\.model-details-drawer \.text-sm\s*\{[^}]*font-size: 13px;[^}]*line-height: 20px;/s
    )
    assert.match(
      styles,
      /\.model-details-drawer \[data-slot='tabs-trigger'\]\s*\{[^}]*font-size: 13px;[^}]*line-height: 18px;[^}]*font-weight: 600;/s
    )
  })

  test('removes promotional uppercase tracking from drawer labels', () => {
    assert.match(
      styles,
      /\.model-details-drawer \.tracking-wider\s*\{[^}]*letter-spacing: 0;/s
    )
    assert.match(
      styles,
      /\.model-details-drawer \.uppercase\s*\{[^}]*text-transform: none;/s
    )
  })
})
