package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Image holds the schema definition for the Image entity.
type Image struct {
	ent.Schema
}

// Fields of the Image.
func (Image) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("image"),
		field.Bool("status"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Optional(),
		field.Time("deleted_at").Optional(),
	}
}

// Edges of the Image.
func (Image) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Product.Type).Ref("images").Unique(),
	}
}
