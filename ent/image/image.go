// Code generated by entc, DO NOT EDIT.

package image

import (
	"time"
)

const (
	// Label holds the string label denoting the image type in the database.
	Label = "image"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the image in the database.
	Table = "images"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "images"
	// OwnerInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	OwnerInverseTable = "products"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "product_images"
)

// Columns holds all SQL columns for image fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldImage,
	FieldURL,
	FieldStatus,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "images"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"product_images",
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)
