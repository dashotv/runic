package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *Application) IndexersIndex(c echo.Context, page int, limit int) error {
	list, count, err := a.DB.IndexerList(page, limit)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &Response{Total: count, Result: list})
}

func (a *Application) IndexersCreate(c echo.Context, subject *Indexer) error {
	if err := a.DB.Indexer.Save(subject); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &Response{Error: false, Result: subject})
}

func (a *Application) IndexersShow(c echo.Context, id string) error {
	subject, err := a.DB.Indexer.Get(id, &Indexer{})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &Response{Error: false, Result: subject})
}

func (a *Application) IndexersUpdate(c echo.Context, id string, subject *Indexer) error {
	if err := a.DB.Indexer.Save(subject); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &Response{Error: false, Result: subject})
}

func (a *Application) IndexersSettings(c echo.Context, id string, setting *Setting) error {
	subject, err := a.DB.IndexerGet(id)
	if err != nil {
		return err
	}

	switch setting.Name {
	case "active":
		subject.Active = setting.Value
	}

	if err := a.DB.Indexer.Save(subject); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &Response{Error: false, Result: setting})
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

	return c.JSON(http.StatusOK, &Response{Error: false, Result: subject})
}

func (a *Application) IndexersRefresh(c echo.Context, id string) error {
	if id == "all" {
		if err := a.Workers.Enqueue(&ParseActive{}); err != nil {
			return c.JSON(http.StatusInternalServerError, &Response{Error: true, Message: err.Error()})
		}
		return c.JSON(http.StatusOK, &Response{Error: false})
	}

	indexer, err := a.DB.IndexerByName(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{Error: true, Message: err.Error()})
	}

	if err := a.Workers.Enqueue(&ParseIndexer{ID: indexer.ID.Hex(), Title: indexer.Name}); err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{Error: true, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, &Response{Error: false})
}
