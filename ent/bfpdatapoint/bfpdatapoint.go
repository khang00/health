// Code generated by ent, DO NOT EDIT.

package bfpdatapoint

import (
	"time"
)

const (
	// Label holds the string label denoting the bfpdatapoint type in the database.
	Label = "bfp_data_point"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFatPercentage holds the string denoting the fat_percentage field in the database.
	FieldFatPercentage = "fat_percentage"
	// FieldTotalWeight holds the string denoting the total_weight field in the database.
	FieldTotalWeight = "total_weight"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the bfpdatapoint in the database.
	Table = "bfp_data_points"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "bfp_data_points"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_bf_ps"
)

// Columns holds all SQL columns for bfpdatapoint fields.
var Columns = []string{
	FieldID,
	FieldFatPercentage,
	FieldTotalWeight,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "bfp_data_points"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_bf_ps",
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
	// FatPercentageValidator is a validator for the "fat_percentage" field. It is called by the builders before save.
	FatPercentageValidator func(float64) error
	// TotalWeightValidator is a validator for the "total_weight" field. It is called by the builders before save.
	TotalWeightValidator func(int) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt time.Time
)