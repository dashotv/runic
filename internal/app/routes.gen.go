// Code generated by github.com/dashotv/golem. DO NOT EDIT.
package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dashotv/fae"
	"github.com/dashotv/golem/plugins/router"
	"github.com/labstack/echo/v4"
)

func init() {
	initializers = append(initializers, setupRoutes)
	healthchecks["routes"] = checkRoutes
	starters = append(starters, startRoutes)
}

func checkRoutes(app *Application) error {
	// TODO: check routes
	return nil
}

func startRoutes(ctx context.Context, app *Application) error {
	go func() {
		app.Routes()
		app.Log.Info("starting routes...")
		if err := app.Engine.Start(fmt.Sprintf(":%d", app.Config.Port)); err != nil {
			app.Log.Errorf("routes: %s", err)
		}
	}()
	return nil
}

func setupRoutes(app *Application) error {
	logger := app.Log.Named("routes").Desugar()
	e, err := router.New(logger)
	if err != nil {
		return fae.Wrap(err, "router plugin")
	}
	app.Engine = e
	// unauthenticated routes
	app.Default = app.Engine.Group("")
	// authenticated routes (if enabled, otherwise same as default)
	app.Router = app.Engine.Group("")

	// TODO: fix auth
	if app.Config.Auth {
		clerkSecret := app.Config.ClerkSecretKey
		if clerkSecret == "" {
			app.Log.Fatal("CLERK_SECRET_KEY is not set")
		}
		clerkToken := app.Config.ClerkToken
		if clerkToken == "" {
			app.Log.Fatal("CLERK_TOKEN is not set")
		}

		app.Router.Use(router.ClerkAuth(clerkSecret, clerkToken))
	}

	return nil
}

type Setting struct {
	Name  string `json:"name"`
	Value bool   `json:"value"`
}

type SettingsBatch struct {
	IDs   []string `json:"ids"`
	Name  string   `json:"name"`
	Value bool     `json:"value"`
}

type Response struct {
	Error   bool        `json:"error"`
	Message string      `json:"message,omitempty"`
	Result  interface{} `json:"result,omitempty"`
	Total   int64       `json:"total,omitempty"`
}

func (a *Application) Routes() {
	a.Default.GET("/", a.indexHandler)
	a.Default.GET("/health", a.healthHandler)

	indexers := a.Router.Group("/indexers")
	indexers.GET("/", a.IndexersIndexHandler)
	indexers.POST("/", a.IndexersCreateHandler)
	indexers.GET("/:id", a.IndexersShowHandler)
	indexers.PUT("/:id", a.IndexersUpdateHandler)
	indexers.PATCH("/:id", a.IndexersSettingsHandler)
	indexers.DELETE("/:id", a.IndexersDeleteHandler)
	indexers.GET("/refresh", a.IndexersRefreshHandler)

	popular := a.Router.Group("/popular")
	popular.GET("/:interval", a.PopularIndexHandler)

	releases := a.Router.Group("/releases")
	releases.GET("/", a.ReleasesIndexHandler)
	releases.POST("/", a.ReleasesCreateHandler)
	releases.GET("/:id", a.ReleasesShowHandler)
	releases.PUT("/:id", a.ReleasesUpdateHandler)
	releases.PATCH("/:id", a.ReleasesSettingsHandler)
	releases.DELETE("/:id", a.ReleasesDeleteHandler)
	releases.GET("/search", a.ReleasesSearchHandler)

	sources := a.Router.Group("/sources")
	sources.GET("/", a.SourcesIndexHandler)
	sources.GET("/:id", a.SourcesShowHandler)
	sources.GET("/:id/read", a.SourcesReadHandler)
	sources.GET("/:id/search", a.SourcesSearchHandler)
	sources.GET("/:id/parse", a.SourcesParseHandler)

}

func (a *Application) indexHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, H{
		"name": "runic",
		"routes": H{
			"indexers": "/indexers",
			"popular":  "/popular",
			"releases": "/releases",
			"sources":  "/sources",
		},
	})
}

func (a *Application) healthHandler(c echo.Context) error {
	health, err := a.Health()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, H{"name": "runic", "health": health})
}

