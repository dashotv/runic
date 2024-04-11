import { Link as RouterLink } from 'react-router-dom';

import { Popular } from 'client';

import { Box, Link, Stack, Typography } from '@mui/material';

import { useQueryString } from 'components/hooks';

import './popular.scss';

export const PopularList = ({ data, type }: { data: Popular[]; type: string }) => {
  const url =
    type !== 'anime' ? `http://themoviedb.org/search?query=` : 'https://myanimelist.net/anime.php?cat=anime&q=';

  return (
    <Box>
      <Stack direction="column" spacing={0.5}>
        <Typography variant="h4" borderBottom="1px solid white">
          {type}
        </Typography>
        {data?.map(({ title, count, year = 0 }, index) => (
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
              <SearchLink {...{ title, type, count }} />
            </Typography>
          </Stack>
        ))}
      </Stack>
    </Box>
  );
};

const SearchLink = ({ title, type, count }: { title?: string; type?: string; count?: number }) => {
  const { queryString } = useQueryString();
  const link =
    '/releases/search?' +
    queryString({
      text: title,
      type: type,
      resolution: type === 'movies' ? '1080' : '',
      exact: type === 'movies' ? false : true,
      verified: true,
    });
  return <RouterLink to={title ? link : '#'}>{count}</RouterLink>;
};

export const PopularListOld = ({ data, type }: { data: Popular[]; type: string }) => {
  const url =
    type !== 'anime' ? `http://themoviedb.org/search?query=` : 'https://myanimelist.net/anime.php?cat=anime&q=';

  return (
    <div className="popular">
      <div className="header">{type}</div>
      {data?.map(({ title, count, year = 0 }, index) => (
        <div key={index} className="entry">
          <span className="title">
            <Link href={`${url}${title}${year > 0 ? `+y:${year}` : ''}`} target="_window">
              {title}
              {year > 0 && ` (${year})`}
            </Link>
          </span>
          <span className="number">
            <RouterLink to={title ? SearchURL(title, type) : '#'}>{count}</RouterLink>
          </span>
        </div>
      ))}
    </div>
  );
};

export function SearchURL(text: string, type: string) {
  const { queryString } = useQueryString();
  const base = '/releases/search?';
  const settings = {
    text: text,
    type: type,
    resolution: type === 'movies' ? '1080' : '',
    exact: type === 'movies' ? false : true,
    verified: true,
  };

  return base + queryString(settings);
}
