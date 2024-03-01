package app

import (
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type RunicSourceSimple struct {
	Name string `json:"name"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

func (a *Application) SourcesIndex(c *gin.Context, page int, limit int) {
	out := make([]*Source, 0)

	list := a.Runic.Sources()
	sort.Strings(list)

	for _, n := range list {
		s, ok := a.Runic.Source(n)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errors.New("indexer does not exist")})
			return
		}
		out = append(out, s)
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "results": out})
}

func (a *Application) SourcesCreate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"error": false,
	})
}

func (a *Application) SourcesShow(c *gin.Context, id string) {
	s, ok := a.Runic.Source(id)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.New("indexer does not exist")})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "source": s})
}

func (a *Application) SourcesUpdate(c *gin.Context, id string) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

func (a *Application) SourcesSettings(c *gin.Context, id string) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

func (a *Application) SourcesDelete(c *gin.Context, id string) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
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

func (a *Application) SourcesRead(c *gin.Context, id string) {
	cats, err := parseCategories(QueryString(c, "categories"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	results, err := a.Runic.Read(id, cats)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"source":  id,
		"results": results,
	})
}

func (a *Application) SourcesParse(c *gin.Context, id string) {
	cats, err := parseCategories(QueryString(c, "categories"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	results, err := a.Runic.Parse(id, cats)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"source":  id,
		"results": results,
	})
}

func (a *Application) SourcesSearch(c *gin.Context, id string, query string, searchType string) {
	results, err := a.Runic.Search(id, []int{5000}, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"source":  id,
		"results": results,
	})
}
