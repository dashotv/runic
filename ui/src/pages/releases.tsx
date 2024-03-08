import { ReleaseList, useReleasesQuery } from 'components/releases';

const Releases = () => {
  const { data } = useReleasesQuery(25, 0, '');
  return <>{data && <ReleaseList data={data?.results} />}</>;
};

export default Releases;
