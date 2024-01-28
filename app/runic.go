package app

import (
	"errors"

	"github.com/dashotv/runic/jackett"
	"github.com/dashotv/runic/newznab"
)

type Runic struct {
	sources map[string]*Source
	parser  *Parser

	jackett struct {
		client   *jackett.Jackett
		indexers []*jackett.Indexer
	}
}

func New() *Runic {
	r := &Runic{
		sources: make(map[string]*Source),
		parser:  &Parser{},
	}

	return r
}

func (r *Runic) Parse(source string, categories []int) error {
	list, err := r.Read(source, categories)
	if err != nil {
		return err
	}

	return r.parser.Parse(list)
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
