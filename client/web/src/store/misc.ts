import { defineStore } from 'pinia'

export interface IMiscState {
  authModalVisible: boolean
}

export const useMiscStore = defineStore('misc', {
  state: (): IMiscState => ({
    authModalVisible: false,
  }),
  getters: {
    getAuthModalVisible(): boolean {
      return this.authModalVisible
    },
  },
  actions: {
    setAuthModalVisible(visible: boolean) {
      this.authModalVisible = visible
    },
  },
})
