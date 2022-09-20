import type {
  DOMConversionMap,
  DOMConversionOutput,
  DOMExportOutput,
  LexicalNode,
  NodeKey,
  SerializedTextNode,
  Spread,
} from 'lexical'
import { TextNode } from 'lexical'

export type SerializedMentionNode = Spread<
  {
    mentionName: string
    type: 'mention'
    version: 1
  },
  SerializedTextNode
>

function convertMentionElement(
  domNode: HTMLElement,
): DOMConversionOutput | null {
  const textContent = domNode.textContent

  if (textContent !== null) {
    const node = $createMentionNode(textContent)
    return {
      node,
    }
  }

  return null
}

const mentionStyle = 'color: #18a058; text-decoration: none;'
export class MentionNode extends TextNode {
  __mention: string

  static getType(): string {
    return 'mention'
  }

  static clone(node: MentionNode): MentionNode {
    return new MentionNode(node.__mention, node.__text, node.__key)
  }

  static importJSON(serializedNode: SerializedMentionNode): MentionNode {
    const node = $createMentionNode(serializedNode.mentionName)
    node.setTextContent(serializedNode.text)
    node.setFormat(serializedNode.format)
    node.setDetail(serializedNode.detail)
    node.setMode(serializedNode.mode)
    node.setStyle(serializedNode.style)
    return node
  }

  constructor(mentionName: string, text?: string, key?: NodeKey) {
    super(text ?? mentionName, key)
    this.__mention = mentionName
  }

  exportJSON(): SerializedMentionNode {
    return {
      ...super.exportJSON(),
      mentionName: this.__mention,
      type: 'mention',
      version: 1,
    }
  }

  createDOM(): HTMLElement {
    const element = document.createElement('a')
    element.style.cssText = mentionStyle
    element.setAttribute('href', `/?topic_id=${this.__mention}`)
    element.setAttribute('target', '__blank')
    const inner = document.createElement('span')
    inner.textContent = this.__text
    element.appendChild(inner)
    return element
  }

  exportDOM(): DOMExportOutput {
    const element = document.createElement('span')
    element.setAttribute('data-lexical-mention', 'true')
    element.textContent = this.__text
    return { element }
  }

  static importDOM(): DOMConversionMap | null {
    return {
      span: (domNode: HTMLElement) => {
        if (!domNode.hasAttribute('data-lexical-mention'))
          return null

        return {
          conversion: convertMentionElement,
          priority: 1,
        }
      },
    }
  }

  isTextEntity(): true {
    return true
  }
}

export function $createMentionNode(mentionName: string, text?: string): MentionNode {
  const mentionNode = new MentionNode(mentionName, text)
  mentionNode.setMode('segmented').toggleDirectionless()
  return mentionNode
}

export function $isMentionNode(
  node: LexicalNode | null | undefined,
): node is MentionNode {
  return node instanceof MentionNode
}
