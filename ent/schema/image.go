package schema

import "entgo.io/ent"

// Image holds the schema definition for the Image entity.
type Image struct {
	ent.Schema
}

// Fields of the Image.
func (Image) Fields() []ent.Field {
	return nil
}

// Edges of the Image.
func (Image) Edges() []ent.Edge {
	return nil
}
