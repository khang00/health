package store

import (
	"context"
	"github.com/khang00/health/ent"
	"github.com/khang00/health/ent/meal"
	"github.com/khang00/health/ent/user"
	"time"
)

func (c *Client) GetMealByInterval(ctx context.Context, userID int, from int64, to int64) ([]*ent.Meal, error) {
	return c.client.Meal.
		Query().
		Where(
			meal.And(
				meal.HasUserWith(user.ID(userID)),
				meal.CreatedAtIn(
					time.Unix(from, 0),
					time.Unix(to, 0)))).
		All(ctx)
}
