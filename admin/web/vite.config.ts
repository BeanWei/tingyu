import path from 'path'
import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import svgrPlugin from '@arco-plugins/vite-plugin-svgr'
import vitePluginForArco from '@arco-plugins/vite-react'
import setting from './src/settings.json'

// https://vitejs.dev/config/
export default defineConfig((configEnv) => {
  return {
    resolve: {
      alias: {
        '~/': `${path.resolve(__dirname, 'src')}/`,
      },
    },
    plugins: [
      react(),
      svgrPlugin({
        svgrOptions: {},
      }),
      vitePluginForArco({
        theme: '@arco-themes/react-arco-pro',
        modifyVars: {
          'arcoblue-6': setting.themeColor,
        },
      }),
    ],
    css: {
      preprocessorOptions: {
        less: {
          javascriptEnabled: true,
        },
      },
    },
    server: {
      proxy: configEnv.mode === 'development'
        ? {
            '/api/': {
              target: 'http://localhost:8888/api/',
              changeOrigin: true,
              rewrite: path => path.replace(/^\/api/, ''),
            },
          }
        : undefined,
    },
  }
})
