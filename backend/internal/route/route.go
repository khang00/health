package route

import (
	"github.com/khang00/health/backend/internal/handler"
	"github.com/khang00/health/backend/internal/store"
	"github.com/labstack/echo/v4"
)

func SetUpRoutes(e *echo.Echo, storeClient store.Client) {
	setUpHealthRoutes(e)

	SetUpRequestValidator(e)
	apiGroup := e.Group("/api/v1")
	setUpAPIRoutes(apiGroup, storeClient)
}

func SetUpRequestValidator(e *echo.Echo) {
	e.Validator = handler.NewRequestValidator()
}
