package runic

import (
	"errors"

	"github.com/dashotv/runic/jackett"
	"github.com/dashotv/runic/newznab"
)

type Runic struct {
	sources map[string]*Source

	jackett struct {
		client   *jackett.Jackett
		indexers []*jackett.Indexer
	}
}

func (r *Runic) Read(source string, categories []int) ([]*newznab.NZB, error) {
	s, ok := r.sources[source]
	if !ok {
		return nil, errors.New("indexer does not exist")
	}

	if s.Type == SourceJackett || s.Type == SourceTorznab {
		return s.client.SearchWithQuery(categories, "", "search")
	}

	return s.client.LoadRSSFeed(categories, 100)
}

func (r *Runic) Search(source string, categories []int, query string) ([]*newznab.NZB, error) {
	s, ok := r.sources[source]
	if !ok {
		return nil, errors.New("indexer does not exist")
	}

	return s.client.SearchWithQuery(categories, query, "search")
}
