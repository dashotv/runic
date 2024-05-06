// Code generated by github.com/dashotv/golem. DO NOT EDIT.
import * as newznab from './newznab';
import * as parser from './parser';

export interface Batch {
  type?: string;
  titles?: string[];
}
export interface BatchResult {
  title?: string;
  info?: parser.TorrentInfo;
}
export interface Indexer {
  id?: string;
  created_at?: string;
  updated_at?: string;

  name?: string;
  url?: string;
  active?: boolean;
  categories?: number[];
  processed_at?: string;
}
export interface Popular {
  title?: string;
  year?: number;
  type?: string;
  count?: number;
}
export interface PopularMovie {
  id?: PopularMovieId;
  count?: number;
  verified?: number;
}
export interface PopularMovieId {
  title?: string;
  year?: number;
}
export interface PopularResponse {
  tv?: Popular[];
  anime?: Popular[];
  movies?: Popular[];
}
export interface Release {
  id?: string;
  created_at?: string;
  updated_at?: string;

  type?: string;
  source?: string;
  title?: string;
  year?: number;
  description?: string;
  size?: number;
  view?: string;
  download?: string;
  infohash?: string;
  season?: number;
  episode?: number;
  volume?: number;
  group?: string;
  website?: string;
  verified?: boolean;
  widescreen?: boolean;
  unrated?: boolean;
  uncensored?: boolean;
  bluray?: boolean;
  threeD?: boolean;
  resolution?: string;
  encodings?: string[];
  quality?: string;
  raw?: newznab.NZB;
  info?: parser.TorrentInfo;
  downloader?: string;
  checksum?: string;
  published_at?: string;
}
