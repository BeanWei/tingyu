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

const createReply = (commentId: number, toUserId: number, toReplyId: number) => {
  return function (values: AnyObject) {
    return useAxios(url.createReply, {
      data: {
        ...values,
        comment_id: commentId,
        to_user_id: toUserId,
        to_reply_id: toReplyId,
      },
    }, instance)
  }
}

const handleSubmitSuccess = (data: AnyObject) => {
  editorVisible.value = false
  onReplySuccess?.(data, index)
}
</script>

<template>
  <NThing content-indented>
    <template #avatar>
      <UserAvatar
        :size="24"
        :src="data.user?.avatar"
        class="flex relative border-neutral-200/70"
      />
    </template>
    <template #header>
      <a class="font-medium text-16px color-#252933 decoration-none cursor-pointer"> {{ data.user?.nickname }} </a>
    </template>
    <template #description>
      <Editor :read-only="true" :initial-state="data.content" />
    </template>
    <template #footer>
      <SubjectAction :data="data" :react-action="url.reactReply">
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
      <div v-if="editorVisible" class="m-t-4">
        <Editor
          :placeholder="`回复 ${data.user?.nickname}...`"
          submit-button-text="发布"
          :autofocus="true"
          :on-submit="createReply(data.comment_id, data.user_id, data.id)"
          :on-submit-success="handleSubmitSuccess"
        />
      </div>
    </template>
  </NThing>
</template>
