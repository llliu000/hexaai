import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'
import { test } from 'node:test'

import { Window } from 'happy-dom'

const palette = readFileSync(
  new URL('../dark-palette.css', import.meta.url),
  'utf8'
)

test('monochrome navigation image icons stay visible in dark mode', () => {
  for (const [theme, current, filter] of [
    ['dark', false, 'brightness(0) invert(0.8)'],
    ['dark', true, 'brightness(0) invert(1)'],
    ['light', true, 'none'],
  ] as const) {
    const window = new Window()
    try {
      window.document.documentElement.className = theme
      window.document.head.innerHTML = `<style>.profile-nav-item img { filter: none; }\n${palette}</style>`
      window.document.body.innerHTML = `<a class="profile-nav-item ${current ? 'is-current' : ''}"><img class="profile-nav-icon" alt=""></a>`
      const icon = window.document.querySelector('img')
      assert.ok(icon)
      assert.equal(window.getComputedStyle(icon).filter, filter)
    } finally {
      window.close()
    }
  }
})

for (const fixture of [
  {
    name: 'login form',
    stylesheet: '../auth-pages.css',
    html: '<main class="auth-page"><form class="auth-page-form"><input aria-label="Email"></form></main>',
    selector: 'input',
    lightBackground: '#ffffff',
    darkBackground: '#292929',
    lightColor: '#101a31',
    darkColor: '#edf0f5',
  },
  {
    name: 'account workspace',
    stylesheet: '../../features/profile/profile-design.css',
    html: '<main class="profile-design">Account settings</main>',
    selector: 'main',
    lightBackground: '#f4f7fc',
    darkBackground: '#1e1e1e',
    lightColor: '#14213b',
    darkColor: '#edf0f5',
  },
  {
    name: 'overview card',
    stylesheet: '../overview-lanhu.css',
    html: '<main class="lujiyun-overview-page"><section class="overview-card">Usage</section></main>',
    selector: 'section',
    lightBackground: '#fff',
    darkBackground: '#292929',
    lightColor: '#15213a',
    darkColor: '#edf0f5',
  },
]) {
  test(`${fixture.name} follows theme changes and restores its light colors`, () => {
    const window = new Window()
    try {
      const stylesheet = readFileSync(
        new URL(fixture.stylesheet, import.meta.url),
        'utf8'
      )
      window.document.head.innerHTML = `<style>${palette}\n${stylesheet}</style>`
      window.document.body.innerHTML = fixture.html
      const element = window.document.querySelector(fixture.selector)
      assert.ok(element)
      for (const theme of ['light', 'dark', 'light'] as const) {
        window.document.documentElement.className = theme
        const style = window.getComputedStyle(element)
        assert.equal(
          style.backgroundColor,
          theme === 'dark' ? fixture.darkBackground : fixture.lightBackground
        )
        assert.equal(
          style.color,
          theme === 'dark' ? fixture.darkColor : fixture.lightColor
        )
      }
    } finally {
      window.close()
    }
  })
}

test('dark workspace canvas, pricing panels, and pinned table cells share the neutral palette', () => {
  const workspaceStyles = readFileSync(
    new URL('../index.css', import.meta.url),
    'utf8'
  )
  const darkBodyRule = workspaceStyles.match(/html\.dark body\s*\{[^}]*\}/)?.[0]
  assert.ok(darkBodyRule)

  const window = new Window()
  try {
    window.document.documentElement.className = 'dark'
    window.document.head.innerHTML = `<style>${palette}\n${darkBodyRule}\nbody, .workspace { background: var(--background); }\n.model-panel { background: var(--card); }\n.pinned-cell { background: var(--table-row); }</style>`
    window.document.body.innerHTML =
      '<main class="workspace"><section class="model-panel"><span class="pinned-cell">Edit</span></section></main>'

    const workspace = window.document.querySelector('.workspace')
    const modelPanel = window.document.querySelector('.model-panel')
    const pinnedCell = window.document.querySelector('.pinned-cell')
    assert.ok(workspace)
    assert.ok(modelPanel)
    assert.ok(pinnedCell)

    const background = window.getComputedStyle(
      window.document.body
    ).backgroundColor
    assert.equal(background, '#1e1e1e')
    assert.equal(window.getComputedStyle(workspace).backgroundColor, background)
    assert.equal(window.getComputedStyle(modelPanel).backgroundColor, '#292929')
    assert.equal(window.getComputedStyle(pinnedCell).backgroundColor, '#292929')
  } finally {
    window.close()
  }
})

test('analytics tiles keep muted, readable accent surfaces in the dark workspace', () => {
  const workspaceStyles = readFileSync(
    new URL('../index.css', import.meta.url),
    'utf8'
  )
  const window = new Window()
  try {
    window.document.documentElement.className = 'dark'
    window.document.head.innerHTML = `<style>${palette}\n${workspaceStyles}</style>`
    window.document.body.innerHTML = `<main class="lujiyun-dashboard-models"><div><div>${Array.from({ length: 5 }, (_, index) => `<div data-tile="${index}"></div>`).join('')}</div></div></main>`

    for (const [index, background] of [
      '#233248',
      '#213a32',
      '#332d45',
      '#21383e',
      '#3c3324',
    ].entries()) {
      const tile = window.document.querySelector(`[data-tile="${index}"]`)
      assert.ok(tile)
      assert.equal(window.getComputedStyle(tile).backgroundColor, background)
    }
  } finally {
    window.close()
  }
})

test('public model catalog and homepage use the same neutral dark canvas as the workspace', () => {
  const workspaceStyles = readFileSync(
    new URL('../index.css', import.meta.url),
    'utf8'
  )
  const homepageStyles = readFileSync(
    new URL('../reference-home.css', import.meta.url),
    'utf8'
  )
  const window = new Window()
  try {
    window.document.documentElement.className = 'dark'
    window.document.head.innerHTML = `<style>${palette}\n${workspaceStyles}\n${homepageStyles}</style>`
    window.document.body.innerHTML =
      '<main class="lujiyun-pricing-page"></main>'
    assert.equal(
      window.getComputedStyle(window.document.body).backgroundColor,
      '#1e1e1e'
    )

    window.document.body.innerHTML =
      '<main class="wm-home"><section class="wm-integration"><div class="wm-model-card"></div></section></main>'
    const home = window.document.querySelector('.wm-home')
    const feature = window.document.querySelector('.wm-integration')
    const card = window.document.querySelector('.wm-model-card')
    assert.ok(home)
    assert.ok(feature)
    assert.ok(card)
    assert.equal(window.getComputedStyle(home).backgroundColor, '#1e1e1e')
    assert.equal(window.getComputedStyle(feature).backgroundColor, '#1e1e1e')
    assert.equal(window.getComputedStyle(card).backgroundColor, '#292929')
  } finally {
    window.close()
  }
})
