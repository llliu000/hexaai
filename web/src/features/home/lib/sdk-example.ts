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
export type ExampleLanguage = 'python' | 'javascript' | 'curl'

export function getSdkBaseUrl(address: unknown, fallback: string): string {
  for (const candidate of [address, fallback, 'http://localhost:3000']) {
    if (typeof candidate !== 'string') continue
    try {
      const url = new URL(candidate)
      if (
        !['http:', 'https:'].includes(url.protocol) ||
        url.username ||
        url.password
      ) {
        continue
      }
      url.hash = ''
      url.search = ''
      const pathname = url.pathname.replace(/\/+$/, '')
      url.pathname = pathname.endsWith('/v1') ? pathname : `${pathname}/v1`
      return url.toString().replace(/\/$/, '')
    } catch {
      /* Try the next safe fallback. */
    }
  }
  return 'http://localhost:3000/v1'
}

export function getSdkExample(
  language: ExampleLanguage,
  base: string,
  model: string
): string {
  const url = JSON.stringify(base),
    name = JSON.stringify(model)
  if (language === 'python') {
    return [
      'import os',
      'from openai import OpenAI',
      '',
      'client = OpenAI(',
      '    api_key=os.environ["AI_API_KEY"],',
      `    base_url=${url},`,
      ')',
      '',
      'response = client.chat.completions.create(',
      `    model=${name},`,
      '    messages=[{"role": "user", "content": "Hello!"}],',
      ')',
      'print(response.choices[0].message.content)',
    ].join('\n')
  }
  if (language === 'javascript') {
    return [
      'import OpenAI from "openai";',
      '',
      'const client = new OpenAI({',
      '  apiKey: process.env.AI_API_KEY,',
      `  baseURL: ${url},`,
      '});',
      '',
      'const response = await client.chat.completions.create({',
      `  model: ${name},`,
      '  messages: [{ role: "user", content: "Hello!" }],',
      '});',
      'console.log(response.choices[0].message.content);',
    ].join('\n')
  }
  const quote = (value: string) => `'${value.replaceAll("'", "'\\''")}'`
  return [
    `curl ${quote(`${base}/chat/completions`)} \\`,
    '  -H "Authorization: Bearer $AI_API_KEY" \\',
    '  -H "Content-Type: application/json" \\',
    `  -d ${quote(JSON.stringify({ model, messages: [{ role: 'user', content: 'Hello!' }] }, null, 2))}`,
  ].join('\n')
}
