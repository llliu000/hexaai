// Run against a signed-in profile page with a Playwright-compatible page handle.
export async function verifyProfileLayout(page) {
  return page.evaluate(() => {
    const root = document.querySelector('.profile-design')
    if (!root) throw new Error('Live profile styles are not mounted')
    const rootRect = root.getBoundingClientRect()
    const cards = [...root.querySelectorAll('[data-slot="card"]')]
    if (cards.length < 7) throw new Error('Profile sections are missing')
    for (const card of cards) {
      const rect = card.getBoundingClientRect()
      if (
        rect.width < 200 ||
        rect.left < rootRect.left - 1 ||
        rect.right > rootRect.right + 1
      ) {
        throw new Error('A profile card is collapsed or overflows the page')
      }
    }
    for (const track of root.querySelectorAll('[data-slot="switch"]')) {
      const outer = track.getBoundingClientRect()
      const thumb = track
        .querySelector('[data-slot="switch-thumb"]')
        .getBoundingClientRect()
      if (
        thumb.left < outer.left ||
        thumb.right > outer.right ||
        thumb.top < outer.top ||
        thumb.bottom > outer.bottom
      ) {
        throw new Error('Switch thumb overflows its track')
      }
    }
    if (innerWidth === 1440) {
      const summary = root
        .querySelector('.profile-summary')
        .getBoundingClientRect()
      const side = root.querySelector('.profile-side').getBoundingClientRect()
      if (Math.abs(summary.width - 800) > 1 || Math.abs(side.width - 352) > 1) {
        throw new Error('Desktop columns do not match the reference dimensions')
      }
      const sidebar = document.querySelector('.profile-design-sidebar')
      if (
        !sidebar ||
        Math.abs(sidebar.getBoundingClientRect().width - 202) > 1
      ) {
        throw new Error('Reference sidebar is missing or has the wrong width')
      }
      const modules = root
        .querySelector('.profile-modules-main')
        .getBoundingClientRect()
      if (Math.abs(summary.top - modules.top) > 1) {
        throw new Error('Desktop profile columns are not vertically aligned')
      }
      const personal = root
        .querySelector('.profile-modules-personal')
        .getBoundingClientRect()
      if (
        Math.abs(modules.height - 509.3359375) > 1 ||
        Math.abs(personal.height - 150.18359375) > 1
      ) {
        throw new Error(
          'Sidebar settings cards do not match the reference geometry'
        )
      }
      const expected = [
        ['.profile-summary', 236, 147, 800, 250],
        ['.profile-settings', 236, 415.1796875, 800, 170],
        ['.profile-modules-main', 1056, 147, 352, 509.3359375],
        ['.profile-modules-personal', 1056, 673.16015625, 352, 150.18359375],
      ]
      for (const [selector, x, y, width, height] of expected) {
        const actual = root.querySelector(selector).getBoundingClientRect()
        for (const [key, value] of Object.entries({ x, y, width, height })) {
          if (Math.abs(actual[key] - value) > 1) {
            throw new Error(`${selector}: unexpected ${key}`)
          }
        }
      }
    }
    const mainCards = [
      ...root.querySelectorAll('.profile-main > [data-slot="card"]'),
    ]
    for (let index = 1; index < mainCards.length; index++) {
      if (
        mainCards[index].getBoundingClientRect().top <
        mainCards[index - 1].getBoundingClientRect().bottom
      ) {
        throw new Error('Content-sized profile cards must not overlap')
      }
    }
    const languageRow = root.querySelector(
      '.profile-language [data-slot="card-content"] > div'
    )
    if (
      languageRow &&
      (languageRow.scrollWidth > languageRow.clientWidth + 1 ||
        languageRow.scrollHeight > languageRow.clientHeight + 1)
    ) {
      throw new Error('Language preference content must fit without clipping')
    }
    const icons = root.querySelectorAll('.profile-design-icon')
    if (icons.length < 10) {
      throw new Error('Reference icons have not been restored')
    }
    for (const image of document.querySelectorAll(
      '.profile-design img, .profile-design-sidebar img, .profile-design-header img'
    )) {
      if (!image.complete || image.naturalWidth === 0) {
        throw new Error(
          `Design asset failed to load: ${image.getAttribute('src')}`
        )
      }
    }
    if (innerWidth >= 1200) {
      const circleNames = [
        'settings',
        'language',
        'security',
        'balance',
        'usage',
        'requests',
        'password',
        'token',
        'delete',
        'modules',
        'sessions',
      ]
      for (const name of circleNames) {
        const icon = root.querySelector(`.icon-${name}`)
        if (!icon) throw new Error(`Missing ${name} icon`)
        const rect = icon.getBoundingClientRect()
        const svg = icon.querySelector('svg')
        if (Math.abs(rect.width - 36) > 1 || Math.abs(rect.height - 36) > 1) {
          throw new Error(`${name} icon must use the shared 36px badge size`)
        }
        if (!svg) throw new Error(`${name} icon must use a linear SVG`)
        const svgRect = svg.getBoundingClientRect()
        const style = getComputedStyle(svg)
        if (
          Math.abs(svgRect.width - 20) > 1 ||
          Math.abs(svgRect.height - 20) > 1 ||
          style.fill !== 'none' ||
          Number.parseFloat(style.strokeWidth) !== 2
        ) {
          throw new Error(
            `${name} icon does not match the model analysis linear icons`
          )
        }
      }
      const gift = document
        .querySelector('.profile-gift')
        .getBoundingClientRect()
      if (Math.abs(gift.width - 87) > 1 || Math.abs(gift.height - 87) > 1) {
        throw new Error(
          'Invite illustration does not preserve its reference size'
        )
      }
    }
    for (const action of root.querySelectorAll('.profile-security button')) {
      const arrow = action.querySelector('.profile-security-arrow')
      if (!arrow || arrow.getBoundingClientRect().width === 0) {
        throw new Error('Security action lost its navigation arrow')
      }
    }
    for (const action of root.querySelectorAll('.profile-session-action')) {
      if (!action.textContent.trim()) {
        throw new Error('Login session action has no visible label')
      }
    }
    return { viewport: innerWidth, cards: cards.length, passed: true }
  })
}

