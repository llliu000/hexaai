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
import type { PricingModel } from '../types'

type VendorIdentity = Pick<
  PricingModel,
  'vendor_name' | 'vendor_icon' | 'model_name'
>

const vendorIcons = new Map<string, string>([
  ['anthropic', 'Anthropic'],
  ['openai', 'OpenAI'],
  ['bytedance', 'ByteDance.Color'],
  ['字节跳动', 'ByteDance.Color'],
  ['字节', 'ByteDance.Color'],
  ['google', 'Google.Color'],
  ['googleai', 'Google.Color'],
  ['deepseek', 'DeepSeek.Color'],
  ['深度求索', 'DeepSeek.Color'],
  ['zhipu', 'Zhipu.Color'],
  ['zhipuai', 'Zhipu.Color'],
  ['智谱', 'Zhipu.Color'],
  ['智谱ai', 'Zhipu.Color'],
  ['zai', 'ZAI'],
  ['阿里百炼', 'Bailian.Color'],
  ['阿里云百炼', 'Bailian.Color'],
  ['bailian', 'Bailian.Color'],
  ['alibabacloudbailian', 'Bailian.Color'],
  ['alibaba', 'Alibaba.Color'],
  ['阿里', 'Alibaba.Color'],
  ['阿里云', 'Alibaba.Color'],
  ['moonshot', 'Moonshot'],
  ['moonshotai', 'Moonshot'],
  ['月之暗面', 'Moonshot'],
  ['minimax', 'Minimax.Color'],
  ['xai', 'XAI'],
  ['tencent', 'Tencent.Color'],
  ['腾讯', 'Tencent.Color'],
  ['baidu', 'Baidu.Color'],
  ['百度', 'Baidu.Color'],
  ['meta', 'Meta.Color'],
  ['mistral', 'Mistral.Color'],
  ['mistralai', 'Mistral.Color'],
])

const canonicalIconNames = new Map(
  [
    'Anthropic',
    'OpenAI',
    'ByteDance',
    'Google',
    'DeepSeek',
    'Zhipu',
    'ZAI',
    'Bailian',
    'Alibaba',
    'Moonshot',
    'Kimi',
    'Minimax',
    'XAI',
    'Tencent',
    'Baidu',
    'Meta',
    'Mistral',
  ].map((name) => [name.toLowerCase(), name])
)

const modelFamilies: ReadonlyArray<readonly [RegExp, string]> = [
  [/(?:^|\/)claude(?:[-_.]|$)/i, 'Anthropic'],
  [/(?:^|\/)(?:gpt|chatgpt|o[1-4])(?:[-_.]|$)/i, 'OpenAI'],
  [/(?:^|\/)(?:doubao|seedance|seedream)(?:[-_.]|$)/i, 'ByteDance.Color'],
  [/(?:^|\/)gemini(?:[-_.]|$)/i, 'Google.Color'],
  [/(?:^|\/)deepseek(?:[-_.]|$)/i, 'DeepSeek.Color'],
  [/(?:^|\/)(?:glm|chatglm)(?:[-_.]|$)/i, 'Zhipu.Color'],
  [/(?:^|\/)(?:qwen|qwq|qvq)(?:\d|[-_.]|$)/i, 'Alibaba.Color'],
  [/(?:^|\/)(?:kimi|moonshot)(?:[-_.]|$)/i, 'Moonshot'],
  [/(?:^|\/)(?:minimax|hailuo)(?:[-_.]|$)/i, 'Minimax.Color'],
  [/(?:^|\/)grok(?:[-_.]|$)/i, 'XAI'],
]

export function getPricingVendorIcon(
  model: VendorIdentity
): string | undefined {
  const configuredIcon = model.vendor_icon?.trim()
  if (configuredIcon) {
    const [base, ...path] = configuredIcon.split('.')
    const normalized = [
      canonicalIconNames.get(base.toLowerCase()) ?? base,
      ...path,
    ].join('.')
    return normalized.toLowerCase() === 'kimi.color' ? 'Kimi' : normalized
  }

  const vendorName = model.vendor_name?.trim()
  if (vendorName) {
    // An explicit vendor must not be replaced by the model family's brand.
    return vendorIcons.get(vendorName.toLowerCase().replaceAll(/[\s._-]/g, ''))
  }

  return modelFamilies.find(([pattern]) => pattern.test(model.model_name))?.[1]
}
