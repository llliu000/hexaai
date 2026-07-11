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
import { Percent, Loader2, Plus, Trash2 } from 'lucide-react'
import { useEffect, useState } from 'react'
import { useTranslation } from 'react-i18next'
import { toast } from 'sonner'

import { Dialog } from '@/components/dialog'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'

import { updateUserModelDiscounts } from '../../api'
import type { User, UserModelDiscounts } from '../../types'

interface Props {
  open: boolean
  onOpenChange: (open: boolean) => void
  user: User | null
  onSuccess?: () => void
}

type DiscountRow = {
  id: string
  model: string
  discount: string
}

function createRow(model = '', discount = '10000'): DiscountRow {
  return {
    id: `${Date.now()}-${Math.random()}`,
    model,
    discount,
  }
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

function modelDiscountsToRows(user: User | null): DiscountRow[] {
  return Object.entries(parseUserModelDiscounts(user)).map(
    ([model, discount]) => createRow(model, String(discount))
  )
}

function rowsToModelDiscounts(rows: DiscountRow[]): UserModelDiscounts | null {
  const normalized: UserModelDiscounts = {}
  for (const row of rows) {
    const model = row.model.trim()
    const discount = Number(row.discount)
    if (!model || !Number.isInteger(discount) || discount <= 0) {
      return null
    }
    if (normalized[model] !== undefined) {
      return null
    }
    normalized[model] = discount
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
  const [rows, setRows] = useState<DiscountRow[]>([])
  const [saving, setSaving] = useState(false)

  useEffect(() => {
    if (open) {
      setRows(modelDiscountsToRows(user))
    }
  }, [open, user])

  const updateRow = (
    id: string,
    field: 'model' | 'discount',
    value: string
  ) => {
    setRows((currentRows) =>
      currentRows.map((row) =>
        row.id === id ? { ...row, [field]: value } : row
      )
    )
  }

  const addRow = () => {
    setRows((currentRows) => [...currentRows, createRow()])
  }

  const removeRow = (id: string) => {
    setRows((currentRows) => currentRows.filter((row) => row.id !== id))
  }

  const handleSave = async () => {
    if (!user) return
    const modelDiscounts = rowsToModelDiscounts(rows)
    if (!modelDiscounts) {
      toast.error(
        t(
          'Please enter unique model names. Discount values must be positive integers.'
        )
      )
      return
    }

    setSaving(true)
    try {
      const result = await updateUserModelDiscounts(user.id, modelDiscounts)
      if (result.success) {
        toast.success(t('Discount settings updated'))
        onOpenChange(false)
        onSuccess?.()
      } else {
        toast.error(result.message || t('Failed to update discount settings'))
      }
    } catch {
      toast.error(t('Failed to update discount settings'))
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
          {t('Discount Settings')}
        </>
      }
      description={
        user
          ? t('Configure per-model discounts for {{username}}', {
              username: user.username,
            })
          : undefined
      }
      contentClassName='sm:max-w-2xl'
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
      <div className='space-y-3'>
        <div className='flex items-center justify-between gap-3'>
          <p className='text-muted-foreground text-sm'>
            {t('10000 means 1.0. For example, 8000 means 0.8x price.')}
          </p>
          <Button type='button' size='sm' onClick={addRow}>
            <Plus className='h-4 w-4' />
            {t('Add')}
          </Button>
        </div>

        {rows.length === 0 ? (
          <div className='text-muted-foreground rounded-md border border-dashed p-6 text-center text-sm'>
            {t('No discount settings configured')}
          </div>
        ) : (
          <div className='space-y-2'>
            <div className='text-muted-foreground hidden grid-cols-[1fr_9rem_2rem] gap-2 px-1 text-xs sm:grid'>
              <span>{t('Model Name')}</span>
              <span>{t('Discount Value')}</span>
              <span />
            </div>
            {rows.map((row) => (
              <div
                key={row.id}
                className='grid gap-2 rounded-md border p-2 sm:grid-cols-[1fr_9rem_2rem] sm:border-0 sm:p-0'
              >
                <Input
                  value={row.model}
                  onChange={(event) =>
                    updateRow(row.id, 'model', event.target.value)
                  }
                  placeholder='gpt-4o'
                />
                <Input
                  type='number'
                  min='1'
                  step='1'
                  value={row.discount}
                  onChange={(event) =>
                    updateRow(row.id, 'discount', event.target.value)
                  }
                  placeholder='10000'
                />
                <Button
                  type='button'
                  variant='ghost'
                  size='icon-sm'
                  className='text-destructive hover:text-destructive justify-self-end'
                  onClick={() => removeRow(row.id)}
                  aria-label={t('Delete')}
                >
                  <Trash2 className='h-4 w-4' />
                </Button>
              </div>
            ))}
          </div>
        )}
      </div>
    </Dialog>
  )
}
