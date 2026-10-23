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
  getAppearingSectionTarget,
  getEnteredSectionTarget,
} from './section-scroll-appear'
import { getSectionNudgeTarget, getSectionSettleTarget } from './section-scroll'
import { createSectionScrollGeometry } from './section-scroll-geometry'
import {
  advanceSectionScrollMotion,
  pageTurnScrollMotion,
  type SectionScrollMotion,
} from './section-scroll-motion'
import { getLastVisibleSectionTarget } from './section-scroll-reveal'

const INTERACTIVE =
  'input, textarea, select, [contenteditable]:not([contenteditable="false"]), [role="slider"], [role="dialog"], [role="menu"]'

export function attachHomeSectionPaging(root: HTMLElement): () => void {
  const desktop = window.matchMedia(
    '(min-width: 821px) and (hover: hover) and (pointer: fine)'
  )
  const reducedMotion = window.matchMedia('(prefers-reduced-motion: reduce)')
  const supportsScrollEnd = 'onscrollend' in document
  const quietDelay = supportsScrollEnd ? 140 : 180
  const geometry = createSectionScrollGeometry(root)
  let armed = false
  let moved = false
  let nativeEnded = !supportsScrollEnd
  let origin = window.scrollY
  let lastPosition = origin
  let lastActivity = 0
  let gestureStartedAt = 0
  let gestureDistance = 0
  let timer = 0
  let frame = 0
  let animationTarget: (() => number) | null = null
  let motion: SectionScrollMotion = { position: origin, velocity: 0 }
  let lastFrame = 0
  let revealOrigin = origin
  let revealDirection = 0
  let pageTurnDirection = 0
  let pageTurnProtectUntil = 0

  const stopAnimation = (): void => {
    window.cancelAnimationFrame(frame)
    frame = 0
    animationTarget = null
    pageTurnDirection = 0
    pageTurnProtectUntil = 0
  }
  const clearTimer = (): void => {
    window.clearTimeout(timer)
    timer = 0
  }
  const cancel = (): void => {
    stopAnimation()
    clearTimer()
    armed = false
    moved = false
    gestureDistance = 0
    revealDirection = 0
    lastPosition = window.scrollY
  }
  const enabled = (): boolean =>
    desktop.matches && !reducedMotion.matches && !root.closest('[inert]')

  const headerHeight = (): number =>
    root.querySelector('header')?.getBoundingClientRect().height ?? 0

  const viewport = () => ({
    height: window.innerHeight,
    headerHeight: headerHeight(),
  })

  const tick = (now: number): void => {
    if (!animationTarget) return
    const target = animationTarget()
    motion = advanceSectionScrollMotion(motion, target, now - lastFrame)
    lastFrame = now
    window.scrollTo({ top: motion.position, behavior: 'instant' })
    lastPosition = window.scrollY
    if (motion.position === target && motion.velocity === 0) stopAnimation()
    else frame = window.requestAnimationFrame(tick)
  }

  const startPageTurn = (
    selected: { element: HTMLElement; range: { start: number; end: number } },
    edge: 'start' | 'end'
  ): void => {
    armed = false
    clearTimer()
    animationTarget = () =>
      geometry.read().find((section) => section.element === selected.element)
        ?.range[edge] ?? window.scrollY
    const goal = animationTarget()
    // Protect only gentle same-direction inertia. Fast wheel bursts cancel below.
    pageTurnDirection = revealDirection || Math.sign(goal - origin) || 1
    pageTurnProtectUntil = performance.now() + 420
    motion = pageTurnScrollMotion(window.scrollY, goal)
    lastFrame = performance.now()
    frame = window.requestAnimationFrame(tick)
  }

  const resolveTarget = (
    position: number
  ): {
    selected: { element: HTMLElement; range: { start: number; end: number } }
    edge: 'start' | 'end'
  } | null => {
    const sections = geometry.read()
    const ranges = sections.map((section) => section.range)
    const direction = revealDirection || Math.sign(position - origin)
    const gesture = {
      origin: revealOrigin || origin,
      direction,
    }
    const travel = Math.abs(position - origin)
    // Fast multi-module flings keep native settle/reveal only.
    const fling =
      gestureDistance > Math.min(520, window.innerHeight * 0.65) ||
      travel > Math.min(480, window.innerHeight * 0.6)
    const target = fling
      ? (getSectionSettleTarget(position, ranges, window.innerHeight, origin) ??
        getLastVisibleSectionTarget(position, sections, viewport(), {
          origin: revealOrigin,
          direction,
        }))
      : (getAppearingSectionTarget(position, sections, viewport(), gesture) ??
        getEnteredSectionTarget(position, sections, viewport(), gesture) ??
        getSectionNudgeTarget(position, ranges, window.innerHeight, {
          origin,
          distance: Math.max(gestureDistance, travel),
          durationMs: Math.max(0, lastActivity - gestureStartedAt),
        }) ??
        getSectionSettleTarget(position, ranges, window.innerHeight, origin) ??
        getLastVisibleSectionTarget(position, sections, viewport(), {
          origin: revealOrigin,
          direction,
        }))
    if (target === null) return null
    const selected = sections.find(
      (section) =>
        section.range.start === target || section.range.end === target
    )
    if (!selected) return null
    return {
      selected,
      edge: selected.range.start === target ? 'start' : 'end',
    }
  }

  const settle = (): void => {
    timer = 0
    if (!armed || !moved || !nativeEnded || !enabled()) return
    if (performance.now() - lastActivity < quietDelay) {
      scheduleSettle()
      return
    }
    armed = false
    if (
      window.getComputedStyle(document.body).overflowY === 'hidden' ||
      window.getComputedStyle(document.documentElement).overflowY === 'hidden'
    ) {
      return
    }
    const resolved = resolveTarget(window.scrollY)
    if (!resolved) return
    startPageTurn(resolved.selected, resolved.edge)
  }

  const scheduleSettle = (): void => {
    clearTimer()
    const remaining = Math.max(
      0,
      quietDelay - (performance.now() - lastActivity)
    )
    timer = window.setTimeout(settle, remaining)
  }

  const onWheel = (event: WheelEvent): void => {
    if (event.deltaX === 0 && event.deltaY === 0) return
    // A strong wheel burst means the user is flinging — release page-turn ownership.
    if (animationTarget && Math.abs(event.deltaY) >= 48) {
      stopAnimation()
      clearTimer()
    } else if (
      animationTarget &&
      performance.now() < pageTurnProtectUntil &&
      pageTurnDirection &&
      Math.sign(event.deltaY) === pageTurnDirection
    ) {
      return
    }
    // Never cancel a wheel event: native acceleration and inertia own the gesture.
    stopAnimation()
    clearTimer()
    const target = event.target
    if (
      !enabled() ||
      event.defaultPrevented ||
      event.ctrlKey ||
      event.metaKey ||
      event.altKey ||
      event.shiftKey ||
      Math.abs(event.deltaX) >= Math.abs(event.deltaY) ||
      !(target instanceof Element) ||
      !root.contains(target) ||
      target.closest(INTERACTIVE)
    ) {
      armed = false
      return
    }
    if (!armed) {
      origin = lastPosition
      revealOrigin = origin
      revealDirection = 0
      gestureStartedAt = performance.now()
      gestureDistance = 0
      moved = window.scrollY !== origin
      nativeEnded = !supportsScrollEnd
    }
    if (!revealDirection && event.deltaY !== 0) {
      revealDirection = Math.sign(event.deltaY)
    }
    armed = true
    lastActivity = performance.now()
    scheduleSettle()
  }

  const tryEarlyPageTurn = (
    position: number,
    frameDelta: number
  ): boolean => {
    if (!enabled() || animationTarget) return false
    // Big jumps are flings — do not hijack them.
    if (frameDelta > 80) return false
    if (gestureDistance > 180) return false
    const elapsed = performance.now() - gestureStartedAt
    if (elapsed < 32) return false
    const travel = Math.abs(position - origin)
    if (travel < 48 || travel > 140) return false
    if (gestureDistance - travel > 40) return false
    const direction = revealDirection || Math.sign(position - origin)
    if (!direction) return false
    const sections = geometry.read()
    const ranges = sections.map((section) => section.range)
    // Appear first so gentle downward turns work when the next module already peeks.
    const target =
      getAppearingSectionTarget(position, sections, viewport(), {
        origin: revealOrigin || origin,
        direction,
      }) ??
      getSectionNudgeTarget(position, ranges, window.innerHeight, {
        origin,
        distance: Math.max(gestureDistance, travel),
        durationMs: Math.max(0, elapsed),
      })
    if (target === null) return false
    const selected = sections.find(
      (section) =>
        section.range.start === target || section.range.end === target
    )
    if (!selected) return false
    startPageTurn(
      selected,
      selected.range.start === target ? 'start' : 'end'
    )
    return true
  }

  const onScroll = (): void => {
    const position = window.scrollY
    const previous = lastPosition
    const delta = position - previous
    const distance = Math.abs(delta)
    lastPosition = position
    if (animationTarget) return
    if (!armed || distance === 0) return
    const direction = Math.sign(delta)
    if (!revealDirection || (direction !== revealDirection && distance >= 2)) {
      revealOrigin = previous
      revealDirection = direction
    }
    gestureDistance += distance
    moved = true
    nativeEnded = !supportsScrollEnd
    lastActivity = performance.now()
    if (tryEarlyPageTurn(position, distance)) return
    scheduleSettle()
  }

  const onScrollEnd = (event: Event): void => {
    if (event.target !== document) return
    if (!armed || animationTarget) return
    nativeEnded = true
    scheduleSettle()
  }

  window.addEventListener('wheel', onWheel, { passive: true })
  window.addEventListener('scroll', onScroll, { passive: true })
  document.addEventListener('scrollend', onScrollEnd, { passive: true })
  window.addEventListener('keydown', cancel)
  window.addEventListener('pointerdown', cancel, { passive: true })
  window.addEventListener('touchstart', cancel, { passive: true })
  window.addEventListener('resize', cancel)
  window.addEventListener('hashchange', cancel)
  window.addEventListener('popstate', cancel)
  root.addEventListener('click', cancel)
  document.addEventListener('visibilitychange', cancel)
  desktop.addEventListener('change', cancel)
  reducedMotion.addEventListener('change', cancel)

  return () => {
    cancel()
    geometry.destroy()
    window.removeEventListener('wheel', onWheel)
    window.removeEventListener('scroll', onScroll)
    document.removeEventListener('scrollend', onScrollEnd)
    window.removeEventListener('keydown', cancel)
    window.removeEventListener('pointerdown', cancel)
    window.removeEventListener('touchstart', cancel)
    window.removeEventListener('resize', cancel)
    window.removeEventListener('hashchange', cancel)
    window.removeEventListener('popstate', cancel)
    root.removeEventListener('click', cancel)
    document.removeEventListener('visibilitychange', cancel)
    desktop.removeEventListener('change', cancel)
    reducedMotion.removeEventListener('change', cancel)
  }
}
