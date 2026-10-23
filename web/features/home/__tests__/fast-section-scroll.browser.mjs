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
import {
  openNativeScrollHome,
  freezeScrollClock,
  resetNativePosition,
  sectionCenter,
  moveNatively,
  finishNativeGesture,
  assertPosition,
} from './native-scroll-fixture.browser.mjs'

export async function verifyFastSectionPaging(page, url) {
  await openNativeScrollHome(page, url)
  await freezeScrollClock(page)
  const passed = []
  const models = await sectionCenter(page, '#models')
  const advantages = await sectionCenter(page, '#advantages')
  const integration = await sectionCenter(page, '#integration')

  await resetNativePosition(page)
  await moveNatively(page, [210, 700, 1100, Math.round(advantages - 70)])
  await finishNativeGesture(page)
  await assertPosition(page, await sectionCenter(page, '#advantages'))
  passed.push(
    'a fast swipe crosses multiple modules before aligning near its actual landing'
  )

  await resetNativePosition(page, advantages)
  await moveNatively(page, [1100, 800, Math.round(models + 50)])
  await finishNativeGesture(page)
  await assertPosition(page, await sectionCenter(page, '#models'))
  passed.push(
    'a fast upward swipe chooses the nearby landing module rather than a queued previous page'
  )

  await resetNativePosition(page)
  await moveNatively(page, [
    350,
    1000,
    1450,
    1300,
    Math.round(integration + 70),
  ])
  await finishNativeGesture(page)
  await assertPosition(page, await sectionCenter(page, '#integration'))
  passed.push(
    'a reversal inside one continuous gesture uses its final real position'
  )
  await page.clock.resume()
  return passed
}
