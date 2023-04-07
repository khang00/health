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

type Client struct {
	client *ent.Client
}

func NewStoreClient() *Client {
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

	return &Client{
		client: entClient,
	}
}

func (c *Client) Closed() {
	c.client.Close()
}
