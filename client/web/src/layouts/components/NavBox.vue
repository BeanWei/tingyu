<script setup lang="ts">
import { useAxios } from '@vueuse/integrations/useAxios'
import type { MenuOption } from 'naive-ui'
import { RouterLink } from 'vue-router'
import type { Result } from '~/api'
import { instance, url } from '~/api'
import ICarbonHome from '~icons/carbon/home'
import ICarbonUserMilitary from '~icons/carbon/user-military'
import ICarbonUserHashtag from '~icons/carbon/hashtag'
import { menuLabelRender } from '~/utils/ui'

const userStore = useUserStore()
const route = useRoute()

const menuOptions: MenuOption[] = [
  {
    label: '广场',
    path: '/',
    key: 'index',
    icon: () => h(ICarbonHome),
  },
  {
    label: '关注',
    key: 'following',
    icon: () => h(ICarbonUserMilitary),
    show: !!userStore.info?.id,
  },
  {
    label: '话题',
    path: '/topic',
    key: 'topic',
    icon: () => h(ICarbonUserHashtag),
  },
]

const getActiveKeyByRouteName = (name: string): string => {
  const { topic_id } = route.query
  if (topic_id)
    return topic_id as string
  return route.name as string
}

const activeKey = ref(getActiveKeyByRouteName(route.name as string))

const { data, isFinished, execute } = useAxios<Result<AnyObject[]>>(url.listTopic, {
  params: {
    limit: 20,
    page: 1,
  },
}, instance, { immediate: false })
const mergeMenus = (): MenuOption[] => {
  if (isFinished) {
    menuOptions[1].show = true
    if (data.value?.total) {
      return [
        ...menuOptions,
        ...(data.value?.data || []).map((item) => {
          return {
            label: () => {
              return h(RouterLink, {
                to: {
                  name: 'index',
                  query: {
                    topic_id: item.id,
                  },
                },
              }, {
                default: () => h('div', {
                  style: {
                    'padding-left': '32px',
                    'color': '#8a919f',
                  },
                }, item.title),
              })
            },
            key: String(item.id),
          }
        }),
      ]
    }
  }
  return menuOptions
}

userStore.$subscribe((_, $state) => {
  if ($state.info?.id) {
    execute(url.listTopic, {
      params: {
        limit: 20,
        page: 1,
        user_id: userStore.info?.id,
      },
    })
  }
  else {
    menuOptions[1].show = false
    mergeMenus()
  }
})

watch(route, () => {
  activeKey.value = getActiveKeyByRouteName(route.name as string)
})
</script>

<template>
  <nav class="w-180px bg-#fff box-border z-99 fixed border-rd-4px">
    <NScrollbar class="max-h-[calc(100%-109px)]">
      <NMenu
        :value="activeKey"
        :options="mergeMenus()"
        :expand-icon="() => null"
        :collapsed="false"
        :render-label="menuLabelRender"
        default-expand-all
      />
    </NScrollbar>
  </nav>
</template>
