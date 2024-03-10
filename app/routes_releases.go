package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GET /releases/
func (a *Application) ReleasesIndex(c echo.Context, page int, limit int) error {
	if limit == 0 {
		limit = 25
	}
	if page == 0 {
		page = 1
	}

	total, err := a.DB.Release.Query().Count()
	if err != nil {
		return err
	}

	q := a.DB.Release.Query()

	source := c.QueryParam("source")
	if source != "" {
		q = q.Where("source", source)
	}
	rType := c.QueryParam("type")
	if rType != "" {
		q = q.Where("type", rType)
	}
	resolution := c.QueryParam("resolution")
	if resolution != "" {
		q = q.Where("resolution", resolution)
	}
	group := c.QueryParam("group")
	if group != "" {
		q = q.Where("group", group)
	}
	website := c.QueryParam("website")
	if website != "" {
		q = q.Where("website", website)
	}

	list, err := q.Desc("published_at").Desc("created_at").Limit(limit).Skip((page - 1) * limit).Run()
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
	subject, err := a.DB.Release.Get(id, &Release{})
	if err != nil {
		return c.JSON(http.StatusNotFound, H{"error": true, "message": "not found"})
	}

	subject.Verified = !subject.Verified
	if err := a.DB.Release.Save(subject); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, H{"error": false})
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
