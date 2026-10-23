import { useTranslation } from 'react-i18next'

export function ModelChartEmpty() {
  const { t } = useTranslation()
  return (
    <div className='model-chart-empty' role='status'>
      <div className='model-empty-message'>
        <span className='models-art model-empty-art' aria-hidden='true' />
        <strong>{t('No data', { defaultValue: '暂无数据' })}</strong>
        <p>
          {t('No call data in the current time range.', {
            defaultValue: '当前时间范围内暂无调用数据',
          })}
        </p>
      </div>
    </div>
  )
}
