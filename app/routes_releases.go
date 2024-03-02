package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GET /releases/
func (a *Application) ReleasesIndex(c echo.Context, page int, limit int) error {
	list, total, err := a.DB.ReleaseList(page, limit)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, H{"error": false, "results": list, "total": total})
}

// POST /releases/
func (a *Application) ReleasesCreate(c echo.Context) error {
	// TODO: implement the route
	return c.JSON(http.StatusNotImplemented, H{"error": "not implmented"})
	// return c.JSON(http.StatusOK, H{"error": false})
}

// GET /releases/:id
func (a *Application) ReleasesShow(c echo.Context, id string) error {
	// subject, err := a.DB.Releases.Get(id)
	// if err != nil {
	//     return c.JSON(http.StatusNotFound, H{"error": true, "message": "not found"})
	// }

	// TODO: implement the route
	return c.JSON(http.StatusNotImplemented, H{"error": "not implmented"})
	// return c.JSON(http.StatusOK, H{"error": false})
}

// PUT /releases/:id
func (a *Application) ReleasesUpdate(c echo.Context, id string) error {
	// subject, err := a.DB.Releases.Get(id)
	// if err != nil {
	//     return c.JSON(http.StatusNotFound, H{"error": true, "message": "not found"})
	// }

	// TODO: implement the route
	return c.JSON(http.StatusNotImplemented, H{"error": "not implmented"})
	// return c.JSON(http.StatusOK, H{"error": false})
}

// PATCH /releases/:id
func (a *Application) ReleasesSettings(c echo.Context, id string) error {
	// subject, err := a.DB.Releases.Get(id)
	// if err != nil {
	//     return c.JSON(http.StatusNotFound, H{"error": true, "message": "not found"})
	// }

	// TODO: implement the route
	return c.JSON(http.StatusNotImplemented, H{"error": "not implmented"})
	// return c.JSON(http.StatusOK, H{"error": false})
}

// DELETE /releases/:id
func (a *Application) ReleasesDelete(c echo.Context, id string) error {
	// subject, err := a.DB.Releases.Get(id)
	// if err != nil {
	//     return c.JSON(http.StatusNotFound, H{"error": true, "message": "not found"})
	// }

	// TODO: implement the route
	return c.JSON(http.StatusNotImplemented, H{"error": "not implmented"})
	// return c.JSON(http.StatusOK, H{"error": false})
}
