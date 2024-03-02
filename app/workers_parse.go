package app

import (
	"context"

	"github.com/dashotv/minion"
)

type ParseActive struct {
	minion.WorkerDefaults[*ParseActive]
}

func (j *ParseActive) Kind() string { return "parse_active" }
func (j *ParseActive) Work(ctx context.Context, job *minion.Job[*ParseActive]) error {
	log := app.Log.Named("parse_active")
	list, count, err := app.DB.IndexerList(1, 100)
	if err != nil {
		return err
	}

	log.Debugf("found %d indexers, processing %d", count, len(list))
	return nil
}
