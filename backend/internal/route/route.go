package route

import (
	"github.com/khang00/health/internal/handler"
	"github.com/labstack/echo/v4"
)

func SetUpRoutes(e *echo.Echo) {
	SetUpRequestValidator(e)
	setUpHealthRoutes(e)

	apiGroup := e.Group("/api/v1")
	setUpAPIRoutes(apiGroup)
}

func SetUpRequestValidator(e *echo.Echo) {
	e.Validator = handler.NewRequestValidator()
}
