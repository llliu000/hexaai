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
import { useQuery } from '@tanstack/react-query'
import { Link } from '@tanstack/react-router'
import {
  Activity,
  ArrowRight,
  Check,
  CircleAlert,
  CircleCheck,
  Code2,
  FileText,
  Info,
  KeyRound,
  Volume2,
  WalletCards,
} from 'lucide-react'
import { useMemo, useState } from 'react'

import { getUserQuotaDates, getUptimeStatus } from '@/features/dashboard/api'
import { useAnnouncements } from '@/features/dashboard/hooks/use-status-data'
import { getPreviewText } from '@/features/dashboard/lib'
import type { QuotaDataItem, UptimeMonitor } from '@/features/dashboard/types'
import { getApiKeys } from '@/features/keys/api'
import type { ApiKey } from '@/features/keys/types'
import { getUserLogs } from '@/features/usage-logs/api'
import { LOG_TYPE_ENUM } from '@/features/usage-logs/constants'
import type { UsageLog } from '@/features/usage-logs/data/schema'
import { getUserModels } from '@/lib/api'
import {
  formatNumber,
  formatQuota,
  formatTimestampRelative,
} from '@/lib/format'
import { computeTimeRange, formatDateTimeObject } from '@/lib/time'
import { useAuthStore } from '@/stores/auth-store'

function getPreferredKey(keys: ApiKey[]): ApiKey | null {
  return keys.find((item) => item.status === 1) ?? keys[0] ?? null
}

function useOverviewSetup() {
  const apiKeysQuery = useQuery({
    queryKey: ['dashboard', 'overview', 'api-keys'],
    queryFn: async () => {
      const result = await getApiKeys({ p: 1, size: 10 })
      return result.success ? (result.data?.items ?? []) : []
    },
    staleTime: 60 * 1000,
  })
  const modelsQuery = useQuery({
    queryKey: ['dashboard', 'overview', 'user-models'],
    queryFn: async () => {
      const result = await getUserModels()
      return result.success ? (result.data ?? []) : []
    },
    staleTime: 5 * 60 * 1000,
  })

  return {
    preferredKey: getPreferredKey(apiKeysQuery.data ?? []),
    model: modelsQuery.data?.[0] ?? '',
    loading: apiKeysQuery.isLoading || modelsQuery.isLoading,
  }
}

export function OverviewHeaderActions() {
  return (
    <div className='overview-header-actions'>
      <Link to='/keys' className='overview-header-action is-primary'>
        <KeyRound />
        <span>创建 API 密钥</span>
      </Link>
      <a
        href='https://docs.666api.ai/'
        target='_blank'
        rel='noreferrer'
        className='overview-header-action'
      >
        <FileText />
        <span>使用文档</span>
      </a>
      <Link to='/wallet' className='overview-header-action is-wallet'>
        <WalletCards />
        <span>充值</span>
      </Link>
    </div>
  )
}

export function OverviewSetupStatus() {
  const user = useAuthStore((state) => state.auth.user)
  const { preferredKey, loading } = useOverviewSetup()
  const balanceReady =
    Number(user?.quota ?? 0) > 0 || Number(user?.used_quota ?? 0) > 0
  const complete = Boolean(
    preferredKey && balanceReady && Number(user?.request_count ?? 0) > 0
  )
  let statusText = '初始化设置未完成'
  if (loading) statusText = '正在检查初始化设置'
  else if (complete) statusText = '初始化设置已完成'

  return (
    <span className={`overview-setup-ready${complete ? '' : ' is-pending'}`}>
      {complete ? <Check /> : <CircleAlert />}
      {statusText}
    </span>
  )
}

function buildChartPoints(data: QuotaDataItem[], width = 620, height = 104) {
  if (!data.length) return ''
  const sorted = [...data].sort((a, b) => a.created_at - b.created_at)
  const values = sorted.map((item) => Number(item.quota) || 0)
  const max = Math.max(...values, 1)
  return values
    .map((value, index) => {
      const x =
        sorted.length === 1 ? width / 2 : (index / (sorted.length - 1)) * width
      const y = height - (value / max) * (height - 8)
      return `${x.toFixed(1)},${y.toFixed(1)}`
    })
    .join(' ')
}

