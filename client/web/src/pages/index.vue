<script setup lang="ts">
import { listPost } from '~/api/post'

const pageNumberRef = ref(1)
const pageLoadingRef = ref(true)
const pageDataRef = ref<Record<string, any>[]>()

const loadPostList = async () => {
  const { data, isLoading } = await listPost({
    limit: 20,
    page: pageNumberRef.value,
  })
  pageLoadingRef.value = isLoading.value
  pageDataRef.value = data.value?.data
}

onMounted(() => {
  loadPostList()
})
</script>

<template>
  <div>
    <div class="relative bg-#fff border-rd-4px p-x-4 p-t-2 p-b-4 mb-3">
      <PostEditor />
    </div>
    <NList clickable>
      <div v-if="pageLoadingRef" class="bg-#fff border-rd-4px relative">
        <PostSkeleton />
      </div>
      <div v-else-if="!pageDataRef" class="bg-#fff border-rd-4px relative">
        <NEmpty size="large" description="一篇荒芜:)" />
      </div>
      <NListItem
        v-for="post in pageDataRef"
        v-else
        :key="post.id"
        :show-divider="false"
        class="bg-#f4f5f5 important-p-0 important-p-b-2"
      >
        <PostItem :data="post" />
      </NListItem>
    </NList>
    <div class="bg-#fff border-rd-4px relative">
      <PostItem />
    </div>
    <div text-4xl>
      <div i-carbon-campsite inline-block />
    </div>
  </div>
</template>

<route lang="yaml">
meta:
  layout: default
</route>
