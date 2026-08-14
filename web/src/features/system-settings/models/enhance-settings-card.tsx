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
import { useEffect, useMemo, useRef } from 'react'
import { useForm } from 'react-hook-form'
import { useTranslation } from 'react-i18next'
import { toast } from 'sonner'
import * as z from 'zod'

import { PasswordInput } from '@/components/password-input'
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
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'

import { SettingsForm } from '../components/settings-form-layout'
import { SettingsPageFormActions } from '../components/settings-page-context'
import { SettingsSection } from '../components/settings-section'
import { useUpdateOption } from '../hooks/use-update-option'

const enhanceSchema = z.object({
  enhance: z.object({
    channel: z.literal('volc'),
    base_url: z.string().trim().url('Must be a valid URL'),
    api_key: z.string(),
  }),
})

type EnhanceFormInput = z.input<typeof enhanceSchema>
type EnhanceFormValues = z.output<typeof enhanceSchema>

export type EnhanceSettingsDefaults = {
  'enhance.channel': string
  'enhance.base_url': string
}

const buildFormDefaults = (
  defaults: EnhanceSettingsDefaults
): EnhanceFormInput => ({
  enhance: {
    channel: 'volc',
    base_url: defaults['enhance.base_url'],
    api_key: '',
  },
})

interface EnhanceSettingsCardProps {
  defaultValues: EnhanceSettingsDefaults
}

export function EnhanceSettingsCard(props: EnhanceSettingsCardProps) {
  const { t } = useTranslation()
  const updateOption = useUpdateOption()
  const formDefaults = useMemo(
    () => buildFormDefaults(props.defaultValues),
    [props.defaultValues]
  )
  const form = useForm<EnhanceFormInput, unknown, EnhanceFormValues>({
    resolver: zodResolver(enhanceSchema),
    defaultValues: formDefaults,
  })
  const baselineRef = useRef(props.defaultValues)
  const baselineSerializedRef = useRef(JSON.stringify(props.defaultValues))

  useEffect(() => {
    const serialized = JSON.stringify(props.defaultValues)
    if (serialized === baselineSerializedRef.current) return

    baselineRef.current = props.defaultValues
    baselineSerializedRef.current = serialized
    form.reset(buildFormDefaults(props.defaultValues))
  }, [form, props.defaultValues])

  const onSubmit = async (values: EnhanceFormValues) => {
    const updates: Array<{ key: string; value: string }> = []
    const channel = values.enhance.channel.trim()
    const baseUrl = values.enhance.base_url.trim()
    const apiKey = values.enhance.api_key.trim()

    if (channel !== baselineRef.current['enhance.channel']) {
      updates.push({ key: 'enhance.channel', value: channel })
    }
    if (baseUrl !== baselineRef.current['enhance.base_url']) {
      updates.push({ key: 'enhance.base_url', value: baseUrl })
    }
    if (apiKey) {
      updates.push({ key: 'enhance.api_key', value: apiKey })
    }

    if (updates.length === 0) {
      toast.info(t('No changes to save'))
      return
    }

    for (const update of updates) {
      await updateOption.mutateAsync(update)
    }

    const savedDefaults: EnhanceSettingsDefaults = {
      'enhance.channel': channel,
      'enhance.base_url': baseUrl,
    }
    baselineRef.current = savedDefaults
    baselineSerializedRef.current = JSON.stringify(savedDefaults)
    form.reset(buildFormDefaults(savedDefaults))
  }

  return (
    <SettingsSection title={t('Video Enhancement Settings')}>
      <Form {...form}>
        <SettingsForm
          onSubmit={form.handleSubmit(onSubmit)}
          autoComplete='off'
          className='w-full max-w-2xl lg:grid-cols-1'
        >
          <SettingsPageFormActions
            onSave={form.handleSubmit(onSubmit)}
            isSaving={updateOption.isPending}
          />

          <FormField
            control={form.control}
            name='enhance.channel'
            render={({ field }) => (
              <FormItem className='max-w-md'>
                <FormLabel>{t('Provider')}</FormLabel>
                <Select
                  items={[{ value: 'volc', label: t('Volcengine') }]}
                  value={field.value}
                  onValueChange={field.onChange}
                >
                  <FormControl>
                    <SelectTrigger>
                      <SelectValue />
                    </SelectTrigger>
                  </FormControl>
                  <SelectContent alignItemWithTrigger={false}>
                    <SelectGroup>
                      <SelectItem value='volc'>{t('Volcengine')}</SelectItem>
                    </SelectGroup>
                  </SelectContent>
                </Select>
                <FormDescription>
                  {t('Currently only Volcengine is supported.')}
                </FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name='enhance.base_url'
            render={({ field }) => (
              <FormItem className='max-w-2xl'>
                <FormLabel>{t('Base URL')}</FormLabel>
                <FormControl>
                  <Input
                    {...field}
                    inputMode='url'
                    autoComplete='url'
                    placeholder='https://mediakit.cn-beijing.volces.com'
                  />
                </FormControl>
                <FormDescription>
                  {t('Video enhancement service endpoint.')}
                </FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name='enhance.api_key'
            render={({ field }) => (
              <FormItem className='max-w-2xl'>
                <FormLabel>{t('API Key')}</FormLabel>
                <FormControl>
                  <PasswordInput
                    {...field}
                    autoComplete='new-password'
                    placeholder={t('Enter API Key')}
                  />
                </FormControl>
                <FormDescription>
                  {t('Leave blank to keep the current API key.')}
                </FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />
        </SettingsForm>
      </Form>
    </SettingsSection>
  )
}
