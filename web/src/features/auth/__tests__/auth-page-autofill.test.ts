import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'
import { test } from 'node:test'

const styles = readFileSync(
  new URL('../../../styles/auth-pages.css', import.meta.url),
  'utf8'
)

test('Chrome autofill matches the default login and registration field colors', () => {
  assert.match(
    styles,
    /\.auth-page \.auth-page-form input:-webkit-autofill,[\s\S]*?\.auth-page \.auth-page-form input:-webkit-autofill:active/
  )
  assert.match(
    styles,
    /-webkit-text-fill-color: var\(--design-text, #101a31\);/
  )
  assert.match(styles, /caret-color: var\(--design-text, #101a31\);/)
  assert.match(
    styles,
    /-webkit-box-shadow: 0 0 0 1000px var\(--design-surface, #fff\) inset;/
  )
})
