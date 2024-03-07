import { defineConfig } from 'vite';
import viteTsconfigPaths from 'vite-tsconfig-paths';

import react from '@vitejs/plugin-react';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react(), viteTsconfigPaths()],
  build: {
    outDir: '../static',
  },
  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:9002',
        changeOrigin: true,
        secure: false,
        ws: true,
        rewrite: path => path.replace(/^\/api/, ''),
      },
    },
  },
});
