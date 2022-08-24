import { createPinia } from 'pinia'
import { type UserModule } from '~/types'

// Setup Pinia
// https://pinia.esm.dev/
export const install: UserModule = ({ initialState, app }) => {
  const pinia = createPinia()
  app.use(pinia)
  pinia.state.value = (initialState.pinia) || {}
}
