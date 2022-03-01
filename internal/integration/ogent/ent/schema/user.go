package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("age"),
		field.Enum("favorite_cat_breed").
			Values("siamese", "bengal", "lion", "tiger", "leopard", "other"),
		field.Enum("favorite_dog_breed").
			Values("Kuro").
			Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pets", Pet.Type),
		edge.To("best_friend", User.Type).
			Unique(),
	}
}
