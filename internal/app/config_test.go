package app

import (
	"testing"

	"github.com/smallstep/assert"
)

func TestConfig_IsVerifiedGroup(t *testing.T) {
	assert.True(t, app.Config.IsVerifiedGroup("animexin"))
	assert.True(t, app.Config.IsVerifiedGroup("naruldonghua"))
	assert.False(t, app.Config.IsVerifiedGroup("myanime"))
	assert.False(t, app.Config.IsVerifiedGroup("donghuastream"))
}
