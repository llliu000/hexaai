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
import { after, describe, test } from 'node:test'

import { Window } from 'happy-dom'
import type React from 'react'

import {
  CONSUMPTION_DISTRIBUTION_CHART_OPTIONS,
  MODEL_ANALYTICS_CHART_OPTIONS,
} from '@/features/dashboard/constants'
import { processChartData } from '@/features/dashboard/lib'

const domWindow = new Window()
const domGlobals = [
  'window',
  'document',
  'navigator',
  'HTMLElement',
  'HTMLButtonElement',
  'SVGElement',
  'Node',
  'Element',
  'Event',
  'MouseEvent',
  'PointerEvent',
  'CustomEvent',
  'MutationObserver',
  'ResizeObserver',
  'requestAnimationFrame',
  'cancelAnimationFrame',
  'getComputedStyle',
] as const

for (const key of domGlobals) {
  Object.defineProperty(globalThis, key, {
    configurable: true,
    value: domWindow[key],
  })
}

const { act, useState } = await import('react')
const { createRoot } = await import('react-dom/client')
const { createInstance } = await import('i18next')
const { I18nextProvider, initReactI18next } = await import('react-i18next')
const { ModelChartTabs } = await import('../model-chart-tabs')

const i18n = createInstance()
await i18n.use(initReactI18next).init({ lng: 'en' })

const reactTestGlobals = globalThis as typeof globalThis & {
  IS_REACT_ACT_ENVIRONMENT?: boolean
}
reactTestGlobals.IS_REACT_ACT_ENVIRONMENT = true

interface ChartTabsHostProps<T extends string> {
  initialValue: T
  options: readonly { value: T; labelKey: string }[]
}

function ChartTabsHost<T extends string>(props: ChartTabsHostProps<T>) {
  const [value, setValue] = useState(props.initialValue)

  return (
    <I18nextProvider i18n={i18n}>
      <ModelChartTabs
        value={value}
        options={props.options}
        onValueChange={setValue}
      />
      <output data-testid='chart-type'>{value}</output>
    </I18nextProvider>
  )
}

async function renderTabs(element: React.ReactNode) {
  const container = document.createElement('div')
  document.body.append(container)
  const root = createRoot(container)
  await act(async () => root.render(element))
  return { container, root }
}

async function click(element: Element) {
  await act(async () => {
    ;(element as HTMLElement).click()
  })
}

describe('model chart selection', () => {
  test('chart selections map to different VChart specifications', () => {
    const chartData = processChartData([])

    assert.equal(chartData.spec_line.type, 'bar')
    assert.equal(chartData.spec_area.type, 'area')
    assert.equal(chartData.spec_model_line.type, 'area')
    assert.equal(chartData.spec_rank_bar.type, 'bar')
  })

  test('clicking area chart updates the consumption chart selection', async () => {
    const rendered = await renderTabs(
      <ChartTabsHost
        initialValue='bar'
        options={CONSUMPTION_DISTRIBUTION_CHART_OPTIONS}
      />
    )
    const tabs = rendered.container.querySelectorAll('[role="tab"]')

    await click(tabs[1])

    assert.equal(tabs[0].getAttribute('aria-selected'), 'false')
    assert.equal(tabs[1].getAttribute('aria-selected'), 'true')
    assert.equal(
      rendered.container.querySelector('[data-testid="chart-type"]')
        ?.textContent,
      'area'
    )
    await act(async () => rendered.root.unmount())
    rendered.container.remove()
  })

  test('clicking ranking updates the model analytics selection', async () => {
    const rendered = await renderTabs(
      <ChartTabsHost
        initialValue='trend'
        options={MODEL_ANALYTICS_CHART_OPTIONS}
      />
    )
    const tabs = rendered.container.querySelectorAll('[role="tab"]')

    await click(tabs[2])

    assert.equal(tabs[0].getAttribute('aria-selected'), 'false')
    assert.equal(tabs[2].getAttribute('aria-selected'), 'true')
    assert.equal(
      rendered.container.querySelector('[data-testid="chart-type"]')
        ?.textContent,
      'top'
    )
    await act(async () => rendered.root.unmount())
    rendered.container.remove()
  })
})

after(() => {
  reactTestGlobals.IS_REACT_ACT_ENVIRONMENT = false
  domWindow.close()
})
