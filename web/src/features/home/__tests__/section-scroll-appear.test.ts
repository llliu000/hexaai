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

import {
  getAppearingSectionTarget,
  getEnteredSectionTarget,
} from '../lib/section-scroll-appear.ts'

const viewport = { height: 900, headerHeight: 80 }
const sections = [
  { top: 0, height: 600, range: { start: 0, end: 0 } },
  { top: 600, height: 560, range: { start: 390, end: 390 } },
  { top: 1160, height: 680, range: { start: 1010, end: 1010 } },
  { top: 1840, height: 560, range: { start: 1630, end: 1630 } },
  { top: 2400, height: 740, range: { start: 2280, end: 2280 } },
  { top: 3140, height: 400, range: { start: 2850, end: 2850 } },
]

test('downward scrolling turns the page once the next module starts appearing', () => {
  assert.equal(
    getAppearingSectionTarget(520, sections, viewport, {
      origin: 390,
      direction: 1,
    }),
    1010
  )
})

test('upward scrolling turns the page once the previous module grows into view', () => {
  assert.equal(
    getAppearingSectionTarget(900, sections, viewport, {
      origin: 1010,
      direction: -1,
    }),
    390
  )
})

test('tiny motion before the next module appears stays native', () => {
  assert.equal(
    getAppearingSectionTarget(410, sections, viewport, {
      origin: 390,
      direction: 1,
    }),
    null
  )
})

test('downward page turns still work when the next module already peeks', () => {
  assert.equal(
    getAppearingSectionTarget(470, sections, viewport, {
      origin: 390,
      direction: 1,
    }),
    1010
  )
})

test('downward scrolling can leave a tall module once the next module appears', () => {
  const tall = [
    { top: 0, height: 500, range: { start: 0, end: 0 } },
    { top: 500, height: 1400, range: { start: 420, end: 1100 } },
    { top: 1900, height: 500, range: { start: 1710, end: 1710 } },
  ]
  assert.equal(
    getAppearingSectionTarget(1180, tall, viewport, {
      origin: 600,
      direction: 1,
    }),
    1710
  )
})

test('long flings are not converted into a one-page appear turn', () => {
  assert.equal(
    getAppearingSectionTarget(1100, sections, viewport, {
      origin: 390,
      direction: 1,
    }),
    null
  )
})

test('a downward overshoot into the top of a tall module still finishes that page turn', () => {
  const tall = [
    { top: 0, height: 500, range: { start: 0, end: 0 } },
    { top: 500, height: 560, range: { start: 390, end: 390 } },
    { top: 1060, height: 1400, range: { start: 980, end: 1660 } },
  ]
  assert.equal(
    getEnteredSectionTarget(1100, tall, viewport, {
      origin: 390,
      direction: 1,
    }),
    980
  )
  assert.equal(
    getEnteredSectionTarget(1400, tall, viewport, {
      origin: 390,
      direction: 1,
    }),
    null
  )
})

test('the first and last boundaries do not invent an adjacent page', () => {
  assert.equal(
    getAppearingSectionTarget(40, sections, viewport, {
      origin: 0,
      direction: -1,
    }),
    null
  )
  assert.equal(
    getAppearingSectionTarget(2900, sections, viewport, {
      origin: 2850,
      direction: 1,
    }),
    null
  )
})
