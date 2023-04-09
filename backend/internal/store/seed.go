package store

import (
	atlas_migrate "ariga.io/atlas/sql/migrate"
	atlas_schema "ariga.io/atlas/sql/schema"
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/khang00/health/backend/ent"
	"time"
)

func seedData(ctx context.Context, client *ent.Client) error {
	err := client.User.Create().
		SetUserName("khang").
		SetPassword("123456789").
		Exec(ctx)
	if err != nil {
		return err
	}

	err = client.Achievement.Create().
		SetUserID(1).
		SetBfpGoal(0.75).
		Exec(ctx)
	if err != nil {
		return err
	}

	now := time.Now()
	err = client.Meal.CreateBulk(
		client.Meal.Create().
			SetUserID(1).
			SetMealType("morning").
			SetCreatedAt(now.Add(-time.Hour*4)),
		client.Meal.Create().
			SetUserID(1).
			SetMealType("lunch").
			SetCreatedAt(now.Add(-time.Hour*8)),
		client.Meal.Create().
			SetUserID(1).
			SetMealType("dinner").
			SetCreatedAt(now.Add(-time.Hour*12)),
		client.Meal.Create().
			SetUserID(1).
			SetMealType("breakfast").
			SetCreatedAt(now.Add(-time.Hour*16)),
	).Exec(ctx)
	if err != nil {
		return err
	}

	err = client.BFPDataPoint.CreateBulk(
		client.BFPDataPoint.Create().
			SetUserID(1).
			SetFatPercentage(0.45).
			SetTotalWeight(60).
			SetCreatedAt(now.Add(-time.Hour*24*1)),
		client.BFPDataPoint.Create().
			SetUserID(1).
			SetFatPercentage(0.55).
			SetTotalWeight(60).
			SetCreatedAt(now.Add(-time.Hour*24*2)),
		client.BFPDataPoint.Create().
			SetUserID(1).
			SetFatPercentage(0.65).
			SetTotalWeight(60).
			SetCreatedAt(now.Add(-time.Hour*24*3)),
		client.BFPDataPoint.Create().
			SetUserID(1).
			SetFatPercentage(0.60).
			SetTotalWeight(60).
			SetCreatedAt(now.Add(-time.Hour*24*4)),
		client.BFPDataPoint.Create().
			SetUserID(1).
			SetFatPercentage(0.75).
			SetTotalWeight(60).
			SetCreatedAt(now.Add(-time.Hour*24*5)),
	).Exec(ctx)
	if err != nil {
		return err
	}

	articleSize := 12
	articleCreates := make([]*ent.ArticleCreate, articleSize)
	for i := range articleCreates {
		articleCreates[i] = client.Article.Create().
			SetTitle("Neque porro quisquam est qui dolorem ipsum").
			SetContent("Neque porro quisquam est qui dolorem ipsum")
	}

	articles, err := client.Article.CreateBulk(
		articleCreates...,
	).Save(ctx)

	if err != nil {
		return err
	}

	articleTagsCreates := make([]*ent.TagCreate, articleSize)
	for i := range articleTagsCreates {
		articleTagsCreates[i] = client.Tag.Create().
			SetTagName("Neque porro").
			SetArticle(articles[i])
	}

	err = client.Tag.CreateBulk(articleTagsCreates...).Exec(ctx)
	if err != nil {
		return err
	}

	err = client.Tag.CreateBulk(articleTagsCreates...).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

type DataSeeder = func(ctx context.Context, client *ent.Client) error

func SeedDataWhenCreateTable(dbdialect string, tableName string, dataSeeder DataSeeder) schema.ApplyHook {
	return func(next schema.Applier) schema.Applier {
		return schema.ApplyFunc(func(ctx context.Context, conn dialect.ExecQuerier, plan *atlas_migrate.Plan) error {
			// Search the schema.Change that triggers the data migration.
			hasC := func() bool {
				for _, c := range plan.Changes {
					m, ok := c.Source.(*atlas_schema.AddTable)
					if ok && m.T.Name == tableName {
						return true
					}
				}
				return false
			}()

			// Change was found, apply the data migration.
			if hasC {
				// At this stage, there are three ways to UPDATE the NULL values to "Unknown".
				// Append a custom migrate.Change to migrate.Plan, execute an SQL statement
				// directly on the dialect.ExecQuerier, or use the generated ent.Client.

				err := next.Apply(ctx, conn, plan)
				if err != nil {
					return err
				}

				// Create a temporary client from the migration connection.
				client := ent.NewClient(
					ent.Driver(sql.NewDriver(dbdialect, sql.Conn{ExecQuerier: conn.(*sql.Tx)})),
				)

				return dataSeeder(ctx, client)
			} else {
				return next.Apply(ctx, conn, plan)
			}
		})
	}
}
