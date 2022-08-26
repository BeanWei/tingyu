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
            Enter some text...
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
    >
      发布
    </NButton>
  </NSpace>
</template>
