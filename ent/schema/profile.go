package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Profile holds the schema definition for the Profile entity.
type Profile struct {
	ent.Schema
}

// Fields of the Profile.
func (Profile) Fields() []ent.Field {
	return []ent.Field{
		field.String("address"),
		field.String("phone"),
		field.String("image").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
		field.Time("deleted_at").Optional(),
	}
}

// Edges of the Profile.
func (Profile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("profile").Unique(),
	}
}
