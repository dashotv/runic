package app

import (
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type RunicSourceSimple struct {
	Name string `json:"name"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

func (a *Application) SourcesIndex(c echo.Context, page int, limit int) error {
	out := make([]*Source, 0)

	list := a.Runic.Sources()
	sort.Strings(list)

	for _, n := range list {
		s, ok := a.Runic.Source(n)
		if !ok {
			return errors.New("indexer does not exist")
		}
		out = append(out, s)
	}

	return c.JSON(http.StatusOK, gin.H{"error": false, "results": out})
}

func (a *Application) SourcesCreate(c echo.Context) error {
	return c.JSON(http.StatusOK, gin.H{"error": false})
}

func (a *Application) SourcesShow(c echo.Context, id string) error {
	s, ok := a.Runic.Source(id)
	if !ok {
		return errors.New("indexer does not exist")
	}

	return c.JSON(http.StatusOK, gin.H{"error": false, "source": s})
}

func (a *Application) SourcesUpdate(c echo.Context, id string) error {
	return errors.New("not implemented")
}

func (a *Application) SourcesSettings(c echo.Context, id string) error {
	return errors.New("not implemented")
}

func (a *Application) SourcesDelete(c echo.Context, id string) error {
	return errors.New("not implemented")
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

	results, err := a.Runic.Read(id, cats)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, gin.H{"error": false, "source": id, "results": results})
}

func (a *Application) SourcesParse(c echo.Context, id string) error {
	cats, err := parseCategories(QueryString(c, "categories"))
	if err != nil {
		return err
	}

	results, err := a.Runic.Parse(id, cats)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, gin.H{"error": false, "source": id, "results": results})
}

func (a *Application) SourcesSearch(c echo.Context, id string, query string, searchType string) error {
	results, err := a.Runic.Search(id, []int{5000}, query)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, gin.H{"error": false, "source": id, "results": results})
}
