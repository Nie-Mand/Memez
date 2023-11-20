//go:build unit
// +build unit

package rating

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
func TestCreateAIBasedRatingService(t *testing.T) {
	assert.NotPanics(t, func () {
		m := NewAIBasedRatingService()
		assert.NotNil(t, m)
	})
}

func TestRateInAIBasedRatingService(t *testing.T) {
	assert.NotPanics(t, func () {
		m := NewAIBasedRatingService()
		assert.NotNil(t, m)

		// err := m.Save(nil, nil)
		// assert.NoError(t, err)
	})
}