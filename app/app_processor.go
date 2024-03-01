package app

import (
	"strconv"

	"github.com/dashotv/runic/newznab"
	"github.com/dashotv/runic/parser"
)

func init() {
	initializers = append(initializers, setupProcessor)
}

func setupProcessor(a *Application) error {
	return nil
}

func catsToInt(cats []string) []int {
	out := make([]int, len(cats))
	for i, c := range cats {
		out[i], _ = strconv.Atoi(c)
	}
	return out
}

type Processor struct{}

func (p *Processor) Process(list []*newznab.NZB) ([]*Release, error) {
	releases := []*Release{}

	for _, nzb := range list {
		t := identifyType(catsToInt(nzb.Category))
		r := &Release{
			Type:        t,
			Raw:         nzb,
			Download:    nzb.DownloadURL,
			Description: nzb.Description,
			Size:        nzb.Size,
			PublishedAt: nzb.PubDate,
			Infohash:    nzb.InfoHash,
		}
		if nzb.IMDBTitle != "" {
			r.Title = nzb.IMDBTitle
			r.Year = nzb.IMDBYear
		}

		r.Downloader = "nzb"
		if nzb.IsTorrent {
			r.Downloader = "torrent"
		}

		info, err := parser.Parse(nzb.Title, t)
		if err != nil {
			return nil, err
		}

		r.Info = info

		r.Title = info.Title
		r.Season = info.Season
		r.Episode = info.Episode
		r.Year = info.Year
		r.Group = info.Group
		r.Website = info.Website
		r.Resolution = info.Resolution
		r.Quality = info.Quality
		r.Encodings = info.Encodings
		r.Unrated = info.Unrated
		r.Uncensored = info.Uncensored
		r.ThreeD = info.ThreeD
		r.Bluray = info.Bluray

		releases = append(releases, r)
	}

	return releases, nil
}
