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
  getLoopedModelScrollPosition,
  getModelScrollEdges,
} from '../lib/model-scroll.ts'

test('model navigation disables both arrows when every card fits', () => {
  assert.deepEqual(getModelScrollEdges(0, 1360, 1360), {
    previous: false,
    next: false,
  })
  assert.deepEqual(getModelScrollEdges(0, 1360, 420), {
    previous: false,
    next: false,
  })
})

test('model navigation enables only the direction with remaining content', () => {
  assert.deepEqual(getModelScrollEdges(0, 1360, 3000), {
    previous: false,
    next: true,
  })
  assert.deepEqual(getModelScrollEdges(440, 1360, 3000), {
    previous: true,
    next: true,
  })
  assert.deepEqual(getModelScrollEdges(1639.5, 1360, 3000), {
    previous: true,
    next: false,
  })
  assert.deepEqual(getModelScrollEdges(-8, 1360, 3000), {
    previous: false,
    next: true,
  })
})

test('automatic model scrolling continues into the next copy without jumping backward', () => {
  assert.deepEqual(getLoopedModelScrollPosition(1200, 1200, 440), {
    start: 1200,
    target: 1640,
  })
  assert.deepEqual(getLoopedModelScrollPosition(2360, 1200, 440), {
    start: 2360,
    target: 2800,
  })
  assert.deepEqual(getLoopedModelScrollPosition(2520, 1200, 440), {
    start: 1320,
    target: 1760,
  })
})

test('looped model scrolling preserves the visual position when moving backward', () => {
  assert.deepEqual(getLoopedModelScrollPosition(1200, 1200, -440), {
    start: 1200,
    target: 760,
  })
  assert.deepEqual(getLoopedModelScrollPosition(760, 1200, -440), {
    start: 1960,
    target: 1520,
  })
  assert.deepEqual(getLoopedModelScrollPosition(100, 0, 440), {
    start: 0,
    target: 0,
  })
})
