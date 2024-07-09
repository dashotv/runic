package app

import (
	"context"

	"github.com/dashotv/fae"
	"github.com/dashotv/minion"
)

type IndexerCounts struct {
	minion.WorkerDefaults[*IndexerCounts]
}

func (j *IndexerCounts) Kind() string { return "indexer_counts" }
func (j *IndexerCounts) Work(ctx context.Context, job *minion.Job[*IndexerCounts]) error {
	a := ContextApp(ctx)
	l := a.Workers.Log.Named("indexer_counts")
	list, err := a.DB.IndexerActive()
	if err != nil {
		return fae.Wrap(err, "getting active indexers")
	}

	for _, indexer := range list {
		total, err := a.DB.Release.Query().Where("source", indexer.Name).Count()
		if err != nil {
			l.Warnf("error counting releases for %s: %s", indexer.Name, err)
			continue
		}

		indexer.Count = total
		if err := a.DB.Indexer.Update(indexer); err != nil {
			l.Warnf("error updating indexer %s: %s", indexer.Name, err)
		}
	}
	return nil
}
