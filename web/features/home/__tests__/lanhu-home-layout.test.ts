import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'
import { test } from 'node:test'

import { Window } from 'happy-dom'

const template = readFileSync(
  new URL('../../../../option/home.html', import.meta.url),
  'utf8'
)
const styles = readFileSync(
  new URL('../../../styles/lanhu-home.css', import.meta.url),
  'utf8'
)

test('at phone width, hero and CLI grids allow content to shrink within a single column', () => {
  const window = new Window({ width: 390, height: 844 })
  window.document.head.innerHTML = `<style>${styles}</style>`
  window.document.body.innerHTML = template
  for (const selector of ['.lh-hero-grid', '.lh-cli-grid']) {
    const grid = window.document.querySelector(selector)
    assert.ok(grid)
    assert.equal(
      window.getComputedStyle(grid).gridTemplateColumns,
      'minmax(0, 1fr)'
    )
  }
  window.happyDOM.abort()
})

test('at design width, the footer keeps separate 135px QR codes without transform offsets', () => {
  const window = new Window({ width: 1440, height: 960 })
  window.document.head.innerHTML = `<style>${styles}</style>`
  window.document.body.innerHTML = template
  const images = window.document.querySelectorAll('.lh-contact img')
  assert.equal(images.length, 2)
  for (const image of images) {
    const style = window.getComputedStyle(image)
    assert.equal(style.width, '135px')
    assert.equal(style.height, '135px')
    assert.ok(!style.transform || style.transform === 'none')
  }
  window.happyDOM.abort()
})

for (const width of [1200, 1024, 901, 768, 390, 320]) {
  test(`at ${width}px both footer QR codes stay square and equal in size`, () => {
    const window = new Window({ width, height: 960 })
    window.document.head.innerHTML = `<style>${styles}</style>`
    window.document.body.innerHTML = template
    const images = window.document.querySelectorAll('.lh-contact img')
    assert.equal(images.length, 2)
    for (const image of images) {
      const style = window.getComputedStyle(image)
      assert.equal(style.width, '135px')
      assert.equal(style.height, '135px')
      assert.equal(style.maxWidth, 'none')
    }
    window.happyDOM.abort()
  })
}

test('the legal information area has no background fill', () => {
  const window = new Window()
  window.document.head.innerHTML = `<style>${styles}</style>`
  window.document.body.innerHTML = template
  const bottom = window.document.querySelector('.lh-footer-bottom')
  assert.ok(bottom)
  assert.equal(window.getComputedStyle(bottom).backgroundColor, 'transparent')
  window.happyDOM.abort()
})

test('at design width, button icons retain explicit sizes instead of shrinking in flex rows', () => {
  const window = new Window({ width: 1440, height: 960 })
  window.document.head.innerHTML = `<style>${styles}</style>`
  window.document.body.innerHTML = template
  for (const [name, size] of [
    ['key', '24px'],
    ['copy', '30px'],
  ]) {
    const slot = window.document.querySelector(`[data-lh-icon="${name}"]`)
    assert.ok(slot)
    const svg = window.document.createElementNS(
      'http://www.w3.org/2000/svg',
      'svg'
    )
    slot.replaceWith(svg)
    const style = window.getComputedStyle(svg)
    assert.equal(style.width, size)
    assert.equal(style.height, size)
    assert.equal(style.flexShrink, '0')
  }
  window.happyDOM.abort()
})

test('at design width, avatars and numbered circles keep the design dimensions', () => {
  const window = new Window({ width: 1440, height: 960 })
  window.document.head.innerHTML = `<style>${styles}</style>`
  window.document.body.innerHTML = template
  for (const [selector, size] of [
    ['.lh-avatar', '50px'],
    ['.lh-step b', '36px'],
    ['.lh-dots i', '14px'],
  ]) {
    for (const element of window.document.querySelectorAll(selector)) {
      const style = window.getComputedStyle(element)
      assert.equal(style.width, size)
      assert.equal(style.height, size)
    }
  }
  window.happyDOM.abort()
})
