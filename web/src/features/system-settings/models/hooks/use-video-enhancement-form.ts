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
import { useMutation, useQueryClient } from '@tanstack/react-query'
import { useEffect } from 'react'
import { useForm } from 'react-hook-form'
import { useTranslation } from 'react-i18next'
import { toast } from 'sonner'

import { updateSystemOption } from '../../api'
import type { UpdateOptionRequest, UpdateOptionResponse } from '../../types'
import {
  videoEnhancementSchema,
  type VideoEnhancementFormValues,
} from '../lib/video-enhancement-form'

export function useVideoEnhancementForm(
  defaultValues: VideoEnhancementFormValues
) {
  const { t } = useTranslation()
  const queryClient = useQueryClient()
  const saveOptions = useMutation({
    mutationFn: async (
      updates: UpdateOptionRequest[]
    ): Promise<UpdateOptionResponse> => {
      for (const update of updates) {
        const response = await updateSystemOption(update)
        if (!response.success) return response
      }
      return { success: true, message: '' }
    },
    onSettled: () =>
      queryClient.invalidateQueries({ queryKey: ['system-options'] }),
  })
  const form = useForm<VideoEnhancementFormValues>({
    resolver: zodResolver(videoEnhancementSchema),
    defaultValues,
  })
  const { isDirty, isSubmitting } = form.formState

  useEffect(() => {
    if (!isDirty && !isSubmitting) form.reset(defaultValues)
  }, [form, defaultValues, isDirty, isSubmitting])

  const onSubmit = async (
    values: VideoEnhancementFormValues
  ): Promise<void> => {
    const updates: UpdateOptionRequest[] = []
    for (const [key, value] of Object.entries(values.enhance)) {
      if (key === 'api_key') continue
      if (value !== defaultValues.enhance[key as keyof typeof values.enhance]) {
        updates.push({ key: `enhance.${key}`, value })
      }
    }
    const apiKey = values.enhance.api_key.trim()
    if (apiKey) updates.push({ key: 'enhance.api_key', value: apiKey })
    if (updates.length === 0) {
      toast.info(t('No changes to save'))
      return
    }

    try {
      const response = await saveOptions.mutateAsync(updates)
      if (!response.success) {
        toast.error(response.message || t('Failed to update setting'))
        return
      }
      form.reset({ enhance: { ...values.enhance, api_key: '' } })
      saveOptions.reset()
      toast.success(t('Setting updated successfully'))
    } catch {
      // Request errors are displayed by the global mutation error handler.
    }
  }

  return {
    form,
    handleSubmit: form.handleSubmit(onSubmit),
    handleReset: () => form.reset(defaultValues),
    isDirty,
    isSaving: saveOptions.isPending || isSubmitting,
  }
}
