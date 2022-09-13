import React, { useEffect } from 'react'
import LoginForm from './form'
import styles from './styles/index.module.less'
import Footer from '~/components/Footer'

function Login() {
  useEffect(() => {
    document.body.setAttribute('arco-theme', 'light')
  }, [])

  return (
    <div className={styles.container}>
      <div className={styles.content}>
        <div className={styles['content-inner']}>
          <LoginForm />
        </div>
        <div className={styles.footer}>
          <Footer />
        </div>
      </div>
    </div>
  )
}
Login.displayName = 'LoginPage'

export default Login
