<script setup lang="ts">
import type { StrictUseAxiosReturn } from '@vueuse/integrations/useAxios'
import type { EditorState, LexicalEditor } from 'lexical'
import type { AxiosError } from 'axios'
import { $getRoot } from 'lexical'
import {
  LexicalAutoFocusPlugin,
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
import { MentionNode } from './nodes/MentionNode'
import { EmojiNode } from './nodes/EmojiNode'
import { ImageNode } from './nodes/ImageNode'
import { extractMentionIds } from '~/utils/lexical'

const props = withDefaults(defineProps<{
  initialState?: string
  autofocus?: boolean
  readOnly?: boolean
  placeholder?: string
  submitButtonText?: string
  pluginsConfig?: {
    mentions: {
      triggers: ('@' | '#')[]
    }
  }
  onChange?: (editorState: EditorState, editor: LexicalEditor) => void
  onSubmit?: (values: any) => PromiseLike<StrictUseAxiosReturn<any>>
  onSubmitSuccess?: (data: AnyObject) => void
  onSubmitFailed?: (error: AxiosError) => void
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
    MentionNode,
    EmojiNode,
    ImageNode,
  ],
  readOnly: props.readOnly,
  onError(error: Error) {
    throw error
  },
}

const userStore = useUserStore()

const submitting = ref(false)
const content = ref()
const contentText = ref('')

const onChange = (editorState: EditorState, editor: LexicalEditor) => {
  editorState.read(() => {
    const text = $getRoot().getTextContent()?.trim()
    if (text) {
      content.value = editorState.toJSON()
      contentText.value = text
    }
    else {
      content.value = undefined
      contentText.value = ''
    }
  })
  props.onChange?.(editorState, editor)
}

const handleSubmit = async (): Promise<boolean | undefined> => {
  if (!props.onSubmit)
    return
  submitting.value = true

  const [topic_ids, user_ids] = extractMentionIds(content.value.root)
  const { data, isFinished, error } = await props.onSubmit({
    content: JSON.stringify(content.value),
    content_text: contentText.value,
    topic_ids,
    user_ids,
  })
  if (isFinished) {
    submitting.value = false
    if (error.value) {
      props.onSubmitFailed?.(error.value)
    }
    else {
      content.value = undefined
      contentText.value = ''
      props.onSubmitSuccess?.(data.value.data)
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
      <LexicalRichTextPlugin
        :initial-editor-state="props.initialState"
      >
        <template #contentEditable>
          <LexicalContentEditable
            class="select-text break-all resize-none tab-1 outline-0 text-left border-rd-4px box-border important-focus-bg-#fff focus-b focus-b-color-#18a058"
            :style="props.readOnly ? {} : { 'min-height': '88px', 'background-color': '#f2f3f5', 'padding': '8px 12px' }"
          />
        </template>
        <template #placeholder>
          <div class="color-#999 overflow-hidden absolute text-ellipsis top-0px left-0px select-none inline-block pointer-events-none p-x-3.5 p-y-2">
            {{ props.placeholder }}
          </div>
        </template>
      </LexicalRichTextPlugin>
      <LexicalAutoFocusPlugin v-if="props.autofocus" />
      <LexicalMarkdownShortcutPlugin />
      <LexicalMentionsPlugin :triggers="props.pluginsConfig?.mentions.triggers" />
      <LexicalEmojisPlugin />
      <LexicalImagePlugin />
    </div>
    <NSpace v-if="!!!props.readOnly" justify="space-between" class="m-t-2">
      <NSpace>
        <LexicalImageMenu />
        <LexicalEmojiMenu />
      </NSpace>
      <EditorSubmit
        :disabled="!!!userStore.info || !!!contentText"
        :loading="submitting"
        :button-text="props.submitButtonText"
        @submit="handleSubmit"
      />
    </NSpace>
  </LexicalComposer>
</template>
