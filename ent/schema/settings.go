package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Settings holds the schema definition for the Settings entity.
type Settings struct {
	ent.Schema
}

// Fields of the Settings.
func (Settings) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("keywords"),
		field.String("description"),
		field.String("company"),
		field.String("address"),
		field.String("phone"),
		field.String("fax"),
		field.String("email"),
		field.String("mailServerAddress"),
		field.String("mailServerEmail"),
		field.String("mailServerPassword"),
		field.String("mailServerPort"),
		field.String("facebook"),
		field.String("Instagram"),
		field.String("twitter"),
		field.String("about"),
		field.String("contact"),
		field.String("references"),
		field.Bool("status"),
	}
}

// Edges of the Settings.
func (Settings) Edges() []ent.Edge {
	return nil
}
