// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSurname holds the string denoting the surname field in the database.
	FieldSurname = "surname"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeProfiles holds the string denoting the profiles edge name in mutations.
	EdgeProfiles = "profiles"
	// Table holds the table name of the user in the database.
	Table = "users"
	// ProfilesTable is the table that holds the profiles relation/edge.
	ProfilesTable = "profiles"
	// ProfilesInverseTable is the table name for the Profile entity.
	// It exists in this package in order to avoid circular dependency with the "profile" package.
	ProfilesInverseTable = "profiles"
	// ProfilesColumn is the table column denoting the profiles relation/edge.
	ProfilesColumn = "user_profiles"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldSurname,
	FieldPassword,
	FieldEmail,
	FieldStatus,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
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

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)