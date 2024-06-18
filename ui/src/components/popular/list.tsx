import { Link as RouterLink } from 'react-router-dom';

import { Popular, PopularMovie } from 'client/runic';

import { Box, Link, Stack, Typography } from '@mui/material';

import { useQueryString } from 'components/hooks';

import './popular.scss';

export const PopularList = ({ mount, data, type }: { mount: string; data: Popular[]; type: string }) => {
  const url =
    type !== 'anime' ? `http://themoviedb.org/search?query=` : 'https://myanimelist.net/anime.php?cat=anime&q=';

  return (
    <Box>
      <Stack direction="column" spacing={0.5}>
        <Typography variant="h4" borderBottom="1px solid white">
          {type}
        </Typography>
        {data?.map(({ title, count, verified, year = 0 }, index) => (
          <Stack key={index} direction="row" spacing={0} alignItems="center" justifyContent="space-between">
            <Typography variant="body1" color="primary" noWrap>
              <Link href={`${url}${title}${year > 0 ? `+y:${year}` : ''}`} target="_window">
                {title}
                {year > 0 && ` (${year})`}
              </Link>
            </Typography>
            <Typography
              variant="body1"
              color="action"
              minWidth="50px"
              justifyContent="end"
              textAlign="right"
              sx={{ '& a': { color: '#f0f0f0', textDecoration: 'none' } }}
            >
              <SearchLink {...{ mount, title, type, count, verified }} />
            </Typography>
          </Stack>
        ))}
      </Stack>
    </Box>
  );
};

export const PopularMoviesList = ({ mount, data }: { mount: string; data: PopularMovie[] }) => {
  const popular: Popular[] = data
    .map(({ id, count, verified }) => {
      if (!id) return;
      return { title: id.title, year: id.year, type: 'movies', count, verified };
    })
    .filter(x => x !== undefined) as Popular[];
  return <PopularList mount={mount} data={popular} type="movies" />;
};

const SearchLink = ({
  mount,
  title,
  type,
  count,
  verified,
}: {
  mount: string;
  title?: string;
  type?: string;
  count?: number;
  verified?: number;
}) => {
  const { queryString } = useQueryString();
  const link =
    `${mount}/?` +
    queryString({
      text: title,
      type: type,
      resolution: type === 'movies' ? '1080' : '',
      exact: type === 'movies' ? false : true,
    });
  return (
    <RouterLink to={title ? link : '#'}>
      {verified !== undefined && verified > 0 ? `${verified} (${count})` : count}
    </RouterLink>
  );
};
