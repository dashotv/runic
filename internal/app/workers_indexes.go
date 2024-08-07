package app

import (
	"context"
	"sync"
	"time"

	"go.uber.org/ratelimit"

	"github.com/sourcegraph/conc"

	"github.com/dashotv/fae"
	"github.com/dashotv/minion"
	scry "github.com/dashotv/scry/client"
)

var scryRateLimit = 100 // per second

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
	return 24 * 60 * time.Minute // TODO: this increases as the number of releases increase, not sure timeout is the right way to handle this
}
func (j *UpdateIndexes) Work(ctx context.Context, job *minion.Job[*UpdateIndexes]) error {
	a := ContextApp(ctx)
	if a == nil {
		return fae.New("app not found")
	}

	log := app.Log.Named("update_indexes")
	log.Debug("updating indices")
	// ctx, cancel := context.WithCancel(ctx)
	// defer cancel()

	rl := ratelimit.New(scryRateLimit) // per second
	ch := make(chan *Release, 100)
	wg := conc.NewWaitGroup()

	wg.Go(func() {
		defer close(ch)

		err := app.DB.Release.Query().Desc("published_at").Each(100, func(r *Release) error {
			// log.Debugw("push", "id", r.ID.Hex())
			ch <- r
			return nil
		})
		if err != nil {
			app.Workers.Log.Errorf("batch releases: %s", err)
		}
	})

	wg.Go(func() {
		for {
			select {
			case <-ctx.Done():
				return
			case r, ok := <-ch:
				if !ok {
					return
				}

				// log.Debugw("handle", "id", r.ID.Hex())
				rl.Take()
				if err := app.DB.Release.Update(r); err != nil {
					app.Workers.Log.Errorf("updating release (%s): %s", r.ID.Hex(), err)
				}
			}
		}
	})

	wg.Wait()

	log.Debug("update complete")

	return nil
}

type ResetIndexes struct {
	minion.WorkerDefaults[*ResetIndexes]
}

func (j *ResetIndexes) Kind() string { return "reset_indexes" }
func (j *ResetIndexes) Work(ctx context.Context, job *minion.Job[*ResetIndexes]) error {
	a := ContextApp(ctx)
	index := "runic_dev*"
	if a.Config.Production {
		index = "runic_prod*"
	}

	_, err := a.Scry.Es.Delete(ctx, &scry.EsDeleteRequest{Index: index})
	if err != nil {
		return fae.Wrap(err, "failed to delete media index")
	}

	if err := a.Workers.Enqueue(&UpdateIndexes{}); err != nil {
		return fae.Wrap(err, "failed to enqueue update indexes job")
	}

	return nil
}
