// Code generated by github.com/dashotv/golem. DO NOT EDIT.
import { Response, scryClient } from '.';
import * as search from './search';

export interface RunicIndexRequest {
  start: number;
  limit: number;
  type: string;
  text: string;
  year: number;
  season: number;
  episode: number;
  group: string;
  website: string;
  resolution: number;
  source: string;
  uncensored: boolean;
  bluray: boolean;
  verified: boolean;
  exact: boolean;
}
export interface RunicIndexResponse extends Response {
  result: search.RunicSearchResponse;
  total: number;
}
export const RunicIndex = async (params: RunicIndexRequest) => {
  const response = await scryClient.get(
    `/runic/?start=${params.start}&limit=${params.limit}&type=${params.type}&text=${params.text}&year=${params.year}&season=${params.season}&episode=${params.episode}&group=${params.group}&website=${params.website}&resolution=${params.resolution}&source=${params.source}&uncensored=${params.uncensored}&bluray=${params.bluray}&verified=${params.verified}&exact=${params.exact}`,
  );

  if (!response.data) {
    throw new Error('response empty?');
  }

  if (response.data.error) {
    if (response.data.message) {
      throw new Error(response.data.message);
    }
    throw new Error('unknown error');
  }

  return response.data as RunicIndexResponse;
};
