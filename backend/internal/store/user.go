package store

import (
	"context"
	"github.com/khang00/health/backend/ent"
	"github.com/khang00/health/backend/ent/user"
)

func (c *client) GetUserWithAchievement(ctx context.Context, userName string) (*ent.User, error) {
	return c.client.User.
		Query().
		Where(
			user.UserNameEQ(userName),
		).
		WithAchievements().
		Only(ctx)
}
