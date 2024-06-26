// Code generated by github.com/dashotv/golem. DO NOT EDIT.
package app

import (
	"context"

	"github.com/dashotv/fae"
	"github.com/dashotv/golem/plugins/cache"
	kv "github.com/philippgille/gokv/redis"
	"github.com/redis/go-redis/v9"
)

// Cache is a basic wrapper of a redis cache
// it adds a few helper methods to make it easier to use

func init() {
	initializers = append(initializers, setupCache)
	healthchecks["cache"] = checkCache
}

func setupCache(app *Application) error {
	opts := &kv.Options{Address: app.Config.RedisAddress, DB: app.Config.RedisDatabase}
	c, err := cache.New(app.Log.Named("cache"), opts)
	if err != nil {
		return fae.Wrap(err, "failed to create cache")
	}

	app.Cache = c
	return nil
}

func checkCache(app *Application) error {
	c := redis.NewClient(&redis.Options{Addr: app.Config.RedisAddress})
	if status := c.Ping(context.Background()); status.Err() != nil {
		return fae.Errorf("failed to connect to redis: %s", status.Err())
	}
	return nil
}
