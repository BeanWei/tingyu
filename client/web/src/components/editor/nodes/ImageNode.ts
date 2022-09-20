import type { Component } from 'vue'
import type {
  DOMConversionMap,
  DOMConversionOutput,
  DOMExportOutput,
  LexicalCommand,
  LexicalNode,
  NodeKey,
  SerializedLexicalNode,
  Spread,
} from 'lexical'
import { createCommand } from 'lexical'
import { NImage } from 'naive-ui'
import { DecoratorBlockNode } from 'lexical-vue'

function convertImageElement(domNode: Node): null | DOMConversionOutput {
  if (domNode instanceof HTMLImageElement) {
    const node = $createImageNode(domNode.src)
    return { node }
  }
  return null
}

export type SerializedImageNode = Spread<
  {
    src: string
    text: string
    type: 'image'
    version: 1
  },
  SerializedLexicalNode
>

export const INSERT_IMAGE_COMMAND: LexicalCommand<string> = createCommand()

export class ImageNode extends DecoratorBlockNode {
  __src: string

  static getType(): string {
    return 'image'
  }

  static clone(node: ImageNode): ImageNode {
    return new ImageNode(node.__src, node.__key)
  }

  static importJSON(serializedNode: SerializedImageNode): ImageNode {
    const node = $createImageNode(serializedNode.src)
    return node
  }

  static importDOM(): DOMConversionMap | null {
    return {
      img: () => ({
        conversion: convertImageElement,
        priority: 0,
      }),
    }
  }

  constructor(src: string, key?: NodeKey) {
    super(key)
    this.__src = src
  }

  decorate(): Component {
    return h(NImage, {
      src: this.__src,
      objectFit: 'contain',
      imgProps: {
        style: {
          'width': '100%',
          'height': 'auto',
          'max-height': '200px',
          'position': 'relative',
          'display': 'block',
          'vertical-align': 'middle',
        },
      },
      style: {
        display: 'flex',
      },
    })
  }

  exportJSON(): SerializedImageNode {
    return {
      src: this.__src,
      text: this.__src,
      type: 'image',
      version: 1,
    }
  }

  updateDOM(): boolean {
    return false
  }

  exportDOM(): DOMExportOutput {
    const element = document.createElement('img')
    element.setAttribute('src', this.__src)
    return { element }
  }

  getTextContent(): string {
    return this.__src
  }
}

export function $isImageNode(
  node: LexicalNode | null | undefined,
): node is ImageNode {
  return node instanceof ImageNode
}

export function $createImageNode(src: string): ImageNode {
  return new ImageNode(src)
}
