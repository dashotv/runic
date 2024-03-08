package app

import (
	"net/http"

	"github.com/labstack/echo/v4/middleware"

	"github.com/dashotv/runic/static"
)

func init() {
	initializers = append(initializers, setupStatic)
}

func setupStatic(a *Application) error {
	app.Default.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       "",
		Index:      "index.html", // This is the default html page for your SPA
		Browse:     false,
		HTML5:      true,
		Filesystem: http.FS(static.FS),
	})) // https://echo.labstack.com/docs/middleware/static
	return nil
}
