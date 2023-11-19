package rating

import "insat/devops/core/store"

type AIBasedRatingService struct {}

func NewAIBasedRatingService() *AIBasedRatingService {
	return &AIBasedRatingService{}
}

func (AIBasedRatingService) Rate(m *store.Meme) error {
	m.Rate = "lol"
	return nil
}