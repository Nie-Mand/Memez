package services

import (
	"insat/devops/core/store"

	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
)

type MemesService struct {
	store.MemesRepository
	store.MemesRatingService
}

func NewMemesService(
	memesRepository store.MemesRepository,
	memesRatingService store.MemesRatingService,
) *MemesService {
	return &MemesService{
		MemesRepository:   memesRepository,
		MemesRatingService: memesRatingService,
	}
}

func (svc *MemesService) UploadAndRate(ctx echo.Context, m *store.Meme) error {
	sp := jaegertracing.CreateChildSpan(ctx, "Upload the meme and rate it")
    defer sp.Finish()
	
	err := svc.MemesRepository.Save(ctx, m)
	if err != nil {
		return err
	}

	err = svc.MemesRatingService.Rate(ctx, m)
	if err != nil {
		return err
	}

	return nil
}
