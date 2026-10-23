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

const usageLogsPage = readFileSync(
  new URL('../../index.tsx', import.meta.url),
  'utf8'
)
const profileLayoutStyles = readFileSync(
  new URL('../../../profile/profile-design.css', import.meta.url),
  'utf8'
)
const sectionPageLayout = readFileSync(
  new URL(
    '../../../../components/layout/components/section-page-layout.tsx',
    import.meta.url
  ),
  'utf8'
)
const taskLogsFilterBar = readFileSync(
  new URL('../task-logs-filter-bar.tsx', import.meta.url),
  'utf8'
)
const taskLogsStyles = readFileSync(
  new URL('../../../../styles/task-logs-lanhu.css', import.meta.url),
  'utf8'
)
const usageLogsStyles = readFileSync(
  new URL('../../../../styles/usage-logs-lanhu.css', import.meta.url),
  'utf8'
)
const commonLogsColumns = readFileSync(
  new URL('../columns/common-logs-columns.tsx', import.meta.url),
  'utf8'
)
const detailsDialog = readFileSync(
  new URL('../dialogs/details-dialog.tsx', import.meta.url),
  'utf8'
)

describe('usage logs layout', () => {
  test('common and task logs share a full-width content container', () => {
    assert.match(
      usageLogsPage,
      /lujiyun-usage-\$\{activeCategory\}[^`]*\bw-full\b[^`]*\bmax-w-none\b/
    )
  })

  test('profile navigation pages override centered theme width limits', () => {
    assert.match(
      profileLayoutStyles,
      /\.profile-design-layout\s*>\s*\[data-slot='sidebar-inset'\]\s*>\s*\*\s*\{[^}]*width:\s*100%;[^}]*max-width:\s*none\s*!important;[^}]*margin-inline:\s*0\s*!important;/s
    )
  })

  test('profile navigation pages use 15px left content padding', () => {
    assert.match(
      profileLayoutStyles,
      /\.profile-design-layout\s*>\s*\[data-slot='sidebar-inset'\]\s*\{[^}]*padding:\s*15px 25px 25px 15px;/s
    )
  })

  test('common and task logs remove the header top and inline padding', () => {
    assert.match(
      sectionPageLayout,
      /className=\{cn\([^)]*props\.headerClassName[^)]*\)\}/s
    )
    assert.match(usageLogsPage, /headerClassName='px-0 pt-0 sm:px-0 sm:pt-0'/)
  })

  test('common and task logs remove the content inline padding', () => {
    assert.match(
      sectionPageLayout,
      /className=\{cn\([\s\S]*props\.contentClassName[\s\S]*\)\}/
    )
    assert.match(
      usageLogsPage,
      /contentClassName='usage-logs-content px-0 pb-0 sm:px-0 sm:pb-0'/
    )
    assert.match(
      taskLogsStyles,
      /> div:nth-child\(2\)\s*\{[^}]*padding-inline:\s*0;/s
    )
  })

  test('task logs use the same two-row filter card geometry as usage logs', () => {
    assert.match(
      taskLogsStyles,
      /\.lujiyun-usage-task \.task-logs-filter-card\s*\{[^}]*min-height:\s*140px;[^}]*padding:\s*21px 20px;[^}]*border-radius:\s*18px;/s
    )
    assert.match(
      taskLogsStyles,
      /\.task-logs-filter-card > div:last-child\s*\{[^}]*margin-top:\s*18px;/s
    )
    assert.match(
      taskLogsStyles,
      /div:first-child:has\(> :nth-child\(3\)\)\s*\{[^}]*grid-template-columns:\s*minmax\(0, 1\.72fr\) repeat\(2, minmax\(0, 1fr\)\);/s
    )
    assert.doesNotMatch(
      taskLogsStyles,
      /\.task-logs-filter-card\s*\{[^}]*display:\s*flex;/s
    )
  })

  test('task ID uses the shared plain filter input', () => {
    assert.doesNotMatch(taskLogsFilterBar, /task-id-search-icon|<Search\b/)
    assert.match(
      taskLogsFilterBar,
      /<LogsFilterInput[\s\S]*aria-label=\{t\('Task ID'\)\}/
    )
  })

  test('common and task logs use the API keys footer surface and height', () => {
    assert.match(sectionPageLayout, /data-slot='page-footer'/)
    assert.match(
      usageLogsStyles,
      /body:has\(:is\(\.lujiyun-usage-common, \.lujiyun-usage-task\)\)\s*\[data-slot='page-footer'\]\s*\{[^}]*min-height:\s*87px;[^}]*padding:\s*0 24px;[^}]*border:\s*0;[^}]*border-radius:\s*0 0 20px 20px;[^}]*background:\s*var\(--design-inset, #f8faff\);/s
    )
    assert.doesNotMatch(taskLogsStyles, /main\s*> div:last-child/)
  })

  test('non-admin model cells do not expose request and actual model popover', () => {
    assert.match(
      commonLogsColumns,
      /actualModel=\{isAdmin \? modelInfo\.actualModel : undefined\}/
    )
  })

  test('non-admin log details do not expose the actual model', () => {
    assert.match(
      detailsDialog,
      /\{props\.isAdmin && \(\s*<DetailRow\s*label=\{t\('Actual Model'\)\}/
    )
  })

  test('log pagination uses the API keys rectangular buttons and blue current page', () => {
    assert.match(
      usageLogsStyles,
      /\[data-slot='page-footer'\]\s*\[data-slot='button'\]\s*\{[^}]*width:\s*42px;[^}]*min-width:\s*42px;[^}]*height:\s*35px;[^}]*border-radius:\s*10px;[^}]*background:\s*var\(--design-surface, #fbfcff\);/s
    )
    assert.match(
      usageLogsStyles,
      /\[data-slot='button'\]\[class~='bg-primary'\]\s*\{[^}]*background:\s*#246cf4;[^}]*color:\s*#fff;/s
    )
  })

  test('log pagination uses the API keys page-size selector dimensions', () => {
    assert.match(
      usageLogsStyles,
      /\[data-slot='page-footer'\]\s*\[data-slot='select-trigger'\]\s*\{[^}]*width:\s*73px;[^}]*height:\s*37px;[^}]*border-color:\s*var\(--design-border, #dce4f0\);[^}]*border-radius:\s*9px;/s
    )
  })

  test('desktop pagination keeps totals on the left and page navigation on the right', () => {
    assert.match(
      usageLogsStyles,
      /\[data-slot='page-footer'\]\s*> div\s*> div\s*\{[^}]*width:\s*100%;[^}]*gap:\s*24px;[^}]*justify-content:\s*flex-start;/s
    )
    assert.match(
      usageLogsStyles,
      /\[data-slot='page-footer'\]\s*> div\s*> div\s*> div:nth-child\(3\)\s*\{[^}]*margin-right:\s*15px;[^}]*margin-left:\s*auto;[^}]*gap:\s*12px;/s
    )
  })

  test('narrow screens wrap page navigation so controls remain reachable', () => {
    assert.match(
      usageLogsStyles,
      /@media \(max-width: 767px\)[\s\S]*\[data-slot='page-footer'\]\s*> div\s*> div\s*\{[^}]*flex-wrap:\s*wrap;/
    )
    assert.match(
      usageLogsStyles,
      /@media \(max-width: 767px\)[\s\S]*\[data-slot='page-footer'\]\s*> div\s*> div\s*> div:nth-child\(3\)\s*\{[^}]*width:\s*100%;[^}]*justify-content:\s*flex-end;/
    )
  })
})
