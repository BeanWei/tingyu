<script setup lang="ts">
import { useEditor } from 'lexical-vue'
import type { UploadFileInfo } from 'naive-ui'
import { INSERT_IMAGE_COMMAND } from '../nodes/ImageNode'
import { fileUpload } from '~/api'

const editor = useEditor()

const handleUploadFinish = (options: { file: UploadFileInfo; event?: ProgressEvent }) => {
  editor.update(() => {
    editor.dispatchCommand(INSERT_IMAGE_COMMAND, options.file.url as string)
  })
}
</script>

<template>
  <NUpload
    abstract
    :max="1"
    :custom-request="fileUpload"
    @finish="handleUploadFinish"
  >
    <NUploadTrigger #="{ handleClick }" abstract>
      <NButton quaternary circle @click="handleClick">
        <template #icon>
          <NIcon><ICarbonImage /></NIcon>
        </template>
      </NButton>
    </NUploadTrigger>
  </NUpload>
</template>
