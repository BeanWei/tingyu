<script setup lang="ts">
import type { MenuOption } from 'naive-ui'
import ICarbonTime from '~icons/carbon/time'
import ICarbonFire from '~icons/carbon/fire'
import ICarbonUserMilitary from '~icons/carbon/user-military'
import ICarbonUserHashtag from '~icons/carbon/hashtag'

const {
  activeKey = 'new',
  menuOptions = [
    {
      label: '最新',
      key: 'new',
      icon: () => h(ICarbonTime),
    },
    {
      label: '热门',
      key: 'hot',
      icon: () => h(ICarbonFire),
    },
    {
      label: '关注',
      key: 'following',
      icon: () => h(ICarbonUserMilitary),
    },
    {
      label: '话题',
      key: 'topics',
      icon: () => h(ICarbonUserHashtag),
    },
  ],
} = defineProps<{
  activeKey?: string
  menuOptions?: MenuOption[]
}>()
</script>

<template>
  <nav class="w-180px bg-#fff p-8px box-border z-99 fixed top-80px" border-rd-4px>
    <n-scrollbar class="max-h-[calc(100%-109px)]">
      <n-menu :value="activeKey" :options="menuOptions" />
    </n-scrollbar>
  </nav>
</template>
