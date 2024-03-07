package app

import (
	"os"
	"strconv"

	rift "github.com/dashotv/rift/client"
	"github.com/dashotv/runic/newznab"
	"github.com/dashotv/runic/parser"
	"github.com/dashotv/runic/reader"
)

func init() {
	initializers = append(initializers, setupReader)
	initializers = append(initializers, setupRift)
	initializers = append(initializers, setupProcessor)
}

func setupProcessor(a *Application) error {
	app.Processor = &Processor{
		db:     a.DB,
		reader: a.Reader,
	}
	return nil
}

func setupReader(app *Application) error {
	r := &reader.Reader{}
	app.Reader = r

	if err := r.Add("geek", os.Getenv("NZBGEEK_URL"), os.Getenv("NZBGEEK_KEY"), 0, false); err != nil {
		return err
	}
	if err := r.AddJackett(os.Getenv("JACKETT_URL"), os.Getenv("JACKETT_KEY")); err != nil {
		return err
	}

	return nil
}

func setupRift(app *Application) error {
	r := rift.New(app.Config.RiftURL)
	app.Rift = r
	return nil
}

func catsToInt(cats []string) []int {
	out := make([]int, len(cats))
	for i, c := range cats {
		out[i], _ = strconv.Atoi(c)
	}
	return out
}

type Processor struct {
	db     *Connector
	reader *reader.Reader
}

func (p *Processor) Parse(source string, categories []int) ([]*Release, error) {
	list, err := p.reader.Read(source, categories)
	if err != nil {
		return nil, err
	}

	return p.Process(source, list)
}

func (p *Processor) Process(source string, list []*newznab.NZB) ([]*Release, error) {
	releases := []*Release{}

	for _, nzb := range list {
		t := reader.IdentifyType(catsToInt(nzb.Category))
		r := &Release{
			Source:      source,
			Type:        t,
			Raw:         nzb,
			Download:    nzb.DownloadURL,
			Description: nzb.Description,
			Size:        nzb.Size,
			PublishedAt: nzb.PubDate,
			Infohash:    nzb.InfoHash,
			Checksum:    nzb.ID,
		}

		if r.Checksum == "" && nzb.IsTorrent && nzb.InfoHash != "" {
			r.Checksum = nzb.InfoHash
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
