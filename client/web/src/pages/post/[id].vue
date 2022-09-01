<script setup lang="ts">
import { useAxios } from '@vueuse/integrations/useAxios'
import type { Result } from '~/api'
import { instance, url } from '~/api'

const props = defineProps<{ id: string }>()

const { data: post, isLoading: postLoading } = useAxios<Result<AnyObject>>(url.getPost, {
  params: {
    id: props.id,
  },
}, instance)

const { data: comments, isLoading: commentsLoading, execute: reloadComments } = useAxios<Result<AnyObject[]>>(url.listComment, {
  params: {
    post_id: props.id,
  },
}, instance)

const createComment = (values: AnyObject) => {
  return useAxios(url.createComment, {
    data: {
      ...values,
      post_id: props.id,
    },
  }, instance)
}
</script>

<template>
  <div v-if="postLoading" class="bg-#fff border-rd-4px relative p-20px">
    <Skeleton />
  </div>
  <PostItem v-else :data="post?.data || {}" />
  <div v-if="commentsLoading" class="bg-#fff border-rd-4px relative p-y-20px m-t-16px">
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
      <NSpace justify="space-between" class="p-b-2">
        <span class="text-18px font-600 color-#252933">
          全部评论
        </span>
        <NSwitch v-if="comments?.total">
          <template #checked>
            最新
          </template>
          <template #unchecked>
            最热
          </template>
        </NSwitch>
      </NSpace>
      <NList :show-divider="false">
        <div v-if="!comments?.total" class=" min-h-40 flex items-center justify-center">
          <NEmpty size="large" description="一篇荒芜 :)" />
        </div>
        <NListItem
          v-for="comment in comments?.data"
          v-else
          :key="comment.id"
          :show-divider="false"
        >
          <CommentItem :data="comment" />
        </NListItem>
      </NList>
    </div>
  </div>
</template>

<route lang="yaml">
meta:
  layout: default
</route>
