package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

const DefaultMealImage = "https://upload.wikimedia.org/wikipedia/commons/thumb/0/03/Japanese_SilkyTofu_%28Kinugoshi_Tofu%29.JPG/440px-Japanese_SilkyTofu_%28Kinugoshi_Tofu%29.JPG"

// Meal holds the schema definition for the Meal entity.
type Meal struct {
	ent.Schema
}

// Fields of the Meal.
func (Meal) Fields() []ent.Field {
	return []ent.Field{
		field.String("meal_type").NotEmpty(),
		field.String("image_url").Default(DefaultMealImage),
		field.Time("created_at").Default(time.Now()),
	}
}

// Edges of the Meal.
func (Meal) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("meals").
			Unique(),
	}
}
