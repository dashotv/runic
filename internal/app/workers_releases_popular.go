package app

import (
	"context"

	"github.com/samber/lo"

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

	intervals := lo.Keys(popularIntervals)

	for _, interval := range intervals {
		if err := a.CachePopular(interval); err != nil {
			logger.Errorf("failed to cache popular: %s: %v", interval, err)
			return fae.Wrapf(err, "failed to cache popular: %s", interval)
		}
	}

	if err := a.CachePopularMovies(); err != nil {
		logger.Errorf("failed to cache popular movies: %v", err)
		return fae.Wrap(err, "failed to cache popular movies")
	}

	return nil
}
