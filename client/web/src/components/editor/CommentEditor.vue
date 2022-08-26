<script setup lang="ts">
import {
  LexicalComposer,
  LexicalContentEditable,
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
}

const onError = (error: Error) => {
  throw error
}
</script>

<template>
  <LexicalComposer :initial-config="config" @error="onError">
    <div class="bg-inherit relative">
      <LexicalRichTextPlugin>
        <template #contentEditable>
          <LexicalContentEditable class="min-h-88px select-text break-all resize-none tab-1 outline-0 text-left" />
        </template>
        <template #placeholder>
          <div class="color-#999 overflow-hidden absolute text-ellipsis top-0px left-0px select-none inline-block pointer-events-none">
            输入评论
          </div>
        </template>
      </LexicalRichTextPlugin>
      <MarkdownShortcutPlugin />
    </div>
  </LexicalComposer>
  <n-space justify="space-between">
    <n-space>
      <n-button quaternary circle>
        <template #icon>
          <n-icon><i-carbon-image /></n-icon>
        </template>
      </n-button>
      <n-button quaternary circle>
        <template #icon>
          <n-icon><i-carbon-face-add /></n-icon>
        </template>
      </n-button>
    </n-space>
    <n-button
      strong secondary round type="primary"
    >
      发表评论
    </n-button>
  </n-space>
</template>
