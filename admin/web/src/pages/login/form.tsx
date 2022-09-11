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
import styles from './style/index.module.less'
import { useLocale } from '~/hooks'
import { request, url } from '~/api'
import { useUserStore } from '~/store'

export default function LoginForm() {
  const t = useLocale(locales)
  const formRef = useRef<FormInstance>()

  const { loading, run } = useRequest((data: AnyObject) => {
    return request({
      url: url.userLogin,
      data,
    })
  }, {
    manual: true,
    onSuccess: (result) => {
      useUserStore().updateToken(result.data.token)
      window.location.href = '/'
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
        initialValues={{ userName: 'admin', password: 'admin' }}
      >
        <Form.Item
          field="userName"
          rules={[{ required: true, message: t['login.form.userName.errMsg'] }]}
        >
          <Input
            prefix={<IconUser />}
            placeholder={t['login.form.userName.placeholder']}
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
            {/* <Checkbox checked={rememberPassword} onChange={setRememberPassword}>
              {t['login.form.rememberPassword']}
            </Checkbox> */}
            <Link>{t['login.form.forgetPassword']}</Link>
          </div>
          <Button type="primary" long onClick={handleLogin} loading={loading}>
            {t['login.form.login']}
          </Button>
          <Button
            type="text"
            long
            className={styles['login-form-register-btn']}
          >
            {t['login.form.register']}
          </Button>
        </Space>
      </Form>
    </div>
  )
}
