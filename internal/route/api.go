package route

import (
	"github.com/khang00/health/internal/handler"
	"github.com/khang00/health/internal/handler/user"
	"github.com/khang00/health/internal/store"
	"github.com/labstack/echo/v4"
)

func setUpAPIRoutes(g *echo.Group, storeClient *store.Client) {

	userHandler := user.NewUserHandler(storeClient)

	userGroup := g.Group("/user")
	userGroup.GET("/meal", userHandler.GetUserMeal)
	userGroup.GET("/body_fat_percentage", userHandler.GetUserBFPByInterval)

	g.GET("/articles", handler.GetArticles)
}
