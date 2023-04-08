package store

import (
	"context"
	"github.com/khang00/health/ent"
	"github.com/khang00/health/ent/article"
)

func (c *Client) GetRecentArticle(ctx context.Context, limit int) ([]*ent.Article, error) {
	return c.client.Article.
		Query().
		WithTags().
		Order(ent.Desc(article.FieldCreatedAt)).
		Limit(limit).
		All(ctx)
}

func (c *Client) GetRecommendationArticle(ctx context.Context, limit int) ([]*ent.Article, error) {
	return c.GetRecentArticle(ctx, limit)
}
