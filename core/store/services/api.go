package services

import (
	"insat/devops/core/store"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MemesHandler struct {}

func NewMemesHandler() *MemesHandler {
	return &MemesHandler{}
}

func (MemesHandler) ShowIndex(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func (MemesHandler) UploadAndRate(c echo.Context) error {
	meme := store.Meme{
		ID:      "",
		Content: "",
		Rate:    "lol",
		// ID:      c.FormValue("id"),
		// Content: c.FormValue("content"),
		// Rate:    c.FormValue("rate"),
	}
	return c.Render(http.StatusOK, "index.html", meme)
}