// Indexers (/indexers)
func (a *Application) IndexersIndexHandler(c echo.Context) error {
	page := QueryParamIntDefault(c, "page", "1")
	limit := QueryParamIntDefault(c, "limit", "25")
	return a.IndexersIndex(c, page, limit)
}
func (a *Application) IndexersCreateHandler(c echo.Context) error {
	subject := &Indexer{}
	if err := c.Bind(subject); err != nil {
		return err
	}
	return a.IndexersCreate(c, subject)
}
func (a *Application) IndexersShowHandler(c echo.Context) error {
	id := c.Param("id")
	return a.IndexersShow(c, id)
}
func (a *Application) IndexersUpdateHandler(c echo.Context) error {
	id := c.Param("id")
	subject := &Indexer{}
	if err := c.Bind(subject); err != nil {
		return err
	}
	return a.IndexersUpdate(c, id, subject)
}
func (a *Application) IndexersSettingsHandler(c echo.Context) error {
	id := c.Param("id")
	setting := &Setting{}
	if err := c.Bind(setting); err != nil {
		return err
	}
	return a.IndexersSettings(c, id, setting)
}
func (a *Application) IndexersDeleteHandler(c echo.Context) error {
	id := c.Param("id")
	return a.IndexersDelete(c, id)
}
func (a *Application) IndexersRefreshHandler(c echo.Context) error {
	id := QueryParamString(c, "id")
	return a.IndexersRefresh(c, id)
}

// Popular (/popular)
func (a *Application) PopularIndexHandler(c echo.Context) error {
	interval := c.Param("interval")
	return a.PopularIndex(c, interval)
}

// Releases (/releases)
func (a *Application) ReleasesIndexHandler(c echo.Context) error {
	page := QueryParamIntDefault(c, "page", "1")
	limit := QueryParamIntDefault(c, "limit", "25")
	return a.ReleasesIndex(c, page, limit)
}
func (a *Application) ReleasesCreateHandler(c echo.Context) error {
	subject := &Release{}
	if err := c.Bind(subject); err != nil {
		return err
	}
	return a.ReleasesCreate(c, subject)
}
func (a *Application) ReleasesShowHandler(c echo.Context) error {
	id := c.Param("id")
	return a.ReleasesShow(c, id)
}
func (a *Application) ReleasesUpdateHandler(c echo.Context) error {
	id := c.Param("id")
	subject := &Release{}
	if err := c.Bind(subject); err != nil {
		return err
	}
	return a.ReleasesUpdate(c, id, subject)
}
func (a *Application) ReleasesSettingsHandler(c echo.Context) error {
	id := c.Param("id")
	setting := &Setting{}
	if err := c.Bind(setting); err != nil {
		return err
	}
	return a.ReleasesSettings(c, id, setting)
}
func (a *Application) ReleasesDeleteHandler(c echo.Context) error {
	id := c.Param("id")
	return a.ReleasesDelete(c, id)
}
func (a *Application) ReleasesSearchHandler(c echo.Context) error {
	page := QueryParamIntDefault(c, "page", "1")
	limit := QueryParamIntDefault(c, "limit", "25")
	source := QueryParamString(c, "source")
	kind := QueryParamString(c, "kind")
	resolution := QueryParamString(c, "resolution")
	group := QueryParamString(c, "group")
	website := QueryParamString(c, "website")
	return a.ReleasesSearch(c, page, limit, source, kind, resolution, group, website)
}

// Sources (/sources)
func (a *Application) SourcesIndexHandler(c echo.Context) error {
	page := QueryParamInt(c, "page")
	limit := QueryParamInt(c, "limit")
	return a.SourcesIndex(c, page, limit)
}
func (a *Application) SourcesShowHandler(c echo.Context) error {
	id := c.Param("id")
	return a.SourcesShow(c, id)
}
func (a *Application) SourcesReadHandler(c echo.Context) error {
	id := c.Param("id")
	categories := QueryParamString(c, "categories")
	return a.SourcesRead(c, id, categories)
}
func (a *Application) SourcesSearchHandler(c echo.Context) error {
	id := c.Param("id")
	q := QueryParamString(c, "q")
	t := QueryParamString(c, "t")
	return a.SourcesSearch(c, id, q, t)
}
func (a *Application) SourcesParseHandler(c echo.Context) error {
	id := c.Param("id")
	categories := QueryParamString(c, "categories")
	return a.SourcesParse(c, id, categories)
}
