<script setup lang="ts">
import { useAxios } from '@vueuse/integrations/useAxios'
import type { FormInst, FormRules, UploadFileInfo } from 'naive-ui'
import type { Result } from '~/api'
import { fileUpload, instance, url } from '~/api'

const userStore = useUserStore()

const loading = ref(false)

const profileForm = ref<FormInst>()
const profileModel = ref({
  avatar: userStore.info?.avatar,
  nickname: userStore.info?.nickname,
  headline: userStore.info?.headline,
})
const profileRules: FormRules = {
  nickname: {
    required: true,
    message: '请输入昵称',
  },
}
const handleProfileUpdate = (e: MouseEvent) => {
  e.preventDefault()
  profileForm.value?.validate(async (errors) => {
    if (!errors) {
      loading.value = true
      const { error } = await useAxios<Result<AnyObject>>(url.updateUserInfo, {
        data: {
          avatar: profileModel.value.avatar,
          nickname: profileModel.value.nickname,
          headline: profileModel.value.headline,
        },
      }, instance)
      if (!error && userStore.info) {
        window.$message?.success('修改成功')
        userStore.info.avatar = profileModel.value.avatar || ''
        userStore.info.nickname = profileModel.value.nickname || ''
        userStore.info.headline = profileModel.value.headline || ''
      }
      loading.value = false
    }
  })
}

const handleAvatarUploadFinish = (options: { file: UploadFileInfo; event?: ProgressEvent }) => {
  profileModel.value.avatar = options.file.url as string
  return options.file
}
</script>

<template>
  <div class="bg-#fff box-border border-rd-4px relative p-20px">
    <NTabs
      default-value="profile"
      size="large"
      animated
    >
      <NTabPane name="profile" tab="个人资料">
        <NForm
          ref="profileForm"
          :model="profileModel"
          :rules="profileRules"
        >
          <NFormItemRow label="头像" path="avatar">
            <NUpload
              list-type="image-card"
              :max="1"
              :custom-request="fileUpload"
              :default-file-list="profileModel.avatar ? [{ id: profileModel.avatar, name: profileModel.avatar, status: 'finished', url: profileModel.avatar }] : []"
              @finish="handleAvatarUploadFinish"
            />
          </NFormItemRow>
          <NFormItemRow label="昵称" path="nickname" required>
            <NInput v-model:value="profileModel.nickname" />
          </NFormItemRow>
          <NFormItemRow label="个人简介" path="headline">
            <NInput v-model:value="profileModel.headline" />
          </NFormItemRow>
          <NButton
            :loading="loading"
            type="primary"
            secondary
            strong
            @click="handleProfileUpdate"
          >
            保存修改
          </NButton>
        </NForm>
      </NTabPane>
    </NTabs>
  </div>
</template>
