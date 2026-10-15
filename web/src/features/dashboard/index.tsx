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
import { getRouteApi, useNavigate } from '@tanstack/react-router'
import { RotateCw } from 'lucide-react'
import { useState, useCallback, useMemo, lazy, Suspense } from 'react'

import '@/styles/models-lanhu.css'
import '@/styles/flow-lanhu.css'
import '@/styles/overview-lanhu.css'
import { useTranslation } from 'react-i18next'

import { SectionPageLayout } from '@/components/layout'
import { FadeIn } from '@/components/page-transition'
import { Button } from '@/components/ui/button'
import { Skeleton } from '@/components/ui/skeleton'
import { Tabs, TabsList, TabsTrigger } from '@/components/ui/tabs'
import { ROLE } from '@/lib/roles'
import { cn } from '@/lib/utils'
import { useAuthStore } from '@/stores/auth-store'

import { ModelsChartPreferences } from './components/models/models-chart-preferences'
import { ModelsFilter } from './components/models/models-filter-dialog'
import {
  OverviewDashboard,
  OverviewHeaderActions,
  OverviewSetupStatus,
} from './components/overview/overview-dashboard'
import { DEFAULT_TIME_GRANULARITY } from './constants'
import {
  buildDefaultDashboardFilters,
  getDefaultDays,
  getSavedChartPreferences,
  getSavedGranularity,
  saveChartPreferences,
} from './lib'
import {
  type DashboardSectionId,
  DASHBOARD_DEFAULT_SECTION,
  DASHBOARD_SECTION_IDS,
} from './section-registry'
import type {
  DashboardChartPreferences,
  DashboardFilters,
  QuotaDataItem,
  UserChartsFilters,
} from './types'

const route = getRouteApi('/_authenticated/dashboard/$section')

const LOG_STAT_CARD_FALLBACK_KEYS = [
  'count',
  'quota',
  'tokens',
  'average-rpm',
  'average-tpm',
] as const
const PERFORMANCE_METRIC_FALLBACK_KEYS = [
  'success-rate',
  'average-latency',
  'throughput',
] as const
const PERFORMANCE_MODEL_FALLBACK_KEYS = [
  'primary-model',
  'secondary-model',
] as const

const LazyLogStatCards = lazy(() =>
  import('./components/models/log-stat-cards').then((m) => ({
    default: m.LogStatCards,
  }))
)

const LazyModelCharts = lazy(() =>
  import('./components/models/model-charts').then((m) => ({
    default: m.ModelCharts,
  }))
)

const LazyConsumptionDistributionChart = lazy(() =>
  import('./components/models/consumption-distribution-chart').then((m) => ({
    default: m.ConsumptionDistributionChart,
  }))
)

const LazyPerformanceOverview = lazy(() =>
  import('./components/models/performance-overview').then((m) => ({
    default: m.PerformanceOverview,
  }))
)

const LazyUserCharts = lazy(() =>
  import('./components/users/user-charts').then((m) => ({
    default: m.UserCharts,
  }))
)

const LazyFlowCharts = lazy(() =>
  import('./components/flow/flow-charts').then((m) => ({
    default: m.FlowCharts,
  }))
)

function LogStatCardsFallback() {
  return (
    <div className='overflow-hidden rounded-lg border'>
      <div className='divide-border/60 grid grid-cols-2 divide-x sm:grid-cols-3 lg:grid-cols-5'>
        {LOG_STAT_CARD_FALLBACK_KEYS.map((key, index) => (
          <div
            key={key}
            className={cn(
              'px-2.5 py-1.5 sm:px-5 sm:py-4',
              index === LOG_STAT_CARD_FALLBACK_KEYS.length - 1 &&
                'col-span-2 sm:col-span-1'
            )}
          >
            <div className='flex items-center gap-1.5 sm:gap-2'>
              <Skeleton className='size-4 rounded-sm sm:size-7 sm:rounded-md' />
              <Skeleton className='h-4 w-16' />
            </div>
            <Skeleton className='mt-1 h-5 w-16 sm:mt-2 sm:h-7 sm:w-20' />
            <Skeleton className='mt-1 hidden h-3.5 w-28 md:block' />
          </div>
        ))}
      </div>
    </div>
  )
}

function ModelChartsFallback() {
  return (
    <div className='model-design-chart flex flex-col overflow-hidden rounded-lg border'>
      <div className='flex items-center justify-between border-b px-4 py-3 sm:px-5'>
        <Skeleton className='h-5 w-32' />
        <Skeleton className='h-8 w-72' />
      </div>
      <div className='model-design-plot min-h-0 flex-1 p-2'>
        <Skeleton className='h-full w-full' />
      </div>
    </div>
  )
}

