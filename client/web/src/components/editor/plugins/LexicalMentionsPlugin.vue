<script setup lang="ts">
import type { TextNode } from 'lexical'
import { $getSelection, $isRangeSelection } from 'lexical'
import { useEditor } from 'lexical-vue'
import { useAxios } from '@vueuse/integrations/useAxios'
import type { DropdownOption } from 'naive-ui'
import { NEmpty, NSpin } from 'naive-ui'
import { $createMentionNode } from '../nodes/MentionNode'
import type { Result } from '~/api'
import { instance, url } from '~/api'

const { triggers = ['@'] } = defineProps<{
  triggers?: string[]
}>()

const dropdownState = ref<{
  x: number
  y: number
  show: boolean
}>({
  x: 0,
  y: 0,
  show: false,
})
const newMentionState = ref<{
  anchorNode: TextNode
  trigger: string
  startOffset: number
  selectionOffset: number
}>()

const editor = useEditor()

let removeUpdateListener: () => void

const { data, isFinished, execute } = useAxios<Result<AnyObject[]>>('', {}, instance, { immediate: false })
const genOptions = (): DropdownOption[] => {
  if (!isFinished) {
    return [{
      key: 'loading',
      type: 'render',
      render: () => {
        return h(
          NSpin,
          { size: 'small', style: { width: '100px' } },
        )
      },
    }]
  }
  if (!data.value?.data?.length) {
    return [{
      key: 'loading',
      type: 'render',
      render: () => {
        return h(
          NEmpty,
          { style: { width: '100px' } },
        )
      },
    }]
  }
  return data.value.data.map((item) => {
    return {
      key: item.id,
      label: item.nickname || item.title,
    }
  })
}

const handleSuggestionSelect = (key: string | number, option: DropdownOption) => {
  if (!newMentionState.value || !newMentionState.value.anchorNode)
    return
  editor.update(() => {
    const anchorNode = newMentionState.value?.anchorNode
    if (!anchorNode)
      return
    const startOffset = newMentionState.value?.startOffset || 0
    const selectionOffset = newMentionState.value?.selectionOffset || 0
    let newNode: TextNode
    if (startOffset === 0)
      [newNode] = anchorNode.splitText(selectionOffset)
    else
      [, newNode] = anchorNode.splitText(startOffset, selectionOffset)
    const mentionNode = $createMentionNode(`${key}`, `${newMentionState.value?.trigger || ''}${option.label}`)
    if (newNode)
      newNode.replace(mentionNode)
    mentionNode.select()
    dropdownState.value.show = false
  })
}

onMounted(() => {
  const updateListener = () => {
    editor.getEditorState().read(() => {
      const selection = $getSelection()
      if (!$isRangeSelection(selection)) {
        dropdownState.value.show = false
        return
      }
      const anchor = selection.anchor
      if (anchor.type !== 'text') {
        dropdownState.value.show = false
        return
      }
      const anchorNode = anchor.getNode()
      if (!anchorNode.isSimpleText()) {
        dropdownState.value.show = false
        return
      }
      const text = anchorNode.getTextContent().slice(0, anchor.offset)
      const matchString = text.match(/\B[@#]\w*$/g)?.[0]
      if (!matchString) {
        dropdownState.value.show = false
        return
      }
      const triggerString = matchString[0]
      if (!triggers.includes(triggerString)) {
        dropdownState.value.show = false
        return
      }
      const domElement = editor.getElementByKey(anchorNode.getKey())
      if (!domElement) {
        dropdownState.value.show = false
        return
      }
      const rect = domElement.getBoundingClientRect()
      dropdownState.value = {
        x: rect.x + rect.width,
        y: rect.y + 12,
        show: true,
      }
      // TODO: @-搜索用户, #-搜索话题
      execute(triggerString === '@' ? url.searchTopic : url.searchTopic, {
        params: {
          keyword: matchString.slice(1) || undefined,
        },
      })
      // 计算新node插入所需的偏移量
      const selectionOffset = anchor.offset
      const textContent = anchorNode.getTextContent().slice(0, selectionOffset)
      let queryOffset = matchString.length
      for (let i = queryOffset; i <= (matchString.length - 1); i++) {
        if (textContent.slice(-i) === matchString.slice(1, i))
          queryOffset = i
      }
      const startOffset = selectionOffset - queryOffset
      if (startOffset < 0)
        return
      newMentionState.value = {
        anchorNode,
        trigger: triggerString,
        startOffset,
        selectionOffset,
      }
    })
  }

  removeUpdateListener = editor.registerUpdateListener(updateListener)
})

onUnmounted(() => {
  removeUpdateListener?.()
})
</script>

<template>
  <NDropdown
    placement="bottom-start"
    trigger="manual"
    :x="dropdownState.x"
    :y="dropdownState.y"
    :show="dropdownState.show"
    :options="genOptions()"
    @select="handleSuggestionSelect"
    @clickoutside="() => {
      dropdownState.show = false
    }"
  />
</template>
