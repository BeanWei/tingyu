import {
  Button,
  Form,
  Input,
  Link,
  Space,
} from '@arco-design/web-react'
import type { FormInstance } from '@arco-design/web-react/es/Form'
import { IconLock, IconUser } from '@arco-design/web-react/icon'
import { useRef } from 'react'
import { useRequest } from 'ahooks'
import locales from './locales'
import styles from './styles/index.module.less'
import { useLocale } from '~/hooks'
import { request, url } from '~/api'
import { useUserStore } from '~/store'

export default function LoginForm() {
  const t = useLocale(locales)
  const formRef = useRef<FormInstance>()
  const userStore = useUserStore()

  const { loading, run } = useRequest((data: AnyObject) => {
    return request({
      url: url.userLogin,
      data,
    })
  }, {
    manual: true,
    onSuccess: (result) => {
      userStore.updateToken(result.data.data.token)
      window.location.href = '/admin'
    },
  })

  function handleLogin() {
    formRef.current.validate().then((values) => {
      run(values)
    })
  }

  return (
    <div className={styles['login-form-wrapper']}>
      <div className={styles['login-form-title']}>{t['login.form.title']}</div>
      <div className={styles['login-form-error-msg']} />
      <Form
        className={styles['login-form']}
        layout="vertical"
        ref={formRef}
      >
        <Form.Item
          field="username"
          rules={[{ required: true, message: t['login.form.username.errMsg'] }]}
        >
          <Input
            prefix={<IconUser />}
            placeholder={t['login.form.username.placeholder']}
            onPressEnter={handleLogin}
          />
        </Form.Item>
        <Form.Item
          field="password"
          rules={[{ required: true, message: t['login.form.password.errMsg'] }]}
        >
          <Input.Password
            prefix={<IconLock />}
            placeholder={t['login.form.password.placeholder']}
            onPressEnter={handleLogin}
          />
        </Form.Item>
        <Space size={16} direction="vertical">
          <div className={styles['login-form-password-actions']}>
            <Link>{t['login.form.forgetPassword']}</Link>
          </div>
          <Button type="primary" long onClick={handleLogin} loading={loading}>
            {t['login.form.login']}
          </Button>
        </Space>
      </Form>
    </div>
  )
}
