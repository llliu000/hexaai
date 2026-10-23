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

import { parse } from 'postcss'

const fields = readFileSync(
  new URL('../hero-slide-fields.tsx', import.meta.url),
  'utf8'
)
const styles = readFileSync(
  new URL('../../../../styles/home-carousel-editor.css', import.meta.url),
  'utf8'
)

function declarations(selector: string): Record<string, string> {
  const result: Record<string, string> = {}
  parse(styles).walkRules(selector, (rule) => {
    if (rule.parent?.type !== 'root') return
    rule.walkDecls((declaration) => {
      result[declaration.prop] = declaration.value
    })
  })
  return result
}

test('global carousel controls stay in one row without wrapping', () => {
  const global = declarations('.hero-config-global')
  assert.equal(global.display, 'grid')
  assert.match(
    global['grid-template-columns'],
    /max-content max-content max-content/
  )
  assert.equal(global['overflow-x'], 'auto')
  assert.equal(declarations('.hero-config-global > *')['white-space'], 'nowrap')
})

test('primary and secondary actions use aligned grid order', () => {
  assert.match(
    fields,
    /primaryLabel[\s\S]*secondaryLabel[\s\S]*hero-config-action-row[\s\S]*primaryAction\.url[\s\S]*primaryAction\.newTab/
  )
  assert.match(
    fields,
    /hero-config-action-row[\s\S]*secondaryAction\.url[\s\S]*secondaryEnabled[\s\S]*secondaryAction\.newTab/
  )
  assert.equal(
    declarations('.hero-config-fields')['grid-template-columns'],
    'repeat(2, minmax(0, 1fr))'
  )
  assert.equal(
    declarations('.hero-config-action-row')['grid-template-columns'],
    'repeat(2, minmax(0, 1fr))'
  )
})
