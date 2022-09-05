<script setup lang="ts">
import { useAxios } from '@vueuse/integrations/useAxios'
import type { Result } from '~/api'
import { instance, url } from '~/api'

const tabName = ref('followed')
const modalVisible = ref(false)
const isLoading = ref(true)

const pageNumber = ref(1)
const topics = ref<AnyObject[]>([])
const hasMore = ref(false)

const userStore = useUserStore()

const loadTopic = () => {
  const params: AnyObject = {
    limit: 20,
    page: pageNumber.value,
  }
  if (tabName.value === 'followed')
    params.user_id = userStore.info?.id
  else if (tabName.value === 'recommend')
    params.is_rec = true
  return useAxios<Result<AnyObject[]>>(url.listTopic, {
    params,
  }, instance)
}

const handleIntersect = async ($state: {
  loaded: () => void
  complete: () => void
}) => {
  if (hasMore.value) {
    pageNumber.value++
    const { data, isFinished } = await loadTopic()
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

const handleTabChange = async (name: string) => {
  tabName.value = name
  pageNumber.value = 1
  isLoading.value = true

  const { data, isFinished } = await loadTopic()
  if (isFinished) {
    topics.value = data.value?.data || []
    isLoading.value = false
  }
}

onMounted(async () => {
  const { data, isFinished } = await loadTopic()
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
      <NTabs type="bar" @update-value="handleTabChange">
        <NTab v-if="userStore.info?.id" name="followed">
          关注
        </NTab>
        <NTab name="recommend">
          推荐
        </NTab>
        <NTab name="more">
          更多
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
  <div v-if="tabName === 'more'" class="bg-#fff p-x-5 p-t-4 p-b-1">
    <NSpace align="center">
      <NTabs type="segment" size="small" class="w-30">
        <NTab name="hot" tab="热门" />
        <NTab name="new" tab="最新" />
      </NTabs>
      <NPopselect
        trigger="hover"
        :options="[{
          label: '全部分类',
          value: 'all',
          key: 'all',
        }]"
        scrollable
      >
        <NButton quaternary icon-placement="right">
          全部分类
          <template #icon>
            <n-icon>
              <ICarbonChevronDown />
            </n-icon>
          </template>
        </NButton>
      </NPopselect>
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
      <TopicItem :data="topic" :is-followed="tabName === 'followed'" />
    </NListItem>
  </NList>
  <InfiniteScroll v-if="!isLoading && topics.length" @intersect="handleIntersect" />
  <AddTopicModal :show="modalVisible" @close="() => { modalVisible = false }" />
</template>
