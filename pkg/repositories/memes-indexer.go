package repositories

import (
	"insat/devops/core/store"
	"log"
	"time"

	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
)

type MemesIndexer struct {}

func NewMemesIndexer() *MemesIndexer {
	return &MemesIndexer{}
}

func (MemesIndexer) Save(ctx echo.Context, m *store.Meme) error {
	log.Println("[*] Indexing the meme...")

	sp := jaegertracing.CreateChildSpan(ctx, "Save the Meme and Index it")
    defer sp.Finish()

	time.Sleep(1000 * time.Millisecond)
	return nil
}
