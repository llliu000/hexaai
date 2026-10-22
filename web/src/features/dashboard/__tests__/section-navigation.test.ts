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

const sharedStyles = readFileSync(
  new URL('../../../styles/index.css', import.meta.url),
  'utf8'
)
const flowStyles = readFileSync(
  new URL('../../../styles/flow-lanhu.css', import.meta.url),
  'utf8'
)
const modelStyles = readFileSync(
  new URL('../../../styles/models-lanhu.css', import.meta.url),
  'utf8'
)
const dashboard = readFileSync(new URL('../index.tsx', import.meta.url), 'utf8')

test('dashboard heading follows the active section tab', () => {
  assert.match(
    dashboard,
    /<SectionPageLayout.Title>[\s\S]*\{t\(activeSectionMeta\.titleKey\)\}/
  )
  assert.match(dashboard, /className='dashboard-section-tabs w-full sm:w-auto'/)
  assert.match(dashboard, /value=\{activeSection\}/)
  assert.match(dashboard, /onValueChange=\{handleSectionChange\}/)
  assert.match(
    modelStyles,
    /\.dashboard-section-tabs[\s\S]*?\[data-slot='tabs-list'\][\s\S]*?border: 1px solid var\(--design-border, #e2e8f0\);[\s\S]*?background: transparent;/
  )
  assert.match(
    modelStyles,
    /\.dashboard-section-tabs[\s\S]*?\[data-slot='tabs-trigger'\]\[data-active\][\s\S]*?background: var\(--design-blue-surface, #e6efff\);[\s\S]*?color: var\(--design-link, #216aff\);/
  )
  assert.match(
    modelStyles,
    /\.dashboard-section-tabs[\s\S]*?\[data-slot='tabs-trigger'\][\s\S]*?border-radius: 18px;/
  )
})

test('user analytics uses the shared dashboard section tab styles', () => {
  assert.match(
    sharedStyles,
    /\.lujiyun-dashboard-users > div:first-child \[data-slot='tabs-list'\]/
  )
  assert.match(
    sharedStyles,
    /\.lujiyun-dashboard-users[\s\S]*?> div:first-child[\s\S]*?\[data-slot='tabs-trigger'\]\[data-state='active'\]/
  )
})

test('user analytics matches the desktop flow tab selection styles', () => {
  assert.match(
    flowStyles,
    /\.lujiyun-dashboard-users\.lujiyun-dashboard-users[\s\S]*?> div:first-child[\s\S]*?\[data-slot='tabs-list'\]/
  )
  assert.match(
    flowStyles,
    /\.lujiyun-dashboard-users[\s\S]*?> div:first-child[\s\S]*?\[data-slot='tabs-trigger'\]\[aria-selected='true'\]/
  )
})
