<script setup lang="ts">
import { useAxios } from '@vueuse/integrations/useAxios'
import type { Result } from '~/api'
import { instance, url } from '~/api'

const tabName = ref('followed')
const modalVisible = ref(false)
const isLoading = ref(true)

const reqParams = ref<{
  limit: number
  page: number
  sort_type?: number
  category_id?: number | string
}>({
  limit: 20,
  page: 1,
})
const topics = ref<AnyObject[]>([])
const hasMore = ref(false)

const userStore = useUserStore()

const loadTopic = () => {
  return useAxios<Result<AnyObject[]>>(url.listTopic, {
    params: {
      ...(tabName.value === 'more'
        ? reqParams.value
        : {
            limit: reqParams.value.limit,
            page: reqParams.value.page,
          }),
      user_id: tabName.value === 'followed' ? userStore.info?.id : undefined,
      is_rec: tabName.value === 'recommend' ? true : undefined,
    },
  }, instance)
}

const reloadTopic = async () => {
  reqParams.value.page = 1
  isLoading.value = true

  const { data, isFinished } = await loadTopic()
  if (isFinished) {
    topics.value = data.value?.data || []
    isLoading.value = false
    hasMore.value = topics.value.length < (data.value?.total || 0)
  }
}

const handleIntersect = async ($state: {
  loaded: () => void
  complete: () => void
}) => {
  if (hasMore.value) {
    reqParams.value.page++
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

const handleTabChange = (name: string) => {
  tabName.value = name
  reloadTopic()
}

const handleCategoryChange = (value: number | string) => {
  reqParams.value.category_id = value
  reloadTopic()
}

const handleSortChange = (sort: string) => {
  reqParams.value.sort_type = ['hot', 'new'].indexOf(sort) + 1
  reloadTopic()
}

const { data: categories } = useAxios<Result<AnyObject[]>>(url.listCategory, {}, instance)

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
      <NTabs
        type="segment"
        size="small"
        class="w-30"
        :on-update:value="handleSortChange"
        :value="reqParams?.sort_type ? ['hot', 'new'][reqParams?.sort_type - 1] : 'hot'"
      >
        <NTab name="hot" tab="热门" />
        <NTab name="new" tab="最新" />
      </NTabs>
      <NPopselect
        v-model:value="reqParams.category_id"
        :on-update:value="handleCategoryChange"
        trigger="hover"
        :options="[{
          label: '全部分类',
          value: 0,
          key: 0,
        }, ...(categories?.data?.map(item => {
          return {
            label: item.name,
            value: item.id,
            key: item.id,
          }
        }) || [])]"
        scrollable
      >
        <NButton quaternary icon-placement="right">
          {{ categories?.data?.find(item => item.id === reqParams.category_id)?.name || '全部分类' }}
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
