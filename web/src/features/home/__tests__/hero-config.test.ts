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
import { test } from 'node:test'

import {
  createDefaultHeroConfig,
  duplicateHeroSlide,
  getHeroText,
  getVisibleHeroSlides,
  heroConfigSchema,
  isSafeHeroUrl,
  parseHeroConfig,
} from '../lib/hero-config.ts'
import { getHomepageModelIcon } from '../lib/model-brand.ts'
import { getSdkBaseUrl, getSdkExample } from '../lib/sdk-example.ts'

test('default slides validate and initialize independent drafts', () => {
  const first = createDefaultHeroConfig(),
    second = createDefaultHeroConfig()
  assert.equal(heroConfigSchema.safeParse(first).success, true)
  first.slides[0].content.zh.titlePrefix = 'Changed'
  assert.notEqual(
    first.slides[0].content.zh.titlePrefix,
    second.slides[0].content.zh.titlePrefix
  )
  assert.equal(first.slides.length, 4)
})

test('model cards resolve provider logos when the catalog has no icon metadata', () => {
  const model = {
    id: 1,
    model_name: 'claude-sonnet-4-5',
    vendor_name: 'Anthropic',
    quota_type: 0,
    model_ratio: 1,
    completion_ratio: 1,
    enable_groups: [],
  }
  assert.equal(getHomepageModelIcon(model), 'Anthropic')
  assert.equal(
    getHomepageModelIcon({
      ...model,
      model_name: 'deepseek-chat',
      vendor_name: 'DeepSeek',
      icon: 'Openai',
    }),
    'DeepSeek.Color'
  )
  assert.equal(
    getHomepageModelIcon({ ...model, vendor_icon: 'Openai' }),
    'OpenAI'
  )
})

test('Kimi logos use the theme-aware variant so their white artwork cannot disappear on light cards', () => {
  const model = {
    id: 1,
    model_name: 'kimi-k2',
    vendor_name: 'Moonshot',
    quota_type: 0,
    model_ratio: 1,
    completion_ratio: 1,
    enable_groups: [],
  }
  assert.equal(getHomepageModelIcon(model), 'Kimi')
  assert.equal(
    getHomepageModelIcon({ ...model, vendor_icon: 'Kimi.Color' }),
    'Kimi'
  )
})

test('disabled or empty configurations fall back to one usable slide', () => {
  const config = createDefaultHeroConfig()
  config.enabled = false
  assert.equal(getVisibleHeroSlides(config).length, 1)
  config.enabled = true
  config.slides = []
  assert.equal(getVisibleHeroSlides(config).length, 1)
  config.slides = createDefaultHeroConfig().slides.map((slide) => ({
    ...slide,
    enabled: false,
  }))
  assert.equal(getVisibleHeroSlides(config).length, 1)
})

test('language selection resolves regional/default fallbacks and the configured brand', () => {
  const slide = createDefaultHeroConfig().slides[0]
  assert.ok(
    getHeroText(slide, 'zh-CN', 'zh', 'Test Brand').titlePrefix.includes(
      'Test Brand'
    )
  )
  assert.equal(
    getHeroText(slide, 'fr', 'en', 'Brand').highlight,
    slide.content.en.highlight
  )
})

test('unsafe links and image URLs cannot be published', () => {
  for (const value of [
    'javascript:alert(1)',
    '//evil.test',
    '/\\evil.test',
    'https://user:pass@evil.test',
    '/unsafe\npath',
    'data:image/png;base64,123',
  ]) {
    assert.equal(isSafeHeroUrl(value), false, value)
  }
  for (const value of [
    '/keys',
    '/pricing?view=card',
    'https://docs.example.com',
  ]) {
    assert.equal(isSafeHeroUrl(value), true, value)
  }
  const config = createDefaultHeroConfig()
  config.slides[0].desktopImage = 'javascript:alert(1)'
  assert.equal(parseHeroConfig(config), null)
})

test('duplicate slide IDs and missing default translations fail validation', () => {
  const config = createDefaultHeroConfig()
  config.slides[1].id = config.slides[0].id
  assert.equal(parseHeroConfig(config), null)
  const translated = createDefaultHeroConfig()
  translated.defaultLocale = 'vi'
  assert.equal(parseHeroConfig(translated), null)
  assert.equal(parseHeroConfig('{bad json'), null)
})

test('copying a slide keeps the source untouched and starts disabled', () => {
  const slide = createDefaultHeroConfig().slides[0],
    copy = duplicateHeroSlide(slide, 'new-id')
  assert.equal(copy.id, 'new-id')
  assert.equal(copy.enabled, false)
  copy.content.zh.highlight = 'Edited'
  assert.notEqual(copy.content.zh.highlight, slide.content.zh.highlight)
})

test('SDK examples normalize base URLs and use environment credentials', () => {
  assert.equal(
    getSdkBaseUrl('https://api.example.com/v1/', 'https://fallback.test'),
    'https://api.example.com/v1'
  )
  assert.equal(
    getSdkBaseUrl('javascript:bad', 'https://fallback.test'),
    'https://fallback.test/v1'
  )
  for (const language of ['python', 'javascript', 'curl'] as const) {
    const code = getSdkExample(
      language,
      'https://api.example.com/v1',
      'model"name'
    )
    assert.ok(code.includes('AI_API_KEY'))
    assert.ok(code.includes('model\\"name'))
    assert.ok(code.includes('https://api.example.com/v1'))
  }
})