function UsageChart(props: { data: QuotaDataItem[] }) {
  const points = buildChartPoints(props.data)
  const areaPoints = points ? `0,104 ${points} 620,104` : ''
  const sorted = [...props.data].sort((a, b) => a.created_at - b.created_at)
  const maxQuota = Math.max(...sorted.map((item) => Number(item.quota) || 0), 0)
  const chartTicks = [
    { key: 'start', ratio: 0 },
    { key: 'quarter', ratio: 0.25 },
    { key: 'middle', ratio: 0.5 },
    { key: 'three-quarters', ratio: 0.75 },
    { key: 'end', ratio: 1 },
  ].map((tick) => ({
    ...tick,
    index: Math.min(
      sorted.length - 1,
      Math.max(0, Math.round((sorted.length - 1) * tick.ratio))
    ),
  }))
  const dateFormatter = new Intl.DateTimeFormat('zh-CN', {
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit',
    hour12: false,
  })

  return (
    <div
      className='overview-usage-chart'
      role='img'
      aria-label='真实用量趋势图'
    >
      <div className='overview-chart-y-labels'>
        <span>{formatQuota(maxQuota)}</span>
        <span>{formatQuota(maxQuota * 0.67)}</span>
        <span>{formatQuota(maxQuota * 0.33)}</span>
        <span>{formatQuota(0)}</span>
      </div>
      <svg viewBox='0 0 620 104' preserveAspectRatio='none' aria-hidden='true'>
        <defs>
          <linearGradient id='overviewUsageFill' x1='0' y1='0' x2='0' y2='1'>
            <stop offset='0%' stopColor='#2d70ff' stopOpacity='.2' />
            <stop offset='100%' stopColor='#2d70ff' stopOpacity='.01' />
          </linearGradient>
        </defs>
        {[8, 40, 72, 103].map((y) => (
          <line key={y} x1='0' x2='620' y1={y} y2={y} />
        ))}
        {areaPoints && (
          <polygon points={areaPoints} fill='url(#overviewUsageFill)' />
        )}
        {points && <polyline points={points} />}
      </svg>
      {!points && <span className='overview-chart-empty'>暂无用量数据</span>}
      <div className='overview-chart-x-labels'>
        {sorted.length ? (
          chartTicks.map((tick) => (
            <span key={tick.key}>
              {dateFormatter.format(
                new Date((sorted[tick.index]?.created_at ?? 0) * 1000)
              )}
            </span>
          ))
        ) : (
          <span>—</span>
        )}
      </div>
    </div>
  )
}

function UsageCard(props: {
  usedQuota: number
  requestCount: number
  periodDays: number
  setPeriodDays: (days: number) => void
  data: QuotaDataItem[]
  loading: boolean
}) {
  const periodUsage = props.data.reduce(
    (sum, item) => sum + (Number(item.quota) || 0),
    0
  )
  const periodRequests = props.data.reduce(
    (sum, item) => sum + (Number(item.count) || 0),
    0
  )
  const periodLabel =
    props.periodDays === 1 ? '近 24 小时消耗' : `近 ${props.periodDays} 天消耗`

  return (
    <section className='overview-card overview-usage-card'>
      <div className='overview-card-heading'>
        <h3>用量与请求趋势</h3>
        <div className='overview-period-tabs' aria-label='时间范围'>
          {[
            [1, '近24小时'],
            [7, '7天'],
            [30, '30天'],
          ].map(([days, label]) => (
            <button
              className={props.periodDays === days ? 'is-active' : ''}
              type='button'
              key={days}
              onClick={() => props.setPeriodDays(Number(days))}
            >
              {label}
            </button>
          ))}
          <button type='button' disabled>
            消耗金额
          </button>
        </div>
      </div>
      <div className='overview-usage-stats'>
        <div>
          <span>{periodLabel}</span>
          <strong>{props.loading ? '—' : formatQuota(periodUsage)}</strong>
          <small>来自实际用量记录</small>
        </div>
        <div>
          <span>历史总消耗</span>
          <strong>{formatQuota(props.usedQuota)}</strong>
          <small>账户累计数据</small>
        </div>
        <div>
          <span>请求次数</span>
          <strong>{formatNumber(props.requestCount)}</strong>
          <small>
            {props.loading
              ? '正在读取趋势'
              : `当前周期 ${formatNumber(periodRequests)} 次`}
          </small>
        </div>
      </div>
      <UsageChart data={props.data} />
    </section>
  )
}

