<script setup lang="ts">
import type { DropdownOption, MenuOption } from 'naive-ui'
import ICarbonUser from '~icons/carbon/user'
import ICarbonSettings from '~icons/carbon/settings'
import ICarbonLogout from '~icons/carbon/logout'
import { menuLabelRender } from '~/utils/ui'

const menuOptions: MenuOption[] = [
  {
    label: '首页',
    path: '/',
    key: 'index',
  },
  {
    label: '关于',
    key: 'about',
  },
]

const userOptions: DropdownOption[] = [
  {
    label: '我的主页',
    key: 'profile',
    icon: () => h(ICarbonUser),
  },
  {
    label: '我的设置',
    key: 'settings',
    icon: () => h(ICarbonSettings),
  },
  {
    label: '退出登录',
    key: 'logout',
    icon: () => h(ICarbonLogout),
  },
]

const userStore = useUserStore()
const miscStore = useMiscStore()
const route = useRoute()
const router = useRouter()

const activeKey = ref(route.name as string)
const searchWord = ref('')

const handleSearch = () => {
  if (searchWord.value)
    router.push(`/?keyword=${encodeURIComponent(searchWord.value)}`)
  else
    router.push('/')
}

const handleUserDropdownSelect = (key: string) => {
  switch (key) {
    case 'profile':
      router.push(`/user/${userStore.info?.id}`)
      return
    case 'settings':
      router.push('/user/settings')
      return
    case 'logout':
      userStore.resetAll()
      router.push('/')
  }
}

watch(route, () => {
  activeKey.value = route.name as string
})
</script>

<template>
  <div class="relative h-54px">
    <header class="border-b-1 border-b-#f1f1f1 color-#909090 z-250 fixed top-0 left-0 right-0 bg-#fff">
      <div class="p-y-6px m-auto max-w-1240px flex items-center h-full w-full relative">
        <nav class="h-full flex-auto">
          <NSpace justify="space-between" align="center">
            <NMenu
              :value="activeKey"
              mode="horizontal"
              :options="menuOptions"
              :render-label="menuLabelRender"
            />
            <NSpace align="center">
              <NInput
                v-model:value="searchWord"
                clearable
                round
                placeholder="搜索"
                @keyup.enter.prevent="handleSearch"
              >
                <template #suffix>
                  <ICarbonSearch />
                </template>
              </NInput>
              <NSpace v-if="!!userStore.info">
                <NButton quaternary circle>
                  <template #icon>
                    <NIcon><ICarbonNotification /></NIcon>
                  </template>
                </NButton>
                <NDropdown
                  trigger="click"
                  :options="userOptions"
                  placement="bottom-end"
                  @select="handleUserDropdownSelect"
                >
                  <UserAvatar :src="userStore.info?.avatar" :size="32" />
                </NDropdown>
              </NSpace>
              <NButton
                v-else
                strong
                secondary
                round
                type="primary"
                @click="miscStore.setAuthModalVisible(true)"
              >
                登录/注册
              </NButton>
            </NSpace>
          </NSpace>
        </nav>
      </div>
    </header>
  </div>
</template>
