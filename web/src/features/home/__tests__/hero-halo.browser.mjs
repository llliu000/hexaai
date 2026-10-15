import assert from 'node:assert/strict'

// Supply a Playwright page with isolated API fixtures. These checks do not
// construct a browser, contact a backend, or modify the site's saved settings.
export async function verifyHeroHaloSurface(page, renderer = 'webgl') {
  const canvas = page.locator('.wm-hero-halo-canvas')
  assert.equal(await canvas.count(), 1)
  assert.equal(await page.locator('.wm-hero-slide canvas').count(), 0)
  assert.equal(await canvas.getAttribute('data-renderer'), renderer)
  const surface = await page.locator('.wm-hero-halo').evaluate((host) => {
    const canvas = host.querySelector('canvas')
    const hero = host.closest('.wm-hero')
    const bounds = host.getBoundingClientRect()
    return {
      hidden: host.getAttribute('aria-hidden'),
      pointerEvents: getComputedStyle(canvas).pointerEvents,
      matchesHero: Math.abs(bounds.width - hero.clientWidth) < 1,
      overflow: bounds.left < -1 || bounds.right > innerWidth + 1,
      pageOverflow: document.documentElement.scrollWidth > innerWidth,
    }
  })
  assert.equal(surface.hidden, 'true')
  assert.equal(surface.pointerEvents, 'none')
  assert.equal(surface.matchesHero, true)
  assert.equal(surface.overflow, false)
  assert.equal(surface.pageOverflow, false)
}

// The fixture needs at least three slides with different theme colors.
export async function verifyHeroHaloSelection(page) {
  const dots = page.locator('.wm-hero-dots button')
  assert.ok((await dots.count()) >= 3)
  const canvas = await page.locator('.wm-hero-halo-canvas').elementHandle()
  assert.ok(canvas)
  for (const index of [1, 0, 2]) {
    await dots.nth(index).dispatchEvent('click')
  }
  await page.waitForFunction(
    () =>
      document
        .querySelectorAll('.wm-hero-slide')[2]
        ?.getAttribute('aria-hidden') === 'false'
  )
  assert.equal(await dots.nth(2).getAttribute('aria-current'), 'true')
  assert.equal(
    await page.locator('.wm-hero-slide[aria-hidden="false"]').count(),
    1
  )
  assert.equal(await page.locator('.wm-hero h1').count(), 1)
  const states = await page.locator('.wm-hero-slide').evaluateAll((slides) =>
    slides.map((slide) => ({
      hidden: slide.getAttribute('aria-hidden'),
      inert: slide.inert,
    }))
  )
  states.forEach((state, index) => {
    assert.equal(state.hidden, String(index !== 2))
    assert.equal(state.inert, index !== 2)
  })
  assert.equal(
    await canvas.evaluate(
      (element) => element === document.querySelector('.wm-hero-halo-canvas')
    ),
    true
  )
  assert.equal(await page.locator('.wm-hero canvas').count(), 1)
  await canvas.dispose()
}

export async function verifyHeroHaloFallback(page) {
  await verifyHeroHaloSurface(page, 'fallback')
  assert.equal(await page.locator('.wm-hero-halo-fallback').isVisible(), true)
  await verifyHeroHaloSelection(page)
  await page
    .locator('.wm-hero-slide[aria-hidden="false"] h1')
    .waitFor({ state: 'visible' })
}

export async function verifyHeroHaloContextRecovery(page) {
  const canvas = page.locator('.wm-hero-halo-canvas')
  const supported = await canvas.evaluate((element) => {
    const extension = element
      .getContext('webgl')
      ?.getExtension('WEBGL_lose_context')
    if (!extension) return false
    element.__testContextExtension = extension
    extension.loseContext()
    return true
  })
  if (!supported) return { supported: false }
  await page.waitForFunction(
    () =>
      document.querySelector('.wm-hero-halo-canvas')?.dataset.renderer ===
      'fallback'
  )
  assert.equal(await page.locator('.wm-hero-halo-fallback').isVisible(), true)
  await canvas.evaluate((element) =>
    element.__testContextExtension.restoreContext()
  )
  await page.waitForFunction(
    () =>
      document.querySelector('.wm-hero-halo-canvas')?.dataset.renderer ===
      'webgl'
  )
  assert.equal(await page.locator('.wm-hero canvas').count(), 1)
  return { supported: true }
}

// Install Playwright's clock before page load, and emulate reduced motion before
// loading the fixture. Compositor screenshots preserve WebGL output correctly.
export async function verifyHeroHaloReducedMotion(page, intervalMs = 3000) {
  const canvas = page.locator('.wm-hero-halo-canvas')
  const before = await canvas.screenshot()
  const selected = await page
    .locator('.wm-hero-dots button[aria-current="true"]')
    .getAttribute('aria-label')
  await page.clock.runFor(intervalMs + 100)
  assert.deepEqual(await canvas.screenshot(), before)
  assert.equal(
    await page
      .locator('.wm-hero-dots button[aria-current="true"]')
      .getAttribute('aria-label'),
    selected
  )
  await page.locator('.wm-hero-dots button').nth(1).click()
  await page.waitForFunction(
    () =>
      document
        .querySelectorAll('.wm-hero-slide')[1]
        ?.getAttribute('aria-hidden') === 'false'
  )
  const changed = await canvas.screenshot()
  assert.notDeepEqual(changed, before)
  await page.clock.runFor(500)
  assert.deepEqual(await canvas.screenshot(), changed)
}
