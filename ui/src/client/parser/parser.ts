// TODO: implement interfaces/types for parser
export interface TorrentInfo {
  title: string;
  season: number;
  episode: number;
  year: number;
  group: string;
  website: string;
  resolution: string;
  quality: string;
  encodings: string[];
  unrated: boolean;
  uncensored: boolean;
  threeD: boolean;
  bluray: boolean;
}
