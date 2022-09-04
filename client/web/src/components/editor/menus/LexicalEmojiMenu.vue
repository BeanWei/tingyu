<script setup lang="ts">
import { useEditor } from 'lexical-vue'
import { INSERT_EMOJI_COMMAND } from '../nodes/EmojiNode'

const Emojis = [
  'grinning-face-with-smiling-eyes',
  'smiling-face-with-sunglasses',
  'kissing-face-with-smiling-eyes',
  'smiling-face-with-heart-eyes',
  'clapping-hands',
  'thumbs-up',
  'thumbs-down',
  'hand-with-index-finger-and-thumb-crossed',
  'anguished-face',
]

const editor = useEditor()

const handleEmojiSelect = (emojiName: string) => {
  editor.update(() => {
    editor.dispatchCommand(INSERT_EMOJI_COMMAND, emojiName)
  })
}
</script>

<template>
  <NPopover trigger="focus" placement="bottom-start">
    <template #trigger>
      <NButton quaternary circle>
        <template #icon>
          <NIcon><ICarbonFaceAdd /></NIcon>
        </template>
      </NButton>
    </template>
    <NGrid :cols="8" x-gap="12" y-gap="8">
      <NGi v-for="emoji in Emojis" :key="emoji">
        <div
          :class="`i-fluent-emoji-flat-${emoji} w-8 h-8 cursor-pointer`"
          @click="() => handleEmojiSelect(`i-fluent-emoji-flat-${emoji}`)"
        />
      </NGi>
    </NGrid>
  </NPopover>
</template>
