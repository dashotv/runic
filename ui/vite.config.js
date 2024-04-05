var __assign = (this && this.__assign) || function () {
    __assign = Object.assign || function(t) {
        for (var s, i = 1, n = arguments.length; i < n; i++) {
            s = arguments[i];
            for (var p in s) if (Object.prototype.hasOwnProperty.call(s, p))
                t[p] = s[p];
        }
        return t;
    };
    return __assign.apply(this, arguments);
};
import { defineConfig } from 'vite';
import viteTsconfigPaths from 'vite-tsconfig-paths';
import federation from '@originjs/vite-plugin-federation';
import react from '@vitejs/plugin-react';
// @ts-ignore
import pkg from './package.json';
var dependencies = pkg.dependencies;
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
            },
            shared: __assign(__assign({}, dependencies), { react: {
                    requiredVersion: dependencies['react'],
                }, 'react-dom': {
                    requiredVersion: dependencies['react-dom'],
                } }),
        }),
    ],
    build: {
        target: 'esnext', //browsers can handle the latest ES features
        outDir: '../static',
    },
    server: {
        proxy: {
            '/api/runic': {
                target: 'http://localhost:59002',
                changeOrigin: true,
                secure: false,
                ws: true,
                rewrite: function (path) { return path.replace(/^\/api\/runic/, ''); },
            },
            '/api/scry': {
                target: 'http://localhost:59003',
                changeOrigin: true,
                secure: false,
                ws: true,
                rewrite: function (path) { return path.replace(/^\/api\/scry/, ''); },
            },
        },
    },
});