function AssetCard(props: {
  balance: number
  recentUsage: number
  monthUsage: number
}) {
  return (
    <section className='overview-asset-card'>
      <div className='overview-asset-title'>
        <h3>账户资产</h3>
        <span>
          <i />
          正常
        </span>
      </div>
      <p>可用余额</p>
      <strong>{formatQuota(props.balance)}</strong>
      <div className='overview-asset-divider' />
      <div className='overview-asset-metrics'>
        <div>
          <span>近 24 小时消耗</span>
          <b>{formatQuota(props.recentUsage)}</b>
        </div>
        <div>
          <span>本月累计消耗</span>
          <b>{formatQuota(props.monthUsage)}</b>
        </div>
      </div>
      <div className='overview-asset-actions'>
        <Link to='/wallet'>充值账户</Link>
        <Link to='/wallet'>查看账单</Link>
      </div>
    </section>
  )
}

function ApiRouteCard(props: {
  preferredKey: ApiKey | null
  model: string
  loading: boolean
}) {
  const configured = Boolean(props.preferredKey && props.model)
  let routeStatus = '未配置'
  if (props.loading) routeStatus = '读取中'
  else if (configured) routeStatus = '已配置'
  return (
    <section className='overview-card overview-api-card'>
      <div className='overview-small-heading'>
        <span className='overview-round-icon overview-route-heading-icon'>
          <Code2 />
        </span>
        <h3>API 路由</h3>
        <span
          className={`overview-muted-pill${configured ? ' is-configured' : ''}`}
        >
          {routeStatus}
        </span>
      </div>
      <h4>{configured ? props.model : '尚未配置默认路由'}</h4>
      <p>
        {configured
          ? `当前使用密钥：${props.preferredKey?.name}`
          : '配置 API 密钥后，即可通过统一接口调用模型。'}
      </p>
      <Link to='/keys' className='overview-dark-button'>
        {configured ? '查看配置' : '立即配置'} <ArrowRight />
      </Link>
      <ApiRouteDecoration />
    </section>
  )
}

export function ApiRouteDecoration() {
  return (
    <span className='overview-route-decoration' aria-hidden='true'>
      <i className='overview-route-segment is-top'>
        <span className='overview-route-face is-top' />
        <span className='overview-route-face is-left' />
        <span className='overview-route-face is-right' />
      </i>
      <i className='overview-route-segment is-middle'>
        <span className='overview-route-face is-top' />
        <span className='overview-route-face is-left' />
        <span className='overview-route-face is-right' />
      </i>
      <i className='overview-route-segment is-bottom'>
        <span className='overview-route-face is-top' />
        <span className='overview-route-face is-left' />
        <span className='overview-route-face is-right' />
      </i>
    </span>
  )
}

function ServiceCard(props: { monitors: UptimeMonitor[]; loading: boolean }) {
  const visible = props.monitors.slice(0, 3)
  const allHealthy =
    visible.length > 0 && visible.every((monitor) => monitor.status === 1)
  const averageUptime = props.monitors.length
    ? props.monitors.reduce(
        (sum, monitor) => sum + (Number(monitor.uptime) || 0),
        0
      ) / props.monitors.length
    : null
  let statusText = '暂无监控数据'
  if (props.loading) statusText = '正在读取服务状态'
  else if (allHealthy) statusText = '所有服务运行正常'
  else if (visible.length) statusText = '部分服务状态异常'

  return (
    <section className='overview-card overview-service-card'>
      <div className='overview-small-heading'>
        <span className='overview-round-icon is-green'>
          <Activity />
        </span>
        <h3>服务状态</h3>
        <a href='https://docs.666api.ai/' target='_blank' rel='noreferrer'>
          查看详情 <ArrowRight />
        </a>
      </div>
      <div className={`overview-service-ok${allHealthy ? '' : ' is-warning'}`}>
        {allHealthy ? <CircleCheck /> : <CircleAlert />}
        {statusText}
      </div>
      <ul>
        {visible.map((monitor) => (
          <li key={`${monitor.group ?? ''}-${monitor.name}`}>
            <span>{monitor.name}</span>
            <b>
              <i className={monitor.status === 1 ? '' : 'is-error'} />
              {monitor.status === 1 ? '正常' : '异常'}
            </b>
          </li>
        ))}
        <li>
          <span>平均可用性</span>
          <strong>
            {averageUptime === null
              ? '—'
              : `${(averageUptime * 100).toFixed(2)}%`}
          </strong>
        </li>
      </ul>
    </section>
  )
}

