import assert from 'node:assert/strict'
import { describe, test } from 'node:test'

import { getHaloBackgroundColor } from '../lib/hero-halo-palette'
import {
  getHaloPlacement,
  getHaloReveal,
  getHaloWaveRadius,
  getHaloWaveOpacity,
  HALO_TRANSITION_DURATION,
  HALO_WAVE_FEATHER,
  HALO_INK_DISTORTION,
} from '../lib/hero-halo-transition'

describe('single watercolor background reveal', () => {
  for (const [width, height] of [
    [1920, 602],
    [1440, 602],
    [960, 602],
    [390, 620],
    [320, 620],
  ]) {
    test(`at ${width}px the wave starts outside and finishes with every background pixel in the new color`, () => {
      const placement = getHaloPlacement(width, height)
      for (let x = 0; x <= 20; x += 1) {
        for (let y = 0; y <= 16; y += 1) {
          const distance = Math.hypot(
            ((x / 20 - placement.origin[0]) * width) / height,
            y / 16 - placement.origin[1]
          )
          for (const distortion of [
            -HALO_INK_DISTORTION,
            0,
            HALO_INK_DISTORTION,
          ]) {
            assert.equal(
              getHaloReveal(
                distance + distortion,
                getHaloWaveRadius(0, placement)
              ),
              0
            )
            assert.equal(
              getHaloReveal(
                distance + distortion,
                getHaloWaveRadius(HALO_TRANSITION_DURATION, placement)
              ),
              1
            )
          }
        }
      }
    })
  }

  test('the swept area changes immediately while the area ahead retains the previous background', () => {
    const placement = getHaloPlacement(1440, 602)
    const radius = getHaloWaveRadius(0.8, placement)
    assert.equal(getHaloReveal(radius - HALO_WAVE_FEATHER * 2, radius), 1)
    assert.equal(getHaloReveal(radius + HALO_WAVE_FEATHER * 2, radius), 0)
    assert.equal(getHaloReveal(radius, radius), 0.5)
  })

  test('each background point changes once without echoes or contraction', () => {
    const placement = getHaloPlacement(1440, 602)
    let previousRadius = -Infinity
    let previousReveal = 0
    for (let frame = 0; frame <= 70; frame += 1) {
      const radius = getHaloWaveRadius(frame / 30, placement)
      const reveal = getHaloReveal(1.3, radius)
      assert.ok(radius >= previousRadius)
      assert.ok(reveal >= previousReveal)
      previousRadius = radius
      previousReveal = reveal
    }
    assert.equal(previousReveal, 1)
    assert.equal(getHaloWaveRadius(20, placement), placement.exitRadius)
  })

  test('paper distortion never reverses the revealed background at any point', () => {
    const placement = getHaloPlacement(1440, 602)
    for (const distortion of [
      -HALO_INK_DISTORTION,
      -0.12,
      0,
      0.12,
      HALO_INK_DISTORTION,
    ]) {
      let previous = 0
      for (let frame = 0; frame <= 80; frame += 1) {
        const reveal = getHaloReveal(
          1.3 + distortion,
          getHaloWaveRadius(frame / 30, placement)
        )
        assert.ok(reveal >= previous)
        previous = reveal
      }
      assert.equal(previous, 1)
    }
  })

  test('the single highlight fades at the end without resetting the revealed colors', () => {
    const placement = getHaloPlacement(390, 620)
    assert.equal(getHaloWaveOpacity(0), 0)
    assert.ok(getHaloWaveOpacity(1) > 0)
    assert.equal(getHaloWaveOpacity(HALO_TRANSITION_DURATION), 0)
    assert.equal(
      getHaloReveal(1, getHaloWaveRadius(HALO_TRANSITION_DURATION, placement)),
      1
    )
  })

  test('the finished background preserves the original theme gradient colors', () => {
    assert.deepEqual(getHaloBackgroundColor('blue', false), [
      44 / 255,
      115 / 255,
      206 / 255,
    ])
    assert.deepEqual(getHaloBackgroundColor('blue', true), [
      123 / 255,
      181 / 255,
      250 / 255,
    ])
  })
})
