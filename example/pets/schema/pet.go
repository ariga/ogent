package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Pet holds the schema definition for the Pet entity.
type Pet struct {
	ent.Schema
}

// Fields of the Pet.
func (Pet) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("weight"),
		field.Time("birthday"),
	}
}

// Edges of the Pet.
func (Pet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("categories", Category.Type).
			Ref("pets"),
		edge.From("owner", User.Type).
			Ref("pets").
			Unique(),
		edge.To("friends", Pet.Type),
	}
}
