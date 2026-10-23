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

export async function verifyHomeScrollPerformance(page, url) {
  await page.addInitScript(() => {
    const stats = {
      enabled: false,
      rectReads: 0,
      sectionRectReads: 0,
      styleReads: 0,
      callbacks: 0,
    }
    window.__homeScrollWork = stats
    const rect = Element.prototype.getBoundingClientRect
    Element.prototype.getBoundingClientRect = function () {
      if (stats.enabled) {
        stats.rectReads += 1
        if (this.matches('.wm-home main > section, .wm-home > header')) {
          stats.sectionRectReads += 1
        }
      }
      return rect.call(this)
    }
    const style = window.getComputedStyle
    window.getComputedStyle = function (...args) {
      if (stats.enabled) stats.styleReads += 1
      return style.apply(this, args)
    }
    const frame = window.requestAnimationFrame.bind(window)
    window.__homeScrollWorkFrame = frame
    window.requestAnimationFrame = (callback) =>
      frame((time) => {
        if (stats.enabled) stats.callbacks += 1
        callback(time)
      })
  })
  await page.setViewportSize({ width: 1440, height: 900 })
  await page.emulateMedia({ reducedMotion: 'no-preference' })
  await page.goto(url, { waitUntil: 'domcontentloaded' })
  await page.locator('.wm-home main > section').first().waitFor()
  await page.evaluate(() => document.fonts.ready)
  const work = await page.evaluate(async () => {
    const root = document.querySelector('.wm-home main')
    const stats = window.__homeScrollWork
    window.dispatchEvent(new PointerEvent('pointerdown'))
    window.scrollTo({ top: 0, behavior: 'instant' })
    stats.enabled = true
    for (let frame = 0; frame < 180; frame += 1) {
      const step = frame % 24
      if (step < 4) {
        const direction = Math.floor(frame / 72) % 2 === 0 ? 1 : -1
        root.dispatchEvent(
          new WheelEvent('wheel', {
            bubbles: true,
            cancelable: true,
            deltaY: [100, 60, 30, 65][step] * direction,
          })
        )
        // Wheel events created by a test have no browser default scroll action.
        window.scrollBy({
          top: [100, 60, 30, 65][step] * direction,
          behavior: 'instant',
        })
      }
      await new Promise((resolve) => window.__homeScrollWorkFrame(resolve))
    }
    stats.enabled = false
    return {
      callbacks: stats.callbacks,
      rectReads: stats.rectReads,
      sectionRectReads: stats.sectionRectReads,
      frames: 180,
    }
  })
  // Work budgets protect smooth scrolling without machine-specific FPS checks.
  assert.ok(
    work.callbacks <= work.frames * 12,
    `Animation work grew with layer count: ${JSON.stringify(work)}`
  )
  assert.ok(
    work.sectionRectReads < work.frames,
    `Stable scrolling repeatedly remeasured layout: ${JSON.stringify(work)}`
  )

  const before = await page.evaluate(() => {
    window.dispatchEvent(new PointerEvent('pointerdown'))
    const section = document.querySelector('#advantages')
    const rect = section.getBoundingClientRect()
    const header = document.querySelector('.wm-header').offsetHeight
    const target =
      scrollY + rect.top + rect.height / 2 - (innerHeight + header) / 2
    window.scrollTo({ top: target, behavior: 'instant' })
    return target
  })
  await page.evaluate(
    () =>
      new Promise((resolve) =>
        requestAnimationFrame(() => requestAnimationFrame(resolve))
      )
  )
  const transforms = await page.evaluate(() => {
    const heading = document.querySelector(
      '#advantages .wm-home-parallax-heading'
    )
    const backdrop = document.querySelector(
      '#advantages .wm-home-parallax-backdrop-slow'
    )
    return {
      heading: new DOMMatrixReadOnly(getComputedStyle(heading).transform).m42,
      backdrop: new DOMMatrixReadOnly(getComputedStyle(backdrop).transform).m42,
    }
  })
  await page.evaluate(
    (position) => scrollTo({ top: position + 80, behavior: 'instant' }),
    before
  )
  await page.waitForFunction((initial) => {
    const heading = document.querySelector(
      '#advantages .wm-home-parallax-heading'
    )
    return (
      new DOMMatrixReadOnly(getComputedStyle(heading).transform).m42 <
      initial.heading
    )
  }, transforms)
  const backdrop = await page
    .locator('#advantages .wm-home-parallax-backdrop-slow')
    .evaluate(
      (element) =>
        new DOMMatrixReadOnly(getComputedStyle(element).transform).m42
    )
  assert.ok(backdrop >= transforms.backdrop)
  assert.ok(backdrop <= 56)

  await page.evaluate(() => {
    window.dispatchEvent(new PointerEvent('pointerdown'))
    scrollTo({ top: 2000, behavior: 'instant' })
  })
  await page.evaluate(
    () =>
      new Promise((resolve) =>
        requestAnimationFrame(() => requestAnimationFrame(resolve))
      )
  )
  const outside = await page
    .locator('.wm-hero-parallax-copy')
    .first()
    .evaluate((element) => element.style.transform)
  await page.evaluate(() => scrollTo({ top: 2200, behavior: 'instant' }))
  await page.evaluate(
    () =>
      new Promise((resolve) =>
        requestAnimationFrame(() => requestAnimationFrame(resolve))
      )
  )
  assert.equal(
    await page
      .locator('.wm-hero-parallax-copy')
      .first()
      .evaluate((element) => element.style.transform),
    outside
  )

  await page.emulateMedia({ reducedMotion: 'reduce' })
  await page.reload({ waitUntil: 'domcontentloaded' })
  await page.locator('.wm-home main > section').first().waitFor()
  await page.waitForFunction(() =>
    [
      ...document.querySelectorAll(
        '.wm-home-parallax-heading, .wm-hero-parallax-copy'
      ),
    ].every((element) => element.style.transform === '')
  )
  return {
    ...work,
    parallaxPreserved: true,
    reducedMotionDisablesParallax: true,
  }
}
