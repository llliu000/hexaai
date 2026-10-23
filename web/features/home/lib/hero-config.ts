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
import { z } from 'zod'

export const HERO_OPTION_KEY = 'HomeHeroCarousel'
export const HERO_QUERY_KEY = ['home-hero'] as const
export const HERO_THEMES = ['violet', 'blue', 'teal', 'rose'] as const
export const HERO_ICONS = [
  'platform',
  'openai',
  'deepseek',
  'qwen',
  'zhipu',
  'moonshot',
  'custom',
] as const

export function isSafeHeroUrl(value: string, optional = false): boolean {
  if (!value) return optional
  if (
    value !== value.trim() ||
    value.includes('\\') ||
    [...value].some((character) => character.charCodeAt(0) < 32)
  ) {
    return false
  }
  if (value.startsWith('/')) return !value.startsWith('//')
  try {
    const url = new URL(value)
    return (
      ['http:', 'https:'].includes(url.protocol) &&
      Boolean(url.hostname) &&
      !url.username &&
      !url.password
    )
  } catch {
    return false
  }
}

const optionalUrl = z
  .string()
  .max(2048)
  .refine((value) => isSafeHeroUrl(value, true), 'homepage.invalidUrl')
const actionSchema = z.object({
  url: z
    .string()
    .max(2048)
    .refine((value) => isSafeHeroUrl(value), 'homepage.invalidUrl'),
  newTab: z.boolean(),
})
export const heroTextSchema = z
  .object({
    badge: z.string().max(80),
    titlePrefix: z.string().max(160),
    highlight: z.string().max(160),
    titleSuffix: z.string().max(160),
    description: z.string().max(500),
    primaryLabel: z.string().trim().min(1).max(40),
    secondaryLabel: z.string().max(40),
  })
  .refine(
    (value) =>
      (value.titlePrefix + value.highlight + value.titleSuffix).trim().length >
        0 &&
      (value.titlePrefix + value.highlight + value.titleSuffix).length <= 160,
    'homepage.invalidTitle'
  )

export const heroSlideSchema = z.object({
  id: z.string().regex(/^[a-zA-Z0-9_-]{1,64}$/),
  internalName: z.string().max(80).optional(),
  enabled: z.boolean(),
  theme: z.enum(HERO_THEMES),
  icon: z.enum(HERO_ICONS),
  logoUrl: optionalUrl,
  desktopImage: optionalUrl,
  mobileImage: optionalUrl,
  imagePosition: z.enum(['center', 'left', 'right']),
  content: z
    .record(z.string().regex(/^[a-zA-Z0-9_-]{1,64}$/), heroTextSchema)
    .refine(
      (value) =>
        Object.keys(value).length > 0 && Object.keys(value).length <= 10
    ),
  primaryAction: actionSchema,
  secondaryAction: actionSchema.optional(),
})

export const heroConfigSchema = z
  .object({
    schemaVersion: z.literal(1),
    enabled: z.boolean(),
    autoplay: z.boolean(),
    loop: z.boolean(),
    intervalMs: z.number().int().min(3000).max(15000),
    defaultLocale: z.string().regex(/^[a-zA-Z0-9_-]{1,64}$/),
    slides: z.array(heroSlideSchema).max(12),
  })
  .superRefine((value, ctx) => {
    const ids = new Set<string>()
    value.slides.forEach((slide, index) => {
      if (ids.has(slide.id)) {
        ctx.addIssue({
          code: 'custom',
          path: ['slides', index, 'id'],
          message: 'homepage.duplicateId',
        })
      }
      ids.add(slide.id)
      if (!slide.content[value.defaultLocale]) {
        ctx.addIssue({
          code: 'custom',
          path: ['slides', index, 'content'],
          message: 'homepage.missingLocale',
        })
      }
      if (slide.secondaryAction) {
        for (const [locale, text] of Object.entries(slide.content)) {
          if (!text.secondaryLabel.trim()) {
            ctx.addIssue({
              code: 'custom',
              path: ['slides', index, 'content', locale, 'secondaryLabel'],
              message: 'Required',
            })
          }
        }
      }
    })
  })

export type HeroText = z.infer<typeof heroTextSchema>
export type HeroSlide = z.infer<typeof heroSlideSchema>
export type HeroConfig = z.infer<typeof heroConfigSchema>

