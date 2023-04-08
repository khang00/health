package route

import (
	"github.com/khang00/health/internal/handler/article"
	"github.com/khang00/health/internal/handler/user"
	"github.com/khang00/health/internal/store"
	"github.com/labstack/echo/v4"
)

func setUpAPIRoutes(g *echo.Group, storeClient *store.Client) {

	userHandler := user.NewUserHandler(storeClient)
	userGroup := g.Group("/user")
	userGroup.GET("", userHandler.GetUser)
	userGroup.GET("/meal", userHandler.GetUserMeal)
	userGroup.GET("/body_fat_percentage", userHandler.GetUserBFPByInterval)

	articleHandler := article.NewArticleHandler(storeClient)
	articleGroup := g.Group("/article")
	articleGroup.GET("", articleHandler.GetArticle)
}
