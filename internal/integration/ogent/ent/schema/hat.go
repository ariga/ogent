package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Hat holds the schema definition for the Hat entity.
type Hat struct {
	ent.Schema
}

// Fields of the Hat.
func (Hat) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Enum("type").
			Values("dad", "trucker", "snapback").
			Immutable(),
	}
}

// Edges of the Hat.
func (Hat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("wearer", User.Type).
			Ref("favorite_hat").
			Unique().
			Immutable(),
	}
}
