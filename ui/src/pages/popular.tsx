import { useParams } from 'react-router-dom';

import { Grid } from '@mui/material';

import { Container, LoadingIndicator } from '@dashotv/components';

import { usePopularQuery } from 'components/popular';
import { PopularList } from 'components/popular/list';

const Popular = () => {
  const { interval } = useParams();
  const { isFetching, data } = usePopularQuery(interval || 'daily');

  return (
    <Container>
      {isFetching ? <LoadingIndicator /> : null}
      <Grid container spacing={2}>
        <Grid item xs={4}>
          {data?.result.tv ? <PopularList type="tv" data={data?.result.tv} /> : null}
        </Grid>
        <Grid item xs={4}>
          {data?.result.anime ? <PopularList type="anime" data={data?.result.anime} /> : null}
        </Grid>
        <Grid item xs={4}>
          {data?.result.movies ? <PopularList type="movies" data={data?.result.movies} /> : null}
        </Grid>
      </Grid>
    </Container>
  );
};
export default Popular;
