package main

import (
	"insat/devops/core/common"
	"insat/devops/core/config"
	"insat/devops/core/store"
	"insat/devops/core/store/services"
	"insat/devops/pkg/rating"
	"insat/devops/pkg/repositories"

	"github.com/labstack/echo/v4"
)

func main() {
	err := common.DisplayLogo()
	if err != nil {
		return
	}

	e := echo.New()
	config.InjectRenderer(e)

	repository := repositories.NewMemesIndexer()
	rater := rating.NewAIBasedRatingService()

	memesService := services.NewMemesService(repository, rater)

	var memezHandler store.APIServer = services.NewMemesHandler(memesService)

	e.GET("/", memezHandler.ShowIndex)
	e.POST("/", memezHandler.UploadAndRate)

	e.Logger.Fatal(e.Start(":1323"))
}
