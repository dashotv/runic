import React from 'react';
import * as ReactDOM from 'react-dom/client';
import { HelmetProvider } from 'react-helmet-async';
import { BrowserRouter as Router } from 'react-router-dom';

// import { SnackbarProvider } from 'notistack';
import CssBaseline from '@mui/material/CssBaseline';

import App from './pages/app.tsx';

// import './index.css';
ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <Router>
      <CssBaseline />
      <HelmetProvider>
        <App />
      </HelmetProvider>
    </Router>
  </React.StrictMode>,
);
