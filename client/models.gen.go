// Code generated by github.com/dashotv/golem. DO NOT EDIT.
package client

import (
	"time"

	"github.com/dashotv/grimoire"
	"github.com/dashotv/runic/internal/newznab"
	"github.com/dashotv/runic/internal/parser"
)

type Indexer struct { // model
	grimoire.Document `bson:",inline"` // includes default model settings
	//ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	//CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	//UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	Name        string    `bson:"name" json:"name"`
	URL         string    `bson:"url" json:"url"`
	Active      bool      `bson:"active" json:"active"`
	Categories  []int     `bson:"categories" json:"categories"`
	ProcessedAt time.Time `bson:"processed_at" json:"processed_at"`
}

type Popular struct { // struct
	Title string `bson:"_id" json:"title"`
	Year  int    `bson:"year" json:"year"`
	Type  string `bson:"type" json:"type"`
	Count int    `bson:"count" json:"count"`
}

type PopularMovie struct { // struct
	ID       *PopularMovieId `bson:"_id" json:"id"`
	Count    int             `bson:"count" json:"count"`
	Verified int             `bson:"verified" json:"verified"`
}

type PopularMovieId struct { // struct
	Title string `bson:"title" json:"title"`
	Year  int    `bson:"year" json:"year"`
}

type PopularResponse struct { // struct
	Tv     []*Popular `bson:"tv" json:"tv"`
	Anime  []*Popular `bson:"anime" json:"anime"`
	Movies []*Popular `bson:"movies" json:"movies"`
}

type Release struct { // model
	grimoire.Document `bson:",inline"` // includes default model settings
	//ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	//CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	//UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	Type        string              `bson:"type" json:"type"`
	Source      string              `bson:"source" json:"source"`
	Title       string              `bson:"title" json:"title"`
	Year        int                 `bson:"year" json:"year"`
	Description string              `bson:"description" json:"description"`
	Size        int64               `bson:"size" json:"size"`
	View        string              `bson:"view" json:"view"`
	Download    string              `bson:"download" json:"download"`
	Infohash    string              `bson:"infohash" json:"infohash"`
	Season      int                 `bson:"season" json:"season"`
	Episode     int                 `bson:"episode" json:"episode"`
	Volume      int                 `bson:"volume" json:"volume"`
	Group       string              `bson:"group" json:"group"`
	Website     string              `bson:"website" json:"website"`
	Verified    bool                `bson:"verified" json:"verified"`
	Widescreen  bool                `bson:"widescreen" json:"widescreen"`
	Unrated     bool                `bson:"unrated" json:"unrated"`
	Uncensored  bool                `bson:"uncensored" json:"uncensored"`
	Bluray      bool                `bson:"bluray" json:"bluray"`
	ThreeD      bool                `bson:"threeD" json:"threeD"`
	Resolution  string              `bson:"resolution" json:"resolution"`
	Encodings   []string            `bson:"encodings" json:"encodings"`
	Quality     string              `bson:"quality" json:"quality"`
	Raw         *newznab.NZB        `bson:"raw" json:"raw"`
	Info        *parser.TorrentInfo `bson:"info" json:"info"`
	Downloader  string              `bson:"downloader" json:"downloader"`
	Checksum    string              `bson:"checksum" json:"checksum"`
	PublishedAt time.Time           `bson:"published_at" json:"published_at"`
}
