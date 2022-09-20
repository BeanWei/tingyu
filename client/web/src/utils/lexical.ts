import type { LexicalNode } from 'lexical'
import { $createParagraphNode, $getRoot, $getSelection, $isNodeSelection, $isRangeSelection } from 'lexical'

export function extractMentionIds(obj: AnyObject): [number[], number[]] {
  const userIds = []
  const topicIds = []
  if (obj.children) {
    for (const child of obj.children) {
      if (child.text && child.type === 'mention' && child.mentionName) {
        if (child.text.startsWith('#'))
          topicIds.push(parseInt(child.mentionName))
        else if (child.text.startsWith('@'))
          userIds.push(parseInt(child.mentionName))
      }
      if (child.children) {
        const ids = extractMentionIds(child)
        topicIds.push(...ids[0])
        userIds.push(...ids[1])
      }
    }
  }
  return [topicIds, userIds]
}

export function $insertNodeToNearestRoot<T extends LexicalNode>(node: T): T {
  const selection = $getSelection()
  if ($isRangeSelection(selection)) {
    const focusNode = selection.focus.getNode()
    focusNode.getTopLevelElementOrThrow().insertAfter(node)
  }
  else if ($isNodeSelection(selection)) {
    const nodes = selection.getNodes()
    nodes[nodes.length - 1].getTopLevelElementOrThrow().insertAfter(node)
  }
  else {
    const root = $getRoot()
    root.append(node)
  }
  const paragraphNode = $createParagraphNode()
  node.insertAfter(paragraphNode)
  paragraphNode.select()
  return node.getLatest()
}
