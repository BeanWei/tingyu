<script setup lang="ts">
import { useAxios } from '@vueuse/integrations/useAxios'
import type { Result } from '~/api'
import { instance, url } from '~/api'

const props = defineProps<{ id: string }>()

const userStore = useUserStore()

const commentsPageNumber = ref(1)
const commentsLoading = ref(true)
const comments = ref<AnyObject[]>([])
const hasMoreComments = ref(false)

const { data: post, isLoading: postLoading } = useAxios<Result<AnyObject>>(url.getPost, {
  params: {
    id: props.id,
  },
}, instance)

const loadComments = () => {
  return useAxios<Result<AnyObject[]>>(url.listComment, {
    params: {
      limit: 20,
      page: commentsPageNumber.value,
      post_id: props.id,
    },
  }, instance)
}

const handleIntersect = async ($state: {
  loaded: () => void
  complete: () => void
}) => {
  if (hasMoreComments.value) {
    commentsPageNumber.value++
    const { data, isFinished } = await loadComments()
    if (isFinished) {
      comments.value.push(...(data.value?.data || []))
      hasMoreComments.value = comments.value.length < (data.value?.total || 0)
      $state.loaded()
    }
  }
  else {
    $state.complete()
  }
}

const reloadComments = async () => {
  commentsPageNumber.value = 1
  commentsLoading.value = true
  const { data, isFinished } = await loadComments()
  if (isFinished) {
    comments.value = data.value?.data || []
    hasMoreComments.value = comments.value.length < (data.value?.total || 0)
  }
}

const createComment = (values: AnyObject) => {
  return useAxios(url.createComment, {
    data: {
      ...values,
      post_id: props.id,
    },
  }, instance)
}

const handleNewReply = (reply: AnyObject, commentIdx: number) => {
  comments.value[commentIdx] = {
    ...comments.value[commentIdx],
    reply_count: (comments.value[commentIdx].reply_count || 0) + 1,
    edges: {
      ...comments.value[commentIdx].edges,
      comment_replies: [
        ...comments.value[commentIdx].edges.comment_replies,
        {
          ...reply,
          edges: {
            user: {
              nickname: userStore.info?.nickname,
            },
          },
        },
      ],
    },
  }
}

onMounted(async () => {
  const { data, isFinished } = await loadComments()
  if (isFinished) {
    commentsLoading.value = false
    if (data.value?.total)
      comments.value.push(...(data.value.data || []))
    hasMoreComments.value = comments.value.length < (data.value?.total || 0)
  }
})
</script>

<template>
  <div v-if="postLoading" class="bg-#fff border-rd-4px relative p-20px">
    <Skeleton />
  </div>
  <PostItem v-else :data="post?.data || {}">
    <template #header-extra>
      <NButton strong secondary round type="primary">
        <template #icon>
          <n-icon>
            <ICarbonAdd />
          </n-icon>
        </template>
        关注
      </NButton>
    </template>
  </PostItem>
  <div v-if="commentsLoading" class="bg-#fff border-rd-4px relative p-20px m-t-16px">
    <Skeleton />
  </div>
  <div v-else class="bg-#fff border-rd-4px relative p-y-20px m-t-16px">
    <div class="p-x-20px">
      <Editor
        placeholder="输入评论"
        submit-button-text="发表评论"
        @submit="createComment"
        @submit-success="() => reloadComments()"
      />
    </div>
    <NDivider />
    <div class="p-x-20px">
      <NSpace justify="space-between" align="center" class="p-b-2">
        <span class="text-18px font-600 color-#252933">
          全部评论
        </span>
        <NTabs v-if="comments.length" type="segment" size="small" class="w-30">
          <NTab name="new" tab="最新" />
          <NTab name="hot" tab="最热" />
        </NTabs>
      </NSpace>
      <NList :show-divider="false">
        <div v-if="!comments.length" class=" min-h-40 flex items-center justify-center">
          <NEmpty size="large" description="一篇荒芜 :)" />
        </div>
        <NListItem
          v-for="(comment, index) in comments"
          v-else
          :key="comment.id"
          :show-divider="false"
        >
          <CommentItem :data="comment" :index="index" @reply-success="handleNewReply" />
        </NListItem>
      </NList>
      <InfiniteScroll v-if="!commentsLoading && comments.length" @intersect="handleIntersect" />
    </div>
  </div>
</template>

<route lang="yaml">
meta:
  layout: default
</route>
