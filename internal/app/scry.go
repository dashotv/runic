package app

import (
	scry "github.com/dashotv/scry/client"
)

func init() {
	initializers = append(initializers, setupScry)
}

func setupScry(app *Application) error {
	app.Scry = scry.New(app.Config.ScryURL)
	return nil
}
