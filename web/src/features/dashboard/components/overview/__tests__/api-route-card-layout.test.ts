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

const component = readFileSync(
  new URL('../overview-dashboard.tsx', import.meta.url),
  'utf8'
)
const styles = readFileSync(
  new URL('../../../../../styles/overview-lanhu.css', import.meta.url),
  'utf8'
)

test('API route decoration uses three faceted cubes instead of a code icon', () => {
  const decoration = component.match(
    /export function ApiRouteDecoration\(\)[\s\S]*?^}/m
  )?.[0]

  assert.ok(decoration)
  assert.match(decoration, /aria-hidden='true'/)
  assert.equal((decoration.match(/overview-route-segment/g) ?? []).length, 3)
  assert.equal(
    (decoration.match(/overview-route-face is-top/g) ?? []).length,
    3
  )
  assert.equal(
    (decoration.match(/overview-route-face is-left/g) ?? []).length,
    3
  )
  assert.equal(
    (decoration.match(/overview-route-face is-right/g) ?? []).length,
    3
  )
  assert.doesNotMatch(decoration, /Code2|<svg/)
})

test('overview header actions keep their original widths and match the recharge account height', () => {
  const headerActionStyles = styles.match(
    /\.overview-header-action \{[\s\S]*?^}/m
  )?.[0]
  const assetActionStyles = styles.match(
    /\.overview-asset-actions a \{[\s\S]*?^}/m
  )?.[0]

  assert.ok(headerActionStyles)
  assert.ok(assetActionStyles)
  assert.match(headerActionStyles, /height: 46px/)
  assert.match(assetActionStyles, /height: 46px/)
  assert.match(
    styles,
    /\.overview-header-action\.is-primary \{[^}]*min-width: 157px/s
  )
  assert.match(
    styles,
    /\.overview-header-action\.is-wallet \{[^}]*min-width: 125px/s
  )
})

test('API route decoration keeps the design dimensions and cannot intercept input', () => {
  const decorationStyles = styles.match(
    /\.overview-route-decoration \{[\s\S]*?^}/m
  )?.[0]

  assert.ok(decorationStyles)
  assert.match(decorationStyles, /right: 20px/)
  assert.match(decorationStyles, /bottom: 10px/)
  assert.match(decorationStyles, /width: 93px/)
  assert.match(decorationStyles, /height: 111px/)
  assert.match(decorationStyles, /pointer-events: none/)
})

test('API route decoration segments overlap into the compact reference silhouette', () => {
  const segmentStyles = styles.match(
    /\.overview-route-segment \{[\s\S]*?^}/m
  )?.[0]
  const middleStyles = styles.match(
    /\.overview-route-segment\.is-middle \{[\s\S]*?^}/m
  )?.[0]

  assert.ok(segmentStyles)
  assert.ok(middleStyles)
  assert.match(segmentStyles, /width: 56px/)
  assert.match(segmentStyles, /height: 57px/)
  assert.match(middleStyles, /top: 27px/)
})

test('API route decoration uses light, shade, and depth to preserve its 3D form', () => {
  const segmentStyles = styles.match(
    /\.overview-route-segment \{[\s\S]*?^}/m
  )?.[0]
  const faceStyles = styles.match(/\.overview-route-face \{[\s\S]*?^}/m)?.[0]
  const topFaceStyles = styles.match(
    /\.overview-route-face\.is-top \{[\s\S]*?^}/m
  )?.[0]
  const leftFaceStyles = styles.match(
    /\.overview-route-face\.is-left \{[\s\S]*?^}/m
  )?.[0]
  const rightFaceStyles = styles.match(
    /\.overview-route-face\.is-right \{[\s\S]*?^}/m
  )?.[0]

  assert.ok(segmentStyles)
  assert.ok(faceStyles)
  assert.ok(topFaceStyles)
  assert.ok(leftFaceStyles)
  assert.ok(rightFaceStyles)
  assert.match(segmentStyles, /filter: drop-shadow/)
  assert.match(faceStyles, /linear-gradient/)
  assert.match(topFaceStyles, /clip-path: polygon/)
  assert.match(leftFaceStyles, /clip-path: polygon/)
  assert.match(rightFaceStyles, /clip-path: polygon/)
})

test('API route content aligns with the service and announcement cards', () => {
  const cardStyles = styles.match(/\.overview-api-card \{[\s\S]*?^}/m)?.[0]
  const headingStyles = styles.match(
    /\.overview-api-card \.overview-small-heading \{[\s\S]*?^}/m
  )?.[0]

  assert.ok(cardStyles)
  assert.ok(headingStyles)
  assert.match(cardStyles, /padding: 0 18px/)
  assert.match(headingStyles, /padding-inline: 0/)
})
