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

export async function openNativeScrollHome(page, url) {
  await page.clock.install()
  await page.setViewportSize({ width: 1440, height: 900 })
  await page.emulateMedia({ reducedMotion: 'no-preference' })
  await page.goto(url, { waitUntil: 'domcontentloaded' })
  await page.locator('.wm-home main > section').first().waitFor()
  await page.evaluate(() => document.fonts.ready)
}

export async function freezeScrollClock(page) {
  await page.clock.pauseAt(await page.evaluate(() => Date.now() + 1000))
}

export async function resetNativePosition(page, position = 0) {
  await page.evaluate((top) => {
    window.dispatchEvent(new PointerEvent('pointerdown'))
    window.scrollTo({ top, behavior: 'instant' })
    window.dispatchEvent(new Event('scroll'))
  }, position)
  await page.clock.runFor(32)
}

export async function sectionCenter(page, selector) {
  return page.locator(selector).evaluate((section) => {
    const rect = section.getBoundingClientRect()
    const header = document.querySelector('.wm-header').offsetHeight
    return Math.max(
      0,
      Math.min(
        document.documentElement.scrollHeight - innerHeight,
        scrollY + rect.top + rect.height / 2 - (innerHeight + header) / 2
      )
    )
  })
}

export async function moveNatively(page, positions, options = {}) {
  for (const position of positions) {
    const result = await page.evaluate(
      ({ top, init }) => {
        const event = new WheelEvent('wheel', {
          bubbles: true,
          cancelable: true,
          deltaY: top - scrollY,
          ...init,
        })
        document.querySelector('.wm-home main').dispatchEvent(event)
        // Synthetic wheel events have no default action: model the native compositor separately.
        window.scrollTo({ top, behavior: 'instant' })
        window.dispatchEvent(new Event('scroll'))
        return event.defaultPrevented
      },
      { top: position, init: options }
    )
    assert.equal(result, false, 'Native wheel input must never be canceled')
    await page.clock.runFor(16)
    assert.ok(
      Math.abs((await page.evaluate(() => scrollY)) - position) <= 1,
      'The page must follow its actual position throughout a gesture'
    )
  }
}

export async function finishNativeGesture(page) {
  await page.evaluate(() => document.dispatchEvent(new Event('scrollend')))
  await page.clock.runFor(1200)
}

export async function assertPosition(page, expected) {
  const actual = await page.evaluate(() => scrollY)
  assert.ok(
    Math.abs(actual - expected) <= 2,
    `Expected resting position ${expected}, got ${actual}`
  )
}
