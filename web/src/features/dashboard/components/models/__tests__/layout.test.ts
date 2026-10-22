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
import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'
import { test } from 'node:test'

const styles = readFileSync(
  new URL('../../../../../styles/models-lanhu.css', import.meta.url),
  'utf8'
)
const chartTabs = readFileSync(
  new URL('../model-chart-tabs.tsx', import.meta.url),
  'utf8'
)
const dashboard = readFileSync(
  new URL('../../../index.tsx', import.meta.url),
  'utf8'
)
const performanceOverview = readFileSync(
  new URL('../performance-overview.tsx', import.meta.url),
  'utf8'
)
const consumptionChart = readFileSync(
  new URL('../consumption-distribution-chart.tsx', import.meta.url),
  'utf8'
)
const modelCharts = readFileSync(
  new URL('../model-charts.tsx', import.meta.url),
  'utf8'
)

test('model analytics fills the available dashboard width like flow', () => {
  const rootStyles = styles.match(
    /\.lujiyun-dashboard-models\.lujiyun-dashboard-models \{[\s\S]*?^}/m
  )?.[0]

  assert.ok(rootStyles)
  assert.match(rootStyles, /width: 100% !important/)
  assert.match(rootStyles, /max-width: none !important/)
  assert.match(rootStyles, /margin-inline: 0 !important/)
})

test('model chart tabs do not create a horizontal scrollbar', () => {
  assert.match(chartTabs, /TabsList className='[^']*overflow-visible/)
  assert.doesNotMatch(chartTabs, /TabsList className='[^']*overflow-x-auto/)
})

test('performance health heading matches the consumption heading scale', () => {
  assert.match(
    performanceOverview,
    /className='size-9 rounded-full \[&>svg\]:size-5'/
  )
  assert.match(
    performanceOverview,
    /className='text-\[13px\] font-semibold whitespace-nowrap'/
  )
})

test('performance health reuses the shared chart empty state when there is no data', () => {
  assert.match(performanceOverview, /import \{ ModelChartEmpty \}/)
  assert.match(performanceOverview, /const showEmpty = !loading && !hasData/)
  assert.match(performanceOverview, /showEmpty && \(\s*<div className='model-design-plot/)
  assert.match(performanceOverview, /<ModelChartEmpty \/>/)
  assert.doesNotMatch(performanceOverview, /No performance data available/)
})

test('model analysis panels share the same heading inset and icon gap', () => {
  for (const panel of [performanceOverview, consumptionChart, modelCharts]) {
    assert.match(panel, /px-\[15px\]/)
    assert.match(panel, /className='flex items-center gap-\[11px\]'/)
  }
})

test('both model charts expand to fill available desktop height without shrinking', () => {
  assert.match(
    dashboard,
    /contentClassName=\{activeSection === 'models' \? 'lg:flex lg:flex-col'/
  )
  assert.equal(
    (dashboard.match(/className='model-stretch-row'/g) ?? []).length,
    2
  )
  assert.match(
    styles,
    /\.model-stretch-row \{\s*flex: 1 1 238px;\s*min-height: 238px;/
  )
  const chartHeightRule = styles.match(
    /\.lujiyun-dashboard-models\.lujiyun-dashboard-models\s*> div:not\(:first-child\)\s*> \.model-design-chart \{[^}]*}/
  )?.[0]
  assert.ok(chartHeightRule)
  assert.match(chartHeightRule, /min-height: 238px/)
  assert.doesNotMatch(chartHeightRule, /(?:^|\s)height:/)
  assert.match(
    styles,
    /\.model-stretch-row\s*> \.model-design-chart \{\s*display: flex;\s*flex-direction: column;\s*height: 100%;/
  )
  assert.match(
    styles,
    /\.model-stretch-row\s*\.model-design-plot \{\s*flex: 1 1 auto;\s*height: auto !important;\s*min-height: 165px;/
  )
})
