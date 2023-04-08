package store

import (
	"context"
	"github.com/khang00/health/ent"
	"github.com/khang00/health/ent/article"
)

func (c *client) GetRecentArticle(ctx context.Context, limit int) ([]*ent.Article, error) {
	return c.client.Article.
		Query().
		WithTags().
		Order(ent.Desc(article.FieldCreatedAt)).
		Limit(limit).
		All(ctx)
}

func (c *client) GetRecommendationArticle(ctx context.Context, limit int) ([]*ent.Article, error) {
	return c.GetRecentArticle(ctx, limit)
}
