export interface IndexersResponse {
  results: Indexer[];
  count: number;
}
export interface RunicSourcesResponse {
  results: RunicSource[];
  error: boolean;
}
export interface RunicSourceResponse {
  error: boolean;
  source: RunicSource;
}
export interface RunicReadResponse {
  error: boolean;
  source: string;
  results: NZB[];
}
export interface RunicParseResponse {
  error: boolean;
  source: string;
  results: RunicRelease[];
}
/*
type Release struct { // model
	grimoire.Document `bson:",inline"` // includes default model settings
	//ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	//CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	//UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	Type        string       `bson:"type" json:"type"`
	Source      string       `bson:"source" json:"source"`
	Title       string       `bson:"title" json:"title"`
	Year        int          `bson:"year" json:"year"`
	Description string       `bson:"description" json:"description"`
	Size        int64        `bson:"size" json:"size"`
	View        string       `bson:"view" json:"view"`
	Download    string       `bson:"download" json:"download"`
	Infohash    string       `bson:"infohash" json:"infohash"`
	Season      int          `bson:"season" json:"season"`
	Episode     int          `bson:"episode" json:"episode"`
	Volume      int          `bson:"volume" json:"volume"`
	Group       string       `bson:"group" json:"group"`
	Verified    bool         `bson:"verified" json:"verified"`
	Widescreen  bool         `bson:"widescreen" json:"widescreen"`
	Uncensored  bool         `bson:"uncensored" json:"uncensored"`
	Bluray      bool         `bson:"bluray" json:"bluray"`
	Nzb         bool         `bson:"nzb" json:"nzb"`
	Resolution  string       `bson:"resolution" json:"resolution"`
	Encoding    string       `bson:"encoding" json:"encoding"`
	Quality     string       `bson:"quality" json:"quality"`
	Raw         *newznab.NZB `bson:"raw" json:"raw"`
	PublishedAt time.Time    `bson:"published_at" json:"published_at"`
}
*/
export interface RunicRelease {
  id: string;
  type: string;
  source: string;
  title: string;
  year: number;
  description: string;
  size: number;
  view: string;
  download: string;
  infohash: string;
  season: number;
  episode: number;
  volume: number;
  group: string;
  website: string;
  verified: boolean;
  widescreen: boolean;
  uncensored: boolean;
  bluray: boolean;
  is_torrent: boolean;
  resolution: string;
  encoding: string;
  quality: string;
  raw: NZB;
  downloader: string;
  published_at: string;
  created_at: string;
  updated_at: string;
}
/*
type NZB struct {
	ID          string    `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Size        int64     `json:"size,omitempty"`
	AirDate     time.Time `json:"air_date,omitempty"`
	PubDate     time.Time `json:"pub_date,omitempty"`
	UsenetDate  time.Time `json:"usenet_date,omitempty"`
	NumGrabs    int       `json:"num_grabs,omitempty"`
	NumComments int       `json:"num_comments,omitempty"`
	Comments    []Comment `json:"comments,omitempty"`
	Poster      string    `json:"poster,omitempty"`
	Password    string    `json:"password,omitempty"`

	SourceEndpoint string `json:"source_endpoint"`
	SourceAPIKey   string `json:"source_apikey"`

	Category []string `json:"category,omitempty"`
	Info     string   `json:"info,omitempty"`
	Genre    string   `json:"genre,omitempty"`
	Group    string   `json:"group,omitempty"`
	Team     string   `json:"team,omitempty"`
	Year     int      `json:"year,omitempty"`

	Resolution string   `json:"resolution,omitempty"`
	Video      string   `json:"video,omitempty"`
	Audio      string   `json:"audio,omitempty"`
	Framerate  string   `json:"framerate,omitempty"`
	Language   string   `json:"language,omitempty"`
	Subs       []string `json:"subs,omitempty"`

	CoverURL         string `json:"coverurl,omitempty"`
	Publisher        string `json:"publisher,omitempty"`
	BackdropCoverURL string `json:"backdropcoverurl,omitempty"`
	Review           string `json:"review,omitempty"`

	// TV Specific stuff
	TVDBID   string `json:"tvdbid,omitempty"`
	TVRageID string `json:"tvrageid,omitempty"`
	TVMazeID string `json:"tvmazeid,omitempty"`
	Season   string `json:"season,omitempty"`
	Episode  string `json:"episode,omitempty"`
	TVTitle  string `json:"tvtitle,omitempty"`
	Rating   int    `json:"rating,omitempty"`

	// Movie Specific stuff
	IMDBID       string   `json:"imdb,omitempty"`
	IMDBTitle    string   `json:"imdbtitle,omitempty"`
	IMDBYear     int      `json:"imdbyear,omitempty"`
	IMDBScore    float32  `json:"imdbscore,omitempty"`
	IMDBActors   []string `json:"imdbactors,omitempty"`
	IMDBDirector string   `json:"imdbdirector,omitempty"`
	IMDBTagline  string   `json:"imdbtagline,omitempty"`
	IMDBPlot     string   `json:"imdbplot,omitempty"`

	// Audio Specific stuff
	Artist string `json:"artist,omitempty"`
	Album  string `json:"album,omitempty"`
	Tracks string `json:"tracks,omitempty"`

	// Book Specific stuff
	Author      string    `json:"author,omitempty"`
	Pages       int       `json:"pages,omitempty"`
	PublishDate time.Time `json:"publish_date,omitempty"`
	BookTitle   string    `json:"booktitle,omitempty"`

	// Torznab specific stuff
	Seeders     int    `json:"seeders,omitempty"`
	Peers       int    `json:"peers,omitempty"`
	InfoHash    string `json:"infohash,omitempty"`
	DownloadURL string `json:"download_url,omitempty"`
	IsTorrent   bool   `json:"is_torrent,omitempty"`
}
*/
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
/*
type Comment struct {
	Title   string    `json:"title,omitempty"`
	Content string    `json:"content,omitempty"`
	PubDate time.Time `json:"pub_date,omitempty"`
}
*/
export interface NZBComment {
  Title: string;
  Content: string;
  PubDate: string;
}
export interface Indexer {
  id: string;
  name: string;
  url: string;
  active: boolean;
  categories: number[];
  processed_at?: string;
  created_at?: string;
  updated_at?: string;
}

