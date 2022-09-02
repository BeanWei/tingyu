<script setup lang="ts">
import { useAxios } from '@vueuse/integrations/useAxios'
import type { FormInst, FormRules } from 'naive-ui'
import { instance, url } from '~/api'

const props = defineProps<{
  show?: boolean
  onClose?: () => void
}>()

const submitting = ref(false)
const topicForm = ref<FormInst>()
const topicModel = ref({
  title: '',
  icon: '',
  description: '',
})
const topicRules: FormRules = {
  title: {
    required: true,
    message: '请输入标题',
  },
}

const createTopic = (e: MouseEvent) => {
  e.preventDefault()
  topicForm.value?.validate(async (errors) => {
    if (errors)
      return
    submitting.value = true
    const { isFinished, error } = await useAxios(url.createTopic, {
      data: {
        title: topicModel.value.title,
        icon: topicModel.value.icon,
        description: topicModel.value.description,
      },
    }, instance)
    if (isFinished) {
      submitting.value = false
      if (!error.value)
        props.onClose?.()
    }
  })
}
</script>

<template>
  <NModal
    :show="props.show"
    preset="card"
    :mask-closable="false"
    :bordered="false"
    title="创建新话题"
    class="w-400px"
    closable
    @close="props.onClose"
  >
    <NForm
      ref="topicForm"
      :model="topicModel"
      :rules="topicRules"
      label-placement="left"
    >
      <NFormItem path="title">
        <NInput
          v-model:value="topicModel.title"
          placeholder="话题名称"
        />
      </NFormItem>
      <NFormItem path="description">
        <NInput
          v-model:value="topicModel.description"
          type="textarea"
          placeholder="简单介绍一下话题吧"
        />
      </NFormItem>
      <NFormItem attr-type="button">
        <NButton
          :loading="submitting"
          type="primary"
          block
          strong
          @click="createTopic"
        >
          提交
        </NButton>
      </NFormItem>
    </NForm>
  </NModal>
</template>
