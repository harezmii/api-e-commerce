package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("keywords"),
		field.String("description"),
		field.String("image"),
		field.String("url"),
		field.Bool("status"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Optional(),
		field.Time("deleted_at").Optional(),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Category.Type).
			From("parent").
			Unique(),
		edge.To("products", Product.Type),
	}
}