const slides: Array<{
  id: string
  icon: HeroSlide['icon']
  theme: HeroSlide['theme']
  zh: string[]
  en: string[]
}> = [
  {
    id: 'platform',
    icon: 'platform',
    theme: 'violet',
    zh: [
      '统一 AI 接入平台',
      '{{brand}} · ',
      '一个 API，多种可能',
      '文本对话、图像生成、视频创作与语音能力，从一个统一接口开始。',
    ],
    en: [
      'UNIFIED AI PLATFORM',
      '{{brand}} · ',
      'One API. More possibilities.',
      'Connect text, image, video and audio models through one familiar API.',
    ],
  },
  {
    id: 'reasoning',
    icon: 'deepseek',
    theme: 'blue',
    zh: [
      '对话 · 推理 · 代码',
      '{{brand}} · ',
      '对话与推理 API',
      '在熟悉的接口中选择适合任务的模型，构建你的 AI 助手与开发工作流。',
    ],
    en: [
      'CHAT · REASONING · CODE',
      '{{brand}} · ',
      'Chat and reasoning APIs',
      'Choose the right model for your task and build assistants with a familiar API.',
    ],
  },
  {
    id: 'creative',
    icon: 'qwen',
    theme: 'rose',
    zh: [
      '图像 · 视频 · 语音',
      '{{brand}} · ',
      '多模态创作 API',
      '探索多模态模型，将文字、图像与声音融入你的产品体验。',
    ],
    en: [
      'IMAGE · VIDEO · AUDIO',
      '{{brand}} · ',
      'Multimodal AI APIs',
      'Explore multimodal models and bring text, images and audio into your products.',
    ],
  },
  {
    id: 'developers',
    icon: 'openai',
    theme: 'teal',
    zh: [
      '开发者友好 · 统一协议',
      '{{brand}} · ',
      'OpenAI 兼容 API',
      '使用现有 OpenAI SDK，只需配置接入地址、密钥和模型名称，即可开始调用。',
    ],
    en: [
      'BUILT FOR DEVELOPERS',
      '{{brand}} · ',
      'OpenAI-compatible APIs',
      'Use your existing OpenAI SDK with a new base URL, API key and model name.',
    ],
  },
]

export function createDefaultHeroConfig(): HeroConfig {
  return {
    schemaVersion: 1,
    enabled: true,
    autoplay: true,
    loop: true,
    intervalMs: 6000,
    defaultLocale: 'zh',
    slides: slides.map((item) => ({
      id: item.id,
      internalName: item.id,
      enabled: true,
      theme: item.theme,
      icon: item.icon,
      logoUrl: '',
      desktopImage: '',
      mobileImage: '',
      imagePosition: 'center',
      content: Object.fromEntries(
        [
          ['zh', item.zh],
          ['en', item.en],
        ].map(([locale, raw]) => {
          const text = raw as string[]
          return [
            locale,
            {
              badge: text[0],
              titlePrefix: text[1],
              highlight: text[2],
              titleSuffix: '',
              description: text[3],
              primaryLabel: locale === 'zh' ? '获取 API 密钥' : 'Get API Key',
              secondaryLabel: locale === 'zh' ? '探索模型' : 'Explore Models',
            },
          ]
        })
      ),
      primaryAction: { url: '/keys', newTab: false },
      secondaryAction: { url: '/pricing', newTab: false },
    })),
  }
}

export function parseHeroConfig(raw: unknown): HeroConfig | null {
  try {
    const value = typeof raw === 'string' ? JSON.parse(raw) : raw
    const result = heroConfigSchema.safeParse(value)
    return result.success ? result.data : null
  } catch {
    return null
  }
}

export function getVisibleHeroSlides(config: HeroConfig): HeroSlide[] {
  const active = config.enabled
    ? config.slides.filter((slide) => slide.enabled)
    : []
  return active.length ? active : createDefaultHeroConfig().slides.slice(0, 1)
}

export function getHeroText(
  slide: HeroSlide,
  locale: string,
  fallback: string,
  brand: string
): HeroText {
  const text =
    slide.content[locale] ??
    slide.content[locale.split('-')[0]] ??
    slide.content[fallback] ??
    Object.values(slide.content)[0]
  return Object.fromEntries(
    Object.entries(text).map(([key, value]) => [
      key,
      value.replaceAll('{{brand}}', brand),
    ])
  ) as HeroText
}

export function duplicateHeroSlide(slide: HeroSlide, id: string): HeroSlide {
  return {
    ...structuredClone(slide),
    id,
    internalName: `${slide.internalName || slide.id} (copy)`,
    enabled: false,
  }
}
