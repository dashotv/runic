package app

import (
	"context"
	"strconv"

	"go.uber.org/ratelimit"

	"github.com/dashotv/fae"
	"github.com/dashotv/minion"
	"github.com/dashotv/runic/parser"
)

type MigrateTorrents struct {
	minion.WorkerDefaults[*MigrateTorrents]
}

func (j *MigrateTorrents) Kind() string { return "migrate_torrents" }
func (j *MigrateTorrents) Work(ctx context.Context, job *minion.Job[*MigrateTorrents]) error {
	a := ContextApp(ctx)
	if a == nil {
		return fae.Errorf("app not found")
	}
	l := a.Workers.Log.Named("migrate_torrents")
	l.Info("start")

	defer TickTock("migrate_torrents")()

	rl := ratelimit.New(scryRateLimit) // per second
	err := a.DB.Torrent.Query().Desc("created_at").Each(100, func(t *Torrent) error {
		rl.Take()
		r := &Release{}

		r.Type = t.Type
		r.Source = t.Source
		r.Title = t.Title
		r.Year = t.Year
		r.Description = t.Description
		size, _ := strconv.ParseInt(t.Size, 10, 64)
		r.Size = size
		r.View = t.View
		r.Download = t.Download
		r.Infohash = t.Infohash
		r.Season = t.Season
		r.Episode = t.Episode
		r.Quality = t.Quality
		r.Volume = t.Volume
		r.Group = t.Group
		if r.Group == "" && t.Author != "" {
			r.Group = t.Author
		}
		r.Verified = t.Verified
		r.Uncensored = t.Uncensored
		r.Bluray = t.Bluray
		r.Resolution = t.Resolution
		r.Encodings = []string{t.Encoding}
		r.Quality = t.Quality
		info, _ := parser.ParseTitle(t.Raw, t.Type)
		if info != nil {
			r.Info = info
		}
		r.Downloader = "torrent"
		if t.Nzb {
			r.Downloader = "nzb"
		}
		r.Checksum = t.Checksum
		r.PublishedAt = t.PublishedAt
		r.CreatedAt = t.CreatedAt
		r.UpdatedAt = t.UpdatedAt

		if err := a.DB.Release.Save(r); err != nil {
			return fae.Wrapf(err, "saving release")
		}
		return nil
	})
	if err != nil {
		return fae.Wrapf(err, "querying torrents")
	}
	return nil
}
