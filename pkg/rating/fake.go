package rating

import (
	"insat/devops/core/store"
	"log"
	"time"

	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
)

type FakeRatingService struct {}

func NewFakeRatingService() *FakeRatingService {
	return &FakeRatingService{}
}

func (FakeRatingService) Rate(ctx echo.Context, m *store.Meme) error {
	log.Println("[*] Rating the meme...")

	sp := jaegertracing.CreateChildSpan(ctx, "Rate the Meme using a Fake Rating Algorithm")
    defer sp.Finish()
	

	time.Sleep(50 * time.Millisecond)

	// TODO: Implement the AI-based rating algorithm
	m.Rate = "LOL"
	return nil
}