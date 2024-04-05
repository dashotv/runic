package app

import (
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"

	"github.com/dashotv/runic/reader"
)

type RunicSourceSimple struct {
	Name string `json:"name"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

func (a *Application) SourcesIndex(c echo.Context, page int, limit int) error {
	out := make([]reader.Source, 0)

	list := a.Reader.Sources()
	sort.Strings(list)

	for _, n := range list {
		s, ok := a.Reader.Source(n)
		if !ok {
			return errors.New("indexer does not exist")
		}
		out = append(out, s)
	}

	return c.JSON(http.StatusOK, &Response{Error: false, Result: out})
}

func (a *Application) SourcesCreate(c echo.Context) error {
	return c.JSON(http.StatusOK, &Response{Error: false})
}

func (a *Application) SourcesShow(c echo.Context, id string) error {
	s, ok := a.Reader.Source(id)
	if !ok {
		return errors.New("indexer does not exist")
	}

	return c.JSON(http.StatusOK, &Response{Error: false, Result: s})
}

func parseCategories(categories string) ([]int, error) {
	cats := strings.Split(categories, ",")
	catsInt := make([]int, 0)
	for _, cat := range cats {
		if cat == "" {
			continue
		}

		i, err := strconv.Atoi(cat)
		if err != nil {
			return nil, err
		}

		catsInt = append(catsInt, i)
	}
	return catsInt, nil
}

func (a *Application) SourcesRead(c echo.Context, id string) error {
	cats, err := parseCategories(QueryString(c, "categories"))
	if err != nil {
		return err
	}

	results, err := a.Reader.Read(id, cats)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &Response{Error: false, Result: results})
}

func (a *Application) SourcesParse(c echo.Context, id string) error {
	cats, err := parseCategories(QueryString(c, "categories"))
	if err != nil {
		return err
	}

	results, err := a.Processor.Parse(id, cats)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &Response{Error: false, Result: results})
}

func (a *Application) SourcesSearch(c echo.Context, id string, query string, searchType string) error {
	results, err := a.Reader.Search(id, []int{5000}, query)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &Response{Error: false, Result: results})
}
