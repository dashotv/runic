import { PopularIndex } from 'client/runic';

import { useQuery } from '@tanstack/react-query';

export const usePopularQuery = (interval: string) =>
  useQuery({
    queryKey: ['popular', interval],
    queryFn: () => PopularIndex({ interval }),
    placeholderData: previousData => previousData,
    retry: false,
  });
