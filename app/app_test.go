package app

import (
	_ "github.com/joho/godotenv/autoload"
)

var testConnector *Connector

func init() {
	app = &Application{}
	list := []func(*Application) error{
		setupConfig,
		setupLogger,
		setupEvents,
		setupDb,
	}
	for _, f := range list {
		if err := f(app); err != nil {
			panic(err)
		}
	}
	testConnector = app.DB
}
