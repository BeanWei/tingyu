import type {
  DOMExportOutput,
  EditorConfig,
  LexicalCommand,
  LexicalNode,
  NodeKey,
  SerializedLexicalNode,
  Spread,
} from 'lexical'
import { DecoratorNode, createCommand } from 'lexical'

export type SerializedEmojiNode = Spread<
  {
    emojiName: string
    type: 'emoji'
    version: 1
  },
  SerializedLexicalNode
>

export const INSERT_EMOJI_COMMAND: LexicalCommand<string> = createCommand()

const emojiStyle = 'display: inline-block; background-position: center; background-repeat: no-repeat; vertical-align: text-bottom;'
const emojiInnerStyle = 'height: 20px; width: 20px; margin: 0 1px'

export class EmojiNode extends DecoratorNode<null> {
  __className: string

  static getType(): string {
    return 'emoji'
  }

  static clone(node: EmojiNode): EmojiNode {
    return new EmojiNode(node.__className, node.__key)
  }

  static importJSON(serializedNode: SerializedEmojiNode): EmojiNode {
    const node = $createEmojiNode(serializedNode.emojiName)
    return node
  }

  constructor(emojiName: string, key?: NodeKey) {
    super(key)
    this.__className = emojiName
  }

  decorate(): null {
    return null
  }

  exportJSON(): SerializedEmojiNode {
    return {
      emojiName: this.__className,
      type: 'emoji',
      version: 1,
    }
  }

  createDOM(_config: EditorConfig): HTMLElement {
    const dom = document.createElement('span')
    dom.style.cssText = emojiStyle
    const inner = document.createElement('div')
    inner.className = this.__className
    inner.style.cssText = emojiInnerStyle
    dom.appendChild(inner)
    return dom
  }

  exportDOM(): DOMExportOutput {
    const dom = document.createElement('span')
    dom.setAttribute('data-lexical-emoji', 'true')
    dom.style.cssText = emojiStyle
    const inner = document.createElement('div')
    inner.className = this.__className
    inner.style.cssText = emojiInnerStyle
    dom.appendChild(inner)
    return { element: dom }
  }
}

export function $isEmojiNode(
  node: LexicalNode | null | undefined,
): node is EmojiNode {
  return node instanceof EmojiNode
}

export function $createEmojiNode(emojiName: string): EmojiNode {
  return new EmojiNode(emojiName)
}
