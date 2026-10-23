import { Copy, Key } from 'lucide-react'
import { useEffect, useState, type KeyboardEvent, type MouseEvent } from 'react'
import { renderToStaticMarkup } from 'react-dom/server'
import { useTranslation } from 'react-i18next'

import { RichContent } from '@/components/rich-content'

import template from '../../../../option/home.html?raw'
import {
  selectCodeLanguage,
  toggleProviderScrolling,
  watchHeaderScroll,
} from '../lib/lanhu-home-interactions'
import { LanhuHomeProviders } from './lanhu-home-providers'

import styles from '@/styles/lanhu-home.css?raw'

const content = `<style>${styles}</style>${template}`
  .replace(
    '<div data-lh-providers></div>',
    renderToStaticMarkup(<LanhuHomeProviders />)
  )
  .replace(
    '<span data-lh-icon="key" aria-hidden="true"></span>',
    renderToStaticMarkup(<Key aria-hidden strokeWidth={2.5} />)
  )
  .replace(
    '<span data-lh-icon="copy" aria-hidden="true"></span>',
    renderToStaticMarkup(<Copy aria-hidden />)
  )

function findButton(
  event: MouseEvent | KeyboardEvent
): HTMLButtonElement | undefined {
  return event.nativeEvent
    .composedPath()
    .find(
      (element): element is HTMLButtonElement =>
        element instanceof HTMLButtonElement
    )
}

export function LanhuHome() {
  const { t } = useTranslation()
  const [floatingHeader, setFloatingHeader] = useState(false)
  useEffect(() => watchHeaderScroll(window, setFloatingHeader), [])

  async function handleClick(event: MouseEvent) {
    const button = findButton(event)
    if (!button) return
    const root = button.getRootNode() as ShadowRoot
    if (button.hasAttribute('data-pause-providers')) {
      toggleProviderScrolling(button)
    }
    if (button.dataset.codeLanguage) {
      selectCodeLanguage(root, button.dataset.codeLanguage)
    }
    if (button.dataset.copyUrl) {
      const status = root.querySelector('.lh-copy-status')
      try {
        await navigator.clipboard.writeText(button.dataset.copyUrl)
        if (status) status.textContent = t('Copied')
      } catch {
        if (status) status.textContent = t('Copy failed')
      }
    }
  }

  function handleKeyDown(event: KeyboardEvent) {
    const button = findButton(event)
    if (!button?.dataset.codeLanguage) return
    const languages = ['curl', 'python', 'javascript']
    const current = languages.indexOf(button.dataset.codeLanguage)
    let next = current
    if (event.key === 'ArrowRight') {
      next = (current + 1) % languages.length
    } else if (event.key === 'ArrowLeft') {
      next = (current + languages.length - 1) % languages.length
    } else if (event.key === 'Home') {
      next = 0
    } else if (event.key === 'End') {
      next = languages.length - 1
    } else {
      return
    }
    event.preventDefault()
    const root = button.getRootNode() as ShadowRoot
    selectCodeLanguage(root, languages[next])
    root
      .querySelector<HTMLButtonElement>(
        `[data-code-language="${languages[next]}"]`
      )
      ?.focus()
  }

  return (
    // Events bubble out of the isolated template's shadow root.
    // eslint-disable-next-line jsx-a11y/no-static-element-interactions -- delegated events for native buttons
    <div onClick={handleClick} onKeyDown={handleKeyDown}>
      <RichContent
        mode='html'
        htmlVariant='isolated'
        content={content}
        className={`lanhu-home-content${floatingHeader ? ' lh-header-floating' : ''}`}
      />
    </div>
  )
}
