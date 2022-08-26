<script setup lang="ts">
import type { MenuOption } from 'naive-ui'
import { NEllipsis } from 'naive-ui'
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
      children: [
        {
          label: () => h(NEllipsis, {
            class: 'color-#8a919f',
          }, { default: () => '电灯熄灭 物换星移 泥牛入海' }),
          key: 'topics1',
        },
        {
          label: () => h(NEllipsis, {
            class: 'color-#8a919f',
          }, { default: () => '更多' }),
          key: 'topics-more',
        },
      ],
    },
  ],
} = defineProps<{
  activeKey?: string
  menuOptions?: MenuOption[]
}>()
</script>

<template>
  <nav class="w-180px bg-#fff box-border z-99 fixed" border-rd-4px>
    <NScrollbar class="max-h-[calc(100%-109px)]">
      <NMenu
        :value="activeKey"
        :options="menuOptions"
        :expand-icon="() => null"
        :collapsed="false"
        default-expand-all
      />
    </NScrollbar>
  </nav>
</template>
