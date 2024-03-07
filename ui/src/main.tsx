import React from 'react';
import ReactDOM from 'react-dom/client';
import { HelmetProvider } from 'react-helmet-async';
import { BrowserRouter as Router } from 'react-router-dom';

import { SnackbarProvider } from 'notistack';

import CssBaseline from '@mui/material/CssBaseline';
import { ThemeProvider, createTheme } from '@mui/material/styles';

import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { ReactQueryDevtools } from '@tanstack/react-query-devtools';

import App from './App.tsx';
import './index.css';

const darkTheme = createTheme({
  palette: {
    mode: 'dark',
  },
});

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry: 5,
      staleTime: 5 * 1000,
      throwOnError: true,
    },
  },
});

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <ThemeProvider theme={darkTheme}>
      <SnackbarProvider
        maxSnack={3}
        anchorOrigin={{ horizontal: 'center', vertical: 'top' }}
        disableWindowBlurListener={true}
        TransitionProps={{ direction: 'down' }}
        dense={true}
      >
        <Router>
          <CssBaseline />
          <HelmetProvider>
            <QueryClientProvider client={queryClient}>
              <App />
              <ReactQueryDevtools initialIsOpen={false} />
            </QueryClientProvider>
          </HelmetProvider>
        </Router>
      </SnackbarProvider>
    </ThemeProvider>
  </React.StrictMode>,
);
