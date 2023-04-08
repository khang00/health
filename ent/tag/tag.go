// Code generated by ent, DO NOT EDIT.

package tag

import (
	"time"
)

const (
	// Label holds the string label denoting the tag type in the database.
	Label = "tag"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTagName holds the string denoting the tag_name field in the database.
	FieldTagName = "tag_name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeArticle holds the string denoting the article edge name in mutations.
	EdgeArticle = "article"
	// Table holds the table name of the tag in the database.
	Table = "tags"
	// ArticleTable is the table that holds the article relation/edge.
	ArticleTable = "tags"
	// ArticleInverseTable is the table name for the Article entity.
	// It exists in this package in order to avoid circular dependency with the "article" package.
	ArticleInverseTable = "articles"
	// ArticleColumn is the table column denoting the article relation/edge.
	ArticleColumn = "article_tags"
)

// Columns holds all SQL columns for tag fields.
var Columns = []string{
	FieldID,
	FieldTagName,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "tags"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"article_tags",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// TagNameValidator is a validator for the "tag_name" field. It is called by the builders before save.
	TagNameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt time.Time
)