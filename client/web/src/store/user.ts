import { acceptHMRUpdate, defineStore } from 'pinia'
import { ACCESS_TOKEN } from '~/constants'

const token = useStorage(ACCESS_TOKEN, '')

export interface IUserState {
  token: string
  info?: {
    id: string
    username: string
    nickname: string
    avatar: string
  }
}

export const useUserStore = defineStore('user', {
  state: (): IUserState => ({
    token: token.value,
  }),
  getters: {
    getToken(): string {
      return this.token
    },
    getInfo(): IUserState['info'] {
      return this.info
    },
  },
  actions: {
    setToken(token: string) {
      this.token = token
    },
    setInfo(info: IUserState['info']) {
      this.info = info
    },
    resetAll() {
      token.value = ''
      this.token = ''
      this.info = undefined
    },
  },
})

if (import.meta.hot)
  import.meta.hot.accept(acceptHMRUpdate(useUserStore, import.meta.hot))
