import React, { Suspense, useEffect, useMemo, useRef, useState } from 'react'
import { Navigate, Route, Routes, useNavigate } from 'react-router-dom'
import { Breadcrumb, Layout, Menu, Spin } from '@arco-design/web-react'
import cs from 'classnames'
import {
  IconApps,
  IconCheckCircle,
  IconDashboard,
  IconExclamationCircle,
  IconFile,
  IconList,
  IconMenuFold,
  IconMenuUnfold,
  IconSettings,
  IconStorage,
  IconUser,
} from '@arco-design/web-react/icon'
import NProgress from 'nprogress'
import Navbar from '~/components/NavBar'
import Footer from '~/components/Footer'
import { isArray } from '~/utils/is'
import styles from '~/styles/layout.module.less'
import { useUserStore } from '~/store'
import type { IRoute } from '~/routes'
import routes from '~/routes'

const MenuItem = Menu.Item
const SubMenu = Menu.SubMenu

const Sider = Layout.Sider
const Content = Layout.Content

NProgress.configure({ showSpinner: false })

function getIconFromKey(key) {
  switch (key) {
    case 'dashboard':
      return <IconDashboard className={styles.icon} />
    case 'system':
      return <IconStorage className={styles.icon} />
    case 'list':
      return <IconList className={styles.icon} />
    case 'form':
      return <IconSettings className={styles.icon} />
    case 'profile':
      return <IconFile className={styles.icon} />
    case 'visualization':
      return <IconApps className={styles.icon} />
    case 'result':
      return <IconCheckCircle className={styles.icon} />
    case 'exception':
      return <IconExclamationCircle className={styles.icon} />
    case 'user':
      return <IconUser className={styles.icon} />
    default:
      return <div className={styles['icon-empty']} />
  }
}

function getFlattenRoutes(routes) {
  const mod = import.meta.glob('~/pages/**/[a-z[]*.tsx')
  const res = []
  function travel(_routes) {
    _routes.forEach((route) => {
      const visibleChildren = (route.children || []).filter(
        child => !child.ignore,
      )
      if (route.key && (!route.children || !visibleChildren.length)) {
        try {
          const Component = React.lazy(mod[`/src/pages/${route.key}/index.tsx`] as any)
          route.element = (
            <Suspense fallback={<Spin className={styles.spin} />}>
              <Component />
            </Suspense>
          )
          res.push(route)
        }
        catch (e) {
          console.error(e)
        }
      }
      if (isArray(route.children) && route.children.length)
        travel(route.children)
    })
  }
  travel(routes)
  return res
}

