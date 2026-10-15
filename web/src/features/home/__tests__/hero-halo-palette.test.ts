import assert from 'node:assert/strict'
import { describe, test } from 'node:test'

import { getHaloColor, interpolateHaloColor } from '../lib/hero-halo-palette'

describe('halo carousel colors', () => {
  test('color transitions ease into and out of their target without a sudden initial change', () => {
    const start = interpolateHaloColor([0, 0, 0], [1, 1, 1], 0.05)
    const finish = interpolateHaloColor([0, 0, 0], [1, 1, 1], 0.95)
    assert.ok(start[0] > 0 && start[0] < 0.02)
    assert.ok(finish[0] < 1 && finish[0] > 0.98)
  })

  test('each carousel theme has its own halo color', () => {
    const colors = ['violet', 'blue', 'teal', 'rose'].map((theme) =>
      getHaloColor(theme as 'violet' | 'blue' | 'teal' | 'rose')
    )
    assert.equal(new Set(colors.map((color) => color.join(','))).size, 4)
    assert.ok(colors.flat().every((value) => value >= 0 && value <= 1))
  })

  test('a slide transition keeps its starting color and arrives at the next theme', () => {
    const from = getHaloColor('violet')
    const to = getHaloColor('teal')
    assert.deepEqual(interpolateHaloColor(from, to, 0), from)
    const final = interpolateHaloColor(from, to, 1)
    final.forEach((channel, index) =>
      assert.ok(Math.abs(channel - to[index]) < 0.000001)
    )
    const middle = interpolateHaloColor(from, to, 0.5)
    assert.notDeepEqual(middle, from)
    assert.notDeepEqual(middle, to)
  })

  test('rapid retargeting starts from the currently displayed color without a jump', () => {
    const current = interpolateHaloColor(
      getHaloColor('violet'),
      getHaloColor('blue'),
      0.4
    )
    assert.deepEqual(
      interpolateHaloColor(current, getHaloColor('rose'), 0),
      current
    )
  })

  test('interpolating colors never mutates the palette shared by other carousel instances', () => {
    const color = getHaloColor('blue')
    color[0] = 1
    assert.notEqual(getHaloColor('blue')[0], 1)
  })
})