export interface RunicSource {
  Name: string;
  Type: string;
  URL: string;
  Caps: RunicSourceCaps;
}
export interface RunicSourceCaps {
  Server: RunicSourceCapsServer;
  Limits: RunicSourceCapsLimits;
  Searching: RunicSourceCapsSearching;
  Categories: RunicSourceCapsCategories;
}
export interface RunicSourceCapsServer {
  Title: string;
}
export interface RunicSourceCapsLimits {
  Default: string;
  Max: string;
}
export interface RunicSourceCapsSearching {
  Search: RunicSourceCapsSearchingSearch;
  Tvsearch: RunicSourceCapsSearchingSearch;
  Moviesearch: RunicSourceCapsSearchingSearch;
  Musicsearch: RunicSourceCapsSearchingSearch;
  Audiosearch: RunicSourceCapsSearchingSearch;
  Booksearch: RunicSourceCapsSearchingSearch;
}
export interface RunicSourceCapsSearchingSearch {
  Available: string;
  SupportedParams: string[];
  SearchEngine: string;
}
export interface RunicSourceCapsCategories {
  Category: RunicSourceCapsCategoriesCategory[];
}
export interface RunicSourceCapsCategoriesCategory {
  ID: string;
  Name: string;
  Subcat: RunicSourceCapsCategoriesCategorySubcat[];
}
export interface RunicSourceCapsCategoriesCategorySubcat {
  ID: string;
  Name: string;
}
