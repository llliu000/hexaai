import {
  Baidu,
  ByteDance,
  Claude,
  DeepSeek,
  Hunyuan,
  Kimi,
  Meta,
  Minimax,
  OpenAI,
  Qwen,
  Zhipu,
} from '@lobehub/icons'
import { Pause, Play } from 'lucide-react'

import { HOME_MODELS } from '../lib/lanhu-home-models'

const providerIcons = {
  Baidu,
  ByteDance,
  Claude,
  DeepSeek,
  Hunyuan,
  Kimi,
  Meta,
  Minimax,
  OpenAI,
  Qwen,
  Zhipu,
}
const brands = [
  ['qwen', 'Qwen'],
  ['openai', 'OpenAI'],
  ['bytedance', '字节跳动'],
  ['gemini', 'Gemini'],
  ['kuaishou', '快手'],
  ['claude', 'Claude'],
  ['deepseek', 'DeepSeek'],
  ['grok', 'Grok'],
] as const

export function LanhuHomeProviders() {
  return (
    <div
      className='lh-providers'
      role='region'
      aria-label='支持的模型与供应商'
      tabIndex={0}
    >
      <div className='lh-provider-track'>
        {[false, true].map((duplicate) => (
          <div
            className='lh-provider-group'
            aria-hidden={duplicate || undefined}
            key={String(duplicate)}
          >
            {brands.map(([image, label]) => (
              <span className='lh-provider' key={image}>
                <img src={`/lanhu/provider-${image}.png`} alt={label} />
              </span>
            ))}
            {HOME_MODELS.map((model) => {
              const Icon = providerIcons[model.provider]
              return (
                <span
                  className='lh-provider lh-provider-model'
                  key={model.name}
                >
                  <Icon size={24} aria-hidden />
                  <span>{model.name}</span>
                </span>
              )
            })}
          </div>
        ))}
      </div>
      <button
        type='button'
        className='lh-provider-pause'
        data-pause-providers
        aria-label='暂停或继续模型滚动'
        title='暂停或继续模型滚动'
        aria-pressed='false'
      >
        <Pause className='lh-pause-icon' aria-hidden />
        <Play className='lh-play-icon' aria-hidden />
      </button>
    </div>
  )
}
