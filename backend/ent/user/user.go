// Code generated by ent, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserName holds the string denoting the user_name field in the database.
	FieldUserName = "user_name"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// EdgeMeals holds the string denoting the meals edge name in mutations.
	EdgeMeals = "meals"
	// EdgeBFPs holds the string denoting the bfps edge name in mutations.
	EdgeBFPs = "BFPs"
	// EdgeAchievements holds the string denoting the achievements edge name in mutations.
	EdgeAchievements = "achievements"
	// Table holds the table name of the user in the database.
	Table = "users"
	// MealsTable is the table that holds the meals relation/edge.
	MealsTable = "meals"
	// MealsInverseTable is the table name for the Meal entity.
	// It exists in this package in order to avoid circular dependency with the "meal" package.
	MealsInverseTable = "meals"
	// MealsColumn is the table column denoting the meals relation/edge.
	MealsColumn = "user_meals"
	// BFPsTable is the table that holds the BFPs relation/edge.
	BFPsTable = "bfp_data_points"
	// BFPsInverseTable is the table name for the BFPDataPoint entity.
	// It exists in this package in order to avoid circular dependency with the "bfpdatapoint" package.
	BFPsInverseTable = "bfp_data_points"
	// BFPsColumn is the table column denoting the BFPs relation/edge.
	BFPsColumn = "user_bf_ps"
	// AchievementsTable is the table that holds the achievements relation/edge.
	AchievementsTable = "achievements"
	// AchievementsInverseTable is the table name for the Achievement entity.
	// It exists in this package in order to avoid circular dependency with the "achievement" package.
	AchievementsInverseTable = "achievements"
	// AchievementsColumn is the table column denoting the achievements relation/edge.
	AchievementsColumn = "user_achievements"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUserName,
	FieldPassword,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
