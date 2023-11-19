package repositories

import (
	"insat/devops/core/store"
	"log"
)

type MemesIndexer struct {}

func NewMemesIndexer() *MemesIndexer {
	return &MemesIndexer{}
}

func (MemesIndexer) Save(m *store.Meme) error {
	log.Println("[*] Indexing the meme...")
	return nil
}
