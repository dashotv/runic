import { defineConfig } from 'vite';
import viteTsconfigPaths from 'vite-tsconfig-paths';

import federation from '@originjs/vite-plugin-federation';
import react from '@vitejs/plugin-react';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    react(),
    viteTsconfigPaths(),
    federation({
      name: 'runic',
      filename: 'remote.js',
      exposes: {
        './Releases': './src/pages/releases.tsx',
      },
      shared: ['react', 'react-dom', 'react-router-dom'],
    }),
  ],
  build: {
    outDir: '../static',
  },
  server: {
    proxy: {
      '/api/runic': {
        target: 'http://localhost:9002',
        changeOrigin: true,
        secure: false,
        ws: true,
        rewrite: path => path.replace(/^\/api\/runic/, ''),
      },
    },
  },
});
