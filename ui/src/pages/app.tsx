import CssBaseline from '@mui/material/CssBaseline';
import { ThemeProvider, createTheme } from '@mui/material/styles';

import { Container, RoutingTabs, RoutingTabsRoute } from '@dashotv/components';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

import { IndexersList } from 'components/indexers';
import Popular from 'pages/popular';
import Releases from 'pages/releases';
import Search from 'pages/search';

const darkTheme = createTheme({
  palette: {
    mode: 'dark',
  },
  components: {
    MuiLink: {
      styleOverrides: {
        root: {
          textDecoration: 'none',
        },
      },
    },
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

const App = ({ mount }: { mount: string }) => {
  const tabsMap: RoutingTabsRoute[] = [
    {
      label: 'Search',
      to: '',
      element: <Search />,
    },
    {
      label: 'Popular',
      to: 'popular',
      element: <Popular />,
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
          <RoutingTabs data={tabsMap} mount={mount} />
        </Container>
      </QueryClientProvider>
    </ThemeProvider>
  );
};

export default App;
