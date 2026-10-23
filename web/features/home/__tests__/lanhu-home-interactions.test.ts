import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'
import { test } from 'node:test'

import { Window } from 'happy-dom'

import {
  selectCodeLanguage,
  toggleProviderScrolling,
  watchHeaderScroll,
} from '../lib/lanhu-home-interactions.ts'

const template = readFileSync(
  new URL('../../../../option/home.html', import.meta.url),
  'utf8'
)

test('selecting Python updates the example and makes only the selected tab keyboard-focusable', () => {
  const window = new Window()
  window.document.body.innerHTML = template
  const root = window.document as unknown as Document
  selectCodeLanguage(root, 'python')
  assert.match(
    root.querySelector('#lh-code-panel')?.textContent ?? '',
    /from openai import OpenAI/
  )
  assert.equal(
    root.querySelector('#lh-code-panel')?.getAttribute('aria-labelledby'),
    'lh-tab-python'
  )
  assert.equal(
    root.querySelectorAll('[role="tab"][aria-selected="true"]').length,
    1
  )
  assert.equal(
    root.querySelector<HTMLButtonElement>('#lh-tab-python')?.tabIndex,
    0
  )
  assert.equal(
    root.querySelector<HTMLButtonElement>('#lh-tab-curl')?.tabIndex,
    -1
  )
  window.happyDOM.abort()
})

test('returning to CURL after JavaScript restores the highlighted original example', () => {
  const window = new Window()
  window.document.body.innerHTML = template
  const root = window.document as unknown as Document
  const original = root.querySelector('#lh-code-panel')?.innerHTML
  selectCodeLanguage(root, 'javascript')
  assert.match(
    root.querySelector('#lh-code-panel')?.textContent ?? '',
    /await client.chat.completions.create/
  )
  selectCodeLanguage(root, 'curl')
  assert.equal(root.querySelector('#lh-code-panel')?.innerHTML, original)
  window.happyDOM.abort()
})

test('an unsupported language leaves the current example and selected tab unchanged', () => {
  const window = new Window()
  window.document.body.innerHTML = template
  const root = window.document as unknown as Document
  const original = root.querySelector('#lh-code-panel')?.innerHTML
  selectCodeLanguage(root, '__proto__')
  assert.equal(root.querySelector('#lh-code-panel')?.innerHTML, original)
  assert.equal(
    root.querySelector('#lh-tab-curl')?.getAttribute('aria-selected'),
    'true'
  )
  window.happyDOM.abort()
})

test('scrolling past the original threshold floats the header and returning to the top restores it', () => {
  const window = new Window()
  const states: boolean[] = []
  const stop = watchHeaderScroll(
    window as unknown as globalThis.Window,
    (floating) => states.push(floating)
  )
  for (const top of [21, 100, 0]) {
    window.scrollTo(0, top)
    window.dispatchEvent(new window.Event('scroll'))
  }
  assert.deepEqual(states, [false, true, false])
  stop()
  window.scrollTo(0, 100)
  window.dispatchEvent(new window.Event('scroll'))
  assert.deepEqual(states, [false, true, false])
  window.happyDOM.abort()
})

test('the marquee pause button toggles both the animation state and its accessible pressed state', () => {
  const window = new Window()
  window.document.body.innerHTML =
    '<div class="lh-providers"><button data-pause-providers aria-pressed="false"></button></div>'
  const button = window.document.querySelector(
    'button'
  ) as unknown as HTMLButtonElement
  toggleProviderScrolling(button)
  assert.equal(button.getAttribute('aria-pressed'), 'true')
  assert.equal(button.parentElement?.dataset.paused, 'true')
  toggleProviderScrolling(button)
  assert.equal(button.getAttribute('aria-pressed'), 'false')
  assert.equal(button.parentElement?.dataset.paused, 'false')
  window.happyDOM.abort()
})
