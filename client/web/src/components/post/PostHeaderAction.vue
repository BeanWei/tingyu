<script setup lang="ts">
import type { DropdownOption } from 'naive-ui'

const options: DropdownOption[] = [
  {
    label: '编辑',
    key: 'edit',
  },
  {
    label: '删除',
    key: 'delete',
  },
]
</script>

<template>
  <NDropdown
    trigger="hover"
    :options="options"
    :show-arrow="true"
    placement="bottom-end"
  >
    <NButton quaternary circle>
      <template #icon>
        <NIcon><ICarbonOverflowMenuVertical /></NIcon>
      </template>
    </NButton>
  </NDropdown>
</template>
