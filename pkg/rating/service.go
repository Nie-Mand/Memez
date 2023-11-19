package rating

import (
	"insat/devops/core/store"
	"log"
)

type AIBasedRatingService struct {}

func NewAIBasedRatingService() *AIBasedRatingService {
	return &AIBasedRatingService{}
}

func (AIBasedRatingService) Rate(m *store.Meme) error {
	log.Println("[*] Rating the meme...")

	m.Rate = "LOL"
	return nil
}