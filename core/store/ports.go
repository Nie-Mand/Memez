package store

import "github.com/labstack/echo/v4"

// API Server
type APIServer interface {
	ShowIndex(c echo.Context) error
	UploadAndRate(c echo.Context) error
}