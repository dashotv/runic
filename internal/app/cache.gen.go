// Code generated by github.com/dashotv/golem. DO NOT EDIT.
package app

import (
	"context"

	"github.com/dashotv/fae"
	kv "github.com/philippgille/gokv/redis"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

// Cache is a basic wrapper of a redis cache
// it adds a few helper methods to make it easier to use

func init() {
	initializers = append(initializers, setupCache)
	healthchecks["cache"] = checkCache
}

func setupCache(app *Application) error {
	cache, err := NewCache(app.Log.Named("cache"), kv.Options{Address: app.Config.RedisAddress})
	if err != nil {
		return fae.Wrap(err, "failed to create cache")
	}

	app.Cache = cache
	return nil
}

func checkCache(app *Application) error {
	c := redis.NewClient(&redis.Options{Addr: app.Config.RedisAddress})
	if status := c.Ping(context.Background()); status.Err() != nil {
		return fae.Errorf("failed to connect to redis: %s", status.Err())
	}
	return nil
}

func NewCache(log *zap.SugaredLogger, options kv.Options) (*Cache, error) {
	client, err := kv.NewClient(options)
	if err != nil {
		return nil, err
	}
	return &Cache{client: &client}, nil
}

type Cache struct {
	client *kv.Client
	log    *zap.SugaredLogger
}

func (c *Cache) Set(k string, v interface{}) error {
	return c.client.Set(k, v)
}

func (c *Cache) Get(k string, v interface{}) (bool, error) {
	return c.client.Get(k, v)
}

func (c *Cache) Delete(k string) error {
	return c.client.Delete(k)
}

func (c *Cache) Fetch(k string, v interface{}, f func() (interface{}, error)) (bool, error) {
	ok, err := c.client.Get(k, v)
	// there was an error
	if err != nil {
		return ok, err
	}
	// the item was found
	if ok {
		//c.log.Infof("cache: hit: %s", k)
		return ok, nil
	}

	// get the value and set it
	v, err = f()
	if err != nil {
		return false, err
	}
	//c.log.Infof("cache: miss: %s", k)
	return false, c.client.Set(k, v)
}

func (c *Cache) Close() error {
	return c.client.Close()
}
