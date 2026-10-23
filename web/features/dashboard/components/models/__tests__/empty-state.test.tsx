import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'
import { test } from 'node:test'

import { Window } from 'happy-dom'
import { createInstance } from 'i18next'
import { renderToStaticMarkup } from 'react-dom/server'
import { I18nextProvider } from 'react-i18next'

import { ModelChartEmpty } from '../model-chart-empty'
import { ModelChartState } from '../model-chart-state'

const i18n = createInstance()
await i18n.init({
  lng: 'en',
  resources: { en: { translation: { 'No data': 'No data' } } },
})

function emptyMarkup() {
  return renderToStaticMarkup(
    <I18nextProvider i18n={i18n}>
      <ModelChartEmpty />
    </I18nextProvider>
  )
}

test('an empty model chart shows its existing message and artwork without simulated chart axes', () => {
  const window = new Window()
  try {
    window.document.body.innerHTML = emptyMarkup()
    const status = window.document.querySelector('[role="status"]')
    assert.ok(status)
    assert.ok(status.querySelector('.model-empty-art'))
    assert.match(status.textContent || '', /No data/)
    assert.equal(
      Boolean(status.querySelector('.model-empty-grid')),
      false,
      'empty state must not draw grid lines'
    )
    assert.equal(status.querySelector('canvas, svg'), null)
  } finally {
    window.close()
  }
})

test('empty chart artwork is larger while preserving the original icon proportions', () => {
  const window = new Window()
  try {
    const style = window.document.createElement('style')
    style.textContent = readFileSync(
      new URL('../../../../../styles/models-lanhu.css', import.meta.url),
      'utf8'
    )
    window.document.head.append(style)
    window.document.body.innerHTML = emptyMarkup()
    const icon = window.document.querySelector('.model-empty-art')
    assert.ok(icon)
    const computed = window.getComputedStyle(icon)
    const width = parseFloat(computed.width)
    const height = parseFloat(computed.height)
    assert.ok(width > 104 && height > 88)
    assert.ok(Math.abs(width / height - 104 / 88) < 0.001)
  } finally {
    window.close()
  }
})

for (const loading of [true, false]) {
  test(`${loading ? 'loading' : 'empty'} chart state never mounts the chart renderer`, () => {
    function ChartRenderer(): never {
      throw new Error(
        'The chart renderer must not mount without available data'
      )
    }
    const markup = renderToStaticMarkup(
      <I18nextProvider i18n={i18n}>
        <ModelChartState loading={loading} empty>
          <ChartRenderer />
        </ModelChartState>
      </I18nextProvider>
    )
    if (loading) {
      assert.match(markup, /aria-label="Loading..."/)
      assert.doesNotMatch(markup, /model-empty-message/)
    } else assert.match(markup, /No data/)
  })
}

test('nonempty chart state renders the supplied graph without an empty-state message', () => {
  const markup = renderToStaticMarkup(
    <I18nextProvider i18n={i18n}>
      <ModelChartState empty={false}>
        <canvas aria-label='Usage chart' />
      </ModelChartState>
    </I18nextProvider>
  )
  assert.match(markup, /aria-label="Usage chart"/)
  assert.doesNotMatch(markup, /model-empty-message/)
})
