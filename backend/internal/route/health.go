package route

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func setUpHealthRoutes(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	})
}
