package reader

import (
	"github.com/dashotv/runic/internal/jackett"
	"github.com/dashotv/runic/internal/newznab"
)

func dedupCategories(cats []jackett.Category) []jackett.Category {
	var out []jackett.Category
	seen := map[string]bool{}
	for _, cat := range cats {
		if _, ok := seen[cat.ID]; ok {
			continue
		}

		out = append(out, cat)
		seen[cat.ID] = true
		cat.Subcat = dedupSubcats(cat.Subcat)
	}
	return out
}
func dedupSubcats(cats []jackett.Subcat) []jackett.Subcat {
	var out []jackett.Subcat
	seen := map[string]bool{}
	for _, cat := range cats {
		if _, ok := seen[cat.ID]; ok {
			continue
		}

		out = append(out, cat)
		seen[cat.ID] = true
	}
	return out
}

func newznabToJackett(i newznab.Capabilities) *jackett.Capabilities {
	caps := &jackett.Capabilities{}
	caps.Searching.Search.Available = i.Searching.Search.Available
	caps.Searching.Search.SupportedParams = i.Searching.Search.SupportedParams
	caps.Searching.Search.SearchEngine = "raw"
	caps.Searching.TvSearch.Available = i.Searching.Search.Available
	caps.Searching.TvSearch.SupportedParams = i.Searching.Search.SupportedParams
	caps.Searching.TvSearch.SearchEngine = "raw"
	caps.Searching.MovieSearch.Available = i.Searching.Search.Available
	caps.Searching.MovieSearch.SupportedParams = i.Searching.Search.SupportedParams
	caps.Searching.MovieSearch.SearchEngine = "raw"
	caps.Searching.MusicSearch.Available = "no"
	caps.Searching.MusicSearch.SupportedParams = ""
	caps.Searching.MusicSearch.SearchEngine = "raw"
	caps.Searching.AudioSearch.Available = "no"
	caps.Searching.AudioSearch.SupportedParams = ""
	caps.Searching.AudioSearch.SearchEngine = "raw"
	caps.Searching.BookSearch.Available = "no"
	caps.Searching.BookSearch.SupportedParams = ""
	caps.Searching.BookSearch.SearchEngine = "raw"

	for _, cat := range i.Categories.Category {
		c := jackett.Category{
			ID:   cat.ID,
			Name: cat.Name,
		}
		for _, subcat := range cat.Subcat {
			c.Subcat = append(c.Subcat, jackett.Subcat{
				ID:   subcat.ID,
				Name: subcat.Name,
			})
		}
		caps.Categories.Category = append(caps.Categories.Category, c)
	}
	return caps
}
