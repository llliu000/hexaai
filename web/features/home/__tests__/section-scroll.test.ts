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
  getSectionNudgeTarget,
  getSectionScrollRanges,
  getSectionSettleTarget,
} from '../lib/section-scroll.ts'

const ranges = [0, 400, 1050, 1650, 2200].map((position) => ({
  start: position,
  end: position,
}))

test('settling aligns the nearest module to the actual native resting position', () => {
  assert.equal(getSectionSettleTarget(340, ranges, 900, 0), 400)
  assert.equal(getSectionSettleTarget(450, ranges, 900, 0), 400)
})

test('a long native swipe can pass several modules before choosing its nearby destination', () => {
  assert.equal(getSectionSettleTarget(2070, ranges, 900, 0), 2200)
  assert.equal(getSectionSettleTarget(1120, ranges, 900, 2200), 1050)
})

test('overshooting a nearby center settles back a little instead of predicting another page', () => {
  assert.equal(getSectionSettleTarget(1130, ranges, 900, 400), 1050)
})

test('small departures from an aligned module are not pulled back to the starting point', () => {
  assert.equal(getSectionSettleTarget(90, ranges, 900, 0), null)
  assert.equal(getSectionSettleTarget(500, ranges, 900, 400), null)
  assert.equal(getSectionSettleTarget(950, ranges, 900, 1050), null)
})

test('a correction never drags the page back beyond the start of the gesture', () => {
  assert.equal(getSectionSettleTarget(290, ranges, 900, 350), null)
})

test('tiny deliberate adjustments and already centered modules stay where the user left them', () => {
  assert.equal(getSectionSettleTarget(385, ranges, 900, 370), null)
  assert.equal(getSectionSettleTarget(400, ranges, 900, 0), null)
  assert.equal(getSectionSettleTarget(401, ranges, 900, 0), null)
})

test('positions between distant centers remain native rather than jumping across a large gap', () => {
  assert.equal(getSectionSettleTarget(730, ranges, 900, 0), null)
  assert.equal(getSectionSettleTarget(245, ranges, 400, 0), null)
})

test('reading the interior or edge of a tall module never snaps its content away', () => {
  const tall = [
    { start: 0, end: 0 },
    { start: 920, end: 1600 },
    { start: 2200, end: 2200 },
  ]
  assert.equal(getSectionSettleTarget(1150, tall, 800, 400), null)
  assert.equal(getSectionSettleTarget(1500, tall, 800, 2200), null)
  assert.equal(getSectionSettleTarget(920, tall, 800, 400), null)
  assert.equal(getSectionSettleTarget(830, tall, 800, 400), 920)
})

test('scrolling just past the last module eases back so that module stays centered', () => {
  assert.equal(getSectionSettleTarget(2270, ranges, 900, 1650), 2200)
  assert.equal(getSectionSettleTarget(2270, ranges, 900, 2200), 2200)
  assert.equal(getSectionSettleTarget(2500, ranges, 900, 1650), null)
  assert.equal(getSectionSettleTarget(500, [], 900, 0), null)
})

test('equally close modules follow the actual direction of travel', () => {
  const equal = [
    { start: 300, end: 300 },
    { start: 600, end: 600 },
  ]
  assert.equal(getSectionSettleTarget(450, equal, 900, 0), 600)
  assert.equal(getSectionSettleTarget(450, equal, 900, 900), 300)
})

test('short modules center below the fixed header and targets stay within document bounds', () => {
  assert.deepEqual(
    getSectionScrollRanges(
      [
        { top: 0, height: 600 },
        { top: 800, height: 400 },
        { top: 1300, height: 0 },
        { top: 1800, height: 200 },
      ],
      800,
      80,
      1200
    ),
    [
      { start: 0, end: 0 },
      { start: 560, end: 560 },
      { start: 1200, end: 1200 },
    ]
  )
})

test('tall modules expose a full reading range rather than a forced center', () => {
  assert.deepEqual(
    getSectionScrollRanges([{ top: 1000, height: 1400 }], 800, 80, 3000),
    [{ start: 920, end: 1600 }]
  )
})

