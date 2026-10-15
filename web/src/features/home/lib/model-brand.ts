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
import type { PricingModel } from '../../pricing/types'

export function getHomepageModelIcon(model: PricingModel): string {
  if (model.vendor_icon) {
    return model.vendor_icon
      .replace(/^openai(?=\.|$)/i, 'OpenAI')
      .replace(/^Kimi\.Color$/i, 'Kimi')
  }
  const name = `${model.vendor_name ?? ''} ${model.model_name}`.toLowerCase()
  const brands = [
    [/claude|anthropic/, 'Anthropic'],
    [/doubao|seedance|seedream|bytedance|字节/, 'Doubao.Color'],
    [/deepseek/, 'DeepSeek.Color'],
    [/qwen|alibaba|阿里/, 'Qwen.Color'],
    [/gemini|google/, 'Gemini.Color'],
    [/kimi|moonshot|月之暗面/, 'Kimi'],
    [/glm|zhipu|智谱/, 'ZAI'],
    [/minimax/, 'Minimax.Color'],
    [/gpt|openai/, 'OpenAI'],
  ] as const
  return (
    brands.find(([pattern]) => pattern.test(name))?.[1] ??
    model.icon?.replace(/^openai(?=\.|$)/i, 'OpenAI') ??
    ''
  )
}
