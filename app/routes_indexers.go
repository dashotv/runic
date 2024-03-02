package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

func (a *Application) IndexersIndex(c echo.Context, page int, limit int) error {
	indexers, count, err := a.DB.IndexerList(page, limit)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, gin.H{"count": count, "results": indexers})
}

func (a *Application) IndexersCreate(c echo.Context) error {
	indexer := &Indexer{}
	if err := c.Bind(indexer); err != nil {
		return err
	}

	if err := a.DB.Indexer.Save(indexer); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, gin.H{"error": false, "result": indexer})
}

func (a *Application) IndexersShow(c echo.Context, id string) error {
	subject, err := a.DB.Indexer.Get(id, &Indexer{})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, gin.H{"error": false, "result": subject})
}

func (a *Application) IndexersUpdate(c echo.Context, id string) error {
	subject := &Indexer{}
	if err := c.Bind(subject); err != nil {
		return err
	}

	if err := a.DB.Indexer.Save(subject); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, gin.H{"error": false, "result": subject})
}

type Setting struct {
	Key   string `json:"setting"`
	Value bool   `json:"value"`
}

func (a *Application) IndexersSettings(c echo.Context, id string) error {
	subject, err := a.DB.IndexerGet(id)
	if err != nil {
		return err
	}

	setting := &Setting{}
	if err := c.Bind(setting); err != nil {
		return err
	}

	switch setting.Key {
	case "active":
		subject.Active = setting.Value
	}

	if err := a.DB.Indexer.Save(subject); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, gin.H{"error": false})
}

func (a *Application) IndexersDelete(c echo.Context, id string) error {
	// asssuming this is a CRUD route, get the subject from the database
	subject, err := a.DB.Indexer.Get(id, &Indexer{})
	if err != nil {
		return err
	}

	if err := a.DB.Indexer.Delete(subject); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, gin.H{"error": false})
}
