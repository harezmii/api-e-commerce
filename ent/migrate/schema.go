// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "keywords", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "image", Type: field.TypeString},
		{Name: "url", Type: field.TypeString},
		{Name: "status", Type: field.TypeBool},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "category_children", Type: field.TypeInt, Nullable: true},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "categories_categories_children",
				Columns:    []*schema.Column{CategoriesColumns[10]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CommentsColumns holds the columns for the "comments" table.
	CommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "comment", Type: field.TypeString},
		{Name: "rate", Type: field.TypeFloat64},
		{Name: "ip", Type: field.TypeString},
		{Name: "status", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "product_comments", Type: field.TypeInt, Nullable: true},
		{Name: "user_comments", Type: field.TypeInt, Nullable: true},
	}
	// CommentsTable holds the schema information for the "comments" table.
	CommentsTable = &schema.Table{
		Name:       "comments",
		Columns:    CommentsColumns,
		PrimaryKey: []*schema.Column{CommentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comments_products_comments",
				Columns:    []*schema.Column{CommentsColumns[8]},
				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "comments_users_comments",
				Columns:    []*schema.Column{CommentsColumns[9]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FaqsColumns holds the columns for the "faqs" table.
	FaqsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "question", Type: field.TypeString, Unique: true, Size: 300},
		{Name: "answer", Type: field.TypeString},
		{Name: "status", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// FaqsTable holds the schema information for the "faqs" table.
	FaqsTable = &schema.Table{
		Name:       "faqs",
		Columns:    FaqsColumns,
		PrimaryKey: []*schema.Column{FaqsColumns[0]},
	}
	// ImagesColumns holds the columns for the "images" table.
	ImagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "image", Type: field.TypeString},
		{Name: "url", Type: field.TypeString},
		{Name: "status", Type: field.TypeBool},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "product_images", Type: field.TypeInt, Nullable: true},
	}
	// ImagesTable holds the schema information for the "images" table.
	ImagesTable = &schema.Table{
		Name:       "images",
		Columns:    ImagesColumns,
		PrimaryKey: []*schema.Column{ImagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "images_products_images",
				Columns:    []*schema.Column{ImagesColumns[8]},
				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MessagesColumns holds the columns for the "messages" table.
	MessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
		{Name: "subject", Type: field.TypeString},
		{Name: "message", Type: field.TypeString},
		{Name: "ip", Type: field.TypeString},
		{Name: "status", Type: field.TypeBool},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// MessagesTable holds the schema information for the "messages" table.
	MessagesTable = &schema.Table{
		Name:       "messages",
		Columns:    MessagesColumns,
		PrimaryKey: []*schema.Column{MessagesColumns[0]},
	}
	// ProductsColumns holds the columns for the "products" table.
	ProductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "keywords", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "photos", Type: field.TypeJSON},
		{Name: "urls", Type: field.TypeJSON},
		{Name: "status", Type: field.TypeBool},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "category_products", Type: field.TypeInt, Nullable: true},
		{Name: "user_products", Type: field.TypeInt, Nullable: true},
	}
	// ProductsTable holds the schema information for the "products" table.
	ProductsTable = &schema.Table{
		Name:       "products",
		Columns:    ProductsColumns,
		PrimaryKey: []*schema.Column{ProductsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "products_categories_products",
				Columns:    []*schema.Column{ProductsColumns[10]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "products_users_products",
				Columns:    []*schema.Column{ProductsColumns[11]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProfilesColumns holds the columns for the "profiles" table.
	ProfilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "address", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
		{Name: "image", Type: field.TypeString, Nullable: true},
		{Name: "url", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "user_profile", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// ProfilesTable holds the schema information for the "profiles" table.
	ProfilesTable = &schema.Table{
		Name:       "profiles",
		Columns:    ProfilesColumns,
		PrimaryKey: []*schema.Column{ProfilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "profiles_users_profile",
				Columns:    []*schema.Column{ProfilesColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SettingsColumns holds the columns for the "settings" table.
	SettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "keywords", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "company", Type: field.TypeString},
		{Name: "address", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
		{Name: "fax", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "mail_server_address", Type: field.TypeString},
		{Name: "mail_server_email", Type: field.TypeString},
		{Name: "mail_server_password", Type: field.TypeString},
		{Name: "mail_server_port", Type: field.TypeString},
		{Name: "facebook", Type: field.TypeString},
		{Name: "instagram", Type: field.TypeString},
		{Name: "twitter", Type: field.TypeString},
		{Name: "about", Type: field.TypeString},
		{Name: "contact", Type: field.TypeString},
		{Name: "references", Type: field.TypeString},
		{Name: "status", Type: field.TypeBool},
	}
	// SettingsTable holds the schema information for the "settings" table.
	SettingsTable = &schema.Table{
		Name:       "settings",
		Columns:    SettingsColumns,
		PrimaryKey: []*schema.Column{SettingsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "surname", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "status", Type: field.TypeBool},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CategoriesTable,
		CommentsTable,
		FaqsTable,
		ImagesTable,
		MessagesTable,
		ProductsTable,
		ProfilesTable,
		SettingsTable,
		UsersTable,
	}
)

func init() {
	CategoriesTable.ForeignKeys[0].RefTable = CategoriesTable
	CommentsTable.ForeignKeys[0].RefTable = ProductsTable
	CommentsTable.ForeignKeys[1].RefTable = UsersTable
	ImagesTable.ForeignKeys[0].RefTable = ProductsTable
	ProductsTable.ForeignKeys[0].RefTable = CategoriesTable
	ProductsTable.ForeignKeys[1].RefTable = UsersTable
	ProfilesTable.ForeignKeys[0].RefTable = UsersTable
}
