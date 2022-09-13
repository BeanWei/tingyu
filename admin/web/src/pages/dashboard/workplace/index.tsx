import { Space } from '@arco-design/web-react'
import Overview from './overview'
import styles from './styles/index.module.less'

function Workplace() {
  return (
    <div className={styles.wrapper}>
      <Space size={16} direction="vertical">
        <Overview />
      </Space>
    </div>
  )
}

export default Workplace
