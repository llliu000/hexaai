import { ArrowUp } from 'lucide-react'
import { useEffect, useMemo, useState } from 'react'
import { useTranslation } from 'react-i18next'

const TYPE_DELAY = 90
const DELETE_DELAY = 45
const HOLD_DELAY = 1600
const NEXT_PROMPT_DELAY = 280

function TypewriterText(props: { texts: string[] }) {
  const [visibleText, setVisibleText] = useState('')

  useEffect(() => {
    const prompts = props.texts.map((text) => Array.from(text))
    const reduceMotion = window.matchMedia(
      '(prefers-reduced-motion: reduce)'
    ).matches

    if (reduceMotion) {
      setVisibleText(props.texts[0] ?? '')
      return
    }

    let promptIndex = 0
    let characterIndex = 0
    let isDeleting = false
    let timer: number
    setVisibleText('')

    const tick = () => {
      const characters = prompts[promptIndex] ?? []

      if (!isDeleting && characterIndex < characters.length) {
        characterIndex += 1
        setVisibleText(characters.slice(0, characterIndex).join(''))
        timer = window.setTimeout(tick, TYPE_DELAY)
        return
      }

      if (!isDeleting) {
        isDeleting = true
        timer = window.setTimeout(tick, HOLD_DELAY)
        return
      }

      if (characterIndex > 0) {
        characterIndex -= 1
        setVisibleText(characters.slice(0, characterIndex).join(''))
        timer = window.setTimeout(tick, DELETE_DELAY)
        return
      }

      isDeleting = false
      promptIndex = (promptIndex + 1) % prompts.length
      timer = window.setTimeout(tick, NEXT_PROMPT_DELAY)
    }

    timer = window.setTimeout(tick, NEXT_PROMPT_DELAY)

    return () => window.clearTimeout(timer)
  }, [props.texts])

  return (
    <span className='auth-page-typewriter' aria-label={props.texts.join(' ')}>
      {props.texts.map((text) => (
        <span
          key={text}
          className='auth-page-typewriter-sizer'
          aria-hidden='true'
        >
          {text}
        </span>
      ))}
      <span className='auth-page-typewriter-text' aria-hidden='true'>
        {visibleText}
      </span>
    </span>
  )
}

export function AuthPageShowcase(props: { brand: string }) {
  const { t } = useTranslation()
  const prompts = useMemo(
    () =>
      [
        'auth.page.prompt',
        'auth.page.promptExplore',
        'auth.page.promptCreate',
        'auth.page.promptAgent',
      ].map((key) => t(key, { brand: props.brand })),
    [props.brand, t]
  )

  return (
    <aside className='auth-page-showcase'>
      <div className='auth-page-showcase-glow' aria-hidden='true' />
      <div className='auth-page-sculpture' aria-hidden='true'>
        <span />
        <span />
        <span />
      </div>
      <div className='auth-page-showcase-content'>
        <h2>
          <span>{t('auth.page.heroEyebrow')}</span>
          {t('auth.page.heroTitle')}
        </h2>
        <div className='auth-page-prompt'>
          <TypewriterText texts={prompts} />
          <span className='auth-page-prompt-arrow' aria-hidden='true'>
            <ArrowUp size={20} />
          </span>
        </div>
        <p>{t('auth.page.footnote')}</p>
      </div>
    </aside>
  )
}
