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
import { test } from 'node:test'

import { getLastVisibleSectionTarget } from '../lib/section-scroll-reveal.ts'

const viewport = { height: 900, headerHeight: 80 }
const sections = [
  { top: 0, height: 600, range: { start: 0, end: 0 } },
  { top: 600, height: 560, range: { start: 390, end: 390 } },
  { top: 1160, height: 680, range: { start: 1010, end: 1010 } },
  { top: 1840, height: 560, range: { start: 1630, end: 1630 } },
  { top: 2400, height: 740, range: { start: 2280, end: 2280 } },
  { top: 3140, height: 400, range: { start: 2850, end: 2850 } },
]

test('downward scrolling can finish at the last meaningfully revealed module', () => {
  assert.equal(
    getLastVisibleSectionTarget(700, sections, viewport, {
      origin: 0,
      direction: 1,
    }),
    1010
  )
  assert.equal(
    getLastVisibleSectionTarget(2040, sections, viewport, {
      origin: 0,
      direction: 1,
    }),
    2280
  )
})

test('upward scrolling aligns the nearest module still ahead from above', () => {
  assert.equal(
    getLastVisibleSectionTarget(700, sections, viewport, {
      origin: 1600,
      direction: -1,
    }),
    390
  )
  assert.equal(
    getLastVisibleSectionTarget(1200, sections, viewport, {
      origin: 1630,
      direction: -1,
    }),
    1010
  )
})

test('upward page turns still work when the previous module already peeked', () => {
  assert.equal(
    getLastVisibleSectionTarget(850, sections, viewport, {
      origin: 1010,
      direction: -1,
    }),
    390
  )
})

test('a barely visible strip of the following module does not cause an extra jump', () => {
  assert.equal(
    getLastVisibleSectionTarget(1510, sections, viewport, {
      origin: 0,
      direction: 1,
    }),
    1630
  )
})

test('a module already substantially visible at the start is not treated as a new arrival', () => {
  assert.equal(
    getLastVisibleSectionTarget(700, sections, viewport, {
      origin: 500,
      direction: 1,
    }),
    null
  )
})

test('small gestures and mismatched travel directions remain with the existing behavior', () => {
  assert.equal(
    getLastVisibleSectionTarget(700, sections, viewport, {
      origin: 620,
      direction: 1,
    }),
    null
  )
  assert.equal(
    getLastVisibleSectionTarget(700, sections, viewport, {
      origin: 1600,
      direction: 1,
    }),
    null
  )
})

test('the added alignment does not pull readers out of a tall module or the footer', () => {
  const tall = [
    sections[0],
    { top: 600, height: 1400, range: { start: 520, end: 1100 } },
    sections[4],
  ]
  assert.equal(
    getLastVisibleSectionTarget(800, tall, viewport, {
      origin: 0,
      direction: 1,
    }),
    null
  )
  assert.equal(
    getLastVisibleSectionTarget(3000, sections, viewport, {
      origin: 2000,
      direction: 1,
    }),
    null
  )
})

test('revealing a module does not allow an excessive correction across most of a screen', () => {
  assert.equal(
    getLastVisibleSectionTarget(450, sections, viewport, {
      origin: 0,
      direction: 1,
    }),
    null
  )
  assert.equal(
    getLastVisibleSectionTarget(700, [], viewport, { origin: 0, direction: 1 }),
    null
  )
})
