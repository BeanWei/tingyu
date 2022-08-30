<script setup lang="ts">
import { useAxios } from '@vueuse/integrations/useAxios'
import type { Result } from '~/api'
import { instance, url } from '~/api'

const pageNumberRef = ref(1)
const { data, isLoading } = useAxios<Result<API.ListPostResp>>(url.listPost, {
  params: {
    limit: 20,
    page: pageNumberRef.value,
  },
}, instance)
</script>

<template>
  <div>
    <div class="relative bg-#fff border-rd-4px p-x-4 p-t-2 p-b-4 mb-3">
      <PostEditor />
    </div>
    <NList clickable>
      <div v-if="isLoading" class="bg-#fff border-rd-4px relative">
        <PostSkeleton />
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
    <div text-4xl>
      <div i-carbon-campsite inline-block />
    </div>
  </div>
</template>

<route lang="yaml">
meta:
  layout: default
</route>
