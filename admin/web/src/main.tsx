import { useEffect } from 'react'
import { ConfigProvider } from '@arco-design/web-react'
import zhCN from '@arco-design/web-react/es/locale/zh-CN'
import enUS from '@arco-design/web-react/es/locale/en-US'
import ReactDOM from 'react-dom/client'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { useSettingsStore, useUserStore } from './store'
import PageLayout from './layout'
import Login from './pages/login'
import { request, url } from './api'

function App() {
  const settingsStore = useSettingsStore()
  const userStore = useUserStore()

  function getArcoLocale() {
    switch (settingsStore.lang) {
      case 'zh-CN':
        return zhCN
      case 'en-US':
        return enUS
      default:
        return zhCN
    }
  }

  useEffect(() => {
    if (userStore.token) {
      request({
        url: url.getUserInfo,
      }).then((res) => {
        userStore.info = res.data
      })
    }
    else if (window.location.pathname.replace(/\//g, '') !== 'login') { window.location.pathname = '/login' }
  }, [])

  return (
    <BrowserRouter>
      <ConfigProvider
        locale={getArcoLocale()}
        componentConfig={{
          Card: {
            bordered: false,
          },
          List: {
            bordered: false,
          },
          Table: {
            border: false,
          },
        }}
      >
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/" element={<PageLayout />} />
        </Routes>
      </ConfigProvider>
    </BrowserRouter>
  )
}

ReactDOM.createRoot(document.getElementById('root')).render(<App />)
