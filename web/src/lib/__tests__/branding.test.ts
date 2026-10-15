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

import { DEFAULT_LOGO, DEFAULT_SYSTEM_NAME } from '../constants.ts'

const documentTemplate = readFileSync(
  new URL('../../../index.html', import.meta.url),
  'utf8'
)
const buildConfig = readFileSync(
  new URL('../../../rsbuild.config.ts', import.meta.url),
  'utf8'
)

describe('application branding', () => {
  test('uses the system name and icon before runtime configuration loads', () => {
    assert.equal(DEFAULT_SYSTEM_NAME, '陆极云AI')
    assert.equal(DEFAULT_LOGO, '/logo.svg')
    assert.match(documentTemplate, /<title>陆极云AI<\/title>/)
    assert.match(documentTemplate, /<meta name="title" content="陆极云AI" \/>/)
    assert.match(buildConfig, /favicon: '\.\/public\/logo\.svg'/)
    assert.doesNotMatch(documentTemplate, /New API/)
  })
})