function PerformanceOverviewFallback() {
  return (
    <div className='overflow-hidden rounded-lg border'>
      <div className='flex flex-wrap items-center gap-x-6 gap-y-2 px-4 py-3 sm:px-5'>
        <div className='flex items-center gap-2'>
          <Skeleton className='h-4 w-24' />
        </div>
        {PERFORMANCE_METRIC_FALLBACK_KEYS.map((key) => (
          <div key={key} className='flex items-center gap-1.5'>
            <Skeleton className='h-3 w-14' />
            <Skeleton className='h-4 w-16' />
          </div>
        ))}
        <div className='ml-auto flex items-center gap-2'>
          {PERFORMANCE_MODEL_FALLBACK_KEYS.map((key) => (
            <Skeleton key={key} className='h-5 w-28 rounded-full' />
          ))}
        </div>
      </div>
    </div>
  )
}

const SECTION_META: Record<DashboardSectionId, { titleKey: string }> = {
  overview: {
    titleKey: 'Overview',
  },
  models: {
    titleKey: 'Model Call Analytics',
  },
  flow: {
    titleKey: 'Flow',
  },
  users: {
    titleKey: 'User Analytics',
  },
}

const SECTION_DESCRIPTION_KEYS: Partial<Record<DashboardSectionId, string>> = {
  overview: 'View account usage, API status, and recent activity.',
  flow: 'Review model traffic flows to optimize resource allocation and cost.',
}

