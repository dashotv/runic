package app

import (
	"regexp"

	"github.com/dashotv/runic/newznab"
)

var words []string
var replacements = []*Replacement{
	{Regex: regexp.MustCompile(`\'(\w{1,2})\b`), Sub: "$1"},
	{Regex: regexp.MustCompile(`[\W]+`), Sub: " "},
	{Regex: regexp.MustCompile(`\s+`), Sub: " "},
	{Regex: regexp.MustCompile(`^\s+`), Sub: ""},
	{Regex: regexp.MustCompile(`\s+$`), Sub: ""},
}

type Replacement struct {
	Regex *regexp.Regexp
	Sub   string
}

func init() {
	initializers = append(initializers, setupParser)
}

func setupParser(a *Application) error {
	replacements = []*Replacement{
		{Regex: regexp.MustCompile(`\'(\w{1,2})\b`), Sub: "$1"},
		{Regex: regexp.MustCompile(`[\W]+`), Sub: " "},
	}
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
		r := &Release{}
		r.Raw = nzb
		if nzb.IMDBTitle != "" {
			r.Title = nzb.IMDBTitle
			r.Year = nzb.IMDBYear
		} else {
			r.Title = p.parseTitle(nzb.Title)
		}
		r.Description = nzb.Description
		r.Size = nzb.Size
		r.PublishedAt = nzb.PubDate

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
