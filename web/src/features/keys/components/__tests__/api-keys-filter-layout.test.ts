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

const apiKeysTable = readFileSync(
  new URL('../api-keys-table.tsx', import.meta.url),
  'utf8'
)
const apiKeysStyles = readFileSync(
  new URL('../../../../styles/api-keys-lanhu.css', import.meta.url),
  'utf8'
)

describe('API key status filter', () => {
  test('uses the shared single-select structure from usage log filters', () => {
    assert.match(apiKeysTable, /<Select\s+items=\{statusItems\}/)
    assert.match(apiKeysTable, /value=\{selectedStatus\}/)
    assert.match(apiKeysTable, /className='api-key-status-filter'/)
    assert.match(apiKeysTable, /<SelectContent alignItemWithTrigger=\{false\}>/)
    assert.match(
      apiKeysTable,
      /<SelectItem value=\{API_KEY_STATUS_ALL_VALUE\}>/
    )
    assert.doesNotMatch(apiKeysTable, /columnId: 'status',[\s\S]*singleSelect/)
  })

  test('clears the status column filter when all statuses is selected', () => {
    assert.match(
      apiKeysTable,
      /value === API_KEY_STATUS_ALL_VALUE \? undefined : \[value\]/
    )
  })

  test('shows a status placeholder before a status filter is selected', () => {
    assert.match(
      apiKeysTable,
      /const selectedStatus = statusFilter\[0\][\s\S]*<SelectValue placeholder=\{t\('Select status'\)\}/
    )
  })

  test('shows the translated label after a status is selected', () => {
    assert.match(
      apiKeysTable,
      /label: t\(option\.label\)[\s\S]*const selectedStatusLabel = statusItems\.find/
    )
    assert.match(
      apiKeysTable,
      /<SelectValue placeholder=\{t\('Select status'\)\}>\s*\{selectedStatusLabel\}\s*<\/SelectValue>/
    )
  })

  test('matches the API key search control dimensions', () => {
    assert.match(
      apiKeysStyles,
      /\.lujiyun-keys-page \.api-key-status-filter\s*\{[^}]*height:\s*42px;[^}]*border-radius:\s*12px;[^}]*background:\s*var\(--design-surface, #fff\);/s
    )
  })
})
