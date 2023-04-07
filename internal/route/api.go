package route

import (
	"github.com/khang00/health/internal/handler"
	"github.com/khang00/health/internal/store"
	"github.com/labstack/echo/v4"
)

func setUpAPIRoutes(g *echo.Group, storeClient *store.Client) {
	user := g.Group("/user")
	userHandler := handler.NewUserHandler(storeClient)
	user.GET("/user/meal", userHandler.GetUserMeal)
	user.GET("/user/body_fat_percentage`", userHandler.GetUserBFPByInterval)

	g.GET("/articles", handler.GetArticles)
}
