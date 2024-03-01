package app

import (
	"os"

	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/dashotv/runic/jackett"
	"github.com/dashotv/runic/newznab"
)

func init() {
	initializers = append(initializers, setupRunic)
}

func setupRunic(app *Application) error {
	r := &Runic{}
	app.Runic = r

	if err := r.Add("geek", os.Getenv("NZBGEEK_URL"), os.Getenv("NZBGEEK_KEY"), 0, false); err != nil {
		return err
	}
	if err := r.Jackett(os.Getenv("JACKETT_URL"), os.Getenv("JACKETT_KEY")); err != nil {
		return err
	}

	return nil
}

type Runic struct {
	sources   map[string]*Source
	processor *Processor

	jackett struct {
		client   *jackett.Jackett
		indexers []*jackett.Indexer
	}
}

func New() *Runic {
	r := &Runic{
		sources:   make(map[string]*Source),
		processor: &Processor{},
	}

	return r
}

func identifyType(categories []int) string {
	if len(categories) == 0 {
		return ""
	}

	if lo.Contains(categories, 5070) {
		return "anime"
	}
	if lo.Contains(categories, 5000) {
		return "tv"
	}
	if lo.Contains(categories, 2000) {
		return "movies"
	}
	return ""
}

func (r *Runic) Parse(source string, categories []int) ([]*Release, error) {
	list, err := r.Read(source, categories)
	if err != nil {
		return nil, err
	}

	return r.processor.Process(source, list)
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
