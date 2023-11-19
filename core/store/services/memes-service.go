package services

import "insat/devops/core/store"

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

func (svc *MemesService) UploadAndRate(m *store.Meme) error {
	err := svc.MemesRepository.Save(m)
	if err != nil {
		return err
	}

	err = svc.MemesRatingService.Rate(m)
	if err != nil {
		return err
	}

	return nil
}
