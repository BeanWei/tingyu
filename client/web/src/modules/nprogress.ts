import type { UserModule } from '~/types'

export const install: UserModule = ({ router }) => {
  router.beforeEach((to, from) => {
    if (to.path !== from.path)
      window.$nprogress?.start()
  })
  router.afterEach(() => { window.$nprogress?.finish() })
}
