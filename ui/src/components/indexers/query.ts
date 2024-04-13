import * as runic from 'client/runic';

import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';

export const getRunicSources = async (page: number = 1, limit: number = 100) => {
  const response = await runic.SourcesIndex({ page, limit });
  return response;
};

export const getRunicSource = async (id: string) => {
  const response = await runic.SourcesShow({ id });
  return response;
};

export const getRunicRead = async (id: string, cats: number[]) => {
  const categories = cats.join(',');
  const response = await runic.SourcesRead({ id, categories });
  return response;
};

export const getRunicParse = async (id: string, cats: number[]) => {
  const categories = cats.join(',');
  const response = await runic.SourcesParse({ id, categories });
  return response;
};

export const getIndexersAll = async (page: number = 1, limit: number = 50) => {
  const response = await runic.IndexersIndex({ page, limit });
  return response;
};

export const getIndexer = async (id: string) => {
  const response = await runic.IndexersShow({ id });
  return response;
};

export const postIndexer = async (subject: runic.Indexer) => {
  const response = await runic.IndexersCreate({ subject });
  return response;
};

export const putIndexer = async (id: string, subject: runic.Indexer) => {
  const response = await runic.IndexersUpdate({ id, subject });
  return response;
};

export const patchIndexer = async (id: string, setting: runic.Setting) => {
  const response = await runic.IndexersSettings({ id, setting });
  return response;
};

export const deleteIndexer = async (id: string) => {
  const response = await runic.IndexersDelete({ id });
  return response;
};

export const useRunicSourcesQuery = () =>
  useQuery({
    queryKey: ['runic', 'sources'],
    queryFn: () => getRunicSources(),
    placeholderData: previousData => previousData,
    retry: false,
  });

export const useRunicReadQuery = (id: string, categories: number[]) =>
  useQuery({
    queryKey: ['runic', id, 'read', categories],
    queryFn: () => getRunicRead(id, categories),
    placeholderData: previousData => previousData,
    retry: false,
  });

export const useRunicParseQuery = (id: string, categories: number[]) =>
  useQuery({
    queryKey: ['runic', id, 'parse', categories],
    queryFn: () => getRunicParse(id, categories),
    placeholderData: previousData => previousData,
    retry: false,
  });

export const useIndexersAllQuery = () =>
  useQuery({
    queryKey: ['indexers', 'all'],
    queryFn: () => getIndexersAll(),
    placeholderData: previousData => previousData,
    retry: false,
  });

export const useIndexerQuery = id =>
  useQuery({
    queryKey: ['indexers', id],
    queryFn: () => getIndexer(id),
    placeholderData: previousData => previousData,
    retry: false,
  });

export const useIndexerMutation = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (indexer: runic.Indexer) => {
      if (!indexer?.id) {
        throw new Error('Indexer ID is required');
      }
      return putIndexer(indexer.id, indexer);
    },
    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: ['indexers', 'all'] });
    },
  });
};

export const useIndexerCreateMutation = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (indexer: runic.Indexer) => postIndexer(indexer),
    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: ['indexers', 'all'] });
    },
  });
};

export const useIndexerDeleteMutation = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (id: string) => deleteIndexer(id),
    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: ['indexers', 'all'] });
    },
  });
};

export const useIndexerSettingMutation = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (args: { id: string; setting: runic.Setting }) => patchIndexer(args.id, args.setting),
    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: ['indexers', 'all'] });
    },
  });
};
