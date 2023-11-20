//go:build unit
// +build unit

package repositories

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMemesIndexer(t *testing.T) {
	assert.NotPanics(t, func () {
		m := NewMemesIndexer()
		assert.NotNil(t, m)
	})
}

func TestSaveInMemesIndexer(t *testing.T) {
	assert.NotPanics(t, func () {
		m := NewMemesIndexer()
		assert.NotNil(t, m)

		// err := m.Save(nil, nil)
		// assert.NoError(t, err)
	})
}