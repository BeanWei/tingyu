<script setup lang="ts">
import type { StrictUseAxiosReturn } from '@vueuse/integrations/useAxios'
import type { EditorState } from 'lexical'
import { $getRoot } from 'lexical'
import {
  LexicalComposer,
  LexicalContentEditable,
  LexicalOnChangePlugin,
  LexicalRichTextPlugin,
} from 'lexical-vue'
import { HeadingNode, QuoteNode } from '@lexical/rich-text'
import { TableCellNode, TableNode, TableRowNode } from '@lexical/table'
import { ListItemNode, ListNode } from '@lexical/list'
import { CodeHighlightNode, CodeNode } from '@lexical/code'
import { AutoLinkNode, LinkNode } from '@lexical/link'
import { HashtagNode } from '@lexical/hashtag'
import defaultTheme from './themes/default'

const props = withDefaults(defineProps<{
  placeholder?: string
  submitButtonText?: string
  onSubmit?: (values: {
    content: string
    content_text: string
  }) => PromiseLike<StrictUseAxiosReturn<any>>
  onSubmitSuccess?: () => void
  onSubmitFailed?: () => void
}>(), {
  placeholder: '说点什么吧~',
})

const config = {
  theme: defaultTheme,
  nodes: [
    HeadingNode,
    ListNode,
    ListItemNode,
    QuoteNode,
    CodeNode,
    CodeHighlightNode,
    TableNode,
    TableCellNode,
    TableRowNode,
    AutoLinkNode,
    LinkNode,
    HashtagNode,
  ],
  onError(error: Error) {
    throw error
  },
}

const userStore = useUserStore()

const submittingRef = ref(false)
const contentRef = ref()
const contentTextRef = ref('')

const onChange = (editorState: EditorState) => {
  editorState.read(() => {
    const text = $getRoot().getTextContent()?.trim()
    if (text) {
      contentRef.value = editorState.toJSON()
      contentTextRef.value = text
    }
    else {
      contentRef.value = undefined
      contentTextRef.value = ''
    }
  })
}

const handleSubmit = async (): Promise<boolean | undefined> => {
  if (!props.onSubmit)
    return
  submittingRef.value = true

  const { isFinished, error } = await props.onSubmit({
    content: JSON.stringify(contentRef.value),
    content_text: contentTextRef.value,
  })
  if (isFinished) {
    submittingRef.value = false
    if (error.value) {
      props.onSubmitFailed?.()
    }
    else {
      contentRef.value = undefined
      contentTextRef.value = ''
      props.onSubmitSuccess?.()
      return true
    }
  }
}
</script>

<template>
  <LexicalComposer :initial-config="config">
    <div class="bg-inherit relative">
      <LexicalOnChangePlugin
        @change="onChange"
      />
      <LexicalClearEditorPlugin />
      <LexicalRichTextPlugin>
        <template #contentEditable>
          <LexicalContentEditable class="min-h-88px select-text break-all resize-none tab-1 outline-0 text-left" />
        </template>
        <template #placeholder>
          <div class="color-#999 overflow-hidden absolute text-ellipsis top-0px left-0px select-none inline-block pointer-events-none">
            {{ props.placeholder }}
          </div>
        </template>
      </LexicalRichTextPlugin>
      <MarkdownShortcutPlugin />
    </div>
    <NSpace justify="space-between">
      <NSpace>
        <NButton quaternary circle>
          <template #icon>
            <NIcon><ICarbonImage /></NIcon>
          </template>
        </NButton>
        <NButton quaternary circle>
          <template #icon>
            <NIcon><ICarbonFaceAdd /></NIcon>
          </template>
        </NButton>
      </NSpace>
      <EditorSubmit
        :disabled="!!!userStore.info || !!!contentTextRef"
        :loading="submittingRef"
        :button-text="props.submitButtonText"
        @submit="handleSubmit"
      />
    </NSpace>
  </LexicalComposer>
</template>
