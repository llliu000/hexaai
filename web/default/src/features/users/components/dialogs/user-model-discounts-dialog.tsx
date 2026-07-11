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
import { Percent, Loader2 } from 'lucide-react'
import { useEffect, useState } from 'react'
import { useTranslation } from 'react-i18next'
import { toast } from 'sonner'

import { Button } from '@/components/design-system/button'
import { Dialog } from '@/components/dialog'
import { Label } from '@/components/ui/label'
import { Textarea } from '@/components/ui/textarea'

import { updateUserModelDiscounts } from '../../api'
import type { User, UserModelDiscounts } from '../../types'

interface Props {
  open: boolean
  onOpenChange: (open: boolean) => void
  user: User | null
  onSuccess?: () => void
}

function parseUserModelDiscounts(user: User | null): UserModelDiscounts {
  if (!user?.setting) return {}
  try {
    const setting = JSON.parse(user.setting) as {
      model_discounts?: unknown
    }
    if (
      setting.model_discounts &&
      typeof setting.model_discounts === 'object' &&
      !Array.isArray(setting.model_discounts)
    ) {
      return setting.model_discounts as UserModelDiscounts
    }
  } catch {
    return {}
  }
  return {}
}

function formatModelDiscounts(user: User | null): string {
  return JSON.stringify(parseUserModelDiscounts(user), null, 2)
}

function normalizeModelDiscounts(value: string): UserModelDiscounts | null {
  const trimmed = value.trim()
  if (!trimmed) return {}

  let parsed: unknown
  try {
    parsed = JSON.parse(trimmed)
  } catch {
    return null
  }

  if (!parsed || typeof parsed !== 'object' || Array.isArray(parsed)) {
    return null
  }

  const normalized: UserModelDiscounts = {}
  for (const [modelName, discount] of Object.entries(parsed)) {
    const key = modelName.trim()
    if (
      !key ||
      typeof discount !== 'number' ||
      !Number.isInteger(discount) ||
      discount <= 0
    ) {
      return null
    }
    normalized[key] = discount
  }
  return normalized
}

export function UserModelDiscountsDialog({
  open,
  onOpenChange,
  user,
  onSuccess,
}: Props) {
  const { t } = useTranslation()
  const [value, setValue] = useState('{}')
  const [saving, setSaving] = useState(false)

  useEffect(() => {
    if (open) {
      setValue(formatModelDiscounts(user))
    }
  }, [open, user])

  const handleSave = async () => {
    if (!user) return
    const modelDiscounts = normalizeModelDiscounts(value)
    if (!modelDiscounts) {
      toast.error(
        t('Model price config must be a JSON object with positive integer values')
      )
      return
    }

    setSaving(true)
    try {
      const result = await updateUserModelDiscounts(user.id, modelDiscounts)
      if (result.success) {
        toast.success(t('Model price config updated'))
        onOpenChange(false)
        onSuccess?.()
      } else {
        toast.error(result.message || t('Failed to update model price config'))
      }
    } catch {
      toast.error(t('Failed to update model price config'))
    } finally {
      setSaving(false)
    }
  }

  return (
    <Dialog
      open={open}
      onOpenChange={onOpenChange}
      title={
        <>
          <Percent className='h-5 w-5' />
          {t('Model Price Config')}
        </>
      }
      description={
        user
          ? t('Configure per-model price multipliers for {{username}}', {
              username: user.username,
            })
          : undefined
      }
      contentClassName='sm:max-w-xl'
      titleClassName='flex items-center gap-2'
      bodyClassName='space-y-3'
      footer={
        <>
          <Button
            type='button'
            variant='outline'
            onClick={() => onOpenChange(false)}
            disabled={saving}
          >
            {t('Cancel')}
          </Button>
          <Button type='button' onClick={handleSave} disabled={!user || saving}>
            {saving && <Loader2 className='h-4 w-4 animate-spin' />}
            {t('Save')}
          </Button>
        </>
      }
    >
      <div className='space-y-2'>
        <Label htmlFor='user-model-discounts'>
          {t('model_discounts JSON')}
        </Label>
        <Textarea
          id='user-model-discounts'
          rows={10}
          className='font-mono text-xs'
          value={value}
          onChange={(event) => setValue(event.target.value)}
          placeholder='{"gpt-4o": 8000}'
        />
        <p className='text-muted-foreground text-xs'>
          {t('Use 10000 as 1.0. For example, 8000 means 0.8x.')}
        </p>
      </div>
    </Dialog>
  )
}
