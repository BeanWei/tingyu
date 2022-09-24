<script setup lang="ts">
import sortBy from 'lodash/sortBy'
import { instance } from '~/api'

const { data, reactAction } = defineProps<{
  data: Record<string, any>
  reactAction: string
}>()

const reactions = ref<AnyObject[]>(data.reactions || [])

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

const handleReact = (code: string) => {
  instance({
    url: reactAction,
    data: {
      id: data.id,
      code,
    },
  }).then(() => {
    const length = reactions.value.length
    for (let i = 0; i < length; i++) {
      const item = reactions.value[i]
      if (item.code === code) {
        if (item.active) {
          if (item.count === 1)
            reactions.value = reactions.value.filter((item_: AnyObject) => item_.code !== code)
          else
            reactions.value[i] = { code, count: item.count - 1, active: false }
        }
        else {
          reactions.value[i] = { code, count: item.count + 1, active: true }
        }
        return
      }
    }
    reactions.value = [
      ...(reactions.value || []),
      { code, count: 1, active: true },
    ]
  })
}
</script>

<template>
  <NSpace v-if="reactions.length > 0" :size="[4, 0]" class="mb-2">
    <NButton
      v-for="reaction in sortBy(reactions, (obj) => { return codes.indexOf(obj.code) })"
      :key="reaction.code"
      :type="reaction.active ? 'primary' : 'default'"
      ghost
      round
      size="tiny"
      @click="() => handleReact(reaction.code)"
    >
      {{ emojis[reaction.code] }}
      <span class="ml-1">{{ reaction.count > 99 ? '99+' : reaction.count }}</span>
    </NButton>
  </NSpace>
  <NSpace justify="space-between" align="center">
    <NSpace align="center" :size="[2, 0]">
      <slot name="left" />
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
    </NSpace>
    <slot name="right" />
  </NSpace>
</template>
