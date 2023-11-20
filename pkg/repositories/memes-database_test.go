package repositories

import (
	"insat/devops/core/config"
	"insat/devops/core/store"
	"insat/devops/core/store/services"
	"insat/devops/pkg/rating"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func setup() (*echo.Echo, store.APIServer, func()) {
	e := echo.New()
	config.InjectRenderer(e, config.WithTemplatesDir("../../templates"))
	c := jaegertracing.New(e, nil)

	// Initialize Services and Repositories
	db := NewMemesDatabase()

	// rater := rating.NewAIBasedRatingService()
	fake := rating.NewFakeRatingService()
	memesService := services.NewMemesService(db, fake)

	// Initialize API Handlers
	var memezHandler store.APIServer = services.NewMemesHandler(memesService)

	return e, memezHandler, func() {
		defer c.Close()
	}
}

func TestCreateMemesDatabase(t *testing.T) {
	assert.NotPanics(t, func () {
		m := NewMemesDatabase()
		assert.NotNil(t, m)
	})
}

func TestSaveInMemesDatabase(t *testing.T) {
	assert.NotPanics(t, func () {
		e, svr, clean := setup()
		defer clean()
		assert.NotNil(t, svr)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
    	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    	rec := httptest.NewRecorder()
    	ctx := e.NewContext(req, rec)

		m := NewMemesDatabase()
		assert.NotNil(t, m)
		
		err := m.Save(ctx, nil)
		assert.NoError(t, err)
	})
}