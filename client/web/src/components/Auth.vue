<script setup lang="ts">
import type { FormInst, FormRules } from 'naive-ui'

const miscStore = useMiscStore()

const authType = ref<'login' | 'reset'>('login')
const loading = ref(false)
// 登录/注册
const loginRef = ref<FormInst>()
const loginForm = ref({
  username: '',
  password: '',
})
// 重置密码
const resetRef = ref<FormInst>()
const resetForm = ref({
  username: '',
  code: '',
})

const loginRules: FormRules = {
  username: {
    required: true,
    message: '请输入邮箱或手机号',
  },
  password: {
    required: true,
    message: '请输入密码',
  },
}
const resetRules: FormRules = {
  username: {
    required: true,
    message: '请输入邮箱或手机号',
  },
  code: {
    required: true,
    message: '请输入验证码',
  },
}

const onUpdateAuthType = (e: MouseEvent) => {
  authType.value = authType.value === 'login' ? 'reset' : 'login'
}

const onLogin = (e: MouseEvent) => {
  e.preventDefault()
  loginRef.value?.validate((errors) => {
    if (!errors)
      window.$message?.success('valid')
  })
}
const onReset = (e: MouseEvent) => {
  e.preventDefault()
  resetRef.value?.validate((errors) => {
    if (!errors)
      window.$message?.success('valid')
  })
}
</script>

<template>
  <NModal
    :show="miscStore.authModalVisible"
    preset="card"
    :mask-closable="false"
    :bordered="false"
    :title="authType === 'login' ? '密码登录' : '重置密码'"
    class="w-400px"
    closable
    @close="() => miscStore.setAuthModalVisible(false)"
  >
    <NForm
      v-if="authType === 'login'"
      ref="loginRef"
      :model="loginForm"
      :rules="loginRules"
      label-placement="left"
    >
      <NFormItem path="username">
        <NInput
          v-model:value="loginForm.username"
          placeholder="邮箱或手机号"
        />
      </NFormItem>
      <NFormItem path="password">
        <NInput
          v-model:value="loginForm.password"
          type="password"
          show-password-on="click"
          placeholder="密码"
          :maxlength="16"
        />
      </NFormItem>
      <div class="text-right m-b-2">
        <NButton
          text
          tag="a"
          type="primary"
          @click="onUpdateAuthType"
        >
          忘记密码
        </NButton>
      </div>
      <NFormItem
        attr-type="button"
        @click="onLogin"
      >
        <NButton
          :loading="loading"
          type="primary"
          block
          strong
        >
          登录/注册
        </NButton>
      </NFormItem>
      <div class="text-size-12px color-#767676 m-t-2 text-center">
        未注册账号验证后自动登录
      </div>
    </NForm>
    <NForm
      v-else-if="authType === 'reset'"
      ref="resetRef"
      :model="resetForm"
      :rules="resetRules"
      label-placement="left"
    >
      <NFormItem path="username">
        <NInput
          v-model:value="resetForm.username"
          placeholder="邮箱或手机号"
        />
      </NFormItem>
      <NFormItem path="password">
        <NInput
          v-model:value="resetForm.code"
          placeholder="验证码"
          :maxlength="6"
        />
      </NFormItem>
      <div class="text-right m-b-2">
        <NButton
          text
          tag="a"
          type="primary"
          @click="onUpdateAuthType"
        >
          密码登录
        </NButton>
      </div>
      <NFormItem
        attr-type="button"
        @click="onReset"
      >
        <NButton
          :loading="loading"
          type="primary"
          block
          strong
        >
          重置密码
        </NButton>
      </NFormItem>
      <div class="text-size-12px color-#767676 m-t-2 text-center">
        验证码会发送至您的账号，请注意前往邮箱或者短信中查收
      </div>
    </NForm>
  </NModal>
</template>
