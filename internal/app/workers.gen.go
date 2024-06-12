// Code generated by github.com/dashotv/golem. DO NOT EDIT.
package app

import (
	"context"

	"github.com/dashotv/fae"
	"github.com/dashotv/minion"
)

func init() {
	initializers = append(initializers, setupWorkers)
	healthchecks["workers"] = checkWorkers
	starters = append(starters, startWorkers)
}

func checkWorkers(app *Application) error {
	// TODO: workers health check
	return nil
}

func startWorkers(ctx context.Context, app *Application) error {
	ctx = context.WithValue(ctx, "app", app)

	app.Log.Debugf("starting workers (%d)", app.Config.MinionConcurrency)
	go app.Workers.Start(ctx)

	return nil
}

func setupWorkers(app *Application) error {
	mcfg := &minion.Config{
		Logger:      app.Log.Named("minion"),
		Debug:       app.Config.MinionDebug,
		Concurrency: app.Config.MinionConcurrency,
		BufferSize:  app.Config.MinionBufferSize,
		DatabaseURI: app.Config.MinionURI,
		Database:    app.Config.MinionDatabase,
		Collection:  app.Config.MinionCollection,
	}

	m, err := minion.New("runic", mcfg)
	if err != nil {
		return fae.Wrap(err, "creating minion")
	}

	// add something like the below line in app.Start() (before the workers are
	// started) to subscribe to job notifications.
	// minion sends notifications as jobs are processed and change status
	// m.Subscribe(app.MinionNotification)
	// an example of the subscription function and the basic setup instructions
	// are included at the end of this file.

	m.Queue("scraper", 2, 10, 1)

	if err := minion.RegisterWithQueue[*ParseActive](m, &ParseActive{}, "parser"); err != nil {
		return fae.Wrap(err, "registering worker: parse_active (ParseActive)")
	}
	if _, err := m.Schedule("0 */15 * * * *", &ParseActive{}); err != nil {
		return fae.Wrap(err, "scheduling worker: parse_active (ParseActive)")
	}

	if err := minion.RegisterWithQueue[*ParseIndexer](m, &ParseIndexer{}, "parser"); err != nil {
		return fae.Wrap(err, "registering worker: parse_indexer (ParseIndexer)")
	}

	if err := minion.RegisterWithQueue[*ParseRift](m, &ParseRift{}, "parser"); err != nil {
		return fae.Wrap(err, "registering worker: parse_rift (ParseRift)")
	}
	if _, err := m.Schedule("0 */15 * * * *", &ParseRift{}); err != nil {
		return fae.Wrap(err, "scheduling worker: parse_rift (ParseRift)")
	}

	if err := minion.Register[*ParseRiftAll](m, &ParseRiftAll{}); err != nil {
		return fae.Wrap(err, "registering worker: parse_rift_all (ParseRiftAll)")
	}

	if err := minion.Register[*ReleasesPopular](m, &ReleasesPopular{}); err != nil {
		return fae.Wrap(err, "registering worker: releases_popular (ReleasesPopular)")
	}
	if _, err := m.Schedule("0 */5 * * * *", &ReleasesPopular{}); err != nil {
		return fae.Wrap(err, "scheduling worker: releases_popular (ReleasesPopular)")
	}

	if err := minion.Register[*UpdateIndexes](m, &UpdateIndexes{}); err != nil {
		return fae.Wrap(err, "registering worker: update_indexes (UpdateIndexes)")
	}

	app.Workers = m
	return nil
}

// run the following commands to create the events channel and add the necessary models.
//
// > golem add event jobs event id job:*Minion
// > golem add model minion_attempt --struct started_at:time.Time duration:float64 status error 'stacktrace:[]string'
// > golem add model minion queue kind args status 'attempts:[]*MinionAttempt'
//
// then add a Connection configuration that points to the same database connection information
// as the minion database.

// // This allows you to notify other services as jobs change status.
//func (a *Application) MinionNotification(n *minion.Notification) {
//	if n.JobID == "-" {
//		return
//	}
//
//	j := &Minion{}
//	err := app.DB.Minion.Find(n.JobID, j)
//	if err != nil {
//		log.Errorf("finding job: %s", err)
//		return
//	}
//
//	if n.Event == "job:created" {
//		events.Send("runic.jobs", &EventJob{"created", j.ID.Hex(), j})
//		return
//	}
//	events.Send("runic.jobs", &EventJob{"updated", j.ID.Hex(), j})
//}
