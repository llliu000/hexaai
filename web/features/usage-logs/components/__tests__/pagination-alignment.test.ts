import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'
import { test } from 'node:test'

import { Window } from 'happy-dom'

const usageStyles = readFileSync(
  new URL('../../../../styles/usage-logs-lanhu.css', import.meta.url),
  'utf8'
)
const taskStyles = readFileSync(
  new URL('../../../../styles/task-logs-lanhu.css', import.meta.url),
  'utf8'
)
const sharedStyles = readFileSync(
  new URL('../../../../styles/index.css', import.meta.url),
  'utf8'
)

for (const category of ['common', 'task']) {
  for (const width of [1440, 390]) {
    test(`${category} logs at ${width}px keep pagination flush with the table despite legacy shared spacing`, () => {
      const window = new Window({ width, height: 960 })
      try {
        const document = window.document
        const style = document.createElement('style')
        // Load shared styles last to guard against their older min-height and
        // content-padding overrides, not merely the presence of new rules.
        style.textContent = usageStyles + taskStyles + sharedStyles
        document.head.append(style)
        document.body.innerHTML = `
          <div class="profile-design-layout">
            <section data-slot="sidebar-inset"><div><main>
              <div>Heading</div>
              <div class="usage-logs-content">
                <div class="lujiyun-usage-page lujiyun-usage-${category}">
                  <div class="usage-logs-table-layout">
                    <div class="overflow-hidden rounded-lg border">Table</div>
                  </div>
                </div>
              </div>
              <div data-slot="page-footer"><div><div>
                <div>Count</div><div>Page size</div><div>Page navigation</div>
              </div></div></div>
            </main></div></section>
          </div>`
        const footer = document.querySelector('[data-slot="page-footer"]')
        const content = document.querySelector('.usage-logs-content')
        const table = document.querySelector(
          '.overflow-hidden.rounded-lg.border'
        )
        assert.ok(footer)
        assert.ok(content)
        assert.ok(table)
        const footerStyle = window.getComputedStyle(footer)
        const contentStyle = window.getComputedStyle(content)
        const tableStyle = window.getComputedStyle(table)
        assert.equal(footerStyle.marginLeft, '0px')
        assert.equal(footerStyle.marginRight, '0px')
        assert.equal(footerStyle.marginTop, '0px')
        assert.equal(footerStyle.marginBottom, '0px')
        assert.equal(contentStyle.paddingLeft, '0px')
        assert.equal(contentStyle.paddingRight, '0px')
        assert.equal(contentStyle.paddingBottom, '0px')
        assert.equal(contentStyle.flexGrow, '1')
        assert.equal(Number.parseFloat(contentStyle.minHeight), 0)
        assert.equal(Number.parseFloat(tableStyle.minHeight), 0)
        assert.equal(Number.parseFloat(tableStyle.borderBottomWidth), 0)
      } finally {
        window.close()
      }
    })
  }
}
