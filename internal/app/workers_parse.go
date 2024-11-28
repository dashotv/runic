package app

import (
	"context"
	"fmt"
	"time"

	"github.com/dashotv/fae"
	"github.com/dashotv/minion"
	rift "github.com/dashotv/rift/client"
)

type ParseActive struct {
	minion.WorkerDefaults[*ParseActive]
}

func (j *ParseActive) Kind() string { return "parse_active" }
func (j *ParseActive) Work(ctx context.Context, job *minion.Job[*ParseActive]) error {
	a := ContextApp(ctx)
	log := a.Log.Named("parse_active")

	if !a.Config.Production {
		log.Debugf("skipping: not production")
		return nil
	}

	// log := a.Log.Named("parse_active")
	list, err := a.DB.IndexerActive()
	if err != nil {
		return fae.Wrap(err, "getting active indexers")
	}

	// log.Debugf("processing %d indexers", len(list))
	for _, indexer := range list {
		a.Workers.Enqueue(&ParseIndexer{ID: indexer.ID.Hex(), Title: indexer.Name})
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
		return fae.Wrap(err, "getting indexer")
	}

	// log.Debugf("processing indexer: %s", indexer.Name)
	// start := time.Now()
	// defer func() {
	// 	log.Debugf("processing indexer: %s: done %s", indexer.Name, time.Since(start))
	// }()

	results, err := app.Processor.Parse(indexer.Name, indexer.Categories)
	if err != nil {
		return fae.Wrap(err, "parsing results")
	}

	for _, result := range results {
		// TODO: change this to a unique index?
		count, err := app.DB.Release.Query().Where("checksum", result.Checksum).Count()
		if err != nil {
			return fae.Wrap(err, "checking release")
		}
		if count > 0 {
			continue
		}
		if err := app.DB.Release.Save(result); err != nil {
			return fae.Wrap(err, "saving release")
		}
	}

	indexer.ProcessedAt = time.Now()

	if err := app.DB.Indexer.Save(indexer); err != nil {
		return fae.Wrap(err, "saving indexer")
	}

	return nil
}

type ParseRift struct {
	minion.WorkerDefaults[*ParseRift]
}

func (j *ParseRift) Kind() string { return "parse_rift" }
func (j *ParseRift) Work(ctx context.Context, job *minion.Job[*ParseRift]) error {
	a := ContextApp(ctx)
	if a == nil {
		return fae.New("ParseRift: no app in context")
	}

	resp, err := a.getRift(ctx)
	if err != nil {
		return err
	}

	err = a.processRift(resp)
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
	a := ContextApp(ctx)
	if a == nil {
		return fae.New("ParseRift: no app in context")
	}
	return a.getRiftAll(ctx)
}

func (a *Application) getRift(ctx context.Context) (*rift.VideoIndexResponse, error) {
	return app.Rift.Video.Index(ctx, &rift.VideoIndexRequest{Limit: 100, Page: 1})
}

func (a *Application) getRiftAll(ctx context.Context) error {
	count, err := a.Rift.Video.Index(ctx, &rift.VideoIndexRequest{Limit: 0})
	if err != nil {
		return err
	}

	pages := int(count.Total)/100 + 1
	limit := 100
	for page := 1; page <= pages; page++ {
		resp, err := a.Rift.Video.Index(ctx, &rift.VideoIndexRequest{Limit: limit, Page: page})
		if err != nil {
			return err
		}

		err = a.processRift(resp)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Application) processRift(resp *rift.VideoIndexResponse) error {
	for _, video := range resp.Result {
		// app.Log.Debugf("processRift: %s %02dx%02d", video.Title, video.Season, video.Episode)
		if err := a.processRiftVideo(video); err != nil {
			return err
		}
	}

	return nil
}

func (a *Application) processRiftVideo(video *rift.Video) error {
	// TODO: change this to a unique index?
	// Skip if it exists
	count, err := a.DB.Release.Query().Where("checksum", video.DisplayID).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return nil
	}

	// a.Log.Warnf("video: %s [%t] %s", video.Title, a.Config.IsVerifiedGroup(video.Source), video.Source)
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
		Download:    "metube://" + video.Download,
		Website:     video.Source,
		View:        video.View,
		PublishedAt: time.Now(),
		Verified:    a.Config.IsVerifiedGroup(video.Source),
	}

	if err := a.DB.Release.Save(result); err != nil {
		return err
	}

	return nil
}
