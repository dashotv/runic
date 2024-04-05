import Container from '@mui/material/Container';
import CssBaseline from '@mui/material/CssBaseline';
import { ThemeProvider, createTheme } from '@mui/material/styles';

import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

import { RoutingTabs, RoutingTabsRoute } from 'components/common';
import { IndexersList } from 'components/indexers';
import Releases from 'pages/releases';
import Search from 'pages/search';

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

const App = () => {
  const tabsMap: RoutingTabsRoute[] = [
    {
      label: 'Search',
      to: '',
      element: <Search />,
    },
    {
      label: 'Releases',
      to: 'releases',
      element: <Releases />,
    },
    {
      label: 'Indexers',
      to: 'indexers',
      element: <IndexersList />,
    },
  ];
  return (
    <ThemeProvider theme={darkTheme}>
      <QueryClientProvider client={queryClient}>
        <CssBaseline />
        <Container>
          <RoutingTabs data={tabsMap} route={'/'} />
        </Container>
      </QueryClientProvider>
    </ThemeProvider>
  );
};

export default App;
