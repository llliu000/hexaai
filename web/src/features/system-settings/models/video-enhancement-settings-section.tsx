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
import { useTranslation } from 'react-i18next'

import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'

import { FormDirtyIndicator } from '../components/form-dirty-indicator'
import { SettingsForm } from '../components/settings-form-layout'
import { SettingsPageFormActions } from '../components/settings-page-context'
import { SettingsSection } from '../components/settings-section'
import { useVideoEnhancementForm } from './hooks/use-video-enhancement-form'
import type { VideoEnhancementFormValues } from './lib/video-enhancement-form'

type EnhancementField = {
  name: `enhance.${keyof VideoEnhancementFormValues['enhance']}`
  label: string
  description?: string
  placeholder?: string
  type?: 'text' | 'password'
  options?: Array<{ value: string; label: string }>
}

type VideoEnhancementSettingsSectionProps = {
  defaultValues: VideoEnhancementFormValues
}

export function VideoEnhancementSettingsSection(
  props: VideoEnhancementSettingsSectionProps
) {
  const { t } = useTranslation()
  const settingsForm = useVideoEnhancementForm(props.defaultValues)

  const fields: EnhancementField[] = [
    {
      name: 'enhance.channel',
      label: t('Provider'),
      options: [{ value: 'volc', label: t('Volcengine') }],
    },
    {
      name: 'enhance.base_url',
      label: t('Base URL'),
      placeholder: 'https://mediakit.cn-beijing.volces.com',
      description: t('Leave blank to use the default service address.'),
    },
    {
      name: 'enhance.api_key',
      label: t('API Key'),
      type: 'password',
      placeholder: t('Enter new key to update'),
      description: t('Leave blank to keep the existing API key.'),
    },
    {
      name: 'enhance.tool_version',
      label: t('Enhancement version'),
      options: [
        { value: 'standard', label: t('Standard (standard)') },
        { value: 'professional', label: t('Professional (professional)') },
        { value: 'generative', label: t('Generative (generative)') },
        { value: 'fast', label: t('Fast (fast)') },
      ],
    },
    {
      name: 'enhance.scene',
      label: t('Enhancement scene'),
      placeholder: 'aigc',
      description: t(
        'Only used by the standard version. Leave blank to use aigc.'
      ),
    },
    {
      name: 'enhance.bitrate_level',
      label: t('Bitrate level'),
      options: [
        { value: 'low', label: t('Low (low)') },
        { value: 'medium', label: t('Medium (medium)') },
        { value: 'high', label: t('High (high)') },
      ],
    },
  ]

  return (
    <SettingsSection title={t('Video Enhancement')}>
      <FormDirtyIndicator isDirty={settingsForm.isDirty} />
      <Form {...settingsForm.form}>
        <SettingsForm
          onSubmit={settingsForm.handleSubmit}
          autoComplete='off'
          noValidate
        >
          <SettingsPageFormActions
            onSave={settingsForm.handleSubmit}
            onReset={settingsForm.handleReset}
            isSaving={settingsForm.isSaving}
            isSaveDisabled={!settingsForm.isDirty}
            isResetDisabled={!settingsForm.isDirty}
          />
          {fields.map((config) => (
            <FormField
              key={config.name}
              control={settingsForm.form.control}
              name={config.name}
              render={({ field }) => {
                let options = config.options
                if (
                  options &&
                  field.value &&
                  !options.some((item) => item.value === field.value)
                ) {
                  options = [
                    ...options,
                    { value: field.value, label: field.value },
                  ]
                }
                return (
                  <FormItem>
                    <FormLabel>{config.label}</FormLabel>
                    {options ? (
                      <Select
                        items={options}
                        value={field.value}
                        onValueChange={field.onChange}
                        disabled={settingsForm.isSaving}
                      >
                        <FormControl>
                          <SelectTrigger
                            className='w-full'
                            onBlur={field.onBlur}
                            ref={field.ref}
                          >
                            <SelectValue />
                          </SelectTrigger>
                        </FormControl>
                        <SelectContent>
                          {options.map((option) => (
                            <SelectItem key={option.value} value={option.value}>
                              {option.label}
                            </SelectItem>
                          ))}
                        </SelectContent>
                      </Select>
                    ) : (
                      <FormControl>
                        <Input
                          {...field}
                          type={config.type ?? 'text'}
                          placeholder={config.placeholder}
                          autoComplete={
                            config.type === 'password' ? 'new-password' : 'off'
                          }
                          disabled={settingsForm.isSaving}
                        />
                      </FormControl>
                    )}
                    {config.description && (
                      <FormDescription>{config.description}</FormDescription>
                    )}
                    <FormMessage />
                  </FormItem>
                )
              }}
            />
          ))}
        </SettingsForm>
      </Form>
    </SettingsSection>
  )
}
