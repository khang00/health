package store

import (
	"context"
	"github.com/khang00/health/backend/ent"
	"github.com/khang00/health/backend/ent/meal"
	"github.com/khang00/health/backend/ent/user"
	"time"
)

func (c *client) GetMealByInterval(ctx context.Context, userID int, from int64, to int64) ([]*ent.Meal, error) {
	return c.client.Meal.
		Query().
		Where(
			meal.And(
				meal.HasUserWith(user.ID(userID)),
				meal.CreatedAtGTE(time.Unix(from, 0)),
				meal.CreatedAtLTE(time.Unix(to, 0)),
			),
		).Order(ent.Desc(meal.FieldCreatedAt)).
		All(ctx)
}
