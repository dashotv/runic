package runic

import (
	"context"
	"errors"
	"fmt"

	"github.com/dashotv/runic/jackett"
	"github.com/dashotv/runic/newznab"
)

const (
	SourceUnknown = "unknown"
	SourceNewznab = "newznab"
	SourceTorznab = "torznab"
	SourceJackett = "jackett"
)

type Source struct {
	Name     string
	URL      string
	Key      string
	UserID   int
	Insecure bool
	Type     string
	Caps     *jackett.Capabilities
	client   *newznab.Client
}

func (r *Runic) Add(name, URL, key string, userID int, insecure bool) error {
	n := newznab.New(URL, key, userID, insecure)

	caps, err := n.Capabilities()
	if err != nil {
		return err
	}

	s := &Source{
		Name:     name,
		URL:      URL,
		Key:      key,
		UserID:   userID,
		Insecure: insecure,
		Type:     SourceNewznab,
		Caps:     newznabToJackett(caps),
		client:   n,
	}
	return r.addSource(name, s)
}

func (r *Runic) AddTorznab(name, URL, key string, userID int, insecure bool) error {
	n := newznab.New(URL, key, userID, insecure)

	caps, err := n.Capabilities()
	if err != nil {
		return err
	}

	s := &Source{
		Name:     name,
		URL:      URL,
		Key:      key,
		UserID:   userID,
		Insecure: insecure,
		Type:     SourceTorznab,
		Caps:     newznabToJackett(caps),
		client:   n,
	}

	return r.addSource(name, s)
}

func (r *Runic) Sources() []string {
	var sources []string
	for name := range r.sources {
		sources = append(sources, name)
	}
	return sources
}

func (r *Runic) Source(name string) (*Source, bool) {
	s, ok := r.sources[name]
	return s, ok
}

func (r *Runic) Jackett(URL, key string) error {
	j := jackett.NewJackett(&jackett.Settings{ApiURL: URL, ApiKey: key, Client: nil})

	resp, err := j.Indexers(context.Background(), true)
	if err != nil {
		return err
	}

	r.jackett.client = j
	r.jackett.indexers = resp.Indexers

	for _, indexer := range resp.Indexers {
		u := fmt.Sprintf("%s/api/v2.0/indexers/%s/results/torznab", URL, indexer.ID)
		s := &Source{
			Name:     indexer.ID,
			URL:      u,
			Key:      key,
			UserID:   0,
			Insecure: true,
			Type:     SourceJackett,
			Caps:     &indexer.Caps,
			client:   newznab.New(u, key, 0, true),
		}
		if err := r.addSource(indexer.ID, s); err != nil {
			return err
		}
	}

	return nil
}

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

func (r *Runic) addSource(name string, s *Source) error {
	if r.sources == nil {
		r.sources = make(map[string]*Source)
	}

	if _, ok := r.sources[name]; ok {
		return errors.New("indexer already exists")
	}

	s.Caps.Categories.Category = dedupCategories(s.Caps.Categories.Category)
	r.sources[name] = s

	return nil
}
