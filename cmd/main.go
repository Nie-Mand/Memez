package main

import (
	"insat/devops/core/common"
	"insat/devops/core/config"
	"insat/devops/core/store"
	"insat/devops/core/store/services"
	"insat/devops/pkg/rating"
	"insat/devops/pkg/repositories"

	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
)

func main() {
	err := common.DisplayLogo()
	if err != nil {
		return
	}

	// Initialize Echo
	e := echo.New()
	config.InjectRenderer(e)
	c := jaegertracing.New(e, nil)
    defer c.Close()

	// Initialize Services and Repositories
	indexer := repositories.NewMemesIndexer()
	// repository := repositories.NewMemesDatabase()

	// rater := rating.NewAIBasedRatingService()
	fake := rating.NewFakeRatingService()
	memesService := services.NewMemesService(indexer, fake)

	// Initialize API Handlers
	var memezHandler store.APIServer = services.NewMemesHandler(memesService)

	e.GET("/", memezHandler.ShowIndex)
	e.POST("/", memezHandler.UploadAndRate)

	// Start the server
	e.Logger.Fatal(e.Start(":1323"))
}
