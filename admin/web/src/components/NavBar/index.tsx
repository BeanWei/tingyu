import {
  Avatar,
  Dropdown,
  Input,
  Tooltip,
} from '@arco-design/web-react'
import {
  IconLoading,
  IconMoonFill,
  IconSunFill,
} from '@arco-design/web-react/icon'
import styles from './style/index.module.less'
import IconButton from './IconButton'
import Logo from '~/assets/logo.svg'
import { useSettingsStore, useUserStore } from '~/store/index'
import { useLocale } from '~/hooks'

function Navbar() {
  const t = useLocale()
  const settingsStore = useSettingsStore()
  const userStore = useUserStore()

  return (
    <div className={styles.navbar}>
      <div className={styles.left}>
        <div className={styles.logo}>
          <Logo />
          <div className={styles['logo-name']}>听雨</div>
        </div>
      </div>
      <ul className={styles.right}>
        <li>
          <Input.Search
            className={styles.round}
            placeholder="搜索"
          />
        </li>
        <li>
          <Tooltip
            content={
              settingsStore.theme === 'light'
                ? t['settings.navbar.theme.toDark']
                : t['settings.navbar.theme.toLight']
            }
          >
            <IconButton
              icon={settingsStore.theme !== 'dark' ? <IconMoonFill /> : <IconSunFill />}
              onClick={() => settingsStore.changeTheme()}
            />
          </Tooltip>
        </li>
        <li>
          <Dropdown position="br" disabled={!userStore.token}>
            <Avatar size={32} style={{ cursor: 'pointer' }}>
              {userStore.info?.id ? userStore.info?.avatar ? <img alt="avatar" src={userStore.info?.avatar} /> : (userStore.info?.nickname || userStore.info?.username)[0].toUpperCase() : <IconLoading />}
            </Avatar>
          </Dropdown>
        </li>
      </ul>
    </div>
  )
}

export default Navbar
