package services

import "insat/devops/core/store"

type MemesService struct {}

func NewMemesService() *MemesService {
	return &MemesService{}
}

func (MemesService) UploadAndRate(m *store.Meme) error {
	return nil
}
