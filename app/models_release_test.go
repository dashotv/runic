package app

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestModels_ReleaseUpdateEvent(t *testing.T) {
	list, err := app.DB.Release.Query().Desc("created_at").Where("source", "rift").Limit(1).Run()
	require.NoError(t, err)
	require.NotEmpty(t, list)
	require.Len(t, list, 1)

	list[0].UpdatedAt = time.Now()
	err = app.DB.Release.Save(list[0])
	require.NoError(t, err)
}
