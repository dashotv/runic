import React from 'react';
import { Helmet } from 'react-helmet-async';

import Container from '@mui/material/Container';

import { RoutingTabs, RoutingTabsRoute } from 'components/common';
import Releases from 'pages/releases';

function App() {
  // limit, skip, queries, etc

  const tabsMap: RoutingTabsRoute[] = [
    {
      label: 'Releases',
      to: `releases`,
      element: <Releases />,
    },
    {
      label: 'Indexers',
      to: `indexers`,
      element: <div>Indexers</div>,
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
}

export default App;
