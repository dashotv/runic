import React from 'react';

import { ReleaseList, useReleasesQuery } from 'components/releases';

import './App.css';

function App() {
  const { data } = useReleasesQuery(100, 0, '');

  return <>{data && <ReleaseList data={data?.results} />}</>;
}

export default App;
