package newznab

// fork of https://github.com/mrobinsn/go-newznab/tree/master/newznab

import (
	"encoding/json"
	"encoding/xml"
	"time"

	"github.com/pkg/errors"
)

// NZB represents an NZB found on the index
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

// Comment represents a user comment left on an NZB record
type Comment struct {
	Title   string    `json:"title,omitempty"`
	Content string    `json:"content,omitempty"`
	PubDate time.Time `json:"pub_date,omitempty"`
}

// JSONString returns a JSON string representation of this NZB
func (n NZB) JSONString() string {
	jsonString, _ := json.MarshalIndent(n, "", "  ")
	return string(jsonString)
}

// JSONString returns a JSON string representation of this Comment
func (c Comment) JSONString() string {
	jsonString, _ := json.MarshalIndent(c, "", "  ")
	return string(jsonString)
}

// SearchResponse is a RSS version of the response.
type SearchResponse struct {
	Version   string `xml:"version,attr"`
	ErrorCode int    `xml:"code,attr"`
	ErrorDesc string `xml:"description,attr"`
	Channel   struct {
		Title string `xml:"title"`
		Link  struct {
			Href string `xml:"href,attr"`
			Rel  string `xml:"rel,attr"`
			Type string `xml:"type,attr"`
		} `xml:"http://www.w3.org/2005/Atom link"`
		Description string `xml:"description"`
		Language    string `xml:"language,omitempty"`
		Webmaster   string `xml:"webmaster,omitempty"`
		Category    string `xml:"category,omitempty"`
		Image       struct {
			URL         string `xml:"url"`
			Title       string `xml:"title"`
			Link        string `xml:"link"`
			Description string `xml:"description,omitempty"`
			Width       int    `xml:"width,omitempty"`
			Height      int    `xml:"height,omitempty"`
		} `xml:"image"`

		Response struct {
			Offset int `xml:"offset,attr"`
			Total  int `xml:"total,attr"`
		} `xml:"http://www.newznab.com/DTD/2010/feeds/attributes/ response"`

		// All NZBs that match the search query, up to the response limit.
		NZBs []RawNZB `xml:"item"`
	} `xml:"channel"`
}

// RawNZB represents a single NZB item in search results.
type RawNZB struct {
	Title    string `xml:"title,omitempty"`
	Link     string `xml:"link,omitempty"`
	Size     string `xml:"size,omitempty"`
	Category struct {
		Domain string `xml:"domain,attr"`
		Value  string `xml:",chardata"`
	} `xml:"category,omitempty"`

	GUID struct {
		GUID        string `xml:",chardata"`
		IsPermaLink bool   `xml:"isPermaLink,attr"`
	} `xml:"guid,omitempty"`

	Comments    string `xml:"comments"`
	Description string `xml:"description"`
	Author      string `xml:"author,omitempty"`

	Source struct {
		URL   string `xml:"url,attr"`
		Value string `xml:",chardata"`
	} `xml:"source,omitempty"`

	Date Time `xml:"pubDate,omitempty"`

	Enclosure struct {
		URL    string `xml:"url,attr"`
		Length string `xml:"length,attr"`
		Type   string `xml:"type,attr"`
	} `xml:"enclosure,omitempty"`

	Attributes []struct {
		XMLName xml.Name
		Name    string `xml:"name,attr"`
		Value   string `xml:"value,attr"`
	} `xml:"attr"`
}

type Time struct {
	time.Time
}

func (t *Time) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeToken(start); err != nil {
		return errors.Wrap(err, "failed to encode xml token")
	}
	if err := e.EncodeToken(xml.CharData([]byte(t.UTC().Format(time.RFC1123Z)))); err != nil {
		return errors.Wrap(err, "failed to encode xml token")
	}
	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return errors.Wrap(err, "failed to encode xml token")
	}
	return nil
}

func (t *Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var raw string

	err := d.DecodeElement(&raw, &start)
	if err != nil {
		return err
	}
	date, err := time.Parse(time.RFC1123Z, raw)

	if err != nil {
		return err
	}

	*t = Time{date}
	return nil

}

type Capabilities struct {
	Server struct {
		Title string `xml:"title,attr" json:"title,omitempty"`
	} `xml:"server" json:"server,omitempty"`
	Searching struct {
		Search struct {
			Available       string `xml:"available,attr" json:"available,omitempty"`
			SupportedParams string `xml:"supportedParams,attr" json:"supported_params,omitempty"`
		} `xml:"search" json:"search,omitempty"`
		TvSearch struct {
			Available       string `xml:"available,attr" json:"available,omitempty"`
			SupportedParams string `xml:"supportedParams,attr" json:"supported_params,omitempty"`
		} `xml:"tv-search" json:"tv_search,omitempty"`
		MovieSearch struct {
			Available       string `xml:"available,attr" json:"available,omitempty"`
			SupportedParams string `xml:"supportedParams,attr" json:"supported_params,omitempty"`
		} `xml:"movie-search" json:"movie_search,omitempty"`
	} `xml:"searching" json:"searching,omitempty"`
	Categories struct {
		Category []Category `xml:"category" json:"category,omitempty"`
	} `xml:"categories" json:"categories,omitempty"`
}

type Category struct {
	ID     string   `xml:"id,attr" json:"id,omitempty"`
	Name   string   `xml:"name,attr" json:"name,omitempty"`
	Subcat []Subcat `xml:"subcat" json:"subcat,omitempty"`
}
type Subcat struct {
	ID   string `xml:"id,attr" json:"id,omitempty"`
	Name string `xml:"name,attr" json:"name,omitempty"`
}

type Details struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Channel struct {
		Text string `xml:",chardata"`
		Item struct {
			Text  string `xml:",chardata"`
			Title string `xml:"title"`
			Guid  struct {
				Text        string `xml:",chardata"`
				IsPermaLink string `xml:"isPermaLink,attr"`
			} `xml:"guid"`
			Link        string `xml:"link"`
			Comments    string `xml:"comments"`
			PubDate     string `xml:"pubDate"`
			Category    string `xml:"category"`
			Description string `xml:"description"`
			Enclosure   struct {
				Text   string `xml:",chardata"`
				URL    string `xml:"url,attr"`
				Length string `xml:"length,attr"`
				Type   string `xml:"type,attr"`
			} `xml:"enclosure"`
			Attr []struct {
				Text  string `xml:",chardata"`
				Name  string `xml:"name,attr"`
				Value string `xml:"value,attr"`
			} `xml:"attr"`
		} `xml:"item"`
	} `xml:"channel"`
}
