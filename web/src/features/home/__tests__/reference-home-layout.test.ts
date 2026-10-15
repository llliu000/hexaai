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

import { Window, type Element as HappyDOMElement } from 'happy-dom'
import { parse } from 'postcss'

const css = readFileSync(
  new URL('../../../styles/reference-home.css', import.meta.url),
  'utf8'
)
const headerSource = readFileSync(
  new URL('../components/reference-header.tsx', import.meta.url),
  'utf8'
)
const heroSource = readFileSync(
  new URL('../components/hero-carousel.tsx', import.meta.url),
  'utf8'
)
const modelsSource = readFileSync(
  new URL('../components/reference-models.tsx', import.meta.url),
  'utf8'
)
const homeSource = readFileSync(
  new URL('../components/reference-home.tsx', import.meta.url),
  'utf8'
)
const footerSource = readFileSync(
  new URL('../../../components/layout/components/footer.tsx', import.meta.url),
  'utf8'
)
const integrationSource = readFileSync(
  new URL('../components/reference-integration.tsx', import.meta.url),
  'utf8'
)
const parallaxSource = readFileSync(
  new URL('../components/home-parallax.tsx', import.meta.url),
  'utf8'
)
const parallaxHookSource = readFileSync(
  new URL('../hooks/use-home-parallax.ts', import.meta.url),
  'utf8'
)

test('desktop and mobile headers expose the same four primary destinations', () => {
  const desktopNav = headerSource.match(
    /<nav className='wm-desktop-nav'[\s\S]*?<\/nav>/
  )?.[0]
  const mobileNav = headerSource.match(
    /<DropdownMenuContent align='end'>[\s\S]*?<\/DropdownMenuContent>/
  )?.[0]
  assert.ok(desktopNav)
  assert.ok(mobileNav)

  for (const nav of [desktopNav, mobileNav]) {
    assert.match(nav, /to='\/'[\s\S]*t\('Home'\)/)
    assert.match(nav, /to='\/dashboard'[\s\S]*t\('Console'\)/)
    assert.match(nav, /to='\/pricing'[\s\S]*t\('Model Square'\)/)
    assert.match(nav, /href=\{props\.docsUrl\}[\s\S]*t\('Docs'\)/)
    assert.doesNotMatch(nav, /homepage\.nav\.(products|solutions)/)
  }

  const headerActions = headerSource.match(
    /<div className='wm-header-actions'>[\s\S]*?<DropdownMenu>/
  )?.[0]
  assert.ok(headerActions)
  assert.doesNotMatch(headerActions, /to='\/dashboard'/)
  assert.doesNotMatch(headerActions, /wm-console-link/)
})

test('the floating header gently draws its content inward on wider screens', () => {
  const root = parse(css)
  const properties = (selector: string, media?: string) => {
    const values: Record<string, string> = {}
    root.walkRules(selector, (rule) => {
      if (
        media
          ? rule.parent?.type !== 'atrule' || rule.parent.params !== media
          : rule.parent !== root
      ) {
        return
      }
      rule.walkDecls((declaration) => {
        values[declaration.prop] = declaration.value
      })
    })
    return values
  }
  const header = properties('.wm-header')
  assert.equal(header.height, '75px')
  const mobileHeader = properties('.wm-header', '(max-width: 600px)')
  assert.equal(mobileHeader.height, '75px')
  const initial = properties('.wm-header-inner')
  assert.equal(initial['max-width'], '1360px')
  assert.match(
    initial.transition ?? '',
    /width 280ms ease,\s*max-width 280ms ease/
  )
  const scrolled = properties(
    '.wm-header.is-scrolled .wm-header-inner',
    '(min-width: 601px)'
  )
  assert.equal(scrolled.width, 'calc(100% - 96px)')
  assert.equal(scrolled['max-width'], '1280px')
})

test('homepage model catalog selects featured by default and exposes its pressed state', () => {
  assert.match(modelsSource, /useState<HomeModelCategory>\('featured'\)/)
  assert.match(modelsSource, /aria-pressed=\{category === item\.id\}/)

  const root = parse(css)
  let selectedBackground = ''
  root.walkRules(
    ".wm-vendor-tabs button:is(.is-active, [aria-pressed='true'])",
    (rule) => {
      rule.walkDecls('background', (declaration) => {
        selectedBackground = declaration.value
      })
    }
  )
  assert.equal(selectedBackground, '#111827')
})

