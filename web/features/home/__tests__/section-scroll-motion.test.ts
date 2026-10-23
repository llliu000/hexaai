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
  advanceSectionScrollMotion,
  pageTurnScrollMotion,
} from '../lib/section-scroll-motion.ts'

test('a new page transition starts moving on the first animation frame', () => {
  const next = advanceSectionScrollMotion({ position: 0, velocity: 0 }, 600, 16)
  assert.ok(next.position > 8)
  assert.ok(next.position < 600)
})

test('a page-turn seed already glides toward a distant module', () => {
  const seeded = pageTurnScrollMotion(0, 600)
  assert.ok(seeded.velocity > 0)
  const next = advanceSectionScrollMotion(seeded, 600, 16)
  assert.ok(next.position > 20)
  assert.ok(next.position < 600)
})

test('nearby settle corrections stay at rest until the spring pulls them', () => {
  const motion = pageTurnScrollMotion(400, 480)
  assert.equal(motion.position, 400)
  assert.equal(motion.velocity, 0)
  assert.equal(motion.frequency, 20)
})

test('a new forward destination continues moving rather than restarting from rest', () => {
  let motion = { position: 0, velocity: 0 }
  for (let frame = 0; frame < 6; frame += 1) {
    motion = advanceSectionScrollMotion(motion, 600, 16)
  }
  const previous = motion
  motion = advanceSectionScrollMotion(motion, 1200, 16)
  assert.ok(motion.position > previous.position)
  assert.ok(motion.velocity >= previous.velocity)
  for (let frame = 0; frame < 80; frame += 1) {
    const next = advanceSectionScrollMotion(motion, 1200, 16)
    assert.ok(next.position >= motion.position)
    assert.ok(next.position <= 1200)
    motion = next
  }
  assert.equal(motion.position, 1200)
  assert.equal(motion.velocity, 0)
})

test('reversing a fast gesture turns back smoothly and settles at the new destination', () => {
  let motion = { position: 300, velocity: 1800 }
  const first = advanceSectionScrollMotion(motion, 0, 16)
  assert.ok(Math.abs(first.position - motion.position) < 40)
  motion = first
  for (let frame = 0; frame < 80; frame += 1) {
    motion = advanceSectionScrollMotion(motion, 0, 16)
    assert.ok(motion.position >= 0)
  }
  assert.equal(motion.position, 0)
  assert.equal(motion.velocity, 0)
})

test('60 Hz and 120 Hz displays produce the same scroll progress over equal time', () => {
  let standard = { position: 0, velocity: 0 }
  let highRefresh = { position: 0, velocity: 0 }
  for (let frame = 0; frame < 12; frame += 1) {
    standard = advanceSectionScrollMotion(standard, 800, 1000 / 60)
  }
  for (let frame = 0; frame < 24; frame += 1) {
    highRefresh = advanceSectionScrollMotion(highRefresh, 800, 1000 / 120)
  }
  assert.ok(Math.abs(standard.position - highRefresh.position) < 0.001)
})

test('a delayed browser frame cannot jump abruptly to the destination', () => {
  const next = advanceSectionScrollMotion(
    { position: 0, velocity: 0 },
    1200,
    1000
  )
  assert.ok(next.position > 0)
  assert.ok(next.position < 600)
})
