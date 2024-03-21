import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path';
import compression from 'vite-plugin-compression';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue(), compression({
    threshold: 10240, // 文件大小大于 10kb时启用压缩
  }),],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
    }
  },
  server: {
    proxy: {
      "/api": {
        target: "http://127.0.0.1:10001",
        changeOrigin: true,
        rewrite: path => path.replace(/^\/api/, ''),
      },
    },
  },

})
