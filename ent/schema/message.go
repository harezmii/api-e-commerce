package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("email"),
		field.String("phone"),
		field.String("subject"),
		field.String("message"),
		field.String("ip"),
		field.Bool("status"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
		field.Time("deleted_at").Optional(),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return nil
}
