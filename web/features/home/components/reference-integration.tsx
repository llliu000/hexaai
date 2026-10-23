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
import { Link } from '@tanstack/react-router'
import { ArrowRight, Blocks, Check, Code2, Copy, Workflow } from 'lucide-react'
import { useId, useState } from 'react'
import { useTranslation } from 'react-i18next'

import { CodeBlock } from '@/components/ai-elements/code-block'

import { HomeParallaxDecorations } from './home-parallax'
import { useCopyToClipboard } from '@/hooks/use-copy-to-clipboard'
import { useStatus } from '@/hooks/use-status'

import {
  getSdkBaseUrl,
  getSdkExample,
  type ExampleLanguage,
} from '../lib/sdk-example'

const languages = [
  { id: 'python', title: 'Python', syntax: 'python', filename: 'main.py' },
  {
    id: 'javascript',
    title: 'Node.js',
    syntax: 'javascript',
    filename: 'index.mjs',
  },
  { id: 'curl', title: 'cURL', syntax: 'bash', filename: 'request.sh' },
] as const

export function ReferenceIntegration(props: { docsUrl: string }) {
  const { t } = useTranslation()
  const { status } = useStatus()
  const [language, setLanguage] = useState<ExampleLanguage>('python')
  const { copiedText, copyToClipboard } = useCopyToClipboard()
  const id = useId()
  const base = getSdkBaseUrl(status?.server_address, window.location.origin)
  const example = getSdkExample(language, base, 'YOUR_MODEL_NAME')
  const selected =
    languages.find((item) => item.id === language) ?? languages[0]
  return (
    <section className='wm-integration wm-home-parallax-section' id='integration'>
      <HomeParallaxDecorations />
      <div className='wm-container'>
        <div className='wm-section-heading wm-home-parallax-heading'>
          <h2>{t('homepage.integration.title')}</h2>
        </div>
        <div className='wm-integration-grid'>
          <div className='wm-integration-copy wm-home-parallax-panel'>
            <p>{t('homepage.integration.description')}</p>
            {[
              { id: 'sdk', Icon: Blocks },
              { id: 'protocol', Icon: Code2 },
              { id: 'features', Icon: Workflow },
            ].map(({ id: feature, Icon }) => (
              <div className='wm-integration-feature' key={feature}>
                <span>
                  <Icon size={22} />
                </span>
                <div>
                  <h3>{t(`homepage.integration.${feature}.title`)}</h3>
                  <p>{t(`homepage.integration.${feature}.description`)}</p>
                </div>
              </div>
            ))}
            <div className='wm-integration-actions'>
              <Link to='/keys' className='wm-button wm-button-primary'>
                {t('homepage.getKey')}
                <ArrowRight size={18} />
              </Link>
              <a
                className='wm-text-link'
                href={props.docsUrl}
                target='_blank'
                rel='noreferrer'
              >
                {t('homepage.integration.docs')}
                <ArrowRight size={16} />
              </a>
            </div>
          </div>
          <div className='wm-code-panel wm-home-parallax-panel'>
            <div className='wm-code-toolbar'>
              <div
                role='tablist'
                aria-label={t('homepage.integration.language')}
              >
                {languages.map((item, index) => (
                  <button
                    key={item.id}
                    type='button'
                    id={`${id}-${item.id}`}
                    role='tab'
                    aria-selected={language === item.id}
                    aria-controls={`${id}-code`}
                    tabIndex={language === item.id ? 0 : -1}
                    onClick={() => setLanguage(item.id)}
                    onKeyDown={(event) => {
                      if (
                        event.key === 'ArrowRight' ||
                        event.key === 'ArrowLeft'
                      ) {
                        event.preventDefault()
                        const next =
                          languages[
                            (index + (event.key === 'ArrowRight' ? 1 : 2)) % 3
                          ]
                        setLanguage(next.id)
                        document.getElementById(`${id}-${next.id}`)?.focus()
                      }
                    }}
                  >
                    {item.title}
                  </button>
                ))}
              </div>
              <div className='wm-code-toolbar-actions'>
                <span>{selected.filename}</span>
                <button
                  type='button'
                  onClick={() => void copyToClipboard(example)}
                  aria-label={t('Copy code')}
                  title={t('Copy code')}
                >
                  {copiedText === example ? (
                    <Check size={18} />
                  ) : (
                    <Copy size={18} />
                  )}
                </button>
              </div>
            </div>
            <div
              role='tabpanel'
              id={`${id}-code`}
              aria-labelledby={`${id}-${language}`}
            >
              <CodeBlock
                code={example}
                language={selected.syntax}
                showLineNumbers
                enableCollapse={false}
              />
            </div>
          </div>
        </div>
      </div>
    </section>
  )
}
