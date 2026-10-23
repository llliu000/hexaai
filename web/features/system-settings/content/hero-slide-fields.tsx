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
import { useFormContext, type FieldPath } from 'react-hook-form'
import { useTranslation } from 'react-i18next'

import { Input } from '@/components/ui/input'
import { Textarea } from '@/components/ui/textarea'
import {
  HERO_ICONS,
  HERO_THEMES,
  type HeroConfig,
} from '@/features/home/lib/hero-config'

const THEMES = {
  violet: '#7650de',
  blue: '#327de2',
  teal: '#12877e',
  rose: '#cd5487',
}

export function HeroSlideFields(props: {
  index: number
  locale: string
  onLocaleChange: (locale: string) => void
}) {
  const { t } = useTranslation()
  const form = useFormContext<HeroConfig>()
  const slide = form.watch(`slides.${props.index}`)
  const base = `slides.${props.index}` as const
  const input = (
    name: FieldPath<HeroConfig>,
    label: string,
    multiline = false
  ) => {
    const error = form.getFieldState(name, form.formState).error
    return (
      <div className='hero-config-field' key={name}>
        <label htmlFor={name}>{t(label)}</label>
        {multiline ? (
          <Textarea
            id={name}
            {...form.register(name)}
            rows={3}
            aria-invalid={Boolean(error)}
            aria-describedby={error ? `${name}-error` : undefined}
          />
        ) : (
          <Input
            id={name}
            {...form.register(name)}
            aria-invalid={Boolean(error)}
            aria-describedby={error ? `${name}-error` : undefined}
          />
        )}
        {error?.message && (
          <span id={`${name}-error`} className='text-destructive text-xs'>
            {t(
              error.message.startsWith('homepage.')
                ? error.message
                : 'homepage.admin.invalidField'
            )}
          </span>
        )}
      </div>
    )
  }
  const textFields = [
    'badge',
    'titlePrefix',
    'highlight',
    'titleSuffix',
    'description',
    'primaryLabel',
    'secondaryLabel',
  ] as const
  return (
    <div className='hero-config-fields'>
      {input(`${base}.internalName`, 'homepage.admin.name')}
      <label className='hero-config-check'>
        <input type='checkbox' {...form.register(`${base}.enabled`)} />
        {t('Enabled')}
      </label>
      <fieldset>
        <legend>{t('homepage.admin.theme')}</legend>
        <div className='hero-config-swatches'>
          {HERO_THEMES.map((theme) => (
            <button
              key={theme}
              type='button'
              style={{ backgroundColor: THEMES[theme] }}
              aria-label={t(`homepage.theme.${theme}`)}
              title={t(`homepage.theme.${theme}`)}
              aria-pressed={slide.theme === theme}
              onClick={() =>
                form.setValue(`${base}.theme`, theme, { shouldDirty: true })
              }
            />
          ))}
        </div>
      </fieldset>
      <label className='hero-config-field'>
        <span>{t('homepage.admin.icon')}</span>
        <select {...form.register(`${base}.icon`)}>
          {HERO_ICONS.map((icon) => (
            <option key={icon} value={icon}>
              {t(`homepage.icon.${icon}`)}
            </option>
          ))}
        </select>
      </label>
      {input(`${base}.logoUrl`, 'homepage.admin.logo')}
      {input(`${base}.desktopImage`, 'homepage.admin.desktopImage')}
      {input(`${base}.mobileImage`, 'homepage.admin.mobileImage')}
      <label className='hero-config-field'>
        <span>{t('homepage.admin.imagePosition')}</span>
        <select {...form.register(`${base}.imagePosition`)}>
          {['left', 'center', 'right'].map((value) => (
            <option key={value} value={value}>
              {t(`homepage.position.${value}`)}
            </option>
          ))}
        </select>
      </label>
      <div className='hero-config-language'>
        <label>
          <span>{t('Language')}</span>
          <select
            value={props.locale}
            onChange={(event) => props.onLocaleChange(event.target.value)}
          >
            {[
              ['zh', '简体中文'],
              ['en', 'English'],
              ['zh-TW', '繁體中文'],
              ['ja', '日本語'],
              ['fr', 'Français'],
              ['ru', 'Русский'],
              ['vi', 'Tiếng Việt'],
            ].map(([value, label]) => (
              <option key={value} value={value}>
                {label}
              </option>
            ))}
          </select>
        </label>
      </div>
      {slide.content[props.locale] ? (
        textFields
          .filter(
            (field) => field !== 'secondaryLabel' || slide.secondaryAction
          )
          .map((field) =>
            input(
              `${base}.content.${props.locale}.${field}`,
              `homepage.admin.${field}`,
              field === 'description'
            )
          )
      ) : (
        <button
          type='button'
          className='hero-config-add-locale'
          onClick={() =>
            form.setValue(
              `${base}.content.${props.locale}`,
              structuredClone(
                slide.content[form.getValues('defaultLocale')] ||
                  Object.values(slide.content)[0]
              ),
              { shouldDirty: true }
            )
          }
        >
          {t('homepage.admin.addLocale')}
        </button>
      )}
      <div className='hero-config-action-row'>
        {input(`${base}.primaryAction.url`, 'homepage.admin.primaryUrl')}
        <label className='hero-config-check'>
          <input
            type='checkbox'
            {...form.register(`${base}.primaryAction.newTab`)}
          />
          {t('homepage.admin.newTab')}
        </label>
      </div>
      <div className='hero-config-action-row'>
        {slide.secondaryAction ? (
          input(`${base}.secondaryAction.url`, 'homepage.admin.secondaryUrl')
        ) : (
          <div className='hero-config-field' />
        )}
        <div className='hero-config-action-options'>
          <label className='hero-config-check'>
            <input
              type='checkbox'
              checked={Boolean(slide.secondaryAction)}
              onChange={(event) => {
                form.setValue(
                  `${base}.secondaryAction`,
                  event.target.checked
                    ? { url: '/pricing', newTab: false }
                    : undefined,
                  { shouldDirty: true }
                )
              }}
            />
            {t('homepage.admin.secondaryEnabled')}
          </label>
          {slide.secondaryAction && (
            <label className='hero-config-check'>
              <input
                type='checkbox'
                {...form.register(`${base}.secondaryAction.newTab`)}
              />
              {t('homepage.admin.newTab')}
            </label>
          )}
        </div>
      </div>
    </div>
  )
}
