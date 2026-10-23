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
import type { PricingModel } from '@/features/pricing/types'

export const HOME_MODEL_CATEGORIES = [
  { id: 'featured', labelKey: 'homepage.models.vendor.featured' },
  { id: 'zhipu', labelKey: 'homepage.models.vendor.zhipu' },
  { id: 'deepseek', labelKey: 'homepage.models.vendor.deepseek' },
  { id: 'minimax', labelKey: 'homepage.models.vendor.minimax' },
  { id: 'moonshot', labelKey: 'homepage.models.vendor.moonshot' },
  { id: 'qwen', labelKey: 'homepage.models.vendor.qwen' },
] as const

export const HOME_MODEL_PROVIDERS = [
  { id: 'zhipu', name: '智谱 AI', icon: 'Zhipu.Color' },
  { id: 'deepseek', name: 'DeepSeek', icon: 'DeepSeek.Color' },
  { id: 'minimax', name: 'MiniMax', icon: 'Minimax.Color' },
  { id: 'moonshot', name: '月之暗面', icon: 'Kimi' },
  { id: 'qwen', name: '阿里通义', icon: 'Qwen.Color' },
] as const satisfies ReadonlyArray<{
  id: Exclude<(typeof HOME_MODEL_CATEGORIES)[number]['id'], 'featured'>
  name: string
  icon: string
}>

export type HomeModelCategory = (typeof HOME_MODEL_CATEGORIES)[number]['id']

type DomesticCategory = Exclude<HomeModelCategory, 'featured'>

const CATEGORY_PATTERNS: ReadonlyArray<{
  id: DomesticCategory
  pattern: RegExp
}> = [
  {
    id: 'zhipu',
    pattern:
      /zhipu|智谱|zai-org|thudm|chatglm|cogview|cogvideo|(?:^|[/._-])glm(?:$|[-._])/i,
  },
  { id: 'deepseek', pattern: /deepseek/i },
  { id: 'minimax', pattern: /minimax|abab|hailuo/i },
  { id: 'moonshot', pattern: /moonshot|月之暗面|kimi(?:$|[-._])/i },
  {
    id: 'qwen',
    pattern:
      /qwen|qwq(?:$|[-._])|qvq(?:$|[-._])|tongyi|阿里|通义|gte(?:$|[-._])/i,
  },
]

export function getHomeModelCategory(
  model: Pick<PricingModel, 'model_name' | 'vendor_name'>
): DomesticCategory | null {
  const identity = `${model.vendor_name ?? ''} ${model.model_name}`
  return (
    CATEGORY_PATTERNS.find(({ pattern }) => pattern.test(identity))?.id ?? null
  )
}

export function getDomesticHomeModels<T extends PricingModel>(
  models: T[]
): T[] {
  return models.filter((model) => getHomeModelCategory(model) !== null)
}
