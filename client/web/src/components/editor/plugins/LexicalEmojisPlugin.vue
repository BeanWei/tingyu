<script setup lang="ts">
import { $getSelection, $isRangeSelection, COMMAND_PRIORITY_EDITOR } from 'lexical'
import { useEditor } from 'lexical-vue'
import { $createEmojiNode, INSERT_EMOJI_COMMAND } from '../nodes/EmojiNode'

const editor = useEditor()
let removeUpdateListener: () => void

onMounted(() => {
  const updateListener = () => {
    editor.registerCommand(
      INSERT_EMOJI_COMMAND,
      (emojiName: string) => {
        const selection = $getSelection()
        if ($isRangeSelection(selection)) {
          const emojiNode = $createEmojiNode(emojiName)
          selection.insertNodes([emojiNode])
        }
        return true
      },
      COMMAND_PRIORITY_EDITOR,
    )
  }

  removeUpdateListener = editor.registerUpdateListener(updateListener)
})

onUnmounted(() => {
  removeUpdateListener?.()
})
</script>

<template />
