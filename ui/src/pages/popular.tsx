import { Route, Routes, useParams } from 'react-router-dom';
import { Link } from 'react-router-dom';

import { Button, Grid } from '@mui/material';

import { Container, LoadingIndicator } from '@dashotv/components';

import { usePopularQuery } from 'components/popular';
import { PopularList } from 'components/popular/list';

const Popular = ({ mount }: { mount: string }) => {
  return (
    <Routes>
      <Route path="" element={<PopularMap mount={mount} />} />
      <Route path=":interval" element={<PopularMap mount={mount} />} />
    </Routes>
  );
};

const PopularMap = ({ mount }: { mount: string }) => {
  const { interval } = useParams();
  const { isFetching, data } = usePopularQuery(interval || 'daily');
  return (
    <>
      <Container>
        <Grid container spacing={2}>
          <Grid item xs={12} md={6}>
            <Button
              component={Link}
              variant={!interval || interval === 'daily' ? 'contained' : 'text'}
              to={`${mount}/popular/daily`}
            >
              Daily
            </Button>
            <Button
              component={Link}
              variant={interval === 'weekly' ? 'contained' : 'text'}
              to={`${mount}/popular/weekly`}
            >
              Weekly
            </Button>
            <Button
              component={Link}
              variant={interval === 'monthly' ? 'contained' : 'text'}
              to={`${mount}/popular/monthly`}
            >
              Monthly
            </Button>
          </Grid>
          <Grid item xs={12} md={6}></Grid>
        </Grid>
      </Container>
      <Container>
        {isFetching ? <LoadingIndicator /> : null}
        <Grid container spacing={2}>
          <Grid item xs={12} md={4}>
            {data?.result.tv ? <PopularList mount={mount} type="tv" data={data?.result.tv} /> : null}
          </Grid>
          <Grid item xs={12} md={4}>
            {data?.result.anime ? <PopularList mount={mount} type="anime" data={data?.result.anime} /> : null}
          </Grid>
          <Grid item xs={12} md={4}>
            {data?.result.movies ? <PopularList mount={mount} type="movies" data={data?.result.movies} /> : null}
          </Grid>
        </Grid>
      </Container>
    </>
  );
};
export default Popular;
