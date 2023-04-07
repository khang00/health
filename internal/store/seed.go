package store

import (
	atlas_migrate "ariga.io/atlas/sql/migrate"
	atlas_schema "ariga.io/atlas/sql/schema"
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/khang00/health/ent"
)

func seedData(ctx context.Context, client *ent.Client) error {
	err := client.User.Create().
		SetUserName("khang").
		SetPassword("123456789").
		Exec(ctx)
	if err != nil {
		return err
	}

	err = client.Meal.CreateBulk(
		client.Meal.Create().
			SetUserID(1).
			SetMealType("morning"),
		client.Meal.Create().
			SetUserID(1).
			SetMealType("lunch"),
		client.Meal.Create().
			SetUserID(1).
			SetMealType("dinner"),
		client.Meal.Create().
			SetUserID(1).
			SetMealType("breakfast"),
	).Exec(ctx)
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