test('homepage model catalog uses a fixed view-all label without a model count', () => {
  assert.match(modelsSource, /t\('homepage\.models\.explore'\)/)
  assert.doesNotMatch(modelsSource, /homepage\.models\.allWithCount/)
})

test('homepage footer adds one WeChat customer service QR code without replacing existing content', () => {
  assert.match(footerSource, /extraContent\?: ReactNode/)
  assert.match(homeSource, /extraContent=\{/)
  assert.equal(
    (homeSource.match(/src='\/lanhu\/home-service-qr\.png'/g) ?? []).length,
    1
  )
  assert.match(homeSource, /t\('homepage\.footer\.wechatSupport'\)/)
  assert.match(homeSource, /width=\{120\}/)
  assert.match(homeSource, /height=\{120\}/)
  assert.match(
    css,
    /\.wm-footer \.site-footer-main\.site-footer-main-has-extra\s*\{[^}]*grid-template-columns:\s*minmax\(220px, 1fr\) minmax\(0, 2fr\) 132px;/s
  )
  assert.match(
    css,
    /\.wm-footer-wechat\s*\{[^}]*align-items:\s*center;[^}]*text-align:\s*center;/s
  )
})

test('hero slides reserve a stable full-width surface and clip transitions', () => {
  const window = new Window(),
    sheet = new window.CSSStyleSheet()
  sheet.replaceSync(css)
  const rules = [...sheet.cssRules].filter(
    (rule) => rule instanceof window.CSSStyleRule
  )
  const slide = rules.find((rule) => rule.selectorText === '.wm-hero-slide')
  const viewport = rules.find(
    (rule) => rule.selectorText === '.wm-hero-viewport'
  )
  assert.equal(slide?.style.minHeight, '602px')
  assert.equal(slide?.style.minWidth, '0')
  assert.equal(viewport?.style.overflow, 'hidden')
})

test('homepage hero keeps autoplay running on hover without a pause or play control', () => {
  assert.match(heroSource, /embla-carousel-autoplay/)
  assert.match(heroSource, /stopOnMouseEnter: false/)
  assert.doesNotMatch(heroSource, /\bhovered\b/)
  assert.doesNotMatch(heroSource, /onPointerEnter=/)
  assert.doesNotMatch(heroSource, /onPointerLeave=/)
  assert.match(heroSource, /className='wm-hero-dots'/)
  assert.doesNotMatch(heroSource, /wm-hero-pause/)
  assert.doesNotMatch(heroSource, /\b(Pause|Play)\b/)
})

test('every homepage content module retains foreground and background parallax layers', () => {
  for (const source of [homeSource, modelsSource, integrationSource]) {
    assert.match(source, /wm-home-parallax-section/)
    assert.match(source, /HomeParallaxDecorations/)
  }
  assert.match(
    parallaxSource,
    /wm-home-parallax-backdrop-slow/,
    'slow background'
  )
  assert.match(
    parallaxHookSource,
    /wm-home-parallax-panel/,
    'foreground panels'
  )
  assert.match(parallaxHookSource, /reducedMotion/)
  assert.doesNotMatch(parallaxHookSource, /wm-hero-depth-backdrop/)
})

test('homepage hero retains composited parallax layers and reduced-motion support', () => {
  assert.match(heroSource, /props\.preview \|\| reducedMotion/)

  const root = parse(css)
  const art = root.nodes.find(
    (node) => node.type === 'rule' && node.selector === '.wm-hero-parallax-art'
  )
  assert.ok(art)
  assert.match(art.toString(), /will-change: transform/)
})

test('phone layouts stack feature panels and keep model cards horizontally scrollable', () => {
  const window = new Window(),
    sheet = new window.CSSStyleSheet()
  sheet.replaceSync(css)
  const mobile = [...sheet.cssRules].find(
    (rule) =>
      rule instanceof window.CSSMediaRule &&
      rule.conditionText === '(max-width: 600px)'
  )
  assert.ok(mobile instanceof window.CSSMediaRule)
  const grid = [...mobile.cssRules].find(
    (rule) =>
      rule instanceof window.CSSStyleRule &&
      rule.selectorText === '.wm-four-grid'
  )
  assert.ok(grid instanceof window.CSSStyleRule)
  assert.equal(grid.style.gridTemplateColumns, 'minmax(0, 1fr)')
})

