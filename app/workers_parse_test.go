package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRift_GetAll(t *testing.T) {
	err := getRiftAll(nil)
	assert.NoError(t, err)
}
