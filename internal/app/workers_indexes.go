package app

import (
	"context"
	"sync"
	"time"

	"go.uber.org/ratelimit"

	"github.com/dashotv/minion"
)

var scryRateLimit = 150 // per second

type Count struct {
	sync.Mutex
	i int
}

func (c *Count) Inc() {
	c.Lock()
	defer c.Unlock()
	c.i++
}

type UpdateIndexes struct {
	minion.WorkerDefaults[*UpdateIndexes]
}

func (j *UpdateIndexes) Kind() string { return "UpdateIndexes" }
func (j *UpdateIndexes) Timeout(job *minion.Job[*UpdateIndexes]) time.Duration {
	return 60 * time.Minute
}
func (j *UpdateIndexes) Work(ctx context.Context, job *minion.Job[*UpdateIndexes]) error {
	log := app.Log.Named("update_indexes")
	// ctx, cancel := context.WithCancel(ctx)
	// defer cancel()

	rl := ratelimit.New(scryRateLimit) // per second

	count := &Count{}

	total, err := app.DB.Release.Query().Limit(-1).Count()
	if err != nil {
		app.Workers.Log.Errorf("getting series count: %s", err)
		return err
	}
	err = app.DB.Release.Query().Desc("published_at").Batch(1000, func(releases []*Release) error {
		for _, r := range releases {
			rl.Take()
			if err := app.DB.Release.Update(r); err != nil {
				app.Workers.Log.Errorf("updating release (%s): %s", r.ID.Hex(), err)
			}
			count.Inc()
		}
		log.Debugf("index release: %d/%d", count.i, total)
		return nil
	})
	if err != nil {
		app.Workers.Log.Errorf("batch releases: %s", err)
		return err
	}

	return nil
}
