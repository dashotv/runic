// Code generated by github.com/dashotv/golem. DO NOT EDIT.
import { Response, scryClient } from '.';
import * as nzbgeek from './nzbgeek';

export interface NzbsMovieRequest {
  imdbid: string;
  tmdbid: string;
}
export interface NzbsMovieResponse extends Response {
  result: nzbgeek.SearchResult[];
}
export const NzbsMovie = async (params: NzbsMovieRequest) => {
  const response = await scryClient.get(`/nzbs/movie?imdbid=${params.imdbid}&tmdbid=${params.tmdbid}`);

  if (!response.data) {
    throw new Error('response empty?');
  }

  if (response.data.error) {
    if (response.data.message) {
      throw new Error(response.data.message);
    }
    throw new Error('unknown error');
  }

  return response.data as NzbsMovieResponse;
};

export interface NzbsTvRequest {
  tvdbid: string;
  season: number;
  episode: number;
}
export interface NzbsTvResponse extends Response {
  result: nzbgeek.SearchResult[];
}
export const NzbsTv = async (params: NzbsTvRequest) => {
  const response = await scryClient.get(
    `/nzbs/tv?tvdbid=${params.tvdbid}&season=${params.season}&episode=${params.episode}`,
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

  return response.data as NzbsTvResponse;
};
