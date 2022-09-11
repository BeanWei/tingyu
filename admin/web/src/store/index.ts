import create from 'zustand'
import { persist } from 'zustand/middleware'

export const useUserStore = create<{
  token: string
  info?: {
    id: number
    username: string
    nickname: string
    avatar: string
  }
  updateToken: (token: string) => void
  deleteToken: () => void
}>()(set => ({
  token: localStorage.getItem('__tingyu_token'),
  updateToken: (token: string) => {
    set(() => ({ token }))
    localStorage.setItem('__tingyu_token', token)
  },
  deleteToken: () => {
    set(() => ({ token: '' }))
    localStorage.removeItem('__tingyu_token')
  },
}))

export const useSettingsStore = create<{
  theme: 'light' | 'dark'
  lang: string
  changeTheme: () => void
}>()(
  persist(
    set => ({
      theme: 'light',
      lang: 'zh-CN',
      changeTheme: () => set(state => ({ theme: state.theme === 'light' ? 'dark' : 'light' })),
    }),
    {
      name: '__tingyu_admin_settings',
    },
  ),
)
