package services

import (
	"bytes"
	"insat/devops/core/config"
	"insat/devops/core/store"
	"insat/devops/pkg/rating"
	"insat/devops/pkg/repositories"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func setup() (*echo.Echo, store.APIServer, func()) {
	e := echo.New()
	config.InjectRenderer(e, config.WithTemplatesDir("../../../templates"))
	c := jaegertracing.New(e, nil)

	// Initialize Services and Repositories
	indexer := repositories.NewMemesIndexer()

	// rater := rating.NewAIBasedRatingService()
	fake := rating.NewFakeRatingService()
	memesService := NewMemesService(indexer, fake)

	// Initialize API Handlers
	var memezHandler store.APIServer = NewMemesHandler(memesService)

	return e, memezHandler, func() {
		defer c.Close()
	}
}

func TestCreateApiServer(t *testing.T) {
	assert.NotPanics(t, func () {
		_, svr, clean := setup()
		defer clean()
		assert.NotNil(t, svr)
	})
}

func TestShowIndex(t *testing.T) {
	assert.NotPanics(t, func () {
		e, svr, clean := setup()
		defer clean()
		assert.NotNil(t, svr)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
    	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    	rec := httptest.NewRecorder()
    	ctx := e.NewContext(req, rec)
		err := svr.ShowIndex(ctx)
		assert.Nil(t, err)
	})
}

func TestUploadAndRate(t *testing.T) {
	assert.NotPanics(t, func () {
		e, svr, clean := setup()
		defer clean()
		assert.NotNil(t, svr)

		// Generate multipart form
		body := new(bytes.Buffer)
    	writer := multipart.NewWriter(body)
    	writer.WriteField("meme", "knock knock")
    	writer.Close()

		req := httptest.NewRequest(http.MethodPost, "/", body)
    	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    	rec := httptest.NewRecorder()
    	ctx := e.NewContext(req, rec)
		err := svr.UploadAndRate(ctx)
		assert.Nil(t, err)
	})
}