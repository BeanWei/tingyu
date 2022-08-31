<script setup lang="ts">
import { useAxios } from '@vueuse/integrations/useAxios'
import type { Result } from '~/api'
import { instance, url } from '~/api'

const pageNumberRef = ref(1)
const { data, isLoading } = useAxios<Result<AnyObject[]>>(url.listPost, {
  params: {
    limit: 20,
    page: pageNumberRef.value,
  },
}, instance)

const createPost = (values: AnyObject) => {
  return useAxios(url.createPost, {
    data: values,
  }, instance)
}
</script>

<template>
  <div>
    <div class="relative bg-#fff border-rd-4px p-x-4 p-t-4 p-b-4 mb-3">
      <Editor
        placeholder="快和水友一起分享新鲜事~"
        submit-button-text="发布"
        @submit="createPost"
      />
    </div>
    <NList>
      <div v-if="isLoading" class="bg-#fff border-rd-4px relative p-20px">
        <Skeleton />
      </div>
      <div v-else-if="!data?.total" class="bg-#fff border-rd-4px relative min-h-100 flex items-center justify-center">
        <NEmpty size="large" description="一篇荒芜 :)" />
      </div>
      <NListItem
        v-for="post in data.data"
        v-else
        :key="post.id"
        :show-divider="false"
        class="bg-#f4f5f5 important-p-0 important-p-b-2"
      >
        <PostItem :data="post" />
      </NListItem>
    </NList>
  </div>
</template>

<route lang="yaml">
meta:
  layout: default
</route>
