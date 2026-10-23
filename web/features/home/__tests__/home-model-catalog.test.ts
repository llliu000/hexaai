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

import type { PricingModel } from '../../pricing/types.ts'
import {
  getDomesticHomeModels,
  getHomeModelCategory,
  HOME_MODEL_CATEGORIES,
  HOME_MODEL_PROVIDERS,
} from '../lib/home-model-catalog.ts'

const model = (model_name: string, vendor_name: string): PricingModel => ({
  id: 1,
  model_name,
  vendor_name,
  quota_type: 0,
  model_ratio: 1,
  completion_ratio: 1,
  enable_groups: [],
})

test('homepage model categories follow the domestic vendor order in the design', () => {
  assert.deepEqual(
    HOME_MODEL_CATEGORIES.map(({ id }) => id),
    ['featured', 'zhipu', 'deepseek', 'minimax', 'moonshot', 'qwen']
  )
})

test('homepage ecosystem displays only the five supported domestic providers', () => {
  assert.deepEqual(
    HOME_MODEL_PROVIDERS.map(({ id }) => id),
    ['zhipu', 'deepseek', 'minimax', 'moonshot', 'qwen']
  )
  assert.doesNotMatch(
    HOME_MODEL_PROVIDERS.map(({ name, icon }) => `${name} ${icon}`).join(' '),
    /OpenAI|Anthropic|Google|Gemini|Claude/i
  )
})

test('homepage catalog recognizes the five supported domestic model families', () => {
  assert.equal(getHomeModelCategory(model('glm-5', '智谱')), 'zhipu')
  assert.equal(getHomeModelCategory(model('deepseek-v4', 'OpenAI')), 'deepseek')
  assert.equal(
    getHomeModelCategory(model('abab6.5-chat', 'MiniMax')),
    'minimax'
  )
  assert.equal(getHomeModelCategory(model('kimi-k3', '月之暗面')), 'moonshot')
  assert.equal(getHomeModelCategory(model('qwen3.7-plus', '阿里百炼')), 'qwen')
})

test('homepage catalog excludes overseas and unlisted domestic vendors', () => {
  const models = [
    model('gpt-5', 'OpenAI'),
    model('claude-sonnet-4', 'Anthropic'),
    model('gemini-3-pro', 'Google'),
    model('doubao-seed-2', '字节跳动'),
    model('glm-5', '智谱'),
  ]
  assert.deepEqual(
    getDomesticHomeModels(models).map(({ model_name }) => model_name),
    ['glm-5']
  )
})
