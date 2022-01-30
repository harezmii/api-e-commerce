package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Faq holds the schema definition for the Faq entity.
type Faq struct {
	ent.Schema
}

// Fields of the Faq.
func (Faq) Fields() []ent.Field {
	return []ent.Field{
		field.String("question").MaxLen(300).Unique(),
		field.String("answer"),
		field.Bool("status").Default(false),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Optional(),
		field.Time("deleted_at").Optional(),
	}
}

// Edges of the Faq.
func (Faq) Edges() []ent.Edge {
	return nil
}
