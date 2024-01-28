package app

import (
	"os"
)

func init() {
	initializers = append(initializers, setupRunic)
}

func setupRunic(app *Application) error {
	r := &Runic{}
	app.Runic = r

	if err := r.Add("geek", os.Getenv("NZBGEEK_URL"), os.Getenv("NZBGEEK_KEY"), 0, false); err != nil {
		return err
	}
	if err := r.Jackett(os.Getenv("JACKETT_URL"), os.Getenv("JACKETT_KEY")); err != nil {
		return err
	}

	return nil
}