function AnnouncementsCard() {
  const { items, loading } = useAnnouncements()
  const announcements = items.slice(0, 3)
  return (
    <section className='overview-card overview-announcements-card'>
      <div className='overview-small-heading'>
        <span className='overview-round-icon is-blue'>
          <Volume2 />
        </span>
        <h3>最新公告</h3>
        <a href='https://docs.666api.ai/' target='_blank' rel='noreferrer'>
          查看全部 <ArrowRight />
        </a>
      </div>
      <ul>
        {announcements.map((item, index) => (
          <li key={item.id ?? `${item.publishDate ?? ''}-${index}`}>
            <i />
            <span>{getPreviewText(item.content, 28)}</span>
            <time>
              {item.publishDate
                ? formatDateTimeObject(new Date(item.publishDate)).slice(0, 10)
                : '—'}
            </time>
          </li>
        ))}
        {!announcements.length && (
          <li className='overview-empty-row'>
            <span>{loading ? '正在读取公告' : '暂无公告'}</span>
          </li>
        )}
      </ul>
    </section>
  )
}

function getLogStatus(log: UsageLog) {
  if (log.type === LOG_TYPE_ENUM.CONSUME) {
    return { label: '成功', className: 'status-success' }
  }
  if (log.type === LOG_TYPE_ENUM.ERROR) {
    return { label: '失败', className: 'status-error' }
  }
  return { label: '其他', className: 'status-other' }
}

function RecentCallsCard(props: { logs: UsageLog[]; loading: boolean }) {
  return (
    <section className='overview-card overview-recent-card'>
      <div className='overview-table-title'>
        <h3>最近调用记录</h3>
        <Link to='/usage-logs'>
          查看全部 <ArrowRight />
        </Link>
      </div>
      <div className='overview-calls-table'>
        <div className='overview-calls-head'>
          <span>模型</span>
          <span>状态</span>
          <span>Tokens</span>
          <span>消耗</span>
          <span>时间</span>
        </div>
        {props.logs.slice(0, 4).map((log) => {
          const status = getLogStatus(log)
          const tokens =
            (Number(log.prompt_tokens) || 0) +
            (Number(log.completion_tokens) || 0)
          return (
            <div className='overview-call-row' key={log.id}>
              <span className='overview-model-cell'>
                <span className='overview-model-dot'>
                  <Code2 />
                </span>
                {log.model_name || '—'}
              </span>
              <span>
                <em className={`overview-status ${status.className}`}>
                  {status.label}
                </em>
              </span>
              <span>{formatNumber(tokens)}</span>
              <span>{formatQuota(Number(log.quota) || 0)}</span>
              <span>
                {formatTimestampRelative(log.created_at, 'seconds', 'zh-CN')}
              </span>
            </div>
          )
        })}
        {!props.logs.length && (
          <div className='overview-calls-empty'>
            {props.loading ? '正在读取调用记录' : '暂无调用记录'}
          </div>
        )}
      </div>
    </section>
  )
}

function ForecastCard(props: { monthUsage: number; balance: number }) {
  const now = new Date()
  const daysInMonth = new Date(
    now.getFullYear(),
    now.getMonth() + 1,
    0
  ).getDate()
  const elapsedDays = Math.max(1, now.getDate())
  const forecast =
    props.monthUsage > 0 ? (props.monthUsage / elapsedDays) * daysInMonth : 0
  const accountTotal = props.balance + props.monthUsage
  const usedPercent =
    accountTotal > 0
      ? Math.min(100, (props.monthUsage / accountTotal) * 100)
      : 0
  const forecastPercent =
    accountTotal > 0 ? Math.min(100, (forecast / accountTotal) * 100) : 0

  return (
    <section className='overview-card overview-forecast-card'>
      <div className='overview-table-title'>
        <h3>本月费用预测</h3>
        <Link to='/wallet'>
          设置预算提醒 <ArrowRight />
        </Link>
      </div>
      <div className='overview-forecast-body'>
        <div className='overview-forecast-total'>
          <span>预计月末消耗</span>
          <strong>{formatQuota(forecast)}</strong>
          <small>根据本月真实用量估算</small>
        </div>
        <div className='overview-forecast-bars'>
          <div className='overview-bar-row'>
            <span>本月已使用</span>
            <div>
              <small>{formatQuota(props.monthUsage)}</small>
              <i>
                <b style={{ width: `${usedPercent}%` }} />
              </i>
            </div>
            <em>{Math.round(usedPercent)}%</em>
          </div>
          <div className='overview-bar-row'>
            <span>本月预测消耗</span>
            <div>
              <small>{formatQuota(forecast)}</small>
              <i>
                <b style={{ width: `${forecastPercent}%` }} />
              </i>
            </div>
            <em>{Math.round(forecastPercent)}%</em>
          </div>
          <div className='overview-bar-row'>
            <span>预算上限</span>
            <div>
              <small>未设置</small>
              <i />
            </div>
            <em>—</em>
          </div>
        </div>
      </div>
      <div className='overview-budget-note'>
        <Info />
        预测值根据当前月份已产生的真实用量计算。
      </div>
    </section>
  )
}

