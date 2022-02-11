package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.String("surname"),
		field.String("password"),
		field.String("email").Unique(),
		field.Bool("status"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
		field.Time("deleted_at").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profile", Profile.Type).Unique(), // Done
		edge.To("comments", Comment.Type),         //Done
		edge.To("products", Product.Type),
	}
}
