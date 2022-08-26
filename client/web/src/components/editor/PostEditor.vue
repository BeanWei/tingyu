<script setup lang="ts">
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
import MarkdownShortcutPlugin from './plugins/MarkdownShortcutPlugin.vue'
import defaultTheme from './themes/default'

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

const content = ref()

const onChange = (editorState: EditorState) => {
  editorState.read(() => {
    const root = $getRoot()
    if (root.getTextContent()?.trim())
      content.value = editorState.toJSON()
    else
      content.value = undefined
  })
}

const onSubmit = () => {
  window.$message?.info(JSON.stringify(content.value))
}
</script>

<template>
  <LexicalComposer :initial-config="config">
    <div class="bg-inherit relative">
      <LexicalOnChangePlugin
        @change="onChange"
      />
      <LexicalRichTextPlugin>
        <template #contentEditable>
          <LexicalContentEditable class="min-h-88px select-text break-all resize-none tab-1 outline-0 text-left" />
        </template>
        <template #placeholder>
          <div class="color-#999 overflow-hidden absolute text-ellipsis top-0px left-0px select-none inline-block pointer-events-none">
            快和水友一起分享新鲜事~
          </div>
        </template>
      </LexicalRichTextPlugin>
      <MarkdownShortcutPlugin />
    </div>
  </LexicalComposer>
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
    <NButton
      strong secondary round type="primary"
      :disabled="!!!userStore.info || !!!content"
      @click="onSubmit"
    >
      发布
    </NButton>
  </NSpace>
</template>
