<script setup lang="ts">
import { useAxios } from '@vueuse/integrations/useAxios'
import type { Result } from '~/api'
import { instance, url } from '~/api'

const modalVisible = ref(false)
const isLoading = ref(true)
const pageNumber = ref(1)
const topics = ref<AnyObject[]>([])
const hasMore = ref(false)

const loadMoreTopic = () => {
  return useAxios<Result<AnyObject[]>>(url.listTopic, {
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
  if (hasMore.value) {
    pageNumber.value++
    const { data, isFinished } = await loadMoreTopic()
    if (isFinished) {
      topics.value.push(...(data.value?.data || []))
      hasMore.value = topics.value.length < (data.value?.total || 0)
      $state.loaded()
    }
  }
  else {
    $state.complete()
  }
}

onMounted(async () => {
  const { data, isFinished } = await loadMoreTopic()
  if (isFinished) {
    isLoading.value = false
    if (data.value?.total)
      topics.value.push(...(data.value.data || []))
    hasMore.value = topics.value.length < (data.value?.total || 0)
  }
})
</script>

<template>
  <div class="bg-#fff p-t-4 p-x-5 rounded-t-4px border-b-1 border-b-#efeff5">
    <NSpace justify="space-between" align="center" class="m-b-4">
      <div class="text-18px font-600 color-#252933">
        话题广场
      </div>
      <NInput round placeholder="搜索话题名称">
        <template #suffix>
          <ICarbonSearch />
        </template>
      </NInput>
    </NSpace>
    <NSpace justify="space-between">
      <NTabs type="bar">
        <NTab name="followed">
          关注
        </NTab>
        <NTab name="recommend">
          推荐
        </NTab>
        <NTab name="more">
          <NDropdown
            :options="[{
              label: '更多',
              key: 'more',
            }]"
          >
            <div class="items-center">
              更多
              <NIcon>
                <ICarbonChevronDown />
              </NIcon>
            </div>
          </NDropdown>
        </NTab>
      </NTabs>
      <NButton
        text
        type="primary"
        @click="() => { modalVisible = true }"
      >
        <template #icon>
          <n-icon>
            <ICarbonAdd />
          </n-icon>
        </template>
        添加话题
      </NButton>
    </NSpace>
  </div>
  <NList :show-divider="false">
    <div v-if="isLoading" class="bg-#fff border-rd-4px relative p-20px">
      <Skeleton />
    </div>
    <div v-else-if="!topics.length" class="bg-#fff border-rd-4px relative min-h-100 flex items-center justify-center">
      <NEmpty size="large" description="一篇荒芜 :)" />
    </div>
    <NListItem
      v-for="topic in topics"
      v-else
      :key="topic.id"
      class="bg-#f4f5f5 important-p-0 important-p-b-2"
    >
      <TopicItem :data="topic" />
    </NListItem>
  </NList>
  <InfiniteScroll v-if="topics.length" @intersect="handleIntersect" />
  <AddTopicModal :show="modalVisible" @close="() => { modalVisible = false }" />
</template>
