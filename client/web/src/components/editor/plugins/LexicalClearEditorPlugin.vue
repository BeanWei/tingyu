<script setup lang="ts">
import {
  $createParagraphNode,
  $getRoot,
  $getSelection,
  CLEAR_EDITOR_COMMAND,
  COMMAND_PRIORITY_EDITOR,
} from 'lexical'
import { useEditor } from 'lexical-vue'
import { onMounted, onUnmounted } from 'vue'

const editor = useEditor()
let unregisterListener: () => void

onMounted(() => {
  unregisterListener = editor.registerCommand(
    CLEAR_EDITOR_COMMAND,
    (_payload) => {
      editor.update(() => {
        const root = $getRoot()
        const selection = $getSelection()
        const paragraph = $createParagraphNode()
        root.clear()
        root.append(paragraph)
        if (selection !== null)
          paragraph.select()
      })
      return true
    },
    COMMAND_PRIORITY_EDITOR,
  )
})

onUnmounted(() => {
  unregisterListener?.()
})
</script>

<template />
