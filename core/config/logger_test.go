package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoteLog(t *testing.T) {
	assert.NotPanics(t, func() {
		RemoteLog("Test Message")
	})
}