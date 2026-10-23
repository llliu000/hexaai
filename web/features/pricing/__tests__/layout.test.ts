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
import { describe, test } from 'node:test'

const modelCardSource = readFileSync(
  new URL('../components/model-card.tsx', import.meta.url),
  'utf8'
)
const pricingStyles = readFileSync(
  new URL('../../../styles/pricing-lanhu.css', import.meta.url),
  'utf8'
)
const darkPalette = readFileSync(
  new URL('../../../styles/dark-palette.css', import.meta.url),
  'utf8'
)
const pricingPageSource = readFileSync(
  new URL('../index.tsx', import.meta.url),
  'utf8'
)
const filtersSource = readFileSync(
  new URL('../hooks/use-filters.ts', import.meta.url),
  'utf8'
)

describe('pricing page layout', () => {
  test('matches the homepage header size and scroll contraction', () => {
    assert.match(
      pricingStyles,
      /\.lujiyun-pricing-page\.lujiyun-pricing-page\s*\{[^}]*padding: 75px 0 32px;/s
    )
    assert.match(
      pricingStyles,
      /body:has\(\.lujiyun-pricing-page, \.lujiyun-pricing-loading\) header\.fixed\s*\{[^}]*height: 75px;/s
    )
    assert.match(
      pricingStyles,
      /body:has\(\.lujiyun-pricing-page, \.lujiyun-pricing-loading\)[\s\S]*?header\.fixed[\s\S]*?> \.pointer-events-auto\s*\{[^}]*width: calc\(100% - 64px\);[^}]*max-width: 1360px;[^}]*height: 100%;[^}]*padding: 0;[^}]*transition:\s*width 280ms ease,\s*max-width 280ms ease;/s
    )
    assert.match(
      pricingStyles,
      /@media \(min-width: 601px\)[\s\S]*?> \.pointer-events-auto:has\(> nav\.rounded-2xl\)\s*\{[^}]*width: calc\(100% - 96px\);[^}]*max-width: 1280px;/s
    )
    assert.match(
      pricingStyles,
      /@media \(max-width: 600px\)[\s\S]*?\.lujiyun-pricing-page\.lujiyun-pricing-page\s*\{[^}]*padding-top: 75px;[^}]*\}[\s\S]*?header\.fixed\s*\{[^}]*height: 75px;[^}]*\}[\s\S]*?> \.pointer-events-auto\s*\{[^}]*width: calc\(100% - 40px\);/s
    )
    assert.match(
      pricingStyles,
      /\.lujiyun-pricing-loading\s*\{[^}]*padding: 99px 24px 32px;/s
    )
    assert.match(
      pricingStyles,
      /@media \(max-width: 640px\)[\s\S]*?\.lujiyun-pricing-loading\s*\{[^}]*padding: 91px 12px 24px;/s
    )
  })

  test('matches the homepage system name typography', () => {
    assert.match(
      pricingStyles,
      /header\.fixed\s*nav\s*> a:first-child\s*> span\s*\{[^}]*font-size: 21px;[^}]*line-height: 1\.5;[^}]*font-weight: 700;[^}]*letter-spacing: 0;/s
    )
    assert.match(
      pricingStyles,
      /@media \(max-width: 600px\)[\s\S]*?header\.fixed[\s\S]*?nav\s*> a:first-child\s*> span\s*\{[^}]*font-size: 18px;/s
    )
    assert.match(
      pricingStyles,
      /@media \(max-width: 360px\)[\s\S]*?header\.fixed[\s\S]*?nav\s*> a:first-child\s*> span\s*\{[^}]*font-size: 16px;/s
    )
  })

  test('shows the complete desktop banner at its source aspect ratio', () => {
    assert.match(
      pricingStyles,
      /\.lujiyun-pricing-page\.lujiyun-pricing-page > header\s*\{[^}]*height: auto;[^}]*aspect-ratio: 5760 \/ 788;[^}]*min-height: 0;[^}]*background:[^;}]*model-squal\.png[^;}]*\/[^;}]*100% 100% no-repeat;/s
    )
  })

  test('keeps the model-square banner image and its baked-in lettering in dark mode', () => {
    assert.match(pricingStyles, /background:[^;}]*model-squal\.png/)
    assert.match(
      pricingStyles,
      /\.lujiyun-pricing-page > header > \*\s*\{[^}]*visibility: hidden;/s
    )
    assert.match(
      darkPalette,
      /html\.dark \.lujiyun-pricing-page\.lujiyun-pricing-page > header\s*\{[^}]*background-image: url\('\.\.\/assets\/model-squal-dark\.png'\);/s
    )
    assert.doesNotMatch(
      darkPalette,
      /\.dark \.lujiyun-pricing-page > header > \*\s*\{[^}]*visibility: visible;/s
    )
    assert.doesNotMatch(
      darkPalette,
      /html\.dark \.lujiyun-pricing-page\.lujiyun-pricing-page > header\s*\{[^}]*filter:/s
    )
  })

  test('dark banner asset is a sharp two-times desktop image with the original aspect ratio', () => {
    const darkBanner = readFileSync(
      new URL('../../../assets/model-squal-dark.png', import.meta.url)
    )
    assert.equal(darkBanner.subarray(0, 8).toString('hex'), '89504e470d0a1a0a')
    assert.equal(darkBanner.readUInt32BE(16), 2880)
    assert.equal(darkBanner.readUInt32BE(20), 394)
  })

  test('keeps banner lettering readable on narrow screens by cropping the right side', () => {
    assert.match(
      pricingStyles,
      /@media \(max-width: 900px\)[\s\S]*?\.lujiyun-pricing-page\.lujiyun-pricing-page > header\s*\{[^}]*height: 140px;[^}]*aspect-ratio: auto;[^}]*min-height: 140px;[^}]*background-position: left center;[^}]*background-size: cover;/s
    )
    assert.match(
      pricingStyles,
      /@media \(max-width: 640px\)[\s\S]*?\.lujiyun-pricing-page\.lujiyun-pricing-page > header\s*\{[^}]*height: 120px;[^}]*min-height: 120px;/s
    )
  })

  test('uses the homepage typography scale for model content', () => {
    assert.match(
      pricingStyles,
      /\.lujiyun-pricing-page\.lujiyun-pricing-page\s*\{[^}]*font-family:\s*var\(--font-body\)[^}]*font-size: 14px;[^}]*line-height: 1\.5;/s
    )
    assert.match(
      pricingStyles,
      /\.lanhu-model-card h3\s*\{[^}]*font-family: inherit;[^}]*font-size: 18px;[^}]*line-height: 24px;/s
    )
    assert.match(
      pricingStyles,
      /\.lanhu-model-card > p\s*\{[^}]*font-size: 13px;[^}]*line-height: 20px;/s
    )
  })

  test('highlights model cards on hover', () => {
    assert.match(
      pricingStyles,
      /\.lujiyun-pricing-page \.lanhu-model-grid > \.lanhu-model-card\s*\{[^}]*border: 1px solid var\(--design-border, #e3eaf5\);[^}]*transition: border-color 160ms ease;/s
    )
    assert.match(
      pricingStyles,
      /\.lujiyun-pricing-page \.lanhu-model-grid > \.lanhu-model-card:hover\s*\{[^}]*border-color: var\(--design-accent-border, #8db2ff\);/s
    )
  })

  test('makes the whole model card clickable and keyboard accessible', () => {
    assert.match(modelCardSource, /cursor-pointer flex-col/)
    assert.match(modelCardSource, /role='button'/)
    assert.match(modelCardSource, /tabIndex=\{0\}/)
    assert.match(modelCardSource, /onClick=\{props\.onClick\}/)
    assert.match(
      modelCardSource,
      /if \(e\.target !== e\.currentTarget\) return/
    )
    assert.match(
      modelCardSource,
      /if \(e\.key === 'Enter' \|\| e\.key === ' '\)/
    )
    assert.match(modelCardSource, /e\.stopPropagation\(\)/)
  })

  test('clears search and sidebar filters from the pricing toolbar', () => {
    assert.match(
      filtersSource,
      /const hasActiveFilters = useMemo\(\s*\(\) =>\s*Boolean\(searchInput\)/s
    )
    assert.match(
      pricingPageSource,
      /<PricingToolbar[\s\S]*?hasActiveFilters=\{hasActiveFilters\}[\s\S]*?onClearFilters=\{handleClearAll\}/s
    )
    assert.match(
      pricingPageSource,
      /<PricingSidebar[\s\S]*?hasActiveFilters=\{hasActiveFilters\}[\s\S]*?onClearFilters=\{handleClearAll\}/s
    )
  })
})
