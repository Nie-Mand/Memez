//go:build unit
// +build unit

package rating

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateFakeRatingService(t *testing.T) {
	assert.NotPanics(t, func () {
		m := NewFakeRatingService()
		assert.NotNil(t, m)
	})
}

func TestRateInFakeRatingService(t *testing.T) {
	assert.NotPanics(t, func () {
		m := NewFakeRatingService()
		assert.NotNil(t, m)

		// err := m.Save(nil, nil)
		// assert.NoError(t, err)
	})
}