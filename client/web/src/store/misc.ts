import { defineStore } from 'pinia'

export interface IMiscState {
  authModalVisible: boolean
}

export const useMiscStore = defineStore('misc', {
  state: (): IMiscState => ({
    authModalVisible: false,
  }),
  actions: {
    setAuthModalVisible(visible: boolean) {
      this.authModalVisible = visible
    },
  },
})
