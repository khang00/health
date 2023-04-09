package article

import (
	"github.com/khang00/health/backend/ent"
	"github.com/khang00/health/backend/internal/store"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler interface {
	GetArticle(c echo.Context) error
}

type handler struct {
	store store.Client
}

func NewArticleHandler(store store.Client) Handler {
	return &handler{
		store: store,
	}
}

type ArticleQueryType = string

const (
	ArticleRecommendation ArticleQueryType = "ARTICLE_RECOMMENDATION"
	ArticleRecent         ArticleQueryType = "ARTICLE_RECENT"
)

type GetArticleReq struct {
	QueryType ArticleQueryType `json:"query_type" query:"query_type" validate:"required"`
	Limit     int              `json:"limit"`
}

type GetArticleResponse struct {
	Articles []*ent.Article `json:"articles"`
}

func (h *handler) GetArticle(c echo.Context) error {
	getArticleReq := GetArticleReq{}

	if err := c.Bind(&getArticleReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&getArticleReq); err != nil {
		return err
	}
	reqCtx := c.Request().Context()

	articles := make([]*ent.Article, 0)
	var queryError error
	switch getArticleReq.QueryType {
	case ArticleRecent:
		articles, queryError = h.store.GetRecentArticle(reqCtx, getArticleReq.Limit)
		if queryError != nil {
			return queryError
		}
	case ArticleRecommendation:
		articles, queryError = h.store.GetRecommendationArticle(reqCtx, getArticleReq.Limit)
		if queryError != nil {
			return queryError
		}
	}

	return c.JSON(http.StatusOK, GetArticleResponse{
		Articles: articles,
	})
}