export function OverviewDashboard() {
  const user = useAuthStore((state) => state.auth.user)
  const [periodDays, setPeriodDays] = useState(1)
  const { preferredKey, model, loading: setupLoading } = useOverviewSetup()
  const periodRange = useMemo(() => computeTimeRange(periodDays), [periodDays])
  const dailyRange = useMemo(() => computeTimeRange(1), [])
  const monthRange = useMemo(() => {
    const now = new Date()
    return {
      start_timestamp: Math.floor(
        new Date(now.getFullYear(), now.getMonth(), 1).getTime() / 1000
      ),
      end_timestamp: Math.floor(now.getTime() / 1000) + 3600,
    }
  }, [])

  const periodQuery = useQuery({
    queryKey: [
      'dashboard',
      'overview',
      'usage-period',
      periodDays,
      periodRange.start_timestamp,
      periodRange.end_timestamp,
    ],
    queryFn: () =>
      getUserQuotaDates({
        ...periodRange,
        default_time: periodDays === 1 ? 'hour' : 'day',
      }),
    staleTime: 60 * 1000,
  })
  const dailyQuery = useQuery({
    queryKey: [
      'dashboard',
      'overview',
      'usage-day',
      dailyRange.start_timestamp,
      dailyRange.end_timestamp,
    ],
    queryFn: () => getUserQuotaDates({ ...dailyRange, default_time: 'hour' }),
    staleTime: 60 * 1000,
  })
  const monthQuery = useQuery({
    queryKey: [
      'dashboard',
      'overview',
      'usage-month',
      monthRange.start_timestamp,
      monthRange.end_timestamp,
    ],
    queryFn: () => getUserQuotaDates({ ...monthRange, default_time: 'day' }),
    staleTime: 60 * 1000,
  })
  const uptimeQuery = useQuery({
    queryKey: ['dashboard', 'overview', 'uptime'],
    queryFn: getUptimeStatus,
    staleTime: 60 * 1000,
    retry: false,
  })
  const logsQuery = useQuery({
    queryKey: ['dashboard', 'overview', 'recent-logs'],
    queryFn: async () => {
      const result = await getUserLogs({ p: 1, page_size: 20 })
      return (result.data?.items ?? [])
        .filter((item): item is UsageLog => 'model_name' in item)
        .filter(
          (item) =>
            item.type === LOG_TYPE_ENUM.CONSUME ||
            item.type === LOG_TYPE_ENUM.ERROR
        )
        .slice(0, 4)
    },
    staleTime: 60 * 1000,
  })

  const periodData = periodQuery.data?.data ?? []
  const dailyUsage = (dailyQuery.data?.data ?? []).reduce(
    (sum, item) => sum + (Number(item.quota) || 0),
    0
  )
  const monthUsage = (monthQuery.data?.data ?? []).reduce(
    (sum, item) => sum + (Number(item.quota) || 0),
    0
  )
  const monitors = (uptimeQuery.data?.data ?? []).flatMap(
    (group) => group.monitors ?? []
  )
  const balance = Number(user?.quota ?? 0)
  const usedQuota = Number(user?.used_quota ?? 0)
  const requestCount = Number(user?.request_count ?? 0)

  return (
    <div className='lujiyun-overview-page'>
      <div className='overview-top-grid'>
        <UsageCard
          usedQuota={usedQuota}
          requestCount={requestCount}
          periodDays={periodDays}
          setPeriodDays={setPeriodDays}
          data={periodData}
          loading={periodQuery.isLoading}
        />
        <AssetCard
          balance={balance}
          recentUsage={dailyUsage}
          monthUsage={monthUsage}
        />
      </div>
      <div className='overview-middle-grid'>
        <ApiRouteCard
          preferredKey={preferredKey}
          model={model}
          loading={setupLoading}
        />
        <ServiceCard monitors={monitors} loading={uptimeQuery.isLoading} />
        <AnnouncementsCard />
      </div>
      <div className='overview-bottom-grid'>
        <RecentCallsCard
          logs={logsQuery.data ?? []}
          loading={logsQuery.isLoading}
        />
        <ForecastCard monthUsage={monthUsage} balance={balance} />
      </div>
    </div>
  )
}
