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
import { useQuery, useQueryClient } from '@tanstack/react-query'
import { Loader2, RefreshCw } from 'lucide-react'
import { useTranslation } from 'react-i18next'

import { Dialog } from '@/components/dialog'
import { Alert, AlertDescription, AlertTitle } from '@/components/ui/alert'
import { Button } from '@/components/ui/button'
import { Progress } from '@/components/ui/progress'

import { getChannelAssetSyncProgress } from '../../api'
import { channelAssetSyncKeys } from '../../hooks/use-channel-asset-sync'
import type { AssetSyncTask, Channel } from '../../types'

interface AssetSyncDetailsDialogProps {
  channel: Channel
  open: boolean
  onOpenChange: (open: boolean) => void
}

export function AssetSyncDetailsDialog(props: AssetSyncDetailsDialogProps) {
  const { t } = useTranslation()
  const queryClient = useQueryClient()
  const task = queryClient.getQueryData<AssetSyncTask>(
    channelAssetSyncKeys.task(props.channel.id)
  )
  const progress = useQuery({
    queryKey: channelAssetSyncKeys.progress(props.channel.id, task?.task_id),
    queryFn: async ({ signal }) => {
      const response = await getChannelAssetSyncProgress(
        props.channel.id,
        task?.task_id,
        signal
      )
      if (!response.success || !response.data) {
        throw new Error(
          response.message || t('Failed to load asset sync details')
        )
      }
      return response.data
    },
    enabled: props.open,
    staleTime: 0,
    retry: false,
    refetchInterval: (query) => {
      const data = query.state.data
      if (query.state.error || !data) return false
      return data.preparing || data.pending > 0 || data.processing > 0
        ? 5000
        : false
    },
  })
  const data = progress.data
  const completed = data ? data.active + data.failed : 0
  const percentage = data && data.total > 0 ? (completed / data.total) * 100 : 0

  return (
    <Dialog
      open={props.open}
      onOpenChange={props.onOpenChange}
      title={t('Asset Sync Details')}
      description={t('Manual asset sync results for {{name}}', {
        name: props.channel.name,
      })}
      descriptionClassName='break-all'
      contentClassName='sm:max-w-lg'
      bodyClassName='space-y-4'
      footer={
        <>
          <Button
            variant='outline'
            disabled={progress.isFetching}
            onClick={() => void progress.refetch()}
          >
            {progress.isFetching ? (
              <Loader2 className='size-4 animate-spin' aria-hidden='true' />
            ) : (
              <RefreshCw className='size-4' aria-hidden='true' />
            )}
            {t('Refresh')}
          </Button>
          <Button variant='outline' onClick={() => props.onOpenChange(false)}>
            {t('Close')}
          </Button>
        </>
      }
    >
      {progress.isPending && (
        <div
          role='status'
          className='text-muted-foreground flex items-center gap-2 py-6'
        >
          <Loader2 className='size-4 animate-spin' aria-hidden='true' />
          {t('Loading...')}
        </div>
      )}
      {progress.isError && (
        <Alert variant='destructive'>
          <AlertTitle>{t('Failed to load asset sync details')}</AlertTitle>
          <AlertDescription className='break-all'>
            {progress.error.message}
          </AlertDescription>
        </Alert>
      )}
      {data && (
        <>
          {data.preparing && (
            <p
              role='status'
              className='text-muted-foreground flex items-center gap-2 text-sm'
            >
              <Loader2 className='size-4 animate-spin' aria-hidden='true' />
              {t('Preparing assets for sync...')}
            </p>
          )}
          {data.prepare_task?.status === 'failed' && (
            <Alert variant='destructive'>
              <AlertTitle>{t('Asset sync preparation failed')}</AlertTitle>
              <AlertDescription className='break-all'>
                {data.prepare_task.error}
              </AlertDescription>
            </Alert>
          )}
          <dl className='grid grid-cols-2 gap-3'>
            <div className='bg-muted/50 col-span-2 rounded-lg border p-4'>
              <dt className='text-muted-foreground text-sm'>{t('Total')}</dt>
              <dd className='mt-1 text-2xl font-semibold tabular-nums'>
                {data.total}
              </dd>
            </div>
            {[
              { label: t('Queued for upload'), count: data.pending },
              { label: t('Processing'), count: data.processing },
              { label: t('Synced'), count: data.active },
              { label: t('Failed'), count: data.failed },
            ].map((item) => (
              <div key={item.label} className='rounded-lg border p-3'>
                <dt className='text-muted-foreground text-sm'>{item.label}</dt>
                <dd className='mt-1 text-xl font-semibold tabular-nums'>
                  {item.count}
                </dd>
              </div>
            ))}
          </dl>
          <Progress value={percentage} aria-label={t('Asset sync progress')} />
          {!data.preparing && data.total === 0 && (
            <p className='text-muted-foreground text-sm'>
              {t('No assets queued for sync.')}
            </p>
          )}
        </>
      )}
    </Dialog>
  )
}
