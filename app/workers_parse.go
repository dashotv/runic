package app

import (
	"context"
	"fmt"
	"time"

	"github.com/dashotv/minion"
	rift "github.com/dashotv/rift/client"
)

type ParseActive struct {
	minion.WorkerDefaults[*ParseActive]
}

func (j *ParseActive) Kind() string { return "parse_active" }
func (j *ParseActive) Work(ctx context.Context, job *minion.Job[*ParseActive]) error {
	// log := app.Log.Named("parse_active")
	list, err := app.DB.IndexerActive()
	if err != nil {
		return err
	}

	// log.Debugf("processing %d indexers", len(list))
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
	// log := app.Log.Named("parse_indexer")
	indexer, err := app.DB.IndexerGet(id)
	if err != nil {
		return err
	}

	// log.Debugf("processing indexer: %s", indexer.Name)
	// start := time.Now()
	// defer func() {
	// 	log.Debugf("processing indexer: %s: done %s", indexer.Name, time.Since(start))
	// }()

	results, err := app.Processor.Parse(indexer.Name, indexer.Categories)
	if err != nil {
		return err
	}

	for _, result := range results {
		// TODO: change this to a unique index?
		count, err := app.DB.Release.Query().Where("checksum", result.Checksum).Count()
		if err != nil {
			return err
		}
		if count > 0 {
			continue
		}
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

type ParseRift struct {
	minion.WorkerDefaults[*ParseRift]
}

func (j *ParseRift) Kind() string { return "parse_rift" }
func (j *ParseRift) Work(ctx context.Context, job *minion.Job[*ParseRift]) error {
	// url := app.Config.RiftURL
	// log := app.Log.Named("parse_rift")
	// log.Debugf("parsing rift: %s", url)

	resp, err := app.Rift.VideoService.Index(context.Background(), &rift.Request{Limit: 100})
	if err != nil {
		return err
	}

	err = processRift(resp)
	if err != nil {
		return err
	}

	return nil
}

type ParseRiftAll struct {
	minion.WorkerDefaults[*ParseRiftAll]
}

func (j *ParseRiftAll) Kind() string { return "parse_rift_all" }
func (j *ParseRiftAll) Work(ctx context.Context, job *minion.Job[*ParseRiftAll]) error {
	count, err := app.Rift.VideoService.Index(context.Background(), &rift.Request{Limit: 0})
	if err != nil {
		return err
	}

	limit := 100
	for skip := 0; skip < int(count.Total); skip += limit {
		resp, err := app.Rift.VideoService.Index(context.Background(), &rift.Request{Limit: 100, Skip: skip})
		if err != nil {
			return err
		}

		err = processRift(resp)
		if err != nil {
			return err
		}
	}

	return nil
}

func processRift(resp *rift.VideosResponse) error {
	for _, video := range resp.Results {
		// app.Log.Debugf("processRift: %s %02dx%02d", video.Title, video.Season, video.Episode)
		// TODO: change this to a unique index?
		// Skip if it exists
		count, err := app.DB.Release.Query().Where("checksum", video.DisplayID).Count()
		if err != nil {
			return err
		}
		if count > 0 {
			continue
		}

		// log.Debugf("video: %s", video.Title)
		season := 1
		if video.Season != 0 {
			season = video.Season
		}

		result := &Release{
			Title:       video.Title,
			Season:      season,
			Episode:     video.Episode,
			Checksum:    video.DisplayID,
			Size:        video.Size,
			Resolution:  fmt.Sprintf("%d", video.Resolution),
			Source:      "rift",
			Type:        "anime",
			Downloader:  "metube",
			Download:    video.Download,
			View:        video.Source,
			PublishedAt: time.Now(),
		}

		if err := app.DB.Release.Save(result); err != nil {
			return err
		}
	}

	return nil
}
