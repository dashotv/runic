import { SearchResponse } from './search';

/*
type Release struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Year        int       `json:"year"`
	Display     string    `json:"display"`
	Raw         string    `json:"raw"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Season      int       `json:"season"`
	Episode     int       `json:"episode"`
	Volume      int       `json:"volume"`
	Size        string    `json:"size"`
	Encoding    string    `json:"encoding"`
	Resolution  string    `json:"resolution"`
	Quality     string    `json:"quality"`
	Group       string    `json:"group"`
	Author      string    `json:"author"`
	GroupAuthor string    `json:"groupauthor"`
	Verified    bool      `json:"verified"`
	Bluray      bool      `json:"bluray"`
	NZB         bool      `json:"nzb"`
	Uncensored  bool      `json:"uncensored"`
	Checksum    string    `json:"checksum"`
	View        string    `json:"view"`
	Download    string    `json:"download"`
	Source      string    `json:"source"`
	Type        string    `json:"type"`
	Created     time.Time `json:"created_at"`
	Published   time.Time `json:"published_at"`
}
*/
export interface Release {
  id: string;
  name: string;
  year: number;
  display: string;
  raw: string;
  title: string;
  description: string;
  season: number;
  episode: number;
  volume: number;
  size: string;
  encoding: string;
  resolution: string;
  quality: string;
  group: string;
  verified: boolean;
  bluray: boolean;
  nzb: boolean;
  uncensored: boolean;
  checksum: string;
  view: string;
  download: string;
  source: string;
  type: string;
  created_at: string;
  published_at: string;
}
export interface ReleaseSearchResponse extends SearchResponse {
  Releases: Release[];
}
