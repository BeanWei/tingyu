<script setup lang="ts">
import { useAxios } from '@vueuse/integrations/useAxios'
import type { FormInst, FormRules } from 'naive-ui'
import type { Result } from '~/api'
import { instance, url } from '~/api'

const userStore = useUserStore()
const miscStore = useMiscStore()

const authTypeRef = ref<'login' | 'reset'>('login')
const loadingRef = ref(false)
// 登录/注册
const loginRef = ref<FormInst>()
const loginModelRef = ref({
  username: '',
  password: '',
})
// 重置密码
const resetRef = ref<FormInst>()
const resetModelRef = ref({
  username: '',
  code: '',
})
// 表单校验规则
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

const handleUpdateAuthType = (e: MouseEvent) => {
  authTypeRef.value = authTypeRef.value === 'login' ? 'reset' : 'login'
}

const handleLogin = (e: MouseEvent) => {
  e.preventDefault()
  loginRef.value?.validate(async (errors) => {
    if (!errors) {
      loadingRef.value = true
      const { data: res1 } = await useAxios<Result<AnyObject>>(url.userLogin, {
        data: {
          username: loginModelRef.value.username,
          password: loginModelRef.value.password,
        },
      }, instance)
      if (res1.value?.data.token) {
        userStore.setToken(res1.value?.data.token)
        const { data: res2 } = await useAxios<Result<AnyObject>>(url.getUserInfo, instance)
        if (res2.value?.data) {
          userStore.setInfo(res2.value?.data as any)
          loginModelRef.value = {
            username: '',
            password: '',
          }
          miscStore.setAuthModalVisible(false)
        }
      }
      loadingRef.value = false
    }
  })
}

const handleReset = (e: MouseEvent) => {
  e.preventDefault()
  resetRef.value?.validate((errors) => {
    if (!errors) {
      window.$message?.success('valid')
      loadingRef.value = true
      setTimeout(() => {
        loadingRef.value = false
      }, 20000)
    }
  })
}

onMounted(() => {
  if (userStore.token) {
    useAxios<Result<AnyObject>>(url.getUserInfo, instance).then(({ data }) => {
      if (data.value?.data)
        userStore.setInfo(data.value.data as any)
    })
  }
})
</script>

<template>
  <NModal
    :show="miscStore.authModalVisible"
    preset="card"
    :mask-closable="false"
    :bordered="false"
    :title="authTypeRef === 'login' ? '密码登录' : '重置密码'"
    class="w-400px"
    closable
    @close="() => miscStore.setAuthModalVisible(false)"
  >
    <NForm
      v-if="authTypeRef === 'login'"
      ref="loginRef"
      :model="loginModelRef"
      :rules="loginRules"
      label-placement="left"
    >
      <NFormItem path="username">
        <NInput
          v-model:value="loginModelRef.username"
          placeholder="邮箱或手机号"
        />
      </NFormItem>
      <NFormItem path="password">
        <NInput
          v-model:value="loginModelRef.password"
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
          @click="handleUpdateAuthType"
        >
          忘记密码
        </NButton>
      </div>
      <NFormItem attr-type="button">
        <NButton
          :loading="loadingRef"
          type="primary"
          block
          strong
          @click="handleLogin"
        >
          登录/注册
        </NButton>
      </NFormItem>
      <div class="text-size-12px color-#767676 m-t-2 text-center">
        未注册账号验证后自动登录
      </div>
    </NForm>
    <NForm
      v-else-if="authTypeRef === 'reset'"
      ref="resetRef"
      :model="resetModelRef"
      :rules="resetRules"
      label-placement="left"
    >
      <NFormItem path="username">
        <NInput
          v-model:value="resetModelRef.username"
          placeholder="邮箱或手机号"
        />
      </NFormItem>
      <NFormItem path="password">
        <NInput
          v-model:value="resetModelRef.code"
          placeholder="验证码"
          :maxlength="6"
        />
      </NFormItem>
      <div class="text-right m-b-2">
        <NButton
          text
          tag="a"
          type="primary"
          @click="handleUpdateAuthType"
        >
          密码登录
        </NButton>
      </div>
      <NFormItem attr-type="button">
        <NButton
          :loading="loadingRef"
          type="primary"
          block
          strong
          @click="handleReset"
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
