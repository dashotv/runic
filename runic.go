package runic

import (
	"context"
	"errors"
	"fmt"

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

const (
	SourceUnknown = iota
	SourceNewznab
	SourceJackett
)

type Source struct {
	Name     string
	URL      string
	Key      string
	UserID   int
	Insecure bool
	Type     int
	client   *newznab.Client
}

func (r *Runic) addSource(name, URL, key string, userID int, srcType int, insecure bool) error {
	if _, ok := r.sources[name]; ok {
		return errors.New("indexer already exists")
	}

	s := &Source{
		Name:     name,
		URL:      URL,
		Key:      key,
		UserID:   userID,
		Insecure: insecure,
		Type:     srcType,
		client:   newznab.New(URL, key, userID, insecure),
	}

	if r.sources == nil {
		r.sources = make(map[string]*Source)
	}
	r.sources[name] = s

	return nil
}

func (r *Runic) Add(name, URL, key string, userID int, insecure bool) error {
	return r.addSource(name, URL, key, userID, SourceNewznab, insecure)
}

func (r *Runic) AddTorznab(name, URL, key string, userID int, insecure bool) error {
	return r.addSource(name, URL, key, userID, SourceJackett, insecure)
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
		if err := r.addSource(indexer.ID, u, key, 0, SourceJackett, false); err != nil {
			return err
		}
	}

	return nil
}

func (r *Runic) Read(name string, categories []int) ([]*newznab.NZB, error) {
	s, ok := r.sources[name]
	if !ok {
		return nil, errors.New("indexer does not exist")
	}

	fmt.Printf("READ:%+v\n", s)
	if s.Type == SourceJackett {
		return s.client.SearchWithQuery(categories, "", "search")
	}

	return s.client.LoadRSSFeed(categories, 100)
}
