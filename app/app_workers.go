// Code generated by github.com/dashotv/golem. DO NOT EDIT.
package app

import (
	"context"

	"github.com/dashotv/minion"
	"github.com/pkg/errors"
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
	go func() {
		app.Log.Infof("starting workers (%d)...", app.Config.MinionConcurrency)
		app.Workers.Start()
	}()
	return nil
}

func setupWorkers(app *Application) error {
	ctx := context.Background()

	mcfg := &minion.Config{
		Logger:      app.Log.Named("minion"),
		Debug:       app.Config.MinionDebug,
		Concurrency: app.Config.MinionConcurrency,
		BufferSize:  app.Config.MinionBufferSize,
		DatabaseURI: app.Config.MinionURI,
		Database:    app.Config.MinionDatabase,
		Collection:  app.Config.MinionCollection,
	}

	m, err := minion.New(ctx, mcfg)
	if err != nil {
		return errors.Wrap(err, "creating minion")
	}

	// add something like the below line in app.Start() (before the workers are
	// started) to subscribe to job notifications.
	// minion sends notifications as jobs are processed and change status
	// m.Subscribe(app.MinionNotification)
	// an example of the subscription function and the basic setup instructions
	// are included at the end of this file.

	if err := minion.Register[*ParseActive](m, &ParseActive{}); err != nil {
		return errors.Wrap(err, "registering worker: parse_active (ParseActive)")
	}
	if _, err := m.Schedule("0 */15 * * * *", &ParseActive{}); err != nil {
		return errors.Wrap(err, "scheduling worker: parse_active (ParseActive)")
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
