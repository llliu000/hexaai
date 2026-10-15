import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'
import { describe, test } from 'node:test'

const themeStyles = readFileSync(new URL('../theme.css', import.meta.url), 'utf8')
const customStylePaths = [
  '../api-keys-lanhu.css',
  '../auth-pages.css',
  '../flow-lanhu.css',
  '../index.css',
  '../lanhu-home.css',
  '../models-lanhu.css',
  '../pricing-lanhu.css',
  '../reference-home.css',
  '../task-logs-lanhu.css',
  '../wallet-lanhu.css',
  '../../features/profile/profile-design.css',
] as const

describe('system font', () => {
  test('defines the requested font stack as the global sans face', () => {
    assert.match(
      themeStyles,
      /--font-sans:\s*Inter,\s*-apple-system,\s*BlinkMacSystemFont,\s*'Segoe UI',\s*'PingFang SC',\s*'Hiragino Sans GB',\s*'Microsoft YaHei',\s*'Helvetica Neue',\s*Helvetica,\s*Arial,\s*sans-serif;/s
    )
    assert.match(themeStyles, /--font-serif:\s*var\(--font-sans\);/)
  })

  test('does not override the global font on customized pages', () => {
    for (const path of customStylePaths) {
      const styles = readFileSync(new URL(path, import.meta.url), 'utf8')
      assert.doesNotMatch(
        styles,
        /font-family:\s*(?:Arial|'Inter'|Inter,|Georgia|var\(--font-body\),\s*Arial)/
      )
    }
  })
})