function PageLayout() {
  const navigate = useNavigate()
  const pathname = window.location.pathname.replace(/^\/admin/, '')
  const userStore = useUserStore()

  const defaultSelectedKeys = [routes[0].children ? routes[0].children[0].key : routes[0].key]
  const defaultOpenKeys = [routes[0].key]

  const [breadcrumb, setBreadCrumb] = useState([])
  const [collapsed, setCollapsed] = useState<boolean>(false)
  const [selectedKeys, setSelectedKeys]
    = useState<string[]>(defaultSelectedKeys)
  const [openKeys, setOpenKeys] = useState<string[]>(defaultOpenKeys)

  const routeMap = useRef<Map<string, React.ReactNode[]>>(new Map())
  const menuMap = useRef<
    Map<string, { menuItem?: boolean; subMenu?: boolean }>
  >(new Map())

  const navbarHeight = 60
  const menuWidth = collapsed ? 48 : 220

  const flattenRoutes = useMemo(() => getFlattenRoutes(routes) || [], [routes])

  function onClickMenuItem(key) {
    const currentRoute = flattenRoutes.find(r => r.key === key)
    NProgress.start()
    navigate(currentRoute.path ? currentRoute.path : `/admin/${key}`)
    NProgress.done()
  }

  function toggleCollapse() {
    setCollapsed(collapsed => !collapsed)
  }

  const paddingLeft = { paddingLeft: menuWidth }
  const paddingTop = { paddingTop: navbarHeight }
  const paddingStyle = { ...paddingLeft, ...paddingTop }

  function renderMenus() {
    routeMap.current.clear()
    return function travel(_routes: IRoute[], level, parentNode = []) {
      return _routes.map((route) => {
        const { breadcrumb = true, ignore } = route
        const iconDom = getIconFromKey(route.key)
        const titleDom = (
          <>
            {iconDom} {route.name}
          </>
        )

        routeMap.current[`/${route.key}`] = breadcrumb ? [...parentNode, route.name] : []

        const visibleChildren = (route.children || []).filter((child) => {
          const { ignore, breadcrumb = true } = child
          if (ignore || route.ignore)
            routeMap.current[`/${child.key}`] = breadcrumb ? [...parentNode, route.name, child.name] : []
          return !ignore
        })

        if (ignore)
          return ''
        if (visibleChildren.length) {
          menuMap.current[route.key] = { subMenu: true }
          return (
            <SubMenu key={route.key} title={titleDom}>
              {travel(visibleChildren, level + 1, [...parentNode, route.name])}
            </SubMenu>
          )
        }
        menuMap.current[route.key] = { menuItem: true }
        return <MenuItem key={route.key}>{titleDom}</MenuItem>
      })
    }
  }

  function updateMenuStatus() {
    const pathKeys = pathname.split('/')
    const newSelectedKeys: string[] = []
    const newOpenKeys: string[] = []
    while (pathKeys.length > 0) {
      const currentRouteKey = pathKeys.join('/')
      const menuKey = currentRouteKey.replace(/^\//, '')
      const menuType = menuMap.current[menuKey]
      if (menuType && menuType.menuItem)
        newSelectedKeys.push(menuKey)
      if (menuType && menuType.subMenu && !openKeys.includes(menuKey))
        newOpenKeys.push(menuKey)
      pathKeys.pop()
    }
    newSelectedKeys.length > 0 && setSelectedKeys(newSelectedKeys)
    newOpenKeys.length > 0 && setOpenKeys(newOpenKeys)
  }

  const menus = useMemo(() => renderMenus()(routes, 1), [])

  useEffect(() => {
    const routeConfig = routeMap.current[pathname]
    setBreadCrumb(routeConfig || [])
    updateMenuStatus()
  }, [pathname])

  return (
    <Layout className={styles.layout}>
      <div
        className={cs(styles['layout-navbar'], {
          [styles['layout-navbar-hidden']]: false,
        })}
      >
        <Navbar />
      </div>
      {!userStore.info?.id
        ? (
        <Spin className={styles.spin} />
          )
        : (
        <Layout>
          <Sider
            className={styles['layout-sider']}
            width={menuWidth}
            collapsed={collapsed}
            onCollapse={setCollapsed}
            trigger={null}
            collapsible
            breakpoint="xl"
            style={paddingTop}
          >
            <div className={styles['menu-wrapper']}>
              <Menu
                collapse={collapsed}
                onClickMenuItem={onClickMenuItem}
                selectedKeys={selectedKeys}
                openKeys={openKeys}
                onClickSubMenu={(_, openKeys) => {
                  setOpenKeys(openKeys)
                }}
              >
                {menus}
              </Menu>
            </div>
            <div className={styles['collapse-btn']} onClick={toggleCollapse}>
              {collapsed ? <IconMenuUnfold /> : <IconMenuFold />}
            </div>
          </Sider>
          <Layout className={styles['layout-content']} style={paddingStyle}>
            <div className={styles['layout-content-wrapper']}>
              {!!breadcrumb.length && (
                <div className={styles['layout-breadcrumb']}>
                  <Breadcrumb>
                    {breadcrumb.map((node, index) => (
                      <Breadcrumb.Item key={index}>
                        {node}
                      </Breadcrumb.Item>
                    ))}
                  </Breadcrumb>
                </div>
              )}
              <Content>
                <Suspense fallback={<Spin className={styles.spin} />}>
                  <Routes>
                    <Route path="/" element={<Navigate to={`/admin/${defaultSelectedKeys[0]}`} />} />
                    {flattenRoutes.map((route, index) => {
                      return (
                        <Route
                          key={index}
                          path={`/${route.key}`}
                          element={route.element}
                        />
                      )
                    })}
                  </Routes>
                </Suspense>
              </Content>
            </div>
            <Footer />
          </Layout>
        </Layout>
          )}
    </Layout>
  )
}

export default PageLayout
