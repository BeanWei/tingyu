<script setup lang="ts">
import Editor from './Editor.vue'
import { createPost } from '~/api/post'
</script>

<template>
  <Editor
    placeholder="快和水友一起分享新鲜事~"
    submit-button-text="发布"
    @submit="createPost"
  />
</template>
