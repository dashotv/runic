package app

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/dashotv/fae"
)

// GET /popular/:interval
func (a *Application) PopularIndex(c echo.Context, interval string) error {
	result := map[string][]*Popular{}
	ok, err := a.Cache.Get("releases_popular_"+interval, &result)
	if err != nil {
		return err
	}
	if !ok {
		return fae.Errorf("no popular releases found for interval: %s", interval)
	}

	return c.JSON(http.StatusOK, &Response{Error: false, Result: result})
}

func (a *Application) PopularMovies(c echo.Context) error {
	result := []*PopularMovie{}
	ok, err := a.Cache.Get("releases_popular_movies", &result)
	if err != nil {
		return err
	}
	if !ok {
		return fae.New("no popular releases found for movies")
	}

	return c.JSON(http.StatusOK, &Response{Error: false, Result: result})
}
