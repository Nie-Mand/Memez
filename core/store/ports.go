package store

import "github.com/labstack/echo/v4"

// API Server
type APIServer interface {

	ShowIndex(c echo.Context) error
	
	UploadAndRate(c echo.Context) error
}

// MemesService
type MemesService interface {

	UploadAndRate(ctx echo.Context, m *Meme) error
}

// MemesRepository
type MemesRepository interface {

	Save(ctx echo.Context, m *Meme) error
}

// MemesRatingService
type MemesRatingService interface {

	Rate(ctx echo.Context, m *Meme) error
}