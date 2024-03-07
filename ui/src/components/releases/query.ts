import axios from 'axios';

import { useQuery } from '@tanstack/react-query';

import { ReleaseResponse, ReleasesResponse } from './types';

export const getReleases = async (limit: number, skip: number, queryString: string) => {
  const response = await axios.get(`/api/releases/?limit=${limit}&skip=${skip}&${queryString}`);
  return response.data as ReleasesResponse;
};

export const getRelease = async (id: string) => {
  const response = await axios.get(`/api/releases/${id}`);
  return response.data as ReleaseResponse;
};

export const useReleasesQuery = (limit: number, skip: number, queryString: string) =>
  useQuery({
    queryKey: ['releases', limit, skip, queryString],
    queryFn: () => getReleases(limit, skip, queryString),
    placeholderData: previousData => previousData,
    retry: false,
  });

export const useReleaseQuery = (id: string) =>
  useQuery({
    queryKey: ['releases', id],
    queryFn: () => getRelease(id),
  });
