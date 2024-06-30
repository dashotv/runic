package parser

import (
	"fmt"
	"strconv"
)

// TorrentInfo is the resulting structure returned by Parse
type TorrentInfo struct {
	Title      string   `json:"title"`
	Season     int      `json:"season"`
	Episode    int      `json:"episode"`
	Year       int      `json:"year"`
	Group      string   `json:"group"`
	Website    string   `json:"website"`
	Resolution string   `json:"resolution"`
	Quality    string   `json:"quality"`
	Encodings  []string `json:"encodings"`
	Unrated    bool     `json:"unrated"`
	Uncensored bool     `json:"uncensored"`
	ThreeD     bool     `json:"3d"`
	Bluray     bool     `json:"bluray"`
	Checksum   string   `json:"checksum"` // md5 sum
}

func (t *TorrentInfo) String() string {
	return fmt.Sprintf("%s s%02de%02d (%d) [%s/%s] r:%sp q:%s e:%v u:%t/%t 3:%t b:%t",
		t.Title,
		t.Season,
		t.Episode,
		t.Year,
		t.Group,
		t.Website,
		t.Resolution,
		t.Quality,
		t.Encodings,
		t.Unrated,
		t.Uncensored,
		t.ThreeD,
		t.Bluray,
	)
}
func (t *TorrentInfo) setYear(year string) {
	if year != "" {
		t.Year, _ = strconv.Atoi(year)
	}
}

func (t *TorrentInfo) setSeason(season string) {
	if season != "" {
		t.Season, _ = strconv.Atoi(season)
	}
}
func (t *TorrentInfo) setEpisode(episode string) {
	if episode != "" {
		t.Episode, _ = strconv.Atoi(episode)
	}
}
