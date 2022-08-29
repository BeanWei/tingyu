<script setup lang="ts">
import type { MenuOption } from 'naive-ui'

const {
  activeKey = 'index',
  menuOptions = [
    {
      label: '首页',
      key: 'index',
    },
    {
      label: '关于',
      key: 'about',
    },
  ],
} = defineProps<{
  activeKey?: string
  menuOptions?: MenuOption[]
}>()

const userStore = useUserStore()
const miscStore = useMiscStore()
</script>

<template>
  <div class="relative h-54px">
    <header class="border-b-1 border-b-#f1f1f1 color-#909090 z-250 fixed top-0 left-0 right-0 bg-#fff">
      <div class="p-y-6px m-auto max-w-1240px flex items-center h-full w-full relative">
        <nav class="h-full flex-auto">
          <NSpace justify="space-between" align="center">
            <NMenu :value="activeKey" mode="horizontal" :options="menuOptions" />
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
                <NAvatar
                  v-if="userStore.info?.avatar"
                  round
                  :src="userStore.info?.avatar"
                />
                <NAvatar
                  v-else
                  round
                  :style="{
                    'background-color': 'rgba(24, 160, 88, 0.16)',
                    'color': '#18a058',
                  }"
                >
                  <NIcon><ICarbonRainDrop /></NIcon>
                </NAvatar>
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
