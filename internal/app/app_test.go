package app

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
		panic(err)
	}
	if err := appSetup(); err != nil {
		panic(err)
	}
}

var envVars = []string{"CONNECTIONS", "NATS_URL", "MINION_URI", "RIFT_URL"}

func appSetup() error {
	if app != nil {
		fmt.Println("app already setup")
		return nil
	}

	err := envReplaceAll("host.docker.internal", "localhost", envVars)
	if err != nil {
		return err
	}

	app = &Application{}
	list := []func(*Application) error{
		setupConfig,
		setupLogger,
		setupEvents,
		setupDb,
		setupProcessor,
		setupRift,
	}

	for _, f := range list {
		if err := f(app); err != nil {
			return err
		}
	}

	return nil
}

func envReplaceAll(old, new string, vars []string) error {
	for _, v := range vars {
		oldval := os.Getenv(v)
		newval := strings.ReplaceAll(oldval, old, new)
		fmt.Printf("replacing %s: %s -> %s\n", v, oldval, newval)
		if err := os.Setenv(v, newval); err != nil {
			return err
		}
	}
	return nil
}
