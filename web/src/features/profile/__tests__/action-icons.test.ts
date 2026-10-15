import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'
import { describe, test } from 'node:test'

const sessionsSource = readFileSync(
  new URL('../components/login-session-item.tsx', import.meta.url),
  'utf8'
)
const securitySource = readFileSync(
  new URL('../components/profile-security-card.tsx', import.meta.url),
  'utf8'
)
const styles = readFileSync(
  new URL('../profile-design.css', import.meta.url),
  'utf8'
)

describe('profile action icons', () => {
  test('retains the session action label when the decorative exit icon is hidden', () => {
    assert.match(sessionsSource, /className='profile-session-action'/)
    assert.match(sessionsSource, /<DesignIcon name='logout' \/>/)
    assert.match(
      styles,
      /\.profile-session-action \.icon-logout\s*\{[^}]*display: none;/s
    )
    assert.match(sessionsSource, /t\('Sign out'\)/)
    assert.match(sessionsSource, /t\('Revoke'\)/)
  })

  test('renders a right arrow in each security action', () => {
    assert.match(securitySource, /<DesignIcon\s+name='chevron-right'/)
    assert.match(securitySource, /className='profile-security-arrow'/)
    assert.match(styles, /\.profile-security-arrow\s*\{/)
  })
})
