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
</script>

<template>
  <div class="relative h-54px">
    <header class="border-b-1 border-b-#f1f1f1 color-#909090 z-250 fixed top-0 left-0 right-0 bg-#fff">
      <div class="p-y-6px m-auto max-w-1240px flex items-center h-full w-full relative">
        <nav class="h-full flex-auto">
          <n-space justify="space-between" align="center">
            <n-menu :value="activeKey" mode="horizontal" :options="menuOptions" />
            <n-space align="center">
              <n-input round placeholder="搜索">
                <template #suffix>
                  <i-carbon-search />
                </template>
              </n-input>
              <n-button quaternary circle>
                <template #icon>
                  <n-icon><i-carbon-notification /></n-icon>
                </template>
              </n-button>
              <n-avatar
                round
                :size="34"
                src="https://avatars.githubusercontent.com/u/33516660?s=40&v=4"
              />
            </n-space>
          </n-space>
        </nav>
      </div>
    </header>
  </div>
</template>
