package rating

import (
	"insat/devops/core/store"
	"log"
	"time"

	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
)

type AIBasedRatingService struct {}

func NewAIBasedRatingService() *AIBasedRatingService {
	return &AIBasedRatingService{}
}

func (AIBasedRatingService) Rate(ctx echo.Context, m *store.Meme) error {
	log.Println("[*] Rating the meme...")

	sp := jaegertracing.CreateChildSpan(ctx, "Rate the Meme using AI")
    defer sp.Finish()
	

	time.Sleep(1500 * time.Millisecond)

	// Calling the AI-based rating algorithm through a gRPC call
	m.Rate = "Thats a good one!"
	return nil
}