<script setup lang="ts">
import { CLEAR_EDITOR_COMMAND } from 'lexical'
import { useEditor } from 'lexical-vue'

const props = defineProps<{
  disabled?: boolean
  loading?: boolean
  buttonText?: string
  onSubmit?: () => Promise<boolean | undefined>
}>()

const editor = useEditor()

const handleSubmit = async () => {
  if (!props.onSubmit)
    return
  const ok = await props.onSubmit()
  if (ok)
    editor.dispatchCommand(CLEAR_EDITOR_COMMAND, undefined)
}
</script>

<template>
  <NButton
    strong secondary round type="primary"
    :disabled="props.disabled"
    :loading="props.loading"
    @click="handleSubmit"
  >
    {{ props.buttonText || '' }}
  </NButton>
</template>
