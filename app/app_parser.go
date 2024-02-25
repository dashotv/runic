package app

import (
	"regexp"

	"github.com/dashotv/runic/newznab"
	"github.com/dashotv/runic/parser"
)

var words []string
var replacements = []*Replacement{
	{Regex: regexp.MustCompile(`\'(\w{1,2})\b`), Sub: "$1"},
	{Regex: regexp.MustCompile(`[\W]+`), Sub: " "},
}

type Replacement struct {
	Regex *regexp.Regexp
	Sub   string
}

func init() {
	initializers = append(initializers, setupParser)
}

func setupParser(a *Application) error {
	// for k, v := range a.Config.Replacements {
	// 	replacements = append(replacements, &Replacement{
	// 		Regex: regexp.MustCompile(k),
	// 		Sub:   v,
	// 	})
	// }
	for _, w := range a.Config.Words {
		replacements = append(replacements, &Replacement{
			Regex: regexp.MustCompile(`\b` + w + `\b`),
			Sub:   "",
		})
	}

	replacements = append(replacements, []*Replacement{
		{Regex: regexp.MustCompile(`\s+`), Sub: " "},
		{Regex: regexp.MustCompile(`^\s+`), Sub: ""},
		{Regex: regexp.MustCompile(`\s+$`), Sub: ""},
	}...)
	return nil
}

type Parser struct{}

func (p *Parser) Parse(list []*newznab.NZB) ([]*Release, error) {
	releases := []*Release{}

	for _, nzb := range list {
		r := &Release{
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
		r.Description = nzb.Description
		r.Size = nzb.Size
		r.PublishedAt = nzb.PubDate
		r.Downloader = "nzb"
		if nzb.IsTorrent {
			r.Downloader = "torrent"
		}

		info, err := parser.Parse(nzb.Title)
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

func (p *Parser) parseTitle(title string) string {
	for _, r := range replacements {
		title = r.Regex.ReplaceAllString(title, r.Sub)
	}
	return title
}
