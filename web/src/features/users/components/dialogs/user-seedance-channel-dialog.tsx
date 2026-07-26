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
import { Loader2, Video } from 'lucide-react'
import { useEffect, useState } from 'react'
import { useTranslation } from 'react-i18next'
import { toast } from 'sonner'

import { Dialog } from '@/components/dialog'
import { Button } from '@/components/ui/button'
import { Label } from '@/components/ui/label'
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import type { Channel } from '@/features/channels/types'

import { getSeedanceChannels, updateUserSeedanceChannel } from '../../api'
import type { User } from '../../types'

interface Props {
  open: boolean
  onOpenChange: (open: boolean) => void
  user: User | null
  onSuccess?: () => void
}

function parseSeedanceChannelId(user: User | null): string {
  if (!user?.setting) return ''
  try {
    const setting = JSON.parse(user.setting) as {
      seedance_channel_id?: unknown
    }
    const channelId = Number(setting.seedance_channel_id)
    return Number.isInteger(channelId) && channelId > 0 ? String(channelId) : ''
  } catch {
    return ''
  }
}

export function UserSeedanceChannelDialog({
  open,
  onOpenChange,
  user,
  onSuccess,
}: Props) {
  const { t } = useTranslation()
  const [channels, setChannels] = useState<Channel[]>([])
  const [selectedChannelId, setSelectedChannelId] = useState('')
  const [loading, setLoading] = useState(false)
  const [saving, setSaving] = useState(false)

  useEffect(() => {
    if (!open) return
    setSelectedChannelId(parseSeedanceChannelId(user))
    setLoading(true)
    getSeedanceChannels()
      .then((result) => {
        if (result.success) {
          setChannels((result.data ?? []).filter((channel) => channel.status === 1))
        } else {
          toast.error(result.message || t('Failed to load Seedance channels'))
          setChannels([])
        }
      })
      .catch(() => {
        toast.error(t('Failed to load Seedance channels'))
        setChannels([])
      })
      .finally(() => {
        setLoading(false)
      })
  }, [open, t, user])

  const handleSave = async () => {
    if (!user) return
    const channelId = Number(selectedChannelId)
    if (!Number.isInteger(channelId) || channelId <= 0) {
      toast.error(t('Please select a Seedance channel'))
      return
    }

    setSaving(true)
    try {
      const result = await updateUserSeedanceChannel(user.id, channelId)
      if (result.success) {
        toast.success(t('Seedance channel updated'))
        onOpenChange(false)
        onSuccess?.()
      } else {
        toast.error(result.message || t('Failed to update Seedance channel'))
      }
    } catch {
      toast.error(t('Failed to update Seedance channel'))
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
          <Video className='h-5 w-5' />
          {t('Seedance Channel')}
        </>
      }
      description={
        user
          ? t('Configure Seedance channel for {{username}}', {
              username: user.username,
            })
          : undefined
      }
      contentClassName='sm:max-w-lg'
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
          <Button
            type='button'
            onClick={handleSave}
            disabled={!user || loading || saving || !selectedChannelId}
          >
            {saving && <Loader2 className='h-4 w-4 animate-spin' />}
            {t('Save')}
          </Button>
        </>
      }
    >
      <div className='space-y-2'>
        <Label>{t('Doubao Video Channel')}</Label>
        <Select
          value={selectedChannelId}
          onValueChange={(value) => setSelectedChannelId(value ?? '')}
          items={channels.map((channel) => ({
            value: String(channel.id),
            label: `#${channel.id} ${channel.name || t('Unnamed')}`,
          }))}
        >
          <SelectTrigger className='w-full'>
            <SelectValue
              placeholder={
                loading
                  ? t('Loading...')
                  : t('Select Seedance channel')
              }
            />
          </SelectTrigger>
          <SelectContent alignItemWithTrigger={false}>
            <SelectGroup>
              {channels.map((channel) => (
                <SelectItem key={channel.id} value={String(channel.id)}>
                  <span className='font-mono text-xs'>#{channel.id}</span>
                  <span>{channel.name || t('Unnamed')}</span>
                </SelectItem>
              ))}
            </SelectGroup>
          </SelectContent>
        </Select>
        <p className='text-muted-foreground text-sm'>
          {channels.length === 0 && !loading
            ? t('No enabled Doubao video channels available')
            : t('Select an enabled Doubao video channel for Seedance assets.')}
        </p>
      </div>
    </Dialog>
  )
}