export async function verifySidebarDraftCancellation(page) {
  const toggle = page.locator('.profile-module-row [role="switch"]').first()
  const original = await toggle.getAttribute('aria-checked')
  await toggle.click()
  if ((await toggle.getAttribute('aria-checked')) === original) {
    throw new Error('Sidebar setting did not update in the draft')
  }
  const cancel = page.locator('.profile-module-save button').first()
  await cancel.click()
  if ((await toggle.getAttribute('aria-checked')) !== original) {
    throw new Error(
      'Cancelling did not restore the saved sidebar configuration'
    )
  }
  if (await page.locator('.profile-module-save').count()) {
    throw new Error('Save controls remained after cancelling changes')
  }
}

export async function verifySettingsTabs(page) {
  const tabs = page.locator('.profile-design [role="tab"]')
  await tabs.nth(1).click()
  if ((await tabs.nth(1).getAttribute('aria-selected')) !== 'true') {
    throw new Error('Preferences tab did not activate')
  }
  await tabs.nth(0).click()
  if ((await tabs.nth(0).getAttribute('aria-selected')) !== 'true') {
    throw new Error('Bindings tab did not activate')
  }
}

// Exercise with several sessions, including a long IPv6 address/login method,
// at desktop and mobile widths. The session list must grow with its contents.
export async function verifyLoginSessionsLayout(page) {
  return page.evaluate(() => {
    const card = document.querySelector('.profile-sessions')
    if (!card) throw new Error('Login sessions card is not mounted')
    const content = card.querySelector('[data-slot="card-content"]')
    const header = card.querySelector('[data-slot="card-header"]')
    const rows = [...card.querySelectorAll('.profile-session-row')]
    if (!content || !header) {
      throw new Error('Session card structure is missing')
    }
    const contentRect = content.getBoundingClientRect()
    const cardRect = card.getBoundingClientRect()
    if (content.scrollHeight > content.clientHeight + 1) {
      throw new Error(
        'Login sessions are constrained to a cramped inner scroll area'
      )
    }
    if (header.getBoundingClientRect().bottom > contentRect.top + 1) {
      throw new Error('Login session header overlaps the list')
    }
    let previousBottom = contentRect.top - 12
    for (const row of rows) {
      const rect = row.getBoundingClientRect()
      if (rect.top < previousBottom + 11) {
        throw new Error(
          'Login session rows need breathing room between devices'
        )
      }
      if (
        rect.bottom > contentRect.bottom + 1 ||
        rect.bottom > cardRect.bottom + 1
      ) {
        throw new Error('Login session card clips its devices')
      }
      if (row.scrollWidth > row.clientWidth + 1) {
        throw new Error('A login session overflows horizontally')
      }
      const text = [...row.querySelectorAll('p')]
      const action = row.querySelector('[data-slot="button"]')
      if (!action) throw new Error('Login session action is missing')
      const actionRect = action.getBoundingClientRect()
      for (const paragraph of text) {
        const textRect = paragraph.getBoundingClientRect()
        const style = getComputedStyle(paragraph)
        if (
          Number.parseFloat(style.fontSize) < 12 ||
          Number.parseFloat(style.lineHeight) < 18
        ) {
          throw new Error(
            'Login session details are too small or tightly spaced'
          )
        }
        if (paragraph.scrollWidth > paragraph.clientWidth + 1) {
          throw new Error('Long login session metadata does not wrap')
        }
        const overlapsAction =
          textRect.left < actionRect.right &&
          textRect.right > actionRect.left &&
          textRect.top < actionRect.bottom &&
          textRect.bottom > actionRect.top
        if (overlapsAction) {
          throw new Error('Login session details overlap their action')
        }
      }
      previousBottom = rect.bottom
    }
    return { viewport: innerWidth, sessions: rows.length, passed: true }
  })
}
