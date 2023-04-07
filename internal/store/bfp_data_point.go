package store

import (
	"context"
	"github.com/khang00/health/ent"
	"github.com/khang00/health/ent/bfpdatapoint"
	"github.com/khang00/health/ent/user"
	"time"
)

func (c *Client) GetBFPDataPointByInterval(ctx context.Context, userID int, from int64, to int64) ([]*ent.BFPDataPoint, error) {
	return c.client.BFPDataPoint.
		Query().
		Where(
			bfpdatapoint.And(
				bfpdatapoint.HasUserWith(user.ID(userID)),
				bfpdatapoint.CreatedAtGTE(time.Unix(from, 0)),
				bfpdatapoint.CreatedAtLTE(time.Unix(to, 0)),
			),
		).
		All(ctx)
}
