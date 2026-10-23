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
import { readFileSync } from 'node:fs'
import { test } from 'node:test'

import { Window } from 'happy-dom'

import { applyFaviconToDom } from '../dom-utils.ts'
import {
  resolveSystemLogo,
  SYSTEM_LOGO_PATH,
  SYSTEM_LOGO_STOPS,
  SYSTEM_LOGO_URL,
  SYSTEM_LOGO_VIEW_BOX,
} from '../system-logo.ts'

test('missing and legacy raster defaults resolve to the vector mark', () => {
  for (const source of [
    undefined,
    null,
    '',
    ' ',
    '/logo.png',
    'logo.png',
    '/logo.png?v=old',
    '/favicon.ico',
    '/lanhu/home-header-logo.png',
    '/lanhu/home-footer-wordmark.png',
  ]) {
    assert.equal(resolveSystemLogo(source), SYSTEM_LOGO_URL)
  }
})

test('explicit custom branding remains configurable after the vector default upgrade', () => {
  for (const source of [
    'https://brand.example.test/logo.png',
    '/uploads/company.svg',
    'data:image/svg+xml,%3Csvg/%3E',
  ]) {
    assert.equal(resolveSystemLogo(` ${source} `), source)
  }
})

test('the favicon uses the same vector geometry and gradient as the inline component, with no embedded raster', () => {
  const window = new Window()
  try {
    const source = readFileSync(
      new URL('../../../public/logo.svg', import.meta.url),
      'utf8'
    )
    const document = new window.DOMParser().parseFromString(
      source,
      'image/svg+xml'
    )
    assert.equal(
      document.querySelector('svg')?.getAttribute('viewBox'),
      SYSTEM_LOGO_VIEW_BOX
    )
    assert.equal(
      document.querySelector('path')?.getAttribute('d'),
      SYSTEM_LOGO_PATH
    )
    assert.equal(
      document.querySelectorAll('image, script, foreignObject').length,
      0
    )
    assert.deepEqual(
      [...document.querySelectorAll('stop')].map((stop) => ({
        offset: stop.getAttribute('offset'),
        color: stop.getAttribute('stop-color'),
      })),
      SYSTEM_LOGO_STOPS
    )
    assert.equal(new Set(SYSTEM_LOGO_STOPS.map((stop) => stop.color)).size, 3)
    assert.equal(
      document.querySelector('path')?.getAttribute('fill'),
      `url(#${document.querySelector('linearGradient')?.id})`
    )
  } finally {
    window.close()
  }
})

test('a cached PNG favicon is replaced with one SVG icon without duplicate head entries', () => {
  const window = new Window({ url: 'http://localhost/' })
  const previousDocument = Object.getOwnPropertyDescriptor(
    globalThis,
    'document'
  )
  const previousWindow = Object.getOwnPropertyDescriptor(globalThis, 'window')
  Object.defineProperty(globalThis, 'document', {
    configurable: true,
    value: window.document,
  })
  Object.defineProperty(globalThis, 'window', {
    configurable: true,
    value: window,
  })
  try {
    window.document.head.innerHTML =
      '<link rel="icon" href="/logo.png"><link rel="shortcut icon" href="/favicon.ico">'
    applyFaviconToDom('/logo.png')
    applyFaviconToDom('/logo.svg')
    const icons = window.document.querySelectorAll('link[rel~="icon"]')
    assert.equal(icons.length, 1)
    assert.equal(icons[0].getAttribute('href'), SYSTEM_LOGO_URL)
    assert.equal(icons[0].getAttribute('type'), 'image/svg+xml')
  } finally {
    if (previousDocument) {
      Object.defineProperty(globalThis, 'document', previousDocument)
    } else {
      Reflect.deleteProperty(globalThis, 'document')
    }
    if (previousWindow) {
      Object.defineProperty(globalThis, 'window', previousWindow)
    } else {
      Reflect.deleteProperty(globalThis, 'window')
    }
    window.close()
  }
})

test('public and console header branding cannot be replaced by the old PNG background override', () => {
  const styles = readFileSync(
    new URL('../../styles/index.css', import.meta.url),
    'utf8'
  )
  assert.doesNotMatch(styles, /lujiyun-wordmark\.png/)
  assert.doesNotMatch(
    styles,
    /header\.fixed nav > a:first-child > \*\s*\{\s*visibility: hidden/
  )
})
