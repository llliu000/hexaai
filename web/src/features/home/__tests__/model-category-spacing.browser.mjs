import assert from 'node:assert/strict'

// Run at different vertical scroll positions, viewport sizes and categories.
// This checks rendered geometry, including the active parallax transforms.
export async function verifyModelCategorySpacing(page) {
  const layout = await page.evaluate(() => {
    const toolbar = document.querySelector('.wm-model-toolbar')
    const track = document.querySelector('.wm-model-track')
    const card = track?.querySelector('.wm-model-card')
    const selected = toolbar?.querySelector('[aria-pressed="true"]')
    if (!toolbar || !track || !card || !selected) return null
    const toolbarRect = toolbar.getBoundingClientRect()
    const cardRect = card.getBoundingClientRect()
    return {
      gap: cardRect.top - toolbarRect.bottom,
      expectedGap: getComputedStyle(track).paddingTop,
      toolbarVisible: toolbarRect.bottom > 0 && toolbarRect.top < innerHeight,
    }
  })
  assert.ok(layout, 'Model categories and cards must be rendered')
  assert.ok(layout.toolbarVisible, 'Scroll the category toolbar into view')
  const expectedGap = Number.parseFloat(layout.expectedGap)
  assert.ok(expectedGap > 0, 'Model cards need spacing below categories')
  assert.ok(
    layout.gap >= expectedGap - 1,
    `Scrolling must preserve the category/card gap: ${layout.gap}px instead of ${expectedGap}px`
  )
  return layout
}
