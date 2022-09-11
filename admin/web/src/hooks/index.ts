import defaultLocales from '~/locales'
import { useSettingsStore } from '~/store'

function useLocale(locales = null) {
  const lang = useSettingsStore().lang

  return (locales || defaultLocales)[lang] || {}
}

export {
  useLocale,
}
