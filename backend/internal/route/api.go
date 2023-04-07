package route

import (
	"github.com/khang00/health/internal/handler"
	"github.com/labstack/echo/v4"
)

func setUpAPIRoutes(g *echo.Group) {
	user := g.Group("/user")
	user.GET("/user/meal", handler.GetUserMeal)

	user.GET("/user/meal", handler.GetUserBFPByInterval)

	g.GET("/articles", handler.GetArticles)
}
