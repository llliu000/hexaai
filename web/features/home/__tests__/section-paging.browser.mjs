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

export async function verifyHomeSectionPaging(page, url) {
  await openNativeScrollHome(page, url)
  const passed = []
  const models = await sectionCenter(page, '#models')
  const rawPosition = Math.round(models - 70)
  await page.mouse.move(1100, 420)
  await page.mouse.wheel(0, rawPosition)
  await page.waitForFunction((position) => scrollY >= position - 1, rawPosition)
  await assertPosition(page, rawPosition)
  await page.waitForFunction(() => {
    const rect = document.querySelector('#models').getBoundingClientRect()
    const header = document.querySelector('.wm-header').offsetHeight
    return (
      Math.abs(rect.top + rect.height / 2 - (innerHeight + header) / 2) <= 2
    )
  })
  passed.push(
    'a real wheel first scrolls natively, then aligns a nearby module'
  )

  await freezeScrollClock(page)
  await resetNativePosition(page)
  await moveNatively(page, [24])
  await finishNativeGesture(page)
  await assertPosition(page, 24)
  passed.push('tiny adjustments remain native without a page turn or pullback')

  await resetNativePosition(page)
  const integration = await sectionCenter(page, '#integration')
  await moveNatively(page, [400, 850, Math.round(integration + 70)])
  await finishNativeGesture(page)
  await assertPosition(page, await sectionCenter(page, '#integration'))
  passed.push(
    'resting just past a module aligns that module instead of advancing another page'
  )

  await resetNativePosition(page)
  const gap = Math.round((models + integration) / 2)
  await resetNativePosition(page, gap - 20)
  await moveNatively(page, [gap])
  await finishNativeGesture(page)
  await assertPosition(page, gap)
  passed.push(
    'a small adjustment without a newly revealed module is left untouched'
  )

  await page.setViewportSize({ width: 1440, height: 600 })
  await page.clock.runFor(32)
  const readingPosition = await page
    .locator('#integration')
    .evaluate((section) => {
      const rect = section.getBoundingClientRect()
      const start =
        scrollY + rect.top - document.querySelector('.wm-header').offsetHeight
      const end = scrollY + rect.bottom - innerHeight
      return Math.round((start + end) / 2)
    })
  await resetNativePosition(page)
  await moveNatively(page, [readingPosition])
  await finishNativeGesture(page)
  await assertPosition(page, readingPosition)
  passed.push(
    'reading inside a tall module remains native after the gesture stops'
  )

  await page.setViewportSize({ width: 1440, height: 900 })
  await page.clock.runFor(32)
  const end = await page.evaluate(
    () => document.documentElement.scrollHeight - innerHeight
  )
  await resetNativePosition(page)
  await moveNatively(page, [end])
  await finishNativeGesture(page)
  await assertPosition(page, await sectionCenter(page, '.wm-providers'))
  passed.push('scrolling to the end centers the last module')

  const keyboardCanceled = await page.evaluate(() => {
    const event = new KeyboardEvent('keydown', {
      key: 'PageDown',
      bubbles: true,
      cancelable: true,
    })
    document.body.dispatchEvent(event)
    return event.defaultPrevented
  })
  assert.equal(keyboardCanceled, false)
  passed.push('keyboard scrolling is not intercepted')

  await page.emulateMedia({ reducedMotion: 'reduce' })
  await page.clock.runFor(32)
  await resetNativePosition(page)
  await moveNatively(page, [rawPosition])
  await finishNativeGesture(page)
  await assertPosition(page, rawPosition)
  passed.push('reduced-motion mode keeps the exact native resting position')

  await page.emulateMedia({ reducedMotion: 'no-preference' })
  await page.setViewportSize({ width: 390, height: 844 })
  await page.clock.runFor(32)
  await resetNativePosition(page)
  await moveNatively(page, [330])
  await finishNativeGesture(page)
  await assertPosition(page, 330)
  passed.push('mobile scrolling remains fully native')
  await page.clock.resume()
  await page.setViewportSize({ width: 1440, height: 900 })
  await page.evaluate(async () => {
    window.dispatchEvent(new PointerEvent('pointerdown'))
    window.scrollTo({ top: 0, behavior: 'instant' })
    await new Promise((resolve) =>
      requestAnimationFrame(() => requestAnimationFrame(resolve))
    )
  })
  const dynamicStart = Math.round(
    (await sectionCenter(page, '#integration')) - 70
  )
  await page.mouse.move(1100, 420)
  await page.mouse.wheel(0, dynamicStart)
  await page.waitForFunction(
    (position) => scrollY > position + 5 && scrollY < position + 55,
    dynamicStart
  )
  await page.evaluate(() => {
    const section = document.querySelector('#integration')
    section.dataset.originalSettlePadding = section.style.paddingBottom
    section.style.paddingBottom = `${Number.parseFloat(getComputedStyle(section).paddingBottom) + 60}px`
  })
  await page.waitForFunction(() => {
    const rect = document.querySelector('#integration').getBoundingClientRect()
    const header = document.querySelector('.wm-header').offsetHeight
    return (
      Math.abs(rect.top + rect.height / 2 - (innerHeight + header) / 2) <= 2
    )
  })
  await page.evaluate(() => {
    const section = document.querySelector('#integration')
    section.style.paddingBottom = section.dataset.originalSettlePadding
    delete section.dataset.originalSettlePadding
  })
  passed.push(
    'content resizing during the final correction still lands at the updated center'
  )
  return passed
}

export async function verifyHomeSectionScrollFallback(page, url) {
  await page.addInitScript(() => {
    let prototype = document
    while (prototype) {
      if (Object.hasOwn(prototype, 'onscrollend')) {
        delete prototype.onscrollend
      }
      prototype = Object.getPrototypeOf(prototype)
    }
    document.addEventListener(
      'scrollend',
      (event) => event.stopImmediatePropagation(),
      true
    )
  })
  await openNativeScrollHome(page, url)
  assert.equal(await page.evaluate(() => 'onscrollend' in document), false)
  await freezeScrollClock(page)
  await resetNativePosition(page)
  const target = await sectionCenter(page, '#models')
  await moveNatively(page, [120, 260, Math.round(target - 60)])
  await page.clock.runFor(1200)
  await assertPosition(page, await sectionCenter(page, '#models'))
  await page.clock.resume()
  return 'browsers without scrollend settle only after native scrolling becomes quiet'
}
