import assert from 'node:assert/strict'
import { test } from 'node:test'

import { HOME_MODELS } from '../lib/lanhu-home-models.ts'

test('the provider strip contains all original marquee and ecosystem model names once', () => {
  const legacyNames = [
    'DeepSeek R1',
    'Seedance 2.0',
    'Qwen-Max',
    'GPT-4o',
    'Kimi K2',
    'Claude 3.5',
    'Llama 3',
    'Doubao 1.5 Pro',
    'Qwen3',
    'Kimi K2 Think',
    'GLM-5',
    'Hunyuan Turbo S',
    'Hunyuan Vision',
    'ERNIE 4.5',
    'MiniMax M2.1',
    'Claude Sonnet',
    'GPT-4o / GPT-5',
  ]
  assert.deepEqual(
    HOME_MODELS.map((model) => model.name).sort(),
    legacyNames.sort()
  )
  assert.equal(
    new Set(HOME_MODELS.map((model) => model.name)).size,
    HOME_MODELS.length
  )
})
