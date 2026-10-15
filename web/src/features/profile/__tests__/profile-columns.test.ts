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

const source = readFileSync(new URL('../index.tsx', import.meta.url), 'utf8')
const css = readFileSync(
  new URL('../profile-design.css', import.meta.url),
  'utf8'
)

test('profile header shares the content area background', () => {
  const backgrounds: Record<string, string> = {}
  parse(css).walkRules((rule) => {
    if (
      rule.selector !== '.profile-design-header' &&
      rule.selector !== '.profile-design'
    ) {
      return
    }
    rule.walkDecls('background', (decl) => {
      backgrounds[rule.selector] = decl.value
    })
  })
  assert.equal(backgrounds['.profile-design'], '#f4f7fc')
  assert.equal(
    backgrounds['.profile-design-header'],
    backgrounds['.profile-design']
  )
})

test('profile page title stays at 20px across breakpoints', () => {
  const sizes: string[] = []

  parse(css).walkRules('.profile-page-title h1', (rule) => {
    rule.walkDecls('font-size', (decl) => {
      sizes.push(decl.value)
    })
  })

  assert.deepEqual(sizes, ['20px', '20px'])
})

test('profile navigation content keeps 15px left padding at every breakpoint', () => {
  const paddings: string[] = []

  parse(css).walkRules((rule) => {
    if (!rule.selector.includes("[data-slot='sidebar-inset']")) return
    if (!rule.selector.includes('.profile-design-layout')) return

    rule.walkDecls('padding', (decl) => {
      paddings.push(decl.value)
    })
  })

  assert.deepEqual(paddings, [
    '15px 25px 25px 15px',
    '15px 32px 14px 15px',
    '24px 24px 24px 15px',
    '18px 12px 18px 15px',
  ])
})

test('profile reserves the side column only when it contains a card', () => {
  assert.match(
    source,
    /const hasSideContent = checkinEnabled \|\| canConfigureSidebar/
  )
  assert.match(
    source,
    /profile-columns\$\{hasSideContent \? '' : ' profile-columns-full'\}/
  )
  assert.match(
    source,
    /\{hasSideContent && \(\s*<div className='profile-side'>/
  )

  const sheet = parse(css)
  const columns: Record<string, string> = {}
  for (const rule of sheet.nodes) {
    if (
      rule.type === 'rule' &&
      (rule.selector === '.profile-columns' ||
        rule.selector === '.profile-columns.profile-columns-full')
    ) {
      rule.walkDecls('grid-template-columns', (decl) => {
        columns[rule.selector] = decl.value
      })
    }
  }
  assert.equal(columns['.profile-columns'], 'minmax(0, 1fr) 352px')
  assert.equal(
    columns['.profile-columns.profile-columns-full'],
    'minmax(0, 1fr)'
  )
  assert.match(
    css,
    /\.profile-columns\.profile-columns-full\s*\{[^}]*max-width: none/s
  )
})
