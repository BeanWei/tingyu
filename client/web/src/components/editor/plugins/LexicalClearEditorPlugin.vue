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
let removeUpdateListener: () => void

onMounted(() => {
  const updateListener = () => {
    editor.registerCommand(
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
  }

  removeUpdateListener = editor.registerUpdateListener(updateListener)
})

onUnmounted(() => {
  removeUpdateListener?.()
})
</script>

<template />
