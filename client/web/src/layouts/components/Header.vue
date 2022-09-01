<script setup lang="ts">
import type { MenuOption } from 'naive-ui'
import { RouterLink } from 'vue-router'

const menuOptions: MenuOption[] = [
  {
    label: () => {
      return h(
        RouterLink,
        {
          to: {
            name: 'index',
          },
        },
        { default: () => '首页' },
      )
    },
    key: 'index',
  },
  {
    label: () => {
      return h(
        RouterLink,
        {
          to: {
            name: 'index',
          },
        },
        { default: () => '关于' },
      )
    },
    key: 'about',
  },
]

const userStore = useUserStore()
const miscStore = useMiscStore()
const route = useRoute()

const activeKey = ref(route.name as string)

watch(route, () => {
  activeKey.value = route.name as string
})
</script>

<template>
  <div class="relative h-54px">
    <header class="border-b-1 border-b-#f1f1f1 color-#909090 z-250 fixed top-0 left-0 right-0 bg-#fff">
      <div class="p-y-6px m-auto max-w-1240px flex items-center h-full w-full relative">
        <nav class="h-full flex-auto">
          <NSpace justify="space-between" align="center">
            <NMenu
              :value="activeKey"
              mode="horizontal"
              :options="menuOptions"
            />
            <NSpace align="center">
              <NInput round placeholder="搜索">
                <template #suffix>
                  <ICarbonSearch />
                </template>
              </NInput>
              <NSpace v-if="!!userStore.info">
                <NButton quaternary circle>
                  <template #icon>
                    <NIcon><ICarbonNotification /></NIcon>
                  </template>
                </NButton>
                <UserAvatar :src="userStore.info?.avatar" :size="32" />
              </NSpace>
              <NButton
                v-else
                strong
                secondary
                round
                type="primary"
                @click="miscStore.setAuthModalVisible(true)"
              >
                登录/注册
              </NButton>
            </NSpace>
          </NSpace>
        </nav>
      </div>
    </header>
  </div>
</template>
