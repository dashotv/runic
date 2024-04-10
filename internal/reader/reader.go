package reader

import (
	"context"
	"errors"
	"fmt"

	"github.com/samber/lo"

	"github.com/dashotv/runic/internal/jackett"
	"github.com/dashotv/runic/internal/newznab"
)

type Reader struct {
	sources map[string]Source

	jackett struct {
		client   *jackett.Jackett
		indexers []*jackett.Indexer
	}
}

func New() *Reader {
	r := &Reader{
		sources: make(map[string]Source),
	}

	return r
}

func IdentifyType(categories []int) string {
	if len(categories) == 0 {
		return ""
	}

	if lo.Contains(categories, 5070) {
		return "anime"
	}
	for _, c := range categories {
		if c >= 5000 && c < 6000 {
			return "tv"
		}
	}
	for _, c := range categories {
		if c >= 2000 && c < 3000 {
			return "movies"
		}
	}
	return ""
}

func (r *Reader) Read(source string, categories []int) ([]*newznab.NZB, error) {
	s, ok := r.sources[source]
	if !ok {
		return nil, errors.New("indexer does not exist")
	}

	return s.Read(categories)

	//	if s.Type == SourceJackett || s.Type == SourceTorznab {
	//		return s.client.SearchWithQuery(categories, "", "search")
	//	}
	//
	// return s.client.LoadRSSFeed(categories, 100)
}

func (r *Reader) Search(source string, categories []int, query string) ([]*newznab.NZB, error) {
	s, ok := r.sources[source]
	if !ok {
		return nil, errors.New("indexer does not exist")
	}

	return s.Search(categories, query, "search")
}

func (r *Reader) Add(name, URL, key string, userID int, insecure bool) error {
	n := newznab.New(URL, key, userID, insecure)

	caps, err := n.Capabilities()
	if err != nil {
		return err
	}

	s := &NewznabSource{
		Name:     name,
		URL:      URL,
		Key:      key,
		UserID:   userID,
		Insecure: insecure,
		Type:     SourceNewznab,
		Caps:     newznabToJackett(caps),
		client:   n,
	}
	s.processCategories()
	return r.addSource(name, s)
}

func (r *Reader) AddTorznab(name, URL, key string, userID int, insecure bool) error {
	n := newznab.New(URL, key, userID, insecure)

	caps, err := n.Capabilities()
	if err != nil {
		return err
	}

	s := &NewznabSource{
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

func (r *Reader) Sources() []string {
	var sources []string
	for name := range r.sources {
		sources = append(sources, name)
	}
	return sources
}

func (r *Reader) Source(name string) (Source, bool) {
	s, ok := r.sources[name]
	return s, ok
}

func (r *Reader) AddJackett(URL, key string) error {
	j := jackett.NewJackett(&jackett.Settings{ApiURL: URL, ApiKey: key, Client: nil})

	resp, err := j.Indexers(context.Background(), true)
	if err != nil {
		return err
	}

	r.jackett.client = j
	r.jackett.indexers = resp.Indexers

	for _, indexer := range resp.Indexers {
		u := fmt.Sprintf("%s/api/v2.0/indexers/%s/results/torznab", URL, indexer.ID)
		s := &JackettSource{
			Name:     indexer.ID,
			URL:      u,
			Key:      key,
			UserID:   0,
			Insecure: true,
			Type:     SourceJackett,
			Caps:     &indexer.Caps,
			client:   newznab.New(u, key, 0, true),
		}
		s.processCategories()
		if err := r.addSource(indexer.ID, s); err != nil {
			return err
		}
	}

	return nil
}

func (r *Reader) addSource(name string, s Source) error {
	if r.sources == nil {
		r.sources = make(map[string]Source)
	}

	if _, ok := r.sources[name]; ok {
		return errors.New("indexer already exists")
	}

	r.sources[name] = s

	return nil
}
