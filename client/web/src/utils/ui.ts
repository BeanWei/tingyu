import type { MenuGroupOption, MenuOption } from 'naive-ui'
import type { VNodeChild } from 'vue'
import { RouterLink } from 'vue-router'

export const menuLabelRender = (option: MenuOption | MenuGroupOption): VNodeChild => {
  if (option.path) {
    return h(
      RouterLink,
      {
        to: option.path,
      },
      { default: () => option.label },
    )
  }
  if (typeof option.label === 'function')
    return option.label?.()
  return option.label as any
}
