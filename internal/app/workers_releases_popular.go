package app

import (
	"context"

	"github.com/samber/lo"
	"golang.org/x/sync/errgroup"

	"github.com/dashotv/fae"
	"github.com/dashotv/minion"
)

type ReleasesPopular struct {
	minion.WorkerDefaults[*ReleasesPopular]
}

func (j *ReleasesPopular) Kind() string { return "releases_popular" }
func (j *ReleasesPopular) Work(ctx context.Context, job *minion.Job[*ReleasesPopular]) error {
	a := ContextApp(ctx)
	logger := a.Log.Named("releases_popular")
	wg, _ := errgroup.WithContext(ctx)

	intervals := lo.Keys(popularIntervals)

	for _, interval := range intervals {
		wg.Go(func() error {
			resp, err := a.DB.ReleasesPopular(interval)
			if err != nil {
				logger.Errorf("failed to get popular releases: %s: %v", interval, err)
				return fae.Wrapf(err, "failed to get popular releases: %s", interval)
			}

			a.Cache.Set("releases_popular_"+interval, resp)
			return nil
		})
	}

	wg.Go(func() error {
		resp, err := a.DB.ReleasesPopularMovies()
		if err != nil {
			logger.Errorf("failed to get popular movies: %v", err)
			return fae.Wrap(err, "failed to get popular movies")
		}

		a.Cache.Set("releases_popular_movies", resp)
		return nil
	})

	if err := wg.Wait(); err != nil {
		return err
	}

	return nil
}
