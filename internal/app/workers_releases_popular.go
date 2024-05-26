package app

import (
	"context"

	"github.com/samber/lo"
	"github.com/sourcegraph/conc"

	"github.com/dashotv/minion"
)

type ReleasesPopular struct {
	minion.WorkerDefaults[*ReleasesPopular]
}

func (j *ReleasesPopular) Kind() string { return "releases_popular" }
func (j *ReleasesPopular) Work(ctx context.Context, job *minion.Job[*ReleasesPopular]) error {
	a := ContextApp(ctx)
	logger := a.Log.Named("releases_popular")
	wg := conc.NewWaitGroup()

	intervals := lo.Keys(popularIntervals)

	for _, interval := range intervals {
		wg.Go(func() {
			resp, err := a.DB.ReleasesPopular(interval)
			if err != nil {
				logger.Errorf("failed to get popular releases: %s: %v", interval, err)
			}

			a.Cache.Set("releases_popular_"+interval, resp)
		})
	}

	wg.Wait()

	return nil
}
