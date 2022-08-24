import { createRouter, createWebHistory } from 'vue-router'
import { setupLayouts } from 'virtual:generated-layouts'
import generatedRoutes from 'virtual:generated-pages'
import App from './App.vue'
import type { UserModule } from './types'
import './styles/main.css'
import 'uno.css'

const routes = setupLayouts(generatedRoutes)

const app = createApp(App)
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})
app.use(router)
// install all modules under `modules/`
Object.values(import.meta.glob<{ install: UserModule }>('./modules/*.ts', { eager: true })).forEach(i => i.install?.({
  app,
  router,
  initialState: {},
}))
app.mount('#app')
