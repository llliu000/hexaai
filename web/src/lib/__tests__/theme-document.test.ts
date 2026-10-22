import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'
import { test } from 'node:test'

import { Window } from 'happy-dom'

import { applyThemeToDocument } from '../theme-document.ts'

test('system dark mode updates the page, native controls and browser toolbar together', () => {
  const window = new Window()
  try {
    window.document.head.innerHTML = '<meta name="theme-color" content="#fff">'
    window.document.documentElement.className = 'light app-ready'
    applyThemeToDocument(window.document as unknown as Document, 'dark')
    assert.equal(
      window.document.documentElement.classList.contains('dark'),
      true
    )
    assert.equal(
      window.document.documentElement.classList.contains('light'),
      false
    )
    assert.equal(
      window.document.documentElement.classList.contains('app-ready'),
      true
    )
    assert.equal(window.document.documentElement.style.colorScheme, 'dark')
    assert.equal(
      window.document
        .querySelector('meta[name="theme-color"]')
        ?.getAttribute('content'),
      '#1e1e1e'
    )

    applyThemeToDocument(window.document as unknown as Document, 'light')
    assert.equal(
      window.document.documentElement.classList.contains('dark'),
      false
    )
    assert.equal(window.document.documentElement.style.colorScheme, 'light')
    assert.equal(
      window.document
        .querySelector('meta[name="theme-color"]')
        ?.getAttribute('content'),
      '#f4f7fc'
    )
  } finally {
    window.close()
  }
})

test('theme changes also work on documents without a toolbar color meta tag', () => {
  const window = new Window()
  try {
    assert.doesNotThrow(() =>
      applyThemeToDocument(window.document as unknown as Document, 'dark')
    )
    assert.equal(
      window.document.documentElement.classList.contains('dark'),
      true
    )
  } finally {
    window.close()
  }
})

test('the initial HTML resolves the theme before loading the application', () => {
  const source = readFileSync(
    new URL('../../../index.html', import.meta.url),
    'utf8'
  )
  const script = source.match(
    /<script data-theme-init>([\s\S]*?)<\/script>/
  )?.[1]
  assert.ok(script, 'theme initialization must run in the document head')
  for (const [cookie, prefersDark, expected] of [
    ['', true, 'dark'],
    ['', false, 'light'],
    ['vite-ui-theme=system', true, 'dark'],
    ['vite-ui-theme=light', true, 'light'],
    ['vite-ui-theme=dark', false, 'dark'],
    ['vite-ui-theme=invalid', true, 'dark'],
  ] as const) {
    const window = new Window({ url: 'http://localhost/' })
    try {
      window.document.head.innerHTML =
        '<meta name="theme-color" content="#fff">'
      window.document.cookie = cookie
      Object.defineProperty(window, 'matchMedia', {
        value: () => ({ matches: prefersDark }),
      })
      window.eval(script)
      assert.equal(
        window.document.documentElement.classList.contains(expected),
        true,
        `${cookie}, prefersDark=${prefersDark}`
      )
      assert.equal(window.document.documentElement.style.colorScheme, expected)
    } finally {
      window.close()
    }
  }
})
