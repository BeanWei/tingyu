import React from 'react'
import { Button, Result } from '@arco-design/web-react'
import locales from './locales'
import styles from './style/index.module.less'
import { useLocale } from '~/hooks'

function Exception500() {
  const t = useLocale(locales)

  return (
    <div className={styles.wrapper}>
      <Result
        className={styles.result}
        status="500"
        subTitle={t['exception.result.500.description']}
        extra={
          <Button key="back" type="primary">
            {t['exception.result.500.back']}
          </Button>
        }
      />
    </div>
  )
}

export default Exception500
