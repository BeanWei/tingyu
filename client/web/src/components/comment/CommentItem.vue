<script setup lang="ts">
import { useAxios } from '@vueuse/integrations/useAxios'
import { instance, url } from '~/api'

const { data, index, onReplySuccess } = defineProps<{
  data: AnyObject
  index: number
  onReplySuccess?: (reply: AnyObject, index: number) => void
}>()

const editorVisible = ref(false)
const changeEditorVisible = () => {
  editorVisible.value = !editorVisible.value
}

const createReply = (commentId: number) => {
  return function (values: AnyObject) {
    return useAxios(url.createReply, {
      data: {
        ...values,
        comment_id: commentId,
      },
    }, instance)
  }
}
const handleNewReply = (reply: AnyObject) => {
  editorVisible.value = false
  onReplySuccess?.(reply, index)
}
</script>

<template>
  <NThing content-indented>
    <template #avatar>
      <UserAvatar
        :size="24"
        :src="data.edges.user?.avatar"
        class="flex relative border-neutral-200/70"
      />
    </template>
    <template #header>
      <a class="font-medium text-16px color-#252933 decoration-none cursor-pointer"> {{ data.edges.user?.nickname }} </a>
    </template>
    <template #description>
      <Editor :read-only="true" :initial-state="data.content" />
    </template>
    <template #footer>
      <SubjectAction :data="data" :react-action="url.reactComment">
        <template #left>
          <div
            class="flex items-center cursor-pointer"
            @click="changeEditorVisible"
          >
            <NIcon size="18">
              <ICarbonChat />
            </NIcon>
            <span class="m-l-0.3em text-12px font-medium">{{ editorVisible ? '取消回复' : data.reply_count || '回复' }}</span>
          </div>
        </template>
        <template #right>
          <CreationInfo :time="data.created_at" :location="data.ip_loc" />
        </template>
      </SubjectAction>
      <div v-if="editorVisible" class="m-t-4 m-b6">
        <Editor
          :placeholder="`回复 ${data.edges.user?.nickname}...`"
          submit-button-text="发布"
          :autofocus="true"
          :on-submit="createReply(data.id)"
          @submit-success="handleNewReply"
        />
      </div>
      <div v-if="data.edges.comment_replies" class="m-t-3 p-x-4 bg-#f7f8fa border-rd-1">
        <NList :show-divider="false">
          <NListItem
            v-for="reply in data.edges.comment_replies"
            :key="reply.id"
            :show-divider="false"
            class="bg-#f7f8fa"
          >
            <ReplyItem :data="reply" :index="index" :on-reply-success="onReplySuccess" />
          </NListItem>
        </NList>
      </div>
    </template>
  </NThing>
</template>
