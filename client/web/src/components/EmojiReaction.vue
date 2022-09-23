<script setup lang="ts">
import sortBy from 'lodash/sortBy'

const emojis: Record<string, string> = {
  'emoji-thumbs_up': '👍',
  'emoji-thumbs_down': '👎',
  'emoji-laugh': '😄',
  'emoji-hooray': '🎉',
  'emoji-confused': '😕',
  'emoji-heart': '❤️',
  'emoji-rocket': '🚀',
  'emoji-eyes': '👀',
}
const codes = Object.keys(emojis)

const reactions = sortBy([
  { code: 'emoji-heart', count: 99999, active: false },
  { code: 'emoji-thumbs_up', count: 99999, active: true },
  { code: 'emoji-laugh', count: 2, active: false },
], (obj) => {
  return codes.indexOf(obj.code)
})

const handleReact = (code: string) => {
  //
}
</script>

<template>
  <NSpace align="center" :size="[2, 0]">
    <NPopover trigger="focus" placement="top-start" style="padding-left: 4px; padding-right: 4px;">
      <template #trigger>
        <NButton quaternary circle>
          <template #icon>
            <NIcon>
              <ICarbonFaceAdd />
            </NIcon>
          </template>
        </NButton>
      </template>
      <NSpace :size="[2, 0]">
        <NButton
          v-for="(text, code) in emojis"
          :key="code"
          quaternary
          circle
          size="small"
          @click="() => handleReact(code)"
        >
          {{ text }}
        </NButton>
      </NSpace>
    </NPopover>
    <NSpace :size="[8, 0]">
      <NButton
        v-for="reaction in reactions"
        :key="reaction.code"
        :type="reaction.active ? 'primary' : 'default'"
        ghost
        round
        size="tiny"
      >
        {{ emojis[reaction.code] }}
        <span class="ml-1">{{ reaction.count > 99 ? '99+' : reaction.count }}</span>
      </NButton>
    </NSpace>
  </NSpace>
</template>
