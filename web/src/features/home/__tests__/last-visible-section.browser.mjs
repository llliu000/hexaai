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

async function centerOf(page, selector) {
  return page.locator(selector).evaluate((element) => {
    const rect = element.getBoundingClientRect()
    return (
      scrollY +
      rect.top +
      rect.height / 2 -
      (innerHeight + document.querySelector('.wm-header').offsetHeight) / 2
    )
  })
}

async function resetAt(page, position) {
  await page.evaluate(async (top) => {
    window.dispatchEvent(new PointerEvent('pointerdown'))
    scrollTo({ top, behavior: 'instant' })
    await new Promise((resolve) =>
      requestAnimationFrame(() => requestAnimationFrame(resolve))
    )
  }, position)
}

async function waitForCenter(page, selector) {
  await page.waitForFunction((target) => {
    const rect = document.querySelector(target).getBoundingClientRect()
    const center =
      (innerHeight + document.querySelector('.wm-header').offsetHeight) / 2
    return Math.abs(rect.top + rect.height / 2 - center) <= 2
  }, selector)
}

export async function verifyLastVisibleSection(page, url) {
  await page.setViewportSize({ width: 1440, height: 900 })
  await page.emulateMedia({ reducedMotion: 'no-preference' })
  await page.goto(url, { waitUntil: 'domcontentloaded' })
  await page.locator('.wm-home main > section').first().waitFor()
  await page.evaluate(() => document.fonts.ready)
  const models = await centerOf(page, '#models')
  const integration = await centerOf(page, '#integration')
  const advantages = await centerOf(page, '#advantages')
  const gap = Math.round((models + integration) / 2)
  const passed = []

  await page.mouse.move(1100, 420)
  await page.mouse.wheel(0, gap)
  await page.waitForFunction((position) => scrollY >= position - 1, gap)
  const samples = await page.evaluate(
    async ({ destination }) => {
      const points = []
      for (let frame = 0; frame < 150; frame += 1) {
        await new Promise((resolve) => requestAnimationFrame(resolve))
        points.push(scrollY)
        if (Math.abs(scrollY - destination) <= 2) break
      }
      return { points }
    },
    { destination: integration }
  )
  await waitForCenter(page, '#integration')
  assert.ok(
    samples.points.some((value) => value > gap + 5 && value < integration - 5),
    'The final alignment should visibly scroll rather than teleport'
  )
  assert.ok(
    samples.points.length > 3,
    'The page-turn alignment should take multiple frames like a slide'
  )
  passed.push(
    'after native scrolling, the newly visible module slides into alignment'
  )

  await resetAt(page, advantages)
  await page.mouse.move(1100, 420)
  await page.mouse.wheel(0, gap - advantages)
  await waitForCenter(page, '#models')
  passed.push('upward scrolling aligns the nearest module still ahead from above')

  await resetAt(page, 0)
  await page.mouse.move(1100, 420)
  await page.mouse.wheel(0, gap)
  await page.waitForFunction(
    (position) => scrollY > position + 5 && scrollY < position + 100,
    gap
  )
  await page.evaluate(() => {
    window.__revealWheelPrevented = null
    window.addEventListener(
      'wheel',
      (event) => {
        window.__revealWheelPrevented = event.defaultPrevented
      },
      { once: true }
    )
  })
  await page.mouse.wheel(0, -60)
  const stopped = await page.evaluate(async () => {
    const positions = []
    for (let frame = 0; frame < 20; frame += 1) {
      await new Promise((resolve) => requestAnimationFrame(resolve))
      positions.push(scrollY)
    }
    const tail = positions.slice(-10)
    return {
      prevented: window.__revealWheelPrevented,
      drift: Math.max(...tail) - Math.min(...tail),
      position: scrollY,
    }
  })
  assert.equal(stopped.prevented, false)
  assert.ok(stopped.drift <= 2, JSON.stringify(stopped))
  assert.ok(stopped.position < integration - 20)
  passed.push(
    'a fresh gesture immediately cancels the sliding alignment without blocking scrolling'
  )

  await resetAt(page, 0)
  await page.mouse.move(1100, 420)
  await page.mouse.wheel(0, gap)
  await page.waitForFunction(
    (position) => scrollY > position + 5 && scrollY < position + 100,
    gap
  )
  await page.evaluate(() => {
    const section = document.querySelector('#integration')
    section.dataset.beforeRevealPadding = section.style.paddingBottom
    section.style.paddingBottom = `${Number.parseFloat(getComputedStyle(section).paddingBottom) + 60}px`
  })
  await waitForCenter(page, '#integration')
  await page.evaluate(() => {
    const section = document.querySelector('#integration')
    section.style.paddingBottom = section.dataset.beforeRevealPadding
    delete section.dataset.beforeRevealPadding
  })
  passed.push(
    'the native alignment follows the selected module when its content height changes'
  )
  return passed
}

export async function verifyLastVisibleSectionFallback(page, url) {
  await page.addInitScript(() => {
    let prototype = document
    while (prototype) {
      if (Object.hasOwn(prototype, 'onscrollend')) delete prototype.onscrollend
      prototype = Object.getPrototypeOf(prototype)
    }
    document.addEventListener(
      'scrollend',
      (event) => event.stopImmediatePropagation(),
      true
    )
  })
  return verifyLastVisibleSection(page, url)
}
