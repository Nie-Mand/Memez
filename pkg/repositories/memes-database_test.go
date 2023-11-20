//go:build unit
// +build unit

package repositories

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMemesDatabase(t *testing.T) {
	assert.NotPanics(t, func () {
		m := NewMemesDatabase()
		assert.NotNil(t, m)
	})
}

func TestSaveInMemesDatabase(t *testing.T) {
	assert.NotPanics(t, func () {
		m := NewMemesDatabase()
		assert.NotNil(t, m)

		// err := m.Save(nil, nil)
		// assert.NoError(t, err)
	})
}