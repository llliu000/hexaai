const samples: Record<string, string> = {
  python: `from openai import OpenAI

client = OpenAI(
    api_key="YOUR_API_KEY",
    base_url="https://666api.ai/v1"
)

response = client.chat.completions.create(
    model="ai-xlance-3",
    messages=[{"role": "user",
               "content": "Write a haiku about APIs"}],
    temperature=0.7
)
print(response.choices[0].message.content)`,
  javascript: `import OpenAI from "openai";

const client = new OpenAI({
  apiKey: "YOUR_API_KEY",
  baseURL: "https://666api.ai/v1"
});

const response = await client.chat.completions.create({
  model: "ai-xlance-3",
  messages: [{ role: "user",
    content: "Write a haiku about APIs" }],
  temperature: 0.7
});
console.log(response.choices[0].message.content);`,
}

const originalPanels = new WeakMap<Element, Node[]>()

export function watchHeaderScroll(
  viewport: Pick<
    Window,
    'scrollY' | 'addEventListener' | 'removeEventListener'
  >,
  onChange: (floating: boolean) => void
): () => void {
  let previous: boolean | undefined
  const update = () => {
    const floating = viewport.scrollY > 20
    if (floating !== previous) {
      previous = floating
      onChange(floating)
    }
  }
  update()
  viewport.addEventListener('scroll', update, { passive: true })
  return () => viewport.removeEventListener('scroll', update)
}

export function toggleProviderScrolling(button: HTMLButtonElement): void {
  const providers = button.closest<HTMLElement>('.lh-providers')
  if (!providers) return
  const paused = providers.dataset.paused !== 'true'
  providers.dataset.paused = String(paused)
  button.setAttribute('aria-pressed', String(paused))
}

export function selectCodeLanguage(root: ParentNode, language: string): void {
  if (language !== 'curl' && !Object.hasOwn(samples, language)) return
  const panel = root.querySelector('#lh-code-panel')
  if (!panel) return
  if (!originalPanels.has(panel)) {
    originalPanels.set(
      panel,
      [...panel.childNodes].map((node) => node.cloneNode(true))
    )
  }
  root
    .querySelectorAll<HTMLButtonElement>('[data-code-language]')
    .forEach((button) => {
      const selected = button.dataset.codeLanguage === language
      button.setAttribute('aria-selected', String(selected))
      button.tabIndex = selected ? 0 : -1
    })
  panel.setAttribute('aria-labelledby', `lh-tab-${language}`)
  if (language === 'curl') {
    panel.replaceChildren(
      ...(originalPanels.get(panel) ?? []).map((node) => node.cloneNode(true))
    )
  } else {
    panel.textContent = samples[language]
  }
}
