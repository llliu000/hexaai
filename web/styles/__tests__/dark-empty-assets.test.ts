import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'
import { test } from 'node:test'

const modelStyles = readFileSync(
  new URL('../models-lanhu.css', import.meta.url),
  'utf8'
)
const flowStyles = readFileSync(
  new URL('../flow-lanhu.css', import.meta.url),
  'utf8'
)
const usageLogsTable = readFileSync(
  new URL(
    '../../features/usage-logs/components/usage-logs-table.tsx',
    import.meta.url
  ),
  'utf8'
)

test('dashboard model empty artwork swaps the light sprite for a transparent dark asset', () => {
  assert.match(
    modelStyles,
    /html\.dark \.model-empty-art \{[^}]*model-chart-empty-dark\.svg[^}]*background-size: 130px 110px/s
  )
})

test('flow empty artwork swaps the light reference crop for a transparent dark asset', () => {
  assert.match(
    flowStyles,
    /html\.dark \.flow-empty-art \{[^}]*flow-empty-dark\.svg[^}]*background-size: 148px 96px/s
  )
})

test('usage log empty states provide separate light and dark artwork', () => {
  for (const asset of ['usage-logs-empty', 'task-logs-empty']) {
    assert.match(
      usageLogsTable,
      new RegExp(`${asset}\\.png[\\s\\S]*?dark:hidden`)
    )
    assert.match(
      usageLogsTable,
      new RegExp(`${asset}-dark\\.svg[\\s\\S]*?hidden dark:block`)
    )
  }
})
