// Code generated by github.com/dashotv/golem. DO NOT EDIT.
package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"go.infratographer.com/x/echox/echozap"
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
	app.Routes()
	app.Log.Info("starting routes...")
	if err := app.Engine.Start(fmt.Sprintf(":%d", app.Config.Port)); err != nil {
		return errors.Wrap(err, "starting router")
	}
	return nil
}

func setupRoutes(app *Application) error {
	logger := app.Log.Named("routes").Desugar()
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Recover())
	e.Use(echozap.Middleware(logger))

	app.Engine = e
	// unauthenticated routes
	app.Default = app.Engine.Group("")
	// authenticated routes (if enabled, otherwise same as default)
	app.Router = app.Engine.Group("")

	// if app.Config.Auth {
	// 	clerkSecret := app.Config.ClerkSecretKey
	// 	if clerkSecret == "" {
	// 		app.Log.Fatal("CLERK_SECRET_KEY is not set")
	// 	}
	//
	// 	clerkClient, err := clerk.NewClient(clerkSecret)
	// 	if err != nil {
	// 		app.Log.Fatalf("clerk: %s", err)
	// 	}
	//
	// 	app.Router.Use(requireSession(clerkClient))
	// }

	return nil
}

// Enable Auth and uncomment to use Clerk to manage auth
// also add this import: "github.com/clerkinc/clerk-sdk-go/clerk"
//
// requireSession wraps the clerk.RequireSession middleware
// func requireSession(client clerk.Client) HandlerFunc {
// 	requireActiveSession := clerk.RequireSessionV2(client)
// 	return func(gctx *gin.Context) {
// 		var skip = true
// 		var handler http.HandlerFunc = func(http.ResponseWriter, *http.Request) {
// 			skip = false
// 		}
// 		requireActiveSession(handler).ServeHTTP(gctx.Writer, gctx.Request)
// 		switch {
// 		case skip:
// 			gctx.AbortWithStatusJSON(http.StatusBadRequest, H{"error": "session required"})
// 		default:
// 			gctx.Next()
// 		}
// 	}
// }

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

	sources := a.Router.Group("/sources")
	sources.GET("/", a.SourcesIndexHandler)
	sources.POST("/", a.SourcesCreateHandler)
	sources.GET("/:id", a.SourcesShowHandler)
	sources.PUT("/:id", a.SourcesUpdateHandler)
	sources.PATCH("/:id", a.SourcesSettingsHandler)
	sources.DELETE("/:id", a.SourcesDeleteHandler)
	sources.GET("/:id/read", a.SourcesReadHandler)
	sources.GET("/:id/search", a.SourcesSearchHandler)
	sources.GET("/:id/parse", a.SourcesParseHandler)

}

func (a *Application) indexHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, H{
		"name": "runic",
		"routes": H{
			"indexers": "/indexers",
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
	page := QueryInt(c, "page")
	limit := QueryInt(c, "limit")
	return a.IndexersIndex(c, page, limit)
}
func (a *Application) IndexersCreateHandler(c echo.Context) error {
	return a.IndexersCreate(c)
}
func (a *Application) IndexersShowHandler(c echo.Context) error {
	id := c.Param("id")
	return a.IndexersShow(c, id)
}
func (a *Application) IndexersUpdateHandler(c echo.Context) error {
	id := c.Param("id")
	return a.IndexersUpdate(c, id)
}
func (a *Application) IndexersSettingsHandler(c echo.Context) error {
	id := c.Param("id")
	return a.IndexersSettings(c, id)
}
func (a *Application) IndexersDeleteHandler(c echo.Context) error {
	id := c.Param("id")
	return a.IndexersDelete(c, id)
}

// Sources (/sources)
func (a *Application) SourcesIndexHandler(c echo.Context) error {
	page := QueryInt(c, "page")
	limit := QueryInt(c, "limit")
	return a.SourcesIndex(c, page, limit)
}
func (a *Application) SourcesCreateHandler(c echo.Context) error {
	return a.SourcesCreate(c)
}
func (a *Application) SourcesShowHandler(c echo.Context) error {
	id := c.Param("id")
	return a.SourcesShow(c, id)
}
func (a *Application) SourcesUpdateHandler(c echo.Context) error {
	id := c.Param("id")
	return a.SourcesUpdate(c, id)
}
func (a *Application) SourcesSettingsHandler(c echo.Context) error {
	id := c.Param("id")
	return a.SourcesSettings(c, id)
}
func (a *Application) SourcesDeleteHandler(c echo.Context) error {
	id := c.Param("id")
	return a.SourcesDelete(c, id)
}
func (a *Application) SourcesReadHandler(c echo.Context) error {
	id := c.Param("id")
	return a.SourcesRead(c, id)
}
func (a *Application) SourcesSearchHandler(c echo.Context) error {
	id := c.Param("id")
	q := QueryString(c, "q")
	t := QueryString(c, "t")
	return a.SourcesSearch(c, id, q, t)
}
func (a *Application) SourcesParseHandler(c echo.Context) error {
	id := c.Param("id")
	return a.SourcesParse(c, id)
}
