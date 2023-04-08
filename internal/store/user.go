package store

import (
	"context"
	"github.com/khang00/health/ent"
	"github.com/khang00/health/ent/user"
)

func (c *Client) GetUserWithAchievement(ctx context.Context, userName string) (*ent.User, error) {
	return c.client.User.
		Query().
		Where(
			user.UserNameEQ(userName),
		).
		WithAchievements().
		Only(ctx)
}