test('a short small gesture from an aligned module turns to its adjacent module', () => {
  assert.equal(
    getSectionNudgeTarget(90, ranges, 900, {
      origin: 0,
      distance: 90,
      durationMs: 60,
    }),
    400
  )
  assert.equal(
    getSectionNudgeTarget(480, ranges, 900, {
      origin: 400,
      distance: 80,
      durationMs: 120,
    }),
    1050
  )
  assert.equal(
    getSectionNudgeTarget(970, ranges, 900, {
      origin: 1050,
      distance: 80,
      durationMs: 120,
    }),
    400
  )
})

test('tiny movements stay native while a medium slide turns one page either way', () => {
  assert.equal(
    getSectionNudgeTarget(430, ranges, 900, {
      origin: 400,
      distance: 30,
      durationMs: 60,
    }),
    null
  )
  assert.equal(
    getSectionNudgeTarget(650, ranges, 900, {
      origin: 400,
      distance: 250,
      durationMs: 60,
    }),
    1050
  )
  assert.equal(
    getSectionNudgeTarget(800, ranges, 900, {
      origin: 1050,
      distance: 250,
      durationMs: 180,
    }),
    400
  )
})

test('a full-viewport fling is left to settle or reveal instead of a one-page nudge', () => {
  assert.equal(
    getSectionNudgeTarget(850, ranges, 900, {
      origin: 400,
      distance: 450,
      durationMs: 60,
    }),
    null
  )
})

test('a slow deliberate slide between modules still turns one page', () => {
  assert.equal(
    getSectionNudgeTarget(490, ranges, 900, {
      origin: 400,
      distance: 90,
      durationMs: 640,
    }),
    1050
  )
})

test('a very long reading gesture is not mistaken for a page turn', () => {
  assert.equal(
    getSectionNudgeTarget(490, ranges, 900, {
      origin: 400,
      distance: 90,
      durationMs: 1200,
    }),
    null
  )
})

test('a longer gesture returning near its origin does not trigger small paging', () => {
  assert.equal(
    getSectionNudgeTarget(490, ranges, 900, {
      origin: 400,
      distance: 410,
      durationMs: 100,
    }),
    null
  )
  assert.equal(
    getSectionNudgeTarget(470, ranges, 900, {
      origin: 400,
      distance: 160,
      durationMs: 100,
    }),
    null
  )
})

test('small paging only starts from an already aligned module', () => {
  assert.equal(
    getSectionNudgeTarget(530, ranges, 900, {
      origin: 460,
      distance: 70,
      durationMs: 100,
    }),
    null
  )
  assert.equal(
    getSectionNudgeTarget(488, ranges, 900, {
      origin: 408,
      distance: 80,
      durationMs: 100,
    }),
    1050
  )
})

test('small paging does not escape tall content or skip an oversized gap', () => {
  const tall = [
    { start: 0, end: 0 },
    { start: 920, end: 1600 },
    { start: 2200, end: 2200 },
  ]
  assert.equal(
    getSectionNudgeTarget(1000, tall, 900, {
      origin: 920,
      distance: 80,
      durationMs: 100,
    }),
    null
  )
  // About one viewport into the next tall module is allowed for downward entry.
  assert.equal(
    getSectionNudgeTarget(90, tall, 900, {
      origin: 0,
      distance: 90,
      durationMs: 100,
    }),
    920
  )
  const wide = [
    { start: 0, end: 0 },
    { start: 1400, end: 2000 },
    { start: 2600, end: 2600 },
  ]
  assert.equal(
    getSectionNudgeTarget(90, wide, 900, {
      origin: 0,
      distance: 90,
      durationMs: 100,
    }),
    null
  )
})

test('the first and last module boundaries stay native', () => {
  assert.equal(
    getSectionNudgeTarget(-80, ranges, 900, {
      origin: 0,
      distance: 80,
      durationMs: 100,
    }),
    null
  )
  assert.equal(
    getSectionNudgeTarget(2280, ranges, 900, {
      origin: 2200,
      distance: 80,
      durationMs: 100,
    }),
    null
  )
})
