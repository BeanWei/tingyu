<script setup lang="ts">
import { useAxios } from '@vueuse/integrations/useAxios'
import type { Result } from '~/api'
import { instance, url } from '~/api'

const route = useRoute()

const reqParams = ref<{
  limit: number
  page: number
  sort_type?: number
  topic_id?: number | string
}>({
  limit: 20,
  page: 1,
  sort_type: 1,
  topic_id: route.query?.topic_id as string,
})
const isLoading = ref(true)

const posts = ref<AnyObject[]>([])
const hasMorePost = ref(false)

const loadPost = () => {
  return useAxios<Result<AnyObject[]>>(url.listPost, {
    params: reqParams.value,
  }, instance)
}

const handleIntersect = async ($state: {
  loaded: () => void
  complete: () => void
}) => {
  if (hasMorePost.value) {
    reqParams.value.page++
    const { data, isFinished } = await loadPost()
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

const handleTabChange = async (name: string) => {
  reqParams.value.sort_type = name === 'hot' ? 1 : 0
  reqParams.value.page = 1
  isLoading.value = true

  const { data, isFinished } = await loadPost()
  if (isFinished) {
    posts.value = data.value?.data || []
    isLoading.value = false
  }
}

const createPost = (values: AnyObject) => {
  return useAxios(url.createPost, {
    data: values,
  }, instance)
}

onMounted(async () => {
  const { data, isFinished } = await loadPost()
  if (isFinished) {
    isLoading.value = false
    if (data.value?.total)
      posts.value.push(...(data.value.data || []))
    hasMorePost.value = posts.value.length < (data.value?.total || 0)
  }
})

watch(route, async () => {
  isLoading.value = true
  reqParams.value = {
    ...reqParams.value,
    topic_id: route.query?.topic_id as string,
  }
  const { data, isFinished } = await loadPost()
  if (isFinished) {
    isLoading.value = false
    if (data.value?.total)
      posts.value.push(...(data.value.data || []))
    hasMorePost.value = posts.value.length < (data.value?.total || 0)
  }
})
</script>

<template>
  <div class="relative bg-#fff border-rd-4px p-x-4 p-t-4 p-b-4 mb-3">
    <Editor
      placeholder="快和水友一起分享新鲜事~"
      submit-button-text="发布"
      @submit="createPost"
    />
  </div>
  <div class="bg-#fff p-t-4 p-x-5 rounded-t-4px border-b-1 border-b-#efeff5">
    <NTabs type="bar" @update-value="handleTabChange">
      <NTab name="host">
        热门
      </NTab>
      <NTab name="new">
        最新
      </NTab>
    </NTabs>
  </div>
  <NList :show-divider="false">
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
      class="bg-#f4f5f5 important-p-0 important-p-b-2"
    >
      <PostItem :data="post" />
    </NListItem>
  </NList>
  <InfiniteScroll v-if="!isLoading && posts.length" @intersect="handleIntersect" />
</template>

<route lang="yaml">
meta:
  layout: default
</route>
