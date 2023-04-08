package store

import (
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/khang00/health/ent"
	"github.com/khang00/health/ent/user"
	_ "github.com/lib/pq"
	"log"
)

type Client interface {
	GetUserWithAchievement(ctx context.Context, userName string) (*ent.User, error)

	GetMealByInterval(ctx context.Context, userID int, from int64, to int64) ([]*ent.Meal, error)
	GetBFPDataPointByInterval(ctx context.Context, userID int, from int64, to int64) ([]*ent.BFPDataPoint, error)

	GetRecentArticle(ctx context.Context, limit int) ([]*ent.Article, error)
	GetRecommendationArticle(ctx context.Context, limit int) ([]*ent.Article, error)

	Closed()
}

type client struct {
	client *ent.Client
}

func NewStoreClient() Client {
	entClient, err := ent.Open("postgres",
		"sslmode=disable host=localhost port=5432 user=postgres dbname=health password=",
	)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	// Run the auto migration tool.
	err = entClient.Schema.Create(context.Background(),
		schema.WithApplyHook(SeedDataWhenCreateTable(dialect.Postgres, user.Table, seedData)),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return &client{
		client: entClient,
	}
}

func (c *client) Closed() {
	c.client.Close()
}
