import { Link as RouterLink } from 'react-router-dom';

import { Popular } from 'client';

import { Link } from '@mui/material';

import { useQueryString } from 'components/hooks';

import './popular.scss';

export const PopularList = ({ data, type }: { data: Popular[]; type: string }) => {
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
