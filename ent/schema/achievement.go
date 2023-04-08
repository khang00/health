package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Achievement holds the schema definition for the Achievement entity.
type Achievement struct {
	ent.Schema
}

const DefaultAchievementThumbnail = "https://images.unsplash.com/photo-1546069901-ba9599a7e63c?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxleHBsb3JlLWZlZWR8MXx8fGVufDB8fHx8&w=1000&q=80"

// Fields of the Achievement.
func (Achievement) Fields() []ent.Field {
	return []ent.Field{
		field.String("status").Default("current"),
		field.String("thumbnail_img_url").Default(DefaultAchievementThumbnail),
		field.Float("bfp_goal").Range(0, 1),
		field.Time("created_at").Default(time.Now()),
	}
}

// Edges of the Achievement.
func (Achievement) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("achievements").
			Unique(),
	}
}
