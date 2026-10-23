/*
 * The feedback and ring emission shader is adapted from Vanta HALO:
 * https://github.com/tengbao/vanta/blob/master/src/vanta.halo.js
 *
 * Copyright 2020 Teng Bao
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

import { HALO_INK_DISTORTION, HALO_WAVE_FEATHER } from './hero-halo-transition'

export const HALO_VERTEX_SHADER = `
attribute vec2 aPosition;
varying vec2 vUv;
void main() {
  vUv = aPosition * 0.5 + 0.5;
  gl_Position = vec4(aPosition, 0.0, 1.0);
}
`

// A stationary paper field gives each pixel a single arrival time. Advancing
// only the radius keeps the wet edge organic without flicker or echo rings.
// Background snapshots omit wet pigment so interrupted transitions stay clean.
export const HALO_REVEAL_SHADER = `
#ifdef GL_FRAGMENT_PRECISION_HIGH
precision highp float;
#else
precision mediump float;
#endif
varying vec2 vUv;
uniform sampler2D uFrom;
uniform vec2 uResolution;
uniform vec2 uOrigin;
uniform vec3 uTargetColor;
uniform vec3 uSurfaceColor;
uniform vec3 uRingColor;
uniform float uRadius;
uniform float uComplete;
uniform float uRingStrength;
uniform vec2 uInkSeed;

// Small hash operands also work on devices without fragment highp support.
float paperHash(vec2 p) {
  vec3 p3 = fract(vec3(p.xyx) * 0.1031);
  p3 = fract(p3 + dot(p3, p3.yzx + 3.33));
  return fract((p3.x + p3.y) * p3.z * 7.0);
}

float paperNoise(vec2 p) {
  vec2 cell = floor(p);
  vec2 f = fract(p);
  f = f * f * (3.0 - 2.0 * f);
  return mix(
    mix(paperHash(cell), paperHash(cell + vec2(1.0, 0.0)), f.x),
    mix(paperHash(cell + vec2(0.0, 1.0)), paperHash(cell + vec2(1.0)), f.x),
    f.y
  );
}

float paperField(vec2 p) {
  return paperNoise(p) * 0.58
    + paperNoise(p * 2.07 + vec2(7.3, 2.8)) * 0.28
    + paperNoise(p * 4.13 + vec2(1.7, 9.2)) * 0.14;
}

vec3 backgroundAt(vec2 uv) {
  float y = 1.0 - uv.y;
  float tint = 0.0;
  if (y < 0.38) tint = mix(0.12, 0.07, y / 0.38);
  else if (y < 0.72) tint = mix(0.07, 0.02, (y - 0.38) / 0.34);
  else if (y < 0.96) tint = mix(0.02, 0.0, (y - 0.72) / 0.24);
  return mix(uSurfaceColor, uTargetColor, tint);
}

void main() {
  vec2 pixel = (vUv - uOrigin) * vec2(uResolution.x / uResolution.y, 1.0);
  vec2 paper = pixel + uInkSeed;
  vec2 warp = vec2(
    paperNoise(paper * 1.35),
    paperNoise(paper * 1.35 + vec2(5.2, 8.6))
  ) - 0.5;
  // Broad overlapping lobes with finer paper fibres at their edges.
  float pooling = paperField(paper * vec2(2.35, 3.1) + warp * 1.5);
  float fibres = paperNoise(paper * vec2(23.0, 31.0) + warp * 3.0);
  float distortion = ((pooling * 2.0 - 1.0) * 0.82
    + (fibres * 2.0 - 1.0) * 0.18) * ${HALO_INK_DISTORTION.toFixed(3)};
  float distance = length(pixel) + distortion;
  float feather = ${HALO_WAVE_FEATHER.toFixed(3)} * mix(0.55, 1.0, pooling);
  float reveal = 1.0 - smoothstep(uRadius - feather, uRadius + feather, distance);
  reveal = mix(reveal, 1.0, uComplete);
  vec3 background = mix(texture2D(uFrom, vUv).rgb, backgroundAt(vUv), reveal);

  // Dilute pigment pools behind the advancing edge, fading into the final wash.
  // The broad, asymmetric falloff avoids an outlined luminous ring.
  float wetDistance = (distance - uRadius + feather * 0.6) / (feather * 2.4);
  float wetEdge = exp(-wetDistance * wetDistance) * reveal;
  float pigment = mix(0.22, 0.85, pooling) * mix(0.80, 1.0, fibres);
  float verticalFade = 1.0 - smoothstep(0.74, 1.0, 1.0 - vUv.y);
  float highlight = wetEdge * pigment * uRingStrength * verticalFade * (1.0 - uComplete);
  gl_FragColor = vec4(mix(background, uRingColor, highlight), 1.0);
}
`