export function Dashboard() {
  const { t } = useTranslation()
  const navigate = useNavigate()
  const params = route.useParams()
  const userRole = useAuthStore((state) => state.auth.user?.role)
  const activeSection = (params.section ??
    DASHBOARD_DEFAULT_SECTION) as DashboardSectionId
  const activeSectionMeta =
    SECTION_META[activeSection] ?? SECTION_META[DASHBOARD_DEFAULT_SECTION]

  const [modelData, setModelData] = useState<QuotaDataItem[]>([])
  const [dataLoading, setDataLoading] = useState(false)
  const [chartPreferences, setChartPreferences] =
    useState<DashboardChartPreferences>(() => getSavedChartPreferences())
  const [modelFilters, setModelFilters] = useState<DashboardFilters>(() =>
    buildDefaultDashboardFilters(getSavedChartPreferences())
  )
  const [userChartsFilters, setUserChartsFilters] = useState<UserChartsFilters>(
    () => {
      const granularity = getSavedGranularity()
      return {
        timeGranularity: granularity,
        selectedRange: getDefaultDays(granularity),
        topUserLimit: 10,
      }
    }
  )

  const handleFilterChange = useCallback((filters: DashboardFilters) => {
    setModelFilters(filters)
  }, [])

  const handleResetFilters = useCallback(() => {
    setModelFilters(buildDefaultDashboardFilters(chartPreferences))
  }, [chartPreferences])

  const handleDataUpdate = useCallback(
    (data: QuotaDataItem[], loading: boolean) => {
      setModelData(data)
      setDataLoading(loading)
    },
    []
  )

  const handleChartPreferencesChange = useCallback(
    (preferences: DashboardChartPreferences) => {
      setChartPreferences(preferences)
      setModelFilters(buildDefaultDashboardFilters(preferences))
      saveChartPreferences(preferences)
    },
    []
  )

  const isAdmin = Boolean(userRole && userRole >= ROLE.ADMIN)
  const visibleSections = useMemo(
    () =>
      DASHBOARD_SECTION_IDS.filter(
        (section) => section !== 'overview' && (section !== 'users' || isAdmin)
      ),
    [isAdmin]
  )
  const handleSectionChange = useCallback(
    (section: string) => {
      void navigate({
        to: '/dashboard/$section',
        params: { section: section as DashboardSectionId },
      })
    },
    [navigate]
  )
  const showSectionTabs =
    activeSection !== 'overview' && visibleSections.length > 1
  const modelActions =
    activeSection === 'models' ? (
      <>
        <ModelsChartPreferences
          preferences={chartPreferences}
          onPreferencesChange={handleChartPreferencesChange}
        />
        <ModelsFilter
          preferences={chartPreferences}
          currentFilters={modelFilters}
          onFilterChange={handleFilterChange}
          onReset={handleResetFilters}
        />
        <Button
          variant='outline'
          size='sm'
          disabled={dataLoading}
          onClick={() => setModelFilters((previous) => ({ ...previous }))}
        >
          <RotateCw
            size={13}
            className={dataLoading ? 'animate-spin' : undefined}
          />
          {t('Refresh')}
        </Button>
      </>
    ) : null
  const flowFilterAction = (
    <ModelsFilter
      preferences={chartPreferences}
      currentFilters={modelFilters}
      onFilterChange={handleFilterChange}
      onReset={handleResetFilters}
      titleKey='Flow Filters'
      descriptionKey='Filter the traffic flow view by time range and user.'
    />
  )
  const sectionActions = modelActions

  return (
    <SectionPageLayout
      contentClassName={activeSection === 'models' ? 'lg:flex lg:flex-col' : undefined}
    >
      <SectionPageLayout.Title>
        {t(activeSectionMeta.titleKey)}
      </SectionPageLayout.Title>
      {SECTION_DESCRIPTION_KEYS[activeSection] && (
        <SectionPageLayout.Description>
          {t(SECTION_DESCRIPTION_KEYS[activeSection])}
          {activeSection === 'overview' && <OverviewSetupStatus />}
        </SectionPageLayout.Description>
      )}
      {activeSection === 'overview' && (
        <SectionPageLayout.Actions>
          <OverviewHeaderActions />
        </SectionPageLayout.Actions>
      )}
      <SectionPageLayout.Content>
        <div
          className={cn(
            'lujiyun-dashboard-page space-y-3 sm:space-y-4',
            `lujiyun-dashboard-${activeSection}`,
            activeSection === 'models' && 'lg:flex-1'
          )}
        >
          {activeSection !== 'overview' && (
            <div className='dashboard-section-bar flex flex-wrap items-center justify-between gap-1.5 sm:gap-2'>
              {showSectionTabs ? (
                <Tabs
                  className='dashboard-section-tabs w-full sm:w-auto'
                  value={activeSection}
                  onValueChange={handleSectionChange}
                >
                  <TabsList
                    className={cn(
                      'grid w-full items-stretch group-data-horizontal/tabs:h-auto',
                      isAdmin ? 'grid-cols-3' : 'grid-cols-2'
                    )}
                  >
                    {visibleSections.map((section) => (
                      <TabsTrigger key={section} value={section}>
                        {t(SECTION_META[section].titleKey)}
                      </TabsTrigger>
                    ))}
                  </TabsList>
                </Tabs>
              ) : (
                <div />
              )}
              {sectionActions != null && (
                <div className='flex shrink-0 flex-wrap items-center gap-1.5 sm:gap-2'>
                  {sectionActions}
                </div>
              )}
            </div>
          )}
          {activeSection === 'overview' && <OverviewDashboard />}
          {activeSection === 'models' && (
            <>
              <FadeIn>
                <Suspense fallback={<LogStatCardsFallback />}>
                  <LazyLogStatCards
                    filters={modelFilters}
                    onDataUpdate={handleDataUpdate}
                  />
                </Suspense>
              </FadeIn>
              {isAdmin && (
                <FadeIn delay={0.05}>
                  <Suspense fallback={<PerformanceOverviewFallback />}>
                    <LazyPerformanceOverview />
                  </Suspense>
                </FadeIn>
              )}
              <FadeIn delay={0.1} className='model-stretch-row'>
                <Suspense fallback={<ModelChartsFallback />}>
                  <LazyConsumptionDistributionChart
                    data={modelData}
                    loading={dataLoading}
                    defaultChartType={
                      chartPreferences.consumptionDistributionChart
                    }
                    timeGranularity={
                      modelFilters.time_granularity || DEFAULT_TIME_GRANULARITY
                    }
                  />
                </Suspense>
              </FadeIn>
              <FadeIn delay={0.15} className='model-stretch-row'>
                <Suspense fallback={<ModelChartsFallback />}>
                  <LazyModelCharts
                    data={modelData}
                    loading={dataLoading}
                    defaultChartTab={chartPreferences.modelAnalyticsChart}
                    timeGranularity={
                      modelFilters.time_granularity || DEFAULT_TIME_GRANULARITY
                    }
                  />
                </Suspense>
              </FadeIn>
            </>
          )}
          {activeSection === 'users' && (
            <FadeIn>
              <Suspense fallback={<ModelChartsFallback />}>
                <LazyUserCharts
                  filters={userChartsFilters}
                  onFiltersChange={setUserChartsFilters}
                />
              </Suspense>
            </FadeIn>
          )}
          {activeSection === 'flow' && (
            <FadeIn>
              <Suspense fallback={<ModelChartsFallback />}>
                <LazyFlowCharts
                  filters={modelFilters}
                  filterAction={flowFilterAction}
                />
              </Suspense>
            </FadeIn>
          )}
        </div>
      </SectionPageLayout.Content>
    </SectionPageLayout>
  )
}
