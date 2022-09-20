<script setup lang="ts">
import { COMMAND_PRIORITY_EDITOR } from 'lexical'
import { useEditor } from 'lexical-vue'
import { $createImageNode, INSERT_IMAGE_COMMAND } from '../nodes/ImageNode'
import { $insertNodeToNearestRoot } from '~/utils/lexical'

const editor = useEditor()

onMounted(() => {
  return editor.registerCommand(
    INSERT_IMAGE_COMMAND,
    (src: string) => {
      const imageNode = $createImageNode(src)
      $insertNodeToNearestRoot(imageNode)
      return true
    },
    COMMAND_PRIORITY_EDITOR,
  )
})
</script>

<template />
