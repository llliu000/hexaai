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

export async function verifySmallSectionScroll(page, url) {
  await openNativeScrollHome(page, url)
  await freezeScrollClock(page)
  const passed = []

  await resetNativePosition(page)
  await moveNatively(page, [30, 60, 90])
  await assertPosition(page, 90)
  await finishNativeGesture(page)
  await assertPosition(page, await sectionCenter(page, '#models'))
  passed.push(
    'a small downward gesture stays native until release, then turns one page'
  )

  const models = await sectionCenter(page, '#models')
  await resetNativePosition(page, models)
  await moveNatively(page, [Math.round(models + 80)])
  await finishNativeGesture(page)
  await assertPosition(page, await sectionCenter(page, '#integration'))
  passed.push(
    'a small gesture from a centered module advances only to its neighbor'
  )

  const integration = await sectionCenter(page, '#integration')
  await resetNativePosition(page, integration)
  await moveNatively(page, [Math.round(integration - 80)])
  await finishNativeGesture(page)
  await assertPosition(page, await sectionCenter(page, '#models'))
  passed.push('small upward gestures return to the adjacent module')

  await resetNativePosition(page, models)
  const tiny = Math.round(models + 24)
  await moveNatively(page, [tiny])
  await finishNativeGesture(page)
  await assertPosition(page, tiny)
  passed.push('tiny adjustments remain precise without triggering a page turn')

  await resetNativePosition(page, models)
  const reversed = Math.round(models + 90)
  await moveNatively(page, [Math.round(models + 250), reversed])
  await finishNativeGesture(page)
  await assertPosition(page, reversed)
  passed.push(
    'a long gesture ending near its origin is not mistaken for a small page turn'
  )

  const unaligned = Math.round(models + 60)
  await resetNativePosition(page, unaligned)
  await moveNatively(page, [unaligned + 70])
  await finishNativeGesture(page)
  await assertPosition(page, unaligned + 70)
  passed.push(
    'ordinary scrolling between modules keeps the existing nearby-settle behavior'
  )

  await resetNativePosition(page)
  await moveNatively(
    page,
    Array.from({ length: 40 }, (_, index) =>
      Math.round((90 * (index + 1)) / 40)
    )
  )
  await finishNativeGesture(page)
  await assertPosition(page, await sectionCenter(page, '#models'))
  passed.push(
    'a slow deliberate slide from the top turns one page to the next module'
  )

  await resetNativePosition(page)
  await moveNatively(page, [90])
  await page.evaluate(() => document.dispatchEvent(new Event('scrollend')))
  await page.clock.runFor(200)
  const moving = await page.evaluate(() => scrollY)
  assert.ok(moving > 90)
  await moveNatively(page, [moving - 20])
  await page.clock.runFor(80)
  await assertPosition(page, moving - 20)
  passed.push(
    'a new gesture can interrupt the small page-turn animation immediately'
  )

  await page.setViewportSize({ width: 1440, height: 600 })
  await page.clock.runFor(32)
  const reading = await page.locator('#integration').evaluate((section) => {
    const rect = section.getBoundingClientRect()
    return Math.round(
      scrollY +
        rect.top -
        document.querySelector('.wm-header').offsetHeight +
        30
    )
  })
  await resetNativePosition(page, reading)
  await moveNatively(page, [reading + 50])
  await finishNativeGesture(page)
  await assertPosition(page, reading + 50)
  passed.push(
    'small gestures inside tall content remain normal reading scrolls'
  )

  await page.clock.resume()
  return passed
}
