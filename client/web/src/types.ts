import type { DialogProviderInst, LoadingBarProviderInst, MessageProviderInst, NotificationProviderInst } from 'naive-ui'
import type { App } from 'vue'
import type { Router } from 'vue-router'

declare global {
  interface Window {
    $nprogress?: LoadingBarProviderInst
    $dialog?: DialogProviderInst
    $message?: MessageProviderInst
    $notification?: NotificationProviderInst
  }
  type AnyObject = Record<string, any>
}

export type UserModule = (ctx: {
  app: App<Element>
  router: Router
  initialState: Record<string, any>
}) => void