test('new homepage labels have Chinese translations and matching interpolation arguments', () => {
  const dir = new URL('../../../i18n/locales/', import.meta.url)
  const en = JSON.parse(
    readFileSync(new URL('en.json', dir), 'utf8')
  ).translation
  const zh = JSON.parse(
    readFileSync(new URL('zh.json', dir), 'utf8')
  ).translation
  const keys = Object.keys(en).filter((key) => key.startsWith('homepage.'))
  assert.ok(keys.length > 50)
  for (const key of keys) {
    assert.equal(typeof zh[key], 'string', key)
    assert.deepEqual(
      zh[key].match(/\{\{\w+\}\}/g),
      en[key].match(/\{\{\w+\}\}/g),
      key
    )
  }
})

test('hero color fades vertically into the shared body surface instead of ending in a solid tint', () => {
  const root = parse(css)
  const declarations: Record<string, Record<string, string>> = {}
  root.walkRules((rule) => {
    if (rule.parent?.type !== 'root') return
    declarations[rule.selector] = Object.fromEntries(
      rule.nodes
        .filter((node) => node.type === 'decl')
        .map((node) => [node.prop, node.value])
    )
  })
  assert.equal(declarations['.wm-home']['--wm-surface'], '#fff')
  assert.equal(declarations['.wm-home'].background, 'var(--wm-surface)')
  assert.match(
    declarations['.wm-hero-slide'].background,
    /linear-gradient\(\s*180deg/
  )
  assert.match(
    declarations['.wm-hero-slide'].background,
    /var\(--wm-surface\) 100%/
  )
  assert.match(declarations['.wm-hero-art']['mask-image'], /transparent/)
  assert.equal(declarations['.dark .wm-home']['--wm-surface'], '#151d2a')
})

test('hero decorations share the highlighted text color and cannot intercept controls', () => {
  const root = parse(css)
  const read = (selector: string, property: string): string => {
    let value = ''
    root.walkRules(selector, (rule) => {
      if (rule.parent?.type === 'root') {
        rule.walkDecls(property, (declaration) => {
          value = declaration.value
        })
      }
    })
    return value
  }
  assert.equal(
    read('.wm-hero-decoration', 'color'),
    read('.wm-hero-copy strong', 'color')
  )
  assert.equal(read('.wm-hero-decoration', 'pointer-events'), 'none')
  assert.equal(read('.wm-hero-decoration', 'user-select'), 'none')
  assert.ok(
    Number(read('.wm-hero-decoration', 'z-index')) <
      Number(read('.wm-hero-copy', 'z-index'))
  )
})

test('desktop catalog uses wider compact cards and reserves a full title area', () => {
  const window = new Window(),
    sheet = new window.CSSStyleSheet()
  sheet.replaceSync(css)
  const rules = [...sheet.cssRules].filter(
    (rule) => rule instanceof window.CSSStyleRule
  )
  const container = rules.find((rule) => rule.selectorText === '.wm-container')
  const card = rules.find((rule) => rule.selectorText === '.wm-model-card')
  const heading = rules.find(
    (rule) => rule.selectorText === '.wm-model-card h3'
  )
  assert.equal(container?.style.maxWidth, '1360px')
  assert.equal(card?.style.flexBasis, '420px')
  assert.equal(card?.style.minHeight, '190px')
  assert.equal(heading?.style.whiteSpace, 'normal')
})

test('model catalog exposes horizontal scrolling when cards exceed its width', () => {
  const window = new Window(),
    sheet = new window.CSSStyleSheet()
  sheet.replaceSync(css)
  const rules = [...sheet.cssRules].filter(
    (rule) => rule instanceof window.CSSStyleRule
  )
  const track = rules.find((rule) => rule.selectorText === '.wm-model-track')
  assert.equal(track?.style.overflowX, 'auto')
  assert.equal(track?.style.minWidth, '0')
  assert.equal(track?.style.touchAction, 'pan-x')

  let scrollbarWidth = ''
  parse(css).walkRules('.wm-model-track', (rule) => {
    rule.walkDecls('scrollbar-width', (declaration) => {
      scrollbarWidth = declaration.value
    })
  })
  assert.equal(scrollbarWidth, 'none')
})

test('overflowing model catalogs automatically advance without interrupting interaction', () => {
  assert.match(modelsSource, /getLoopedModelScrollPosition/)
  assert.match(modelsSource, /window\.setInterval\(/)
  assert.match(modelsSource, /!loopEnabled/)
  assert.match(modelsSource, /prefers-reduced-motion: reduce/)
  assert.match(modelsSource, /copy: 'leading'/)
  assert.match(modelsSource, /copy: 'trailing'/)
  assert.match(modelsSource, /data-model-copy=\{item\.copy\}/)
  assert.match(modelsSource, /aria-hidden=\{item\.copy !== 'original'/)
  assert.match(
    modelsSource,
    /onMouseEnter=\{\(\) => setAutoScrollPaused\(true\)\}/
  )
  assert.match(
    modelsSource,
    /onTouchStart=\{\(\) => setAutoScrollPaused\(true\)\}/
  )
})

test('customer testimonials use a responsive three-column card layout', () => {
  const window = new Window(),
    sheet = new window.CSSStyleSheet()
  sheet.replaceSync(css)
  const rules = [...sheet.cssRules].filter(
    (rule) => rule instanceof window.CSSStyleRule
  )
  const grid = rules.find(
    (rule) => rule.selectorText === '.wm-testimonial-grid'
  )
  const card = rules.find((rule) => rule.selectorText === '.wm-testimonial')
  assert.equal(grid?.style.gridTemplateColumns, 'repeat(3, minmax(0, 1fr))')
  assert.equal(card?.style.flexDirection, 'column')
  assert.equal(card?.style.minHeight, '236px')
})

test('fixed navigation leaves enough space above anchored content', () => {
  const window = new Window(),
    sheet = new window.CSSStyleSheet()
  sheet.replaceSync(css)
  const rules = [...sheet.cssRules].filter(
    (rule) => rule instanceof window.CSSStyleRule
  )
  const header = rules.find((rule) => rule.selectorText === '.wm-header')
  const sections = rules.find(
    (rule) => rule.selectorText === '.wm-home section[id]'
  )
  assert.equal(header?.style.position, 'fixed')
  assert.equal(sections?.style.scrollMarginTop, '104px')
})

for (const width of [390, 1440]) {
  test(`testimonial cards keep the reference spacing and avatar size at ${width}px`, async () => {
    const window = new Window({ width })
    try {
      window.document.head.innerHTML = `<style>${css}</style>`
      window.document.body.innerHTML = `
        <main class="wm-home">
          <article class="wm-testimonial">
            <blockquote>
              <svg class="wm-testimonial-quote wm-testimonial-quote-open"></svg>
              <p>Customer feedback</p>
              <svg class="wm-testimonial-quote wm-testimonial-quote-close"></svg>
            </blockquote>
            <div class="wm-testimonial-person"><img alt="" /><div><strong>Customer</strong><span>Role</span></div></div>
          </article>
        </main>
      `
      const card = window.document.querySelector('.wm-testimonial')
      const avatar = window.document.querySelector('img')
      const quote = window.document.querySelector('.wm-testimonial-quote-close')
      assert.ok(card)
      assert.ok(avatar)
      assert.ok(quote)
      const cardStyle = window.getComputedStyle(card)
      assert.equal(cardStyle.paddingTop, '32px')
      assert.equal(cardStyle.borderRadius, '16px')
      assert.equal(cardStyle.minHeight, '236px')
      assert.notEqual(cardStyle.boxShadow, 'none')
      assert.equal(window.getComputedStyle(avatar).width, '40px')
      assert.equal(window.getComputedStyle(avatar).height, '40px')
      assert.equal(window.getComputedStyle(quote).alignSelf, 'flex-end')
      for (const [selector, fontSize, lineHeight] of [
        ['blockquote', '14px', '24px'],
        ['strong', '15px', '22px'],
        ['.wm-testimonial-person span', '13px', '20px'],
      ]) {
        const text: HappyDOMElement | null = card.querySelector(selector)
        assert.ok(text)
        const textStyle = window.getComputedStyle(text)
        assert.equal(textStyle.fontSize, fontSize)
        assert.equal(textStyle.lineHeight, lineHeight)
      }
    } finally {
      await window.happyDOM.close()
    }
  })
}
