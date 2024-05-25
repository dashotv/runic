import CssBaseline from '@mui/material/CssBaseline';
import { ThemeProvider, createTheme } from '@mui/material/styles';

import { Container, RoutingTabs, RoutingTabsRoute } from '@dashotv/components';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

import { IndexersList } from 'components/indexers';
import { SearchForm } from 'components/releases';
import Popular from 'pages/popular';
import Releases from 'pages/releases';
import Search from 'pages/search';

import RemoteSearch from './remoteSearch';

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

const formDefaults: SearchForm = {
  text: '',
  year: '',
  season: '',
  episode: '',
  group: '',
  website: '',
  resolution: '',
  source: '',
  type: '',
  uncensored: false,
  bluray: false,
  verified: false,
  exact: false,
};
const selector = (url: string) => {
  console.log(url);
};
const selected = '';

const App = ({ mount }: { mount: string }) => {
  const tabsMap: RoutingTabsRoute[] = [
    {
      label: 'Popular',
      to: 'popular',
      path: 'popular/*',
      element: <Popular mount={mount} />,
    },
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

  if (!import.meta.env.PROD) {
    tabsMap.push({
      label: 'Embedded',
      to: 'embedded',
      element: <RemoteSearch {...{ selector, rawForm: formDefaults, selected }} />,
    });
  }
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
