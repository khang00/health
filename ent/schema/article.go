package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

const DefaultArticleThumbnail = "https://media.istockphoto.com/id/1316145932/photo/table-top-view-of-spicy-food.jpg?b=1&s=170667a&w=0&k=20&c=P3jIQq8gVqlXjd4kP2OrXYyzqEXSWCwwYtwrd81psDY="

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.String("thumbnail_img_url").Default(DefaultArticleThumbnail),
		field.String("title").NotEmpty(),
		field.String("content"),
		field.Time("created_at").Default(time.Now()),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tags", Tag.Type),
	}
}
