package app

import (
	"fmt"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	app := Application{}
	funcs := []func(*Application) error{
		setupConfig,
		setupLogger,
		setupRunic,
		setupProcessor,
	}
	for _, f := range funcs {
		if err := f(&app); err != nil {
			t.Fatal(err)
		}
	}
	list, err := app.Runic.Read("geek", []int{5000})
	if err != nil {
		t.Fatal(err)
	}
	releases, err := app.Runic.processor.Process(list)
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range releases {
		fmt.Printf("%s\n  %s\n", r.Raw.Title, r.Info)
	}
}

func TestParser_Parse2(t *testing.T) {
	app := Application{}
	funcs := []func(*Application) error{
		setupConfig,
		setupLogger,
		setupRunic,
		setupProcessor,
	}
	for _, f := range funcs {
		if err := f(&app); err != nil {
			t.Fatal(err)
		}
	}
	list, err := app.Runic.Read("geek", []int{2000})
	if err != nil {
		t.Fatal(err)
	}
	releases, err := app.Runic.processor.Process(list)
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range releases {
		fmt.Printf("%s\n  %s\n", r.Raw.Title, r.Info)
	}
}
