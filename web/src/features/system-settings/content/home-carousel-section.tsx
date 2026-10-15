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
import { zodResolver } from '@hookform/resolvers/zod'
import {
  ArrowDown,
  ArrowUp,
  Copy,
  Eye,
  Laptop,
  Plus,
  Save,
  Smartphone,
  Trash2,
} from 'lucide-react'
import { useEffect, useMemo, useState } from 'react'
import { FormProvider, useFieldArray, useForm, useWatch } from 'react-hook-form'
import { useTranslation } from 'react-i18next'
import { toast } from 'sonner'

import { Dialog } from '@/components/dialog'
import { Button } from '@/components/ui/button'
import { HeroCarousel } from '@/features/home/components/hero-carousel'
import { useHomeHero } from '@/features/home/hooks/use-home-hero'
import {
  createDefaultHeroConfig,
  duplicateHeroSlide,
  HERO_OPTION_KEY,
  heroConfigSchema,
  parseHeroConfig,
  type HeroConfig,
} from '@/features/home/lib/hero-config'

import { SettingsSection } from '../components/settings-section'
import { useUpdateOption } from '../hooks/use-update-option'
import { HeroSlideFields } from './hero-slide-fields'

import '@/styles/reference-home.css'
import '@/styles/home-carousel-editor.css'

