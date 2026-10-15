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

import { getPricingVendorIcon } from '../lib/model-vendor-icon.ts'

const model = { model_name: 'claude-sonnet', vendor_name: 'Anthropic' }

test('catalog vendors resolve to their own logos when no icon is configured', () => {
  const vendors = [
    ['Anthropic', 'Anthropic'],
    ['字节跳动', 'ByteDance.Color'],
    ['Google', 'Google.Color'],
    ['OpenAI', 'OpenAI'],
    ['DeepSeek', 'DeepSeek.Color'],
    ['智谱', 'Zhipu.Color'],
    ['阿里百炼', 'Bailian.Color'],
    ['Moonshot', 'Moonshot'],
    ['月之暗面', 'Moonshot'],
  ] as const
  for (const [vendor_name, expected] of vendors) {
    assert.equal(getPricingVendorIcon({ ...model, vendor_name }), expected)
  }
})

test('configured vendor logos take precedence and normalize known icon casing', () => {
  assert.equal(
    getPricingVendorIcon({ ...model, vendor_icon: ' Openai ' }),
    'OpenAI'
  )
  assert.equal(
    getPricingVendorIcon({ ...model, vendor_icon: 'MiniMax.Color' }),
    'Minimax.Color'
  )
  assert.equal(
    getPricingVendorIcon({ ...model, vendor_icon: 'Kimi.Color' }),
    'Kimi'
  )
  assert.equal(
    getPricingVendorIcon({ ...model, vendor_icon: 'CustomProvider.Avatar' }),
    'CustomProvider.Avatar'
  )
})

test('blank icon configuration falls back to the actual vendor, not the model family', () => {
  assert.equal(
    getPricingVendorIcon({
      ...model,
      vendor_name: ' Google ',
      vendor_icon: '  ',
    }),
    'Google.Color'
  )
  assert.equal(
    getPricingVendorIcon({ ...model, vendor_name: 'Unknown Host' }),
    undefined
  )
})

test('model families provide a logo only when vendor metadata is absent', () => {
  assert.equal(
    getPricingVendorIcon({ model_name: 'claude-sonnet' }),
    'Anthropic'
  )
  assert.equal(
    getPricingVendorIcon({ model_name: 'openai/gpt-4.1', vendor_name: '' }),
    'OpenAI'
  )
  assert.equal(getPricingVendorIcon({ model_name: 'custom-model' }), undefined)
})
