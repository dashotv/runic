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

	if err := a.updateIndexes(ctx, 24*30); err != nil {
		return fae.Wrap(err, "failed to update indexes")
	}

	return nil
}

type UpdateIndexesAll struct {
	minion.WorkerDefaults[*UpdateIndexesAll]
}

func (j *UpdateIndexesAll) Kind() string { return "update_indexes_all" }
func (j *UpdateIndexesAll) Work(ctx context.Context, job *minion.Job[*UpdateIndexesAll]) error {
	a := ContextApp(ctx)
	if a == nil {
		return fae.New("app not found")
	}

	if err := a.updateIndexes(ctx, 24*30*6); err != nil {
		return fae.Wrap(err, "failed to update indexes")
	}

	return nil
}

type UpdateIndexesDaily struct {
	minion.WorkerDefaults[*UpdateIndexesDaily]
}

func (j *UpdateIndexesDaily) Kind() string { return "update_indexes_daily" }
func (j *UpdateIndexesDaily) Work(ctx context.Context, job *minion.Job[*UpdateIndexesDaily]) error {
	a := ContextApp(ctx)
	if a == nil {
		return fae.New("app not found")
	}

	if err := a.updateIndexes(ctx, 24); err != nil {
		return fae.Wrap(err, "failed to update indexes")
	}

	return nil
}

type UpdateIndexesHourly struct {
	minion.WorkerDefaults[*UpdateIndexesHourly]
}

func (j *UpdateIndexesHourly) Kind() string { return "update_indexes_hourly" }
func (j *UpdateIndexesHourly) Work(ctx context.Context, job *minion.Job[*UpdateIndexesHourly]) error {
	a := ContextApp(ctx)
	if a == nil {
		return fae.New("app not found")
	}

	if err := a.updateIndexes(ctx, 1); err != nil {
		return fae.Wrap(err, "failed to update indexes")
	}

	return nil
}

func (a *Application) updateIndexes(ctx context.Context, hours int) error {
	rl := ratelimit.New(scryRateLimit) // per second
	ch := make(chan *Release, 100)
	wg := conc.NewWaitGroup()

	now := time.Now()
	from := now.Add(-time.Duration(hours) * time.Hour)
	q := a.DB.Release.Query().Desc("published_at").GreaterThan("published_at", from).LessThan("published_at", now)

	count, err := q.Count()
	if err != nil {
		return fae.Wrap(err, "failed to count releases")
	}

	a.Workers.Log.Infow("update indexes", "hours", hours, "count", count)

	wg.Go(func() {
		defer close(ch)

		err := q.Each(100, func(r *Release) error {
			ch <- r
			return nil
		})
		if err != nil {
			a.Workers.Log.Errorf("batch releases: %s", err)
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

				// a.Workers.Log.Debugw("handle", "id", r.ID.Hex())
				rl.Take()
				if err := a.DB.Release.Update(r); err != nil {
					a.Workers.Log.Errorf("updating release (%s): %s", r.ID.Hex(), err)
				}
			}
		}
	})

	wg.Wait()

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
