package parser

// TorrentInfo is the resulting structure returned by Parse
type TorrentInfo struct {
	Title      string   `json:"title"`
	Season     int      `json:"season,omitempty"`
	Episode    int      `json:"episode,omitempty"`
	Year       int      `json:"year,omitempty"`
	Group      string   `json:"group,omitempty"`
	Website    string   `json:"website,omitempty"`
	Resolution string   `json:"resolution,omitempty"`
	Quality    string   `json:"quality,omitempty"`
	Encodings  []string `json:"encodings,omitempty"`
	Unrated    bool     `json:"unrated,omitempty"`
	Uncensored bool     `json:"uncensored,omitempty"`
	ThreeD     bool     `json:"3d,omitempty"`
	Bluray     bool     `json:"bluray,omitempty"`
	Verified   bool     `json:"verified"`
}
