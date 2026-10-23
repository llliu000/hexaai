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
export type SectionScrollMotion = {
  position: number
  velocity: number
  frequency?: number
}

export function advanceSectionScrollMotion(
  motion: SectionScrollMotion,
  target: number,
  elapsed: number
): SectionScrollMotion {
  const seconds = Math.min(64, Math.max(0, elapsed)) / 1000
  const frequency = motion.frequency ?? 20
  const offset = motion.position - target
  const response = motion.velocity + frequency * offset
  const decay = Math.exp(-frequency * seconds)
  const position = target + (offset + response * seconds) * decay
  const velocity = (motion.velocity - frequency * response * seconds) * decay
  const crossedTarget = (target - motion.position) * (target - position) < 0

  if (
    crossedTarget ||
    (Math.abs(position - target) < 0.5 && Math.abs(velocity) < 8)
  ) {
    return { position: target, velocity: 0, frequency: motion.frequency }
  }
  return { position, velocity, frequency: motion.frequency }
}

/** Seed a forward glide so module page turns start moving like a native slide. */
export function pageTurnScrollMotion(
  position: number,
  target: number
): SectionScrollMotion {
  const travel = target - position
  const distance = Math.abs(travel)
  // Longer page turns glide a bit slower so the switch reads as sliding, not a snap.
  const frequency = distance > 420 ? 13 : distance > 180 ? 16 : 20
  if (distance <= 140) return { position, velocity: 0, frequency }
  return {
    position,
    velocity: Math.sign(travel) * Math.min(1200, 320 + distance * 1.1),
    frequency,
  }
}
