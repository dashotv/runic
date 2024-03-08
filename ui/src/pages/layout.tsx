import { Helmet } from 'react-helmet-async';

import Container from '@mui/material/Container';

import { RoutingTabs, RoutingTabsRoute } from 'components/common';
import { IndexersList } from 'components/indexers';
import Releases from 'pages/releases';

const Layout = () => {
  // limit, skip, queries, etc

  const tabsMap: RoutingTabsRoute[] = [
    {
      label: 'Releases',
      to: '',
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
