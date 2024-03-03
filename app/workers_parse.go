package app

import (
	"context"
	"time"

	"github.com/dashotv/minion"
)

type ParseActive struct {
	minion.WorkerDefaults[*ParseActive]
}

func (j *ParseActive) Kind() string { return "parse_active" }
func (j *ParseActive) Work(ctx context.Context, job *minion.Job[*ParseActive]) error {
	log := app.Log.Named("parse_active")
	list, err := app.DB.IndexerActive()
	if err != nil {
		return err
	}

	log.Debugf("processing %d indexers", len(list))
	for _, indexer := range list {
		app.Workers.Enqueue(&ParseIndexer{ID: indexer.ID.Hex(), Title: indexer.Name})
	}
	return nil
}

type ParseIndexer struct {
	minion.WorkerDefaults[*ParseIndexer]
	ID    string `bson:"id" json:"id"`
	Title string `bson:"title" json:"title"`
}

func (j *ParseIndexer) Kind() string { return "parse_indexer" }
func (j *ParseIndexer) Work(ctx context.Context, job *minion.Job[*ParseIndexer]) error {
	id := job.Args.ID
	log := app.Log.Named("parse_indexer")
	indexer, err := app.DB.IndexerGet(id)
	if err != nil {
		return err
	}

	log.Debugf("processing indexer: %s", indexer.Name)
	start := time.Now()
	defer func() {
		log.Debugf("processing indexer: %s: done %s", indexer.Name, time.Since(start))
	}()

	cats := []int{}
	for _, v := range indexer.Categories {
		cats = append(cats, v...)
	}

	results, err := app.Runic.Parse(indexer.Name, cats)
	if err != nil {
		return err
	}

	for _, result := range results {
		if err := app.DB.Release.Save(result); err != nil {
			return err
		}
	}

	indexer.ProcessedAt = time.Now()

	if err := app.DB.Indexer.Save(indexer); err != nil {
		return err
	}

	return nil
}