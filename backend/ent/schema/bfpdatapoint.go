package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// BFPDataPoint holds the schema definition for the BFPDataPoint entity.
type BFPDataPoint struct {
	ent.Schema
}

// Fields of the BFPDataPoint.
func (BFPDataPoint) Fields() []ent.Field {
	return []ent.Field{
		field.Float("fat_percentage").Range(0, 1),
		field.Int("total_weight").Positive(),
		field.Time("created_at").Default(time.Now()),
	}
}

// Edges of the BFPDataPoint.
func (BFPDataPoint) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("BFPs").
			Unique(),
	}
}
