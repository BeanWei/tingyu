import React from 'react'
import { Layout } from '@arco-design/web-react'
import type { FooterProps } from '@arco-design/web-react/es/Layout/interface'
import cs from 'classnames'
import styles from './style/index.module.less'

function Footer(props: FooterProps = {}) {
  const { className, ...restProps } = props
  return (
    <Layout.Footer className={cs(styles.footer, className)} {...restProps}>
      🌧听雨 | Admin
    </Layout.Footer>
  )
}

export default Footer
