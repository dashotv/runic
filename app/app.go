package app

import (
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

	//golem:template:app/app_partial_definitions

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
