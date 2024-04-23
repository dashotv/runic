import { defineConfig } from 'vite';
import viteTsconfigPaths from 'vite-tsconfig-paths';

import federation from '@originjs/vite-plugin-federation';
import react from '@vitejs/plugin-react';

import pkg from './package.json';

const { dependencies } = pkg;

// TODO: replace with -swc
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    react(),
    viteTsconfigPaths(),
    federation({
      name: 'runic',
      filename: 'remote.js',
      exposes: {
        './App': './src/pages/app.tsx',
        './Search': './src/pages/remoteSearch.tsx',
      },
      shared: {
        ...dependencies,
        react: {
          requiredVersion: dependencies['react'],
        },
        'react-dom': {
          requiredVersion: dependencies['react-dom'],
        },
      },
    }),
  ],
  build: {
    target: 'esnext', //browsers can handle the latest ES features
    outDir: '../static',
  },
  server: {
    port: 3002,
    proxy: {
      '/api/runic': {
        target: 'http://localhost:59002',
        changeOrigin: true,
        secure: false,
        ws: true,
        rewrite: path => path.replace(/^\/api\/runic/, ''),
      },
      '/api/scry': {
        target: 'http://localhost:59003',
        changeOrigin: true,
        secure: false,
        ws: true,
        rewrite: path => path.replace(/^\/api\/scry/, ''),
      },
    },
  },
});
