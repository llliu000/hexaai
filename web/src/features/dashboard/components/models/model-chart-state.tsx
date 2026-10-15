import { Loader2 } from 'lucide-react'
import type { ReactNode } from 'react'
import { useTranslation } from 'react-i18next'

import { ModelChartEmpty } from './model-chart-empty'

export function ModelChartState(props: {
  loading?: boolean
  empty: boolean
  children: ReactNode
}) {
  const { t } = useTranslation()
  if (props.loading) {
    return (
      <div
        className='model-chart-loading'
        role='status'
        aria-label={t('Loading...')}
      >
        <Loader2 className='size-6 animate-spin' aria-hidden='true' />
      </div>
    )
  }
  if (props.empty) return <ModelChartEmpty />
  return <>{props.children}</>
}
