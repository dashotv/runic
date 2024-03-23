import { Helmet } from 'react-helmet-async';

import Container from '@mui/material/Container';

import { RoutingTabs, RoutingTabsRoute } from 'components/common';
import { IndexersList } from 'components/indexers';
import Releases from 'pages/releases';
import Search from 'pages/search';

const Layout = () => {
  // limit, skip, queries, etc

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
    <>
      <Helmet>
        <title>Runic</title>
        <meta name="description" content="runic" />
      </Helmet>
      <Container>
        <RoutingTabs data={tabsMap} route={'/'} />
      </Container>
    </>
  );
};

export default Layout;
