import { acceptHMRUpdate, defineStore } from 'pinia'
import { ACCESS_TOKEN } from '~/constants'

const tokenRef = useStorage(ACCESS_TOKEN, '')

export interface IUserState {
  token: string
  info?: {
    id: number
    username: string
    nickname: string
    avatar: string
    headline: string
    count_post: number
    count_topic: number
    count_following: number
    count_follower: number
  }
}

export const useUserStore = defineStore('user', {
  state: (): IUserState => ({
    token: tokenRef.value,
    info: undefined,
  }),
  actions: {
    setToken(token: string) {
      tokenRef.value = token
      this.token = token
    },
    setInfo(info: IUserState['info']) {
      this.info = info
    },
    resetAll() {
      this.setToken('')
      this.setInfo(undefined)
    },
  },
})

if (import.meta.hot)
  import.meta.hot.accept(acceptHMRUpdate(useUserStore, import.meta.hot))
