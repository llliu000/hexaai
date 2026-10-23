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
  new URL('../section-page-layout.tsx', import.meta.url),
  'utf8'
)

test('system page titles use a consistent 20px font size', () => {
  assert.match(
    source,
    /<h2 className='truncate !text-\[20px\] font-bold tracking-tight'>/
  )
  assert.doesNotMatch(
    source,
    /truncate (?:text-base|text-xl|text-\[20px\]) font-bold tracking-tight/
  )
})

test('system page content starts 15px from the left edge', () => {
  assert.match(
    source,
    /const sectionPageInlinePadding =\s*'pr-3 pl-\[15px\] sm:pr-4 sm:pl-\[15px\]'/
  )
  assert.equal(source.match(/sectionPageInlinePadding/g)?.length, 4)
})
