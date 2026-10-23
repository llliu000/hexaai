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

import {
  openNativeScrollHome,
  freezeScrollClock,
  resetNativePosition,
  sectionCenter,
  moveNatively,
  finishNativeGesture,
  assertPosition,
} from './native-scroll-fixture.browser.mjs'

export async function verifyTrackpadPaging(page, url) {
  await openNativeScrollHome(page, url)
  await freezeScrollClock(page)
  const passed = []
  const models = await sectionCenter(page, '#models')
  const near = Math.round(models - 60)

  await resetNativePosition(page)
  await moveNatively(
    page,
    Array.from({ length: 40 }, (_, index) =>
      Math.round((near * (index + 1)) / 40)
    )
  )
  await finishNativeGesture(page)
  await assertPosition(page, await sectionCenter(page, '#models'))
  passed.push(
    'small continuous trackpad deltas follow native positions until scrolling ends'
  )

  await resetNativePosition(page)
  await moveNatively(page, [120])
  await moveNatively(page, [230, near - 10, near], { cancelable: false })
  await page.evaluate(() =>
    document
      .querySelector('.wm-home main')
      .dispatchEvent(
        new WheelEvent('wheel', { bubbles: true, deltaY: 0, deltaX: 0 })
      )
  )
  await finishNativeGesture(page)
  await assertPosition(page, await sectionCenter(page, '#models'))
  passed.push(
    'non-cancelable momentum and a zero-delta ending do not suppress nearby alignment'
  )

  await resetNativePosition(page)
  await moveNatively(page, [near])
  await page.evaluate(() => document.dispatchEvent(new Event('scrollend')))
  await page.clock.runFor(200)
  const correcting = await page.evaluate(() => scrollY)
  assert.ok(correcting > near && correcting < models)
  const stopped = correcting - 35
  await moveNatively(page, [stopped])
  await page.clock.runFor(80)
  await assertPosition(page, stopped)
  await finishNativeGesture(page)
  await assertPosition(page, stopped)
  passed.push(
    'a new upward gesture immediately interrupts correction and stays under user control'
  )

  await resetNativePosition(page)
  const ignored = await page.evaluate(() => {
    const root = document.querySelector('.wm-home main')
    return [
      { deltaX: 100, deltaY: 2 },
      { ctrlKey: true, deltaY: 50 },
    ].map((init) => {
      const event = new WheelEvent('wheel', {
        bubbles: true,
        cancelable: true,
        ...init,
      })
      root.dispatchEvent(event)
      return event.defaultPrevented
    })
  })
  assert.deepEqual(ignored, [false, false])
  await page.clock.runFor(1200)
  await assertPosition(page, 0)
  passed.push(
    'horizontal swipes and zoom are never converted to vertical alignment'
  )
  await page.clock.resume()
  return passed
}
