package app

import (
	"github.com/dashotv/fae"
	rift "github.com/dashotv/rift/client"
)

func onRiftVideo(a *Application, msg *rift.Video) error {
	if !a.Config.Production {
		return nil
	}
	if err := a.processRiftVideo(msg); err != nil {
		return fae.Wrap(err, "processing rift video")
	}
	return nil
}