export function HomeCarouselSection(props: { data: string }) {
  const { t } = useTranslation()
  const initial = useMemo(
    () => parseHeroConfig(props.data) ?? createDefaultHeroConfig(),
    [props.data]
  )
  const form = useForm<HeroConfig>({
    resolver: zodResolver(heroConfigSchema),
    defaultValues: initial,
  })
  const slides = useFieldArray({
    control: form.control,
    name: 'slides',
    keyName: 'formId',
  })
  const config = useWatch({ control: form.control }) as HeroConfig
  const [selected, setSelected] = useState(0)
  const [locale, setLocale] = useState(initial.defaultLocale)
  const [preview, setPreview] = useState<HeroConfig | null>(null)
  const [previewSize, setPreviewSize] = useState<'desktop' | 'mobile'>(
    'desktop'
  )
  const [deleteIndex, setDeleteIndex] = useState<number | null>(null)
  const update = useUpdateOption()
  const backend = useHomeHero()
  useEffect(() => {
    if (!form.formState.isDirty) form.reset(initial)
  }, [initial, form])
  useEffect(() => {
    if (!form.formState.isDirty) return
    const beforeUnload = (event: BeforeUnloadEvent) => {
      event.preventDefault()
    }
    window.addEventListener('beforeunload', beforeUnload)
    return () => window.removeEventListener('beforeunload', beforeUnload)
  }, [form.formState.isDirty])
  const validatePreview = () => {
    const result = heroConfigSchema.safeParse(form.getValues())
    if (result.success) setPreview(result.data)
    else {
      void form.trigger()
      toast.error(t('homepage.admin.invalid'))
    }
  }
  const publish = form.handleSubmit(
    async (values) => {
      if (!backend.data?.available) return
      try {
        const result = await update.mutateAsync({
          key: HERO_OPTION_KEY,
          value: JSON.stringify(values),
        })
        if (result.success) form.reset(values)
      } catch {
        /* The mutation hook displays the failure and keeps the draft. */
      }
    },
    () => toast.error(t('homepage.admin.invalid'))
  )
  const activeCount =
    config.slides?.filter((slide) => slide.enabled).length ?? 0
  return (
    <SettingsSection title={t('homepage.admin.title')}>
      <FormProvider {...form}>
        <form onSubmit={publish} className='hero-config-editor'>
          <div className='hero-config-global'>
            <label className='hero-config-check'>
              <input type='checkbox' {...form.register('enabled')} />
              {t('homepage.admin.enabled')}
            </label>
            <label className='hero-config-check'>
              <input type='checkbox' {...form.register('autoplay')} />
              {t('homepage.admin.autoplay')}
            </label>
            <label className='hero-config-check'>
              <input type='checkbox' {...form.register('loop')} />
              {t('homepage.admin.loop')}
            </label>
            <label className='hero-config-field'>
              <span>{t('homepage.admin.interval')}</span>
              <input
                type='number'
                min={3}
                max={15}
                value={(config.intervalMs || 6000) / 1000}
                onChange={(event) =>
                  form.setValue(
                    'intervalMs',
                    Number(event.target.value) * 1000,
                    { shouldDirty: true }
                  )
                }
              />
            </label>
            <label className='hero-config-field'>
              <span>{t('homepage.admin.defaultLocale')}</span>
              <select {...form.register('defaultLocale')}>
                <option value='zh'>简体中文</option>
                <option value='en'>English</option>
                <option value='zh-TW'>繁體中文</option>
                <option value='ja'>日本語</option>
                <option value='fr'>Français</option>
                <option value='ru'>Русский</option>
                <option value='vi'>Tiếng Việt</option>
              </select>
            </label>
          </div>
          {!backend.isLoading && !backend.data?.available && (
            <p role='alert' className='hero-config-notice'>
              {t('homepage.admin.backendRequired')}
            </p>
          )}
          {(!config.enabled || activeCount === 0) && (
            <p className='hero-config-notice'>{t('homepage.admin.fallback')}</p>
          )}
          <div className='hero-config-toolbar'>
            <Button
              type='button'
              variant='outline'
              onClick={() => {
                const slide = duplicateHeroSlide(
                  createDefaultHeroConfig().slides[0],
                  crypto.randomUUID()
                )
                slide.internalName = t('homepage.admin.newSlide')
                slide.enabled = true
                if (!slide.content[config.defaultLocale]) {
                  slide.content[config.defaultLocale] = structuredClone(
                    slide.content.zh
                  )
                }
                slides.append(slide)
                setSelected(slides.fields.length)
              }}
              disabled={slides.fields.length >= 12}
            >
              <Plus size={16} />
              {t('homepage.admin.add')}
            </Button>
            <Button type='button' variant='outline' onClick={validatePreview}>
              <Eye size={16} />
              {t('Preview')}
            </Button>
            <Button
              type='submit'
              disabled={
                update.isPending ||
                !backend.data?.available ||
                !form.formState.isDirty
              }
            >
              <Save size={16} />
              {t(update.isPending ? 'Saving...' : 'homepage.admin.publish')}
            </Button>
          </div>
          <div className='hero-config-layout'>
            <aside
              className='hero-config-list'
              aria-label={t('homepage.admin.slides')}
            >
              {slides.fields.map((slide, index) => (
                <div
                  className='hero-config-list-row'
                  data-selected={selected === index}
                  key={slide.formId}
                >
                  <button
                    type='button'
                    className='hero-config-select'
                    aria-pressed={selected === index}
                    onClick={() => setSelected(index)}
                  >
                    <strong>
                      {index + 1}.{' '}
                      {config.slides[index]?.internalName || slide.id}
                    </strong>
                    <span>
                      {t(
                        config.slides[index]?.enabled ? 'Enabled' : 'Disabled'
                      )}
                    </span>
                  </button>
                  <div className='hero-config-row-actions'>
                    <Button
                      type='button'
                      variant='ghost'
                      size='icon-sm'
                      title={t('Move up')}
                      aria-label={t('Move up')}
                      disabled={index === 0}
                      onClick={() => {
                        slides.move(index, index - 1)
                        setSelected(index - 1)
                      }}
                    >
                      <ArrowUp size={15} />
                    </Button>
                    <Button
                      type='button'
                      variant='ghost'
                      size='icon-sm'
                      title={t('Move down')}
                      aria-label={t('Move down')}
                      disabled={index === slides.fields.length - 1}
                      onClick={() => {
                        slides.move(index, index + 1)
                        setSelected(index + 1)
                      }}
                    >
                      <ArrowDown size={15} />
                    </Button>
                    <Button
                      type='button'
                      variant='ghost'
                      size='icon-sm'
                      title={t('Duplicate')}
                      aria-label={t('Duplicate')}
                      disabled={slides.fields.length >= 12}
                      onClick={() => {
                        slides.insert(
                          index + 1,
                          duplicateHeroSlide(
                            form.getValues(`slides.${index}`),
                            crypto.randomUUID()
                          )
                        )
                        setSelected(index + 1)
                      }}
                    >
                      <Copy size={15} />
                    </Button>
                    <Button
                      type='button'
                      variant='ghost'
                      size='icon-sm'
                      title={t('Delete')}
                      aria-label={t('Delete')}
                      onClick={() => setDeleteIndex(index)}
                    >
                      <Trash2 size={15} />
                    </Button>
                  </div>
                </div>
              ))}
            </aside>
            <div className='hero-config-edit'>
              {slides.fields[selected] ? (
                <HeroSlideFields
                  key={slides.fields[selected].formId}
                  index={selected}
                  locale={locale}
                  onLocaleChange={setLocale}
                />
              ) : (
                <p className='text-muted-foreground p-6 text-sm'>
                  {t('homepage.admin.empty')}
                </p>
              )}
            </div>
          </div>
          <p className='text-muted-foreground text-xs'>
            {t('homepage.admin.hint', { brand: '{{brand}}' })}
          </p>
        </form>
      </FormProvider>
      <Dialog
        open={deleteIndex !== null}
        onOpenChange={(open) => {
          if (!open) setDeleteIndex(null)
        }}
        title={t('Delete')}
        description={t('homepage.admin.deleteConfirm')}
        footer={
          <>
            <Button variant='outline' onClick={() => setDeleteIndex(null)}>
              {t('Cancel')}
            </Button>
            <Button
              variant='destructive'
              onClick={() => {
                if (deleteIndex !== null) {
                  slides.remove(deleteIndex)
                  setSelected(
                    Math.max(0, Math.min(selected, slides.fields.length - 2))
                  )
                }
                setDeleteIndex(null)
              }}
            >
              {t('Delete')}
            </Button>
          </>
        }
      >
        {null}
      </Dialog>
      <Dialog
        open={preview !== null}
        onOpenChange={(open) => {
          if (!open) setPreview(null)
        }}
        title={t('Preview')}
        contentClassName='sm:max-w-[1200px]'
        contentHeight='auto'
        footer={
          <Button variant='outline' onClick={() => setPreview(null)}>
            {t('Close')}
          </Button>
        }
      >
        <div className='hero-config-preview-controls'>
          <Button
            variant={previewSize === 'desktop' ? 'default' : 'outline'}
            size='icon'
            title={t('Desktop')}
            aria-label={t('Desktop')}
            onClick={() => setPreviewSize('desktop')}
          >
            <Laptop size={18} />
          </Button>
          <Button
            variant={previewSize === 'mobile' ? 'default' : 'outline'}
            size='icon'
            title={t('Mobile')}
            aria-label={t('Mobile')}
            onClick={() => setPreviewSize('mobile')}
          >
            <Smartphone size={18} />
          </Button>
        </div>
        {preview && (
          <div
            className={`wm-home hero-config-preview hero-config-preview-${previewSize}`}
          >
            <HeroCarousel config={preview} preview locale={locale} />
          </div>
        )}
      </Dialog>
    </SettingsSection>
  )
}
