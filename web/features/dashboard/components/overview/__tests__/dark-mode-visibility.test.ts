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

const styles = readFileSync(
  new URL('../../../../../styles/overview-lanhu.css', import.meta.url),
  'utf8'
)

test('overview sparse cards define dedicated high-contrast dark visuals', () => {
  assert.match(
    styles,
    /html\.dark \.overview-route-face\.is-top \{[^}]*#[0-9a-f]{6}/s
  )
  assert.match(
    styles,
    /html\.dark \.overview-route-face\.is-left \{[^}]*#[0-9a-f]{6}/s
  )
  assert.match(
    styles,
    /html\.dark \.overview-route-face\.is-right \{[^}]*#[0-9a-f]{6}/s
  )
  assert.match(
    styles,
    /html\.dark \.overview-service-card li strong \{[^}]*color:/s
  )
  assert.match(
    styles,
    /html\.dark \.overview-announcements-card li \{[^}]*color:/s
  )
  assert.match(
    styles,
    /html\.dark \.overview-announcements-card time \{[^}]*color:/s
  )
})
