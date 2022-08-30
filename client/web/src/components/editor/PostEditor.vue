<script setup lang="ts">
import { useAxios } from '@vueuse/integrations/useAxios'
import Editor from './Editor.vue'
import { instance, url } from '~/api'

const createPost = (values: API.CreatePostReq) => {
  return useAxios(url.createPost, {
    data: values,
  }, instance)
}
</script>

<template>
  <Editor
    placeholder="快和水友一起分享新鲜事~"
    submit-button-text="发布"
    @submit="createPost"
  />
</template>
