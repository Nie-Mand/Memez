package services

import (
	"insat/devops/core/config"
	"insat/devops/core/store"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MemesHandler struct {
	store.MemesService
}

func NewMemesHandler(memesService store.MemesService) *MemesHandler {
	return &MemesHandler{
		MemesService: memesService,
	}
}

func (MemesHandler) ShowIndex(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func (h *MemesHandler) UploadAndRate(c echo.Context) error {
	content := c.FormValue("meme")

	if content == "" {
		return c.Render(http.StatusOK, "index.html", nil)
	}

	meme := store.Meme{
		ID:      "",
		Content: content,
		Rate:    "",
	}

	config.RemoteLog("Uploading and rating the meme...")

	err := h.MemesService.UploadAndRate(c, &meme)
	if err != nil {
		return c.Render(http.StatusOK, "index.html", nil)
	}

	return c.Render(http.StatusOK, "index.html", meme)
}