package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GET /popular/:interval
func (a *Application) PopularIndex(c echo.Context, interval string) error {
	result, err := a.DB.ReleasesPopular(interval)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{Error: true, Message: "error loading Poplular"})
	}
	return c.JSON(http.StatusOK, &Response{Error: false, Result: result})
}

func (a *Application) PopularMovies(c echo.Context) error {
	list, err := a.DB.ReleasesPopularMovies()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{Error: true, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, &Response{Error: false, Result: list})
}
