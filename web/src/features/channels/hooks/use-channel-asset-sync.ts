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
import { useMutation, useQueryClient } from '@tanstack/react-query'
import { useTranslation } from 'react-i18next'
import { toast } from 'sonner'

import { handleServerError } from '@/lib/handle-server-error'

import { syncChannelAssets } from '../api'
import type { Channel } from '../types'

export const channelAssetSyncKeys = {
  task: (channelId: number) => ['channel-asset-sync-task', channelId] as const,
  progress: (channelId: number, taskId?: string) =>
    ['channel-asset-sync-progress', channelId, taskId] as const,
}

export function supportsAssetSync(
  channel: Pick<Channel, 'type'>
): boolean {
  return channel.type === 54
}

export function useChannelAssetSync(channelId: number) {
  const { t } = useTranslation()
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: () => syncChannelAssets(channelId),
    onSuccess: async (response) => {
      if (!response.success || !response.data) {
        toast.error(response.message || t('Failed to start asset sync'))
        return
      }
      const taskKey = channelAssetSyncKeys.task(channelId)
      queryClient.setQueryDefaults(taskKey, { gcTime: Infinity })
      queryClient.setQueryData(taskKey, response.data)
      toast.success(t('Asset sync task submitted'))
      await queryClient.invalidateQueries({
        queryKey: ['channel-asset-sync-progress', channelId],
      })
    },
    onError: handleServerError,
  })
}
