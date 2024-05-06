package app

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/dashotv/runic/internal/parser"
)

// GET /parser/parse
func (a *Application) ParserParse(c echo.Context, title string, type_ string) error {
	info, err := parser.Parse(title, type_)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{Error: true, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, &Response{Error: false, Result: info})
}

// GET /parser/title
// ParserTitle parses the only the title of a file and returns just title, season, episode, year.
func (a *Application) ParserTitle(c echo.Context, title string, type_ string) error {
	info, err := parser.Parse(title, type_)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{Error: true, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, &Response{Error: false, Result: info})
}

// GET /parser/batch
func (a *Application) ParserBatch(c echo.Context, batch *Batch) error {
	out := []*BatchResult{}
	for _, title := range batch.Titles {
		info, err := parser.Parse(title, batch.Type)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &Response{Error: true, Message: err.Error()})
		}
		out = append(out, &BatchResult{Title: title, Info: info})
	}
	return c.JSON(http.StatusOK, &Response{Error: false, Result: out})
}
