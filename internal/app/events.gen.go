// Code generated by github.com/dashotv/golem. DO NOT EDIT.
package app

import (
	"context"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	"go.uber.org/zap"

	"github.com/dashotv/fae"
	"github.com/dashotv/mercury"
	rift "github.com/dashotv/rift/client"
)

func init() {
	initializers = append(initializers, setupEvents)
	healthchecks["events"] = checkEvents
	starters = append(starters, startEvents)
}

type EventsChannel string
type EventsTopic string

func setupEvents(app *Application) error {
	events, err := NewEvents(app)
	if err != nil {
		return err
	}

	app.Events = events
	return nil
}

func startEvents(ctx context.Context, app *Application) error {
	go app.Events.Start(ctx)
	return nil
}

func checkEvents(app *Application) error {
	switch app.Events.Merc.Status() {
	case nats.CONNECTED:
		return nil
	default:
		return fae.Errorf("nats status: %s", app.Events.Merc.Status())
	}
}

type Events struct {
	App       *Application
	Merc      *mercury.Mercury
	Log       *zap.SugaredLogger
	Releases  chan *Release
	RiftVideo chan *rift.Video
}

func NewEvents(app *Application) (*Events, error) {
	m, err := mercury.New("runic", app.Config.NatsURL)
	if err != nil {
		return nil, err
	}

	e := &Events{
		App:       app,
		Merc:      m,
		Log:       app.Log.Named("events"),
		Releases:  make(chan *Release),
		RiftVideo: make(chan *rift.Video),
	}

	if err := e.Merc.Sender("runic.releases", e.Releases); err != nil {
		return nil, err
	}

	if err := e.Merc.Receiver("rift.video", e.RiftVideo); err != nil {
		return nil, err
	}
	return e, nil
}

func (e *Events) Start(ctx context.Context) error {
	e.Log.Debugf("starting events...")
	// receiver: RiftVideo
	for i := 0; i < 1; i++ {
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case m := <-e.RiftVideo:
					onRiftVideo(e.App, m)
				}
			}
		}()
	}

	return nil
}

func (e *Events) Send(topic EventsTopic, data any) error {
	f := func() interface{} { return e.doSend(topic, data) }

	err, ok := WithTimeout(f, time.Second*5)
	if !ok {
		e.Log.Errorf("timeout sending: %s", topic)
		return fmt.Errorf("timeout sending: %s", topic)
	}
	if err != nil {
		e.Log.Errorf("sending: %s", err)
		return fae.Wrap(err.(error), "events.send")
	}
	return nil
}

func (e *Events) doSend(topic EventsTopic, data any) error {
	switch topic {
	case "runic.releases":
		m, ok := data.(*Release)
		if !ok {
			return fae.Errorf("events.send: wrong data type: %t", data)
		}
		e.Releases <- m
	default:
		e.Log.Warnf("events.send: unknown topic: %s", topic)
	}
	return nil
}
