<script setup lang="ts">
import { useAxios } from '@vueuse/integrations/useAxios'
import type { Result } from '~/api'
import { instance, url } from '~/api'

const isLoading = ref(true)
const pageNumber = ref(1)
const posts = ref<AnyObject[]>([])
const hasMorePost = ref(false)

const loadMorePost = () => {
  return useAxios<Result<AnyObject[]>>(url.listPost, {
    params: {
      limit: 20,
      page: pageNumber.value,
    },
  }, instance)
}

const handleIntersect = async ($state: {
  loaded: () => void
  complete: () => void
}) => {
  if (hasMorePost.value) {
    pageNumber.value++
    const { data, isFinished } = await loadMorePost()
    if (isFinished) {
      posts.value.push(...(data.value?.data || []))
      hasMorePost.value = posts.value.length < (data.value?.total || 0)
      $state.loaded()
    }
  }
  else {
    $state.complete()
  }
}

const createPost = (values: AnyObject) => {
  return useAxios(url.createPost, {
    data: values,
  }, instance)
}

onMounted(async () => {
  const { data, isFinished } = await loadMorePost()
  if (isFinished) {
    isLoading.value = false
    if (data.value?.total)
      posts.value.push(...(data.value.data || []))
    hasMorePost.value = posts.value.length < (data.value?.total || 0)
  }
})
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
      <div v-else-if="!posts.length" class="bg-#fff border-rd-4px relative min-h-100 flex items-center justify-center">
        <NEmpty size="large" description="一篇荒芜 :)" />
      </div>
      <NListItem
        v-for="post in posts"
        v-else
        :key="post.id"
        :show-divider="false"
        class="bg-#f4f5f5 important-p-0 important-p-b-2"
      >
        <PostItem :data="post" />
      </NListItem>
    </NList>
    <InfiniteScroll v-if="posts.length" @intersect="handleIntersect" />
  </div>
</template>

<route lang="yaml">
meta:
  layout: default
</route>
