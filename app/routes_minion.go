package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"

	"github.com/dashotv/minion"
)

var workerlist = map[string]minion.Payload{
	"parse_active": &ParseActive{},
}

// GET /minion/
func (a *Application) MinionIndex(c echo.Context, page int, limit int) error {
	list, err := a.DB.MinionList(limit, (page-1)*limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, H{"error": true, "message": "error loading Minion"})
	}
	return c.JSON(http.StatusOK, H{"error": false, "results": list})
}

// POST /minion/
func (a *Application) MinionCreate(c echo.Context) error {
	name := QueryString(c, "name")
	j, ok := workerlist[name]
	if !ok {
		return errors.New("invalid job: " + name)
	}

	if err := app.Workers.Enqueue(j); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, H{"error": false})
}

// GET /minion/:id
func (a *Application) MinionShow(c echo.Context, id string) error {
	// subject, err := a.DB.Minion.Get(id)
	// if err != nil {
	//     return c.JSON(http.StatusNotFound, H{"error": true, "message": "not found"})
	// }

	// TODO: implement the route
	return c.JSON(http.StatusNotImplemented, H{"error": "not implmented"})
	// return c.JSON(http.StatusOK, H{"error": false})
}

// PUT /minion/:id
func (a *Application) MinionUpdate(c echo.Context, id string) error {
	// subject, err := a.DB.Minion.Get(id)
	// if err != nil {
	//     return c.JSON(http.StatusNotFound, H{"error": true, "message": "not found"})
	// }

	// TODO: implement the route
	return c.JSON(http.StatusNotImplemented, H{"error": "not implmented"})
	// return c.JSON(http.StatusOK, H{"error": false})
}

// PATCH /minion/:id
func (a *Application) MinionSettings(c echo.Context, id string) error {
	// subject, err := a.DB.Minion.Get(id)
	// if err != nil {
	//     return c.JSON(http.StatusNotFound, H{"error": true, "message": "not found"})
	// }

	// TODO: implement the route
	return c.JSON(http.StatusNotImplemented, H{"error": "not implmented"})
	// return c.JSON(http.StatusOK, H{"error": false})
}

// DELETE /minion/:id
func (a *Application) MinionDelete(c echo.Context, id string) error {
	// subject, err := a.DB.Minion.Get(id)
	// if err != nil {
	//     return c.JSON(http.StatusNotFound, H{"error": true, "message": "not found"})
	// }

	// TODO: implement the route
	return c.JSON(http.StatusNotImplemented, H{"error": "not implmented"})
	// return c.JSON(http.StatusOK, H{"error": false})
}
