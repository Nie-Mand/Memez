package repositories

import (
	"insat/devops/core/store"
	"log"
	"time"

	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
)

type MemesDatabase struct {}

func NewMemesDatabase() *MemesDatabase {
	return &MemesDatabase{}
}

func (MemesDatabase) Save(ctx echo.Context, m *store.Meme) error {
	log.Println("[*] Save the meme to database...")

	sp := jaegertracing.CreateChildSpan(ctx, "Save the Meme to a database")
    defer sp.Finish()

	time.Sleep(100 * time.Millisecond)
	return nil
}
