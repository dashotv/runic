package app

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRift_GetAll(t *testing.T) {
	err := getRiftAll(nil)
	assert.NoError(t, err)
}

func TestRift_Get(t *testing.T) {
	app.Rift.SetDebug(true)
	resp, err := getRift(nil)
	assert.NoError(t, err)
	assert.Len(t, resp.Result, 100)
	for _, v := range resp.Result {
		fmt.Printf("video: %s %02dx%02d\n", v.Title, v.Season, v.Episode)
	}
}
