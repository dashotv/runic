package jackett

import (
	"encoding/xml"
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type IndexersResponse struct {
	XMLName  xml.Name   `xml:"indexers"`
	Indexers []*Indexer `xml:"indexer"`
}

type Indexer struct {
	ID          string       `xml:"id,attr"`
	Configured  string       `xml:"configured,attr"`
	Title       string       `xml:"title"`
	Description string       `xml:"description"`
	Link        string       `xml:"link"`
	Language    string       `xml:"language"`
	Type        string       `xml:"type"`
	Caps        Capabilities `xml:"caps"`
}

type Capabilities struct {
	Server struct {
		Title string `xml:"title,attr"`
	} `xml:"server"`
	Limits struct {
		Default string `xml:"default,attr"`
		Max     string `xml:"max,attr"`
	} `xml:"limits"`
	Searching struct {
		Search      SearchSettings `xml:"search"`
		TvSearch    SearchSettings `xml:"tv-search"`
		MovieSearch SearchSettings `xml:"movie-search"`
		MusicSearch SearchSettings `xml:"music-search"`
		AudioSearch SearchSettings `xml:"audio-search"`
		BookSearch  SearchSettings `xml:"book-search"`
	} `xml:"searching"`
	Categories Categories `xml:"categories"`
}

type SearchSettings struct {
	Available       string `xml:"available,attr"`
	SupportedParams string `xml:"supportedParams,attr"`
	SearchEngine    string `xml:"searchEngine,attr"`
}

type Categories struct {
	Category []Category `xml:"category"`
}
type Category struct {
	ID     string   `xml:"id,attr"`
	Name   string   `xml:"name,attr"`
	Subcat []Subcat `xml:"subcat"`
}
type Subcat struct {
	ID   string `xml:"id,attr"`
	Name string `xml:"name,attr"`
}

func (j *Jackett) Indexers(ctx context.Context, configuredOnly bool) (*IndexersResponse, error) {
	var resp IndexersResponse

	params := url.Values{}
	params.Add("t", "indexers")
	if configuredOnly {
		params.Add("configured", "true")
	}

	err := j.request(ctx, "/api/v2.0/indexers/all/results/torznab", params, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (j *Jackett) request(ctx context.Context, path string, params url.Values, resp any) error {
	u, err := url.Parse(j.settings.ApiURL)
	if err != nil {
		return errors.Wrapf(err, "failed to parse apiURL %q", j.settings.ApiURL)
	}
	u.Path = path
	params.Set("apikey", j.settings.ApiKey)
	u.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return errors.Wrap(err, "failed to make fetch request")
	}
	req.Header.Set("Accept", "application/json")

	res, err := j.settings.Client.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to invoke fetch request")
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read fetch data")
	}

	err = xml.Unmarshal(data, resp)
	if err != nil {
		return errors.Wrapf(err, "failed to unmarshal fetch data with url=%v and data=%v", u, string(data))
	}

	return nil
}
