package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.String("comment"),
		field.Float("rate"),
		field.String("ip"),
		field.Bool("status").Default(false),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Optional(),
		field.Time("deleted_at").Optional(),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Product.Type).Ref("comments").Unique(),
		edge.From("own", User.Type).Ref("comments").Unique(),
	}
}
