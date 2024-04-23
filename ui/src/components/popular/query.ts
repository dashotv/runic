import { PopularIndex, PopularMovies } from 'client/runic';

import { useQuery } from '@tanstack/react-query';

export const usePopularQuery = (interval: string) =>
  useQuery({
    queryKey: ['popular', interval],
    queryFn: () => PopularIndex({ interval }),
    placeholderData: previousData => previousData,
    retry: false,
  });
export const usePopularMoviesQuery = () =>
  useQuery({
    queryKey: ['popular', 'movies'],
    queryFn: () => PopularMovies(),
    placeholderData: previousData => previousData,
    retry: false,
  });
