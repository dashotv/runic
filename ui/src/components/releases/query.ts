import axios from 'axios';

import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';

import { ReleaseResponse, ReleasesResponse, SearchResponse } from './types';

export interface Setting {
  setting: string;
  value: boolean;
}
export interface SettingsArgs {
  id: string;
  setting: Setting;
}

export const getReleases = async (limit: number, skip: number, queryString: string) => {
  const response = await axios.get(`/api/runic/releases/?limit=${limit}&skip=${skip}&${queryString}`);
  return response.data as ReleasesResponse;
};

export const getRelease = async (id: string) => {
  const response = await axios.get(`/api/runic/releases/${id}`);
  return response.data as ReleaseResponse;
};

export const patchRelease = async (id: string, setting: Setting) => {
  const response = await axios.patch(`/api/runic/releases/${id}`, setting);
  return response.data;
};

export const getSearch = async (limit: number, skip: number, queryString: string) => {
  const response = await axios.get(`/api/scry/runic/?limit=${limit}&skip=${skip}&${queryString}`);
  return response.data as SearchResponse;
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

export const useReleaseSettingMutation = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (args: SettingsArgs) => patchRelease(args.id, args.setting),
    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: ['releases'] });
    },
  });
};

export const useSearchQuery = (limit: number, skip: number, queryString: string) =>
  useQuery({
    queryKey: ['search', limit, skip, queryString],
    queryFn: () => getSearch(limit, skip, queryString),
    placeholderData: previousData => previousData,
    retry: false,
  });
