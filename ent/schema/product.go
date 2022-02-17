package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("keywords"),
		field.String("description"),
		field.String("image"),
		field.Bool("status"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Optional(),
		field.Time("deleted_at").Optional(),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("images", Image.Type),
		edge.From("owner", Category.Type).Ref("products").Unique(),
		edge.From("owner1", User.Type).Ref("products").Unique(),
		edge.To("comments", Comment.Type),
	}
}
