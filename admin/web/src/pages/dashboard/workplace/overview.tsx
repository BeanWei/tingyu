import type { ReactNode } from 'react'
import React from 'react'
import {
  Card,
  Divider,
  Grid,
  Link,
  Skeleton,
  Typography,
} from '@arco-design/web-react'
import { IconCaretUp } from '@arco-design/web-react/icon'
import styles from './styles/overview.module.less'
import IconCalendar from './assets/calendar.svg'
import IconComments from './assets/comments.svg'
import IconContent from './assets/content.svg'
import IconIncrease from './assets/increase.svg'
import locales from './locales'
import { useUserStore } from '~/store'
import { useLocale } from '~/hooks'

const { Row, Col } = Grid

interface StatisticItemType {
  icon?: ReactNode
  title?: ReactNode
  count?: ReactNode
  loading?: boolean
  unit?: ReactNode
}

function StatisticItem(props: StatisticItemType) {
  const { icon, title, count, loading, unit } = props
  return (
    <div className={styles.item}>
      <div className={styles.icon}>{icon}</div>
      <div>
        <Skeleton loading={loading} text={{ rows: 2, width: 60 }} animation>
          <div className={styles.title}>{title}</div>
          <div className={styles.count}>
            {count}
            <span className={styles.unit}>{unit}</span>
          </div>
        </Skeleton>
      </div>
    </div>
  )
}

function Overview() {
  const t = useLocale(locales)
  const userStore = useUserStore()

  return (
    <Card>
      <Typography.Title heading={5}>
        {t['workplace.welcomeBack']}
        {userStore.info?.nickname}
      </Typography.Title>
      <Divider />
      <Row>
        <Col flex={1}>
          <StatisticItem
            icon={<IconCalendar />}
            title={t['workplace.totalOnlyData']}
            count={666}
            unit={t['workplace.pecs']}
          />
        </Col>
        <Divider type="vertical" className={styles.divider} />
        <Col flex={1}>
          <StatisticItem
            icon={<IconContent />}
            title={t['workplace.contentInMarket']}
            count={666}
            unit={t['workplace.pecs']}
          />
        </Col>
        <Divider type="vertical" className={styles.divider} />
        <Col flex={1}>
          <StatisticItem
            icon={<IconComments />}
            title={t['workplace.comments']}
            count={666}
            unit={t['workplace.pecs']}
          />
        </Col>
        <Divider type="vertical" className={styles.divider} />
        <Col flex={1}>
          <StatisticItem
            icon={<IconIncrease />}
            title={t['workplace.growth']}
            count={
              <span>
                {66}{' '}
                <IconCaretUp
                  style={{ fontSize: 18, color: 'rgb(var(--green-6))' }}
                />
              </span>
            }
          />
        </Col>
      </Row>
      <Divider />
      <div>
        <div className={styles.ctw}>
          <Typography.Paragraph
            className={styles['chart-title']}
            style={{ marginBottom: 0 }}
          >
            {t['workplace.contentData']}
            <span className={styles['chart-sub-title']}>
              ({t['workplace.1year']})
            </span>
          </Typography.Paragraph>
          <Link>{t['workplace.seeMore']}</Link>
        </div>
      </div>
    </Card>
  )
}

export default Overview
