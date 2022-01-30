package schema

import "entgo.io/ent"

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return nil
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return nil
}
