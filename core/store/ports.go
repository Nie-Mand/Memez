package store

import "github.com/labstack/echo/v4"

// API Server
type APIServer interface {

	ShowIndex(c echo.Context) error
	
	UploadAndRate(c echo.Context) error
}

// MemesService
type MemesService interface {

	UploadAndRate(m *Meme) error
}

// MemesRepository
type MemesRepository interface {

	Save(m *Meme) error
}

// MemesRatingService
type MemesRatingService interface {

	Rate(m *Meme) error
}