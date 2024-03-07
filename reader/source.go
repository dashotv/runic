package reader

import (
	"github.com/dashotv/runic/jackett"
	"github.com/dashotv/runic/newznab"
)

const (
	SourceUnknown = "unknown"
	SourceNewznab = "newznab"
	SourceTorznab = "torznab"
	SourceJackett = "jackett"
	SourceRift    = "rift"
)

type Source interface {
	Read(categories []int) ([]*newznab.NZB, error)
	Search(categories []int, query string, searchType string) ([]*newznab.NZB, error)
	Categories() []jackett.Category
}

type NewznabSource struct {
	Name     string
	URL      string
	Key      string
	UserID   int
	Insecure bool
	Type     string
	Caps     *jackett.Capabilities
	client   *newznab.Client
}

func (s *NewznabSource) processCategories() {
	s.Caps.Categories.Category = dedupCategories(s.Caps.Categories.Category)
}
func (s *NewznabSource) Read(categories []int) ([]*newznab.NZB, error) {
	return s.client.LoadRSSFeed(categories, 100)
}
func (s *NewznabSource) Search(categories []int, query, searchType string) ([]*newznab.NZB, error) {
	return s.client.SearchWithQuery(categories, query, searchType)
}
func (s *NewznabSource) Categories() []jackett.Category {
	return s.Caps.Categories.Category
}

type TorznabSource struct {
	Name     string
	URL      string
	Key      string
	UserID   int
	Insecure bool
	Type     string
	Caps     *jackett.Capabilities
	client   *newznab.Client
}

func (s *TorznabSource) processCategories() {
	s.Caps.Categories.Category = dedupCategories(s.Caps.Categories.Category)
}
func (s *TorznabSource) Read(categories []int) ([]*newznab.NZB, error) {
	return s.client.SearchWithQuery(categories, "", "search")
}
func (s *TorznabSource) Search(categories []int, query, searchType string) ([]*newznab.NZB, error) {
	return s.client.SearchWithQuery(categories, query, searchType)
}
func (s *TorznabSource) Categories() []jackett.Category {
	return s.Caps.Categories.Category
}

type JackettSource struct {
	Name     string
	URL      string
	Key      string
	UserID   int
	Insecure bool
	Type     string
	Caps     *jackett.Capabilities
	client   *newznab.Client
}

func (s *JackettSource) processCategories() {
	s.Caps.Categories.Category = dedupCategories(s.Caps.Categories.Category)
}
func (s *JackettSource) Read(categories []int) ([]*newznab.NZB, error) {
	return s.client.SearchWithQuery(categories, "", "search")
}
func (s *JackettSource) Search(categories []int, query, searchType string) ([]*newznab.NZB, error) {
	return s.client.SearchWithQuery(categories, query, searchType)
}
func (s *JackettSource) Categories() []jackett.Category {
	return s.Caps.Categories.Category
}
