package app

import (
	"fmt"

	"github.com/dashotv/minion"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var app *Application

type setupFunc func(app *Application) error
type healthFunc func(app *Application) error

var initializers = []setupFunc{setupConfig, setupLogger}
var healthchecks = map[string]healthFunc{}

type Application struct {
	Config *Config
	Log    *zap.SugaredLogger

	//golem:template:app/app_partial_definitions
	// DO NOT EDIT. This section is managed by github.com/dashotv/golem.
	// Routes
	Engine  *gin.Engine
	Default *gin.RouterGroup
	Router  *gin.RouterGroup

	// Models
	DB *Connector

	// Workers
	Workers *minion.Minion

	//golem:template:app/app_partial_definitions

	Runic *Runic
}

func Start() error {
	if app != nil {
		return errors.New("application already started")
	}

	app := &Application{}

	for _, f := range initializers {
		if err := f(app); err != nil {
			return err
		}
	}

	app.Log.Info("starting runic...")

	//golem:template:app/app_partial_start
	// DO NOT EDIT. This section is managed by github.com/dashotv/golem.
	go func() {
		app.Log.Infof("starting workers (%d)...", app.Config.MinionConcurrency)
		app.Workers.Start()
	}()

	app.Routes()
	app.Log.Info("starting routes...")
	if err := app.Engine.Run(fmt.Sprintf(":%d", app.Config.Port)); err != nil {
		return errors.Wrap(err, "starting router")
	}

	//golem:template:app/app_partial_start

	return nil
}

func (a *Application) Health() (map[string]bool, error) {
	resp := make(map[string]bool)
	for n, f := range healthchecks {
		err := f(a)
		resp[n] = err == nil
	}

	return resp, nil
}
