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
