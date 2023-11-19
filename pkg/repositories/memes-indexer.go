package repositories

import "insat/devops/core/store"

type MemesIndexer struct {}

func NewMemesIndexer() *MemesIndexer {
	return &MemesIndexer{}
}

func (MemesIndexer) Index(m *store.Meme) error {
	return nil
}
