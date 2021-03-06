// Code generated by entc, DO NOT EDIT.

package product

import (
	"time"
)

const (
	// Label holds the string label denoting the product type in the database.
	Label = "product"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldKeywords holds the string denoting the keywords field in the database.
	FieldKeywords = "keywords"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldPhotos holds the string denoting the photos field in the database.
	FieldPhotos = "photos"
	// FieldUrls holds the string denoting the urls field in the database.
	FieldUrls = "urls"
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
	// EdgeOwner1 holds the string denoting the owner1 edge name in mutations.
	EdgeOwner1 = "owner1"
	// EdgeComments holds the string denoting the comments edge name in mutations.
	EdgeComments = "comments"
	// Table holds the table name of the product in the database.
	Table = "products"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "products"
	// OwnerInverseTable is the table name for the Category entity.
	// It exists in this package in order to avoid circular dependency with the "category" package.
	OwnerInverseTable = "categories"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "category_products"
	// Owner1Table is the table that holds the owner1 relation/edge.
	Owner1Table = "products"
	// Owner1InverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	Owner1InverseTable = "users"
	// Owner1Column is the table column denoting the owner1 relation/edge.
	Owner1Column = "user_products"
	// CommentsTable is the table that holds the comments relation/edge.
	CommentsTable = "comments"
	// CommentsInverseTable is the table name for the Comment entity.
	// It exists in this package in order to avoid circular dependency with the "comment" package.
	CommentsInverseTable = "comments"
	// CommentsColumn is the table column denoting the comments relation/edge.
	CommentsColumn = "product_comments"
)

// Columns holds all SQL columns for product fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldKeywords,
	FieldDescription,
	FieldPhotos,
	FieldUrls,
	FieldStatus,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "products"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"category_products",
	"user_products",
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
