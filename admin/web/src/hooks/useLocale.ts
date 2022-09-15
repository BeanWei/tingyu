import defaultLocales from '~/locales'
import { useSettingsStore } from '~/store'

export default function useLocale(locales = null) {
  const lang = useSettingsStore().lang

  return (locales || defaultLocales)[lang] || {}
}
