// TODO: implement interfaces/types for newznab
export interface BLARG {
  blarg: string;
}
export interface NZB {
  id: string;
  title: string;
  description: string;
  size: number;
  air_date: string;
  pub_date: string;
  usenetdate: string;
  numgrabs: number;
  numcomments: number;
  comments: NZBComment[];
  poster: string;
  password: string;
  source_endpoint: string;
  source_apikey: string;
  category: string[];
  info: string;
  genre: string;
  group: string;
  team: string;
  year: number;
  resolution: string;
  video: string;
  audio: string;
  framerate: string;
  language: string;
  subs: string[];
  coverurl: string;
  publisher: string;
  backdropcoverurl: string;
  review: string;
  tvdbid: string;
  tvrageid: string;
  tvmazeid: string;
  season: string;
  episode: string;
  tvtitle: string;
  rating: number;
  imdbid: string;
  imdbtitle: string;
  imdbyear: number;
  imdbscore: number;
  imdbactors: string[];
  imdbdirector: string;
  imdbtagline: string;
  imdbplot: string;
  artist: string;
  album: string;
  tracks: string;
  seeders: number;
  peers: number;
  infohash: string;
  download_url: string;
  is_torrent: boolean;
}

export interface NZBComment {
  Title: string;
  Content: string;
  PubDate: string;
}
