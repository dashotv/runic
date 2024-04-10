package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GET /releases/
func (a *Application) ReleasesIndex(c echo.Context, page int, limit int) error {
	q := a.DB.Release.Query()

	total, err := q.Count()
	if err != nil {
		return err
	}

	list, err := q.Desc("published_at").Desc("created_at").Limit(limit).Skip((page - 1) * limit).Run()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &Response{Error: false, Result: list, Total: total})
}

// GET /releases/
func (a *Application) ReleasesSearch(c echo.Context, page int, limit int, source, kind, resolution, group, website string) error {
	q := a.DB.Release.Query()

	if source != "" {
		q = q.Where("source", source)
	}
	if kind != "" {
		q = q.Where("type", kind)
	}
	if resolution != "" {
		q = q.Where("resolution", resolution) // resolution might be a number?
	}
	if group != "" {
		q = q.Where("group", group)
	}
	if website != "" {
		q = q.Where("website", website)
	}

	total, err := q.Count()
	if err != nil {
		return err
	}

	list, err := q.Desc("published_at").Desc("created_at").Limit(limit).Skip((page - 1) * limit).Run()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &Response{Error: false, Result: list, Total: total})
}

// POST /releases/
func (a *Application) ReleasesCreate(c echo.Context, subject *Release) error {
	// TODO: implement the route
	return c.JSON(http.StatusNotImplemented, &Response{Error: true, Message: "not implmented"})
	// return c.JSON(http.StatusOK, &Response{Error: false})
}

// GET /releases/:id
func (a *Application) ReleasesShow(c echo.Context, id string) error {
	// subject, err := a.DB.Releases.Get(id)
	// if err != nil {
	//     return c.JSON(http.StatusNotFound, &Response{Error: true, Message: "not found"})
	// }

	// TODO: implement the route
	return c.JSON(http.StatusNotImplemented, &Response{Error: true, Message: "not implmented"})
	// return c.JSON(http.StatusOK, &Response{Error: false})
}

// PUT /releases/:id
func (a *Application) ReleasesUpdate(c echo.Context, id string, subject *Release) error {
	// subject, err := a.DB.Releases.Get(id)
	// if err != nil {
	//     return c.JSON(http.StatusNotFound, &Response{Error: true, Message: "not found"})
	// }

	// TODO: implement the route
	return c.JSON(http.StatusNotImplemented, &Response{Error: true, Message: "not implmented"})
	// return c.JSON(http.StatusOK, &Response{Error: false})
}

// PATCH /releases/:id
func (a *Application) ReleasesSettings(c echo.Context, id string, setting *Setting) error {
	subject, err := a.DB.Release.Get(id, &Release{})
	if err != nil {
		return c.JSON(http.StatusNotFound, &Response{Error: true, Message: "not found"})
	}

	subject.Verified = !subject.Verified
	if err := a.DB.Release.Save(subject); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &Response{Error: false})
}

// DELETE /releases/:id
func (a *Application) ReleasesDelete(c echo.Context, id string) error {
	// subject, err := a.DB.Releases.Get(id)
	// if err != nil {
	//     return c.JSON(http.StatusNotFound, &Response{Error: true, Message: "not found"})
	// }

	// TODO: implement the route
	return c.JSON(http.StatusNotImplemented, &Response{Error: true, Message: "not implmented"})
	// return c.JSON(http.StatusOK, &Response{Error: false})
}
