package app

import (
	"context"
	"sync"
	"time"

	"go.uber.org/ratelimit"

	"github.com/dashotv/minion"
)

var batchSize = 1000
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
	for i := 0; i < int(total); i += batchSize {
		series, err := app.DB.Release.Query().Desc("created_at").Limit(batchSize).Skip(i).Run()
		if err != nil {
			app.Workers.Log.Errorf("getting series: %s", err)
			return err
		}
		for _, s := range series {
			rl.Take()
			if err := app.DB.Release.Update(s); err != nil {
				app.Workers.Log.Errorf("updating series: %s", err)
			}
			count.Inc()
		}
		log.Debugf("series: %d/%d", count.i, total)
	}

	return nil
}
