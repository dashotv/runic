import * as runic from 'client/runic';
import * as scry from 'client/scry';

import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';

import { SearchForm } from '.';

export interface Setting {
  setting: string;
  value: boolean;
}
export interface SettingsArgs {
  id: string;
  setting: Setting;
}

export const getReleases = async (page: number, limit: number) => {
  const response = await runic.ReleasesIndex({ page, limit });
  return response;
};

export const getReleasesSearch = async (
  page: number,
  limit: number,
  source: string | null,
  kind: string | null,
  resolution: string | null,
  group: string | null,
  website: string | null,
) => {
  const response = await runic.ReleasesSearch({
    page,
    limit,
    source: source || '',
    kind: kind || '',
    resolution: resolution || '',
    group: group || '',
    website: website || '',
  });
  return response;
};

export const getRelease = async (id: string) => {
  const response = await runic.ReleasesShow({ id });
  return response;
};

export const patchRelease = async (id: string, setting: runic.Setting) => {
  const response = await runic.ReleasesSettings({ id, setting });
  return response;
};

export const getSearch = async (limit: number, start: number, search: SearchForm) => {
  const req = {
    start,
    limit,
    text: search.text ? search.text : '',
    year: search.year ? Number(search.year) : -1,
    season: search.season ? Number(search.season) : -1,
    episode: search.episode ? Number(search.episode) : -1,
    group: search.group ? search.group : '',
    website: search.website ? search.website : '',
    type: search.type ? search.type : '',
    source: search.source ? search.source : '',
    resolution: search.resolution ? Number(search.resolution) : -1,
    uncensored: search.uncensored || false,
    bluray: search.bluray || false,
    verified: search.verified || false,
    exact: search.exact || false,
  };

  const response = await scry.RunicIndex(req);
  return response.result;
};
// export const getSearch = async (limit: number, skip: number, queryString: string) => {
//   const response = await axios.get(`/api/scry/runic/?limit=${limit}&skip=${skip}&${queryString}`);
//   return response.data as SearchResponse;
// };

export const useReleasesQuery = (page: number, limit: number) =>
  useQuery({
    queryKey: ['releases', page, limit],
    queryFn: () => getReleases(page, limit),
    placeholderData: previousData => previousData,
    retry: false,
  });

export const useReleasesSearchQuery = (
  page: number,
  limit: number,
  source: string | null,
  kind: string | null,
  resolution: string | null,
  group: string | null,
  website: string | null,
) =>
  useQuery({
    queryKey: ['releases', page, limit, source, kind, resolution, group, website],
    queryFn: () => getReleasesSearch(page, limit, source, kind, resolution, group, website),
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
    mutationFn: (args: { id: string; setting: runic.Setting }) => patchRelease(args.id, args.setting),
    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: ['releases'] });
    },
  });
};

export const useSearchQuery = (limit: number, skip: number, form: SearchForm) =>
  useQuery({
    queryKey: ['search', limit, skip, form],
    queryFn: () => getSearch(limit, skip, form),
    placeholderData: previousData => previousData,
    retry: false,
  });
