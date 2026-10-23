export function applyThemeToDocument(
  document: Document,
  theme: 'light' | 'dark'
): void {
  const root = document.documentElement
  root.classList.remove('light', 'dark')
  root.classList.add(theme)
  root.style.colorScheme = theme
  document
    .querySelector('meta[name="theme-color"]')
    ?.setAttribute('content', theme === 'dark' ? '#1e1e1e' : '#f4f7fc')
}
