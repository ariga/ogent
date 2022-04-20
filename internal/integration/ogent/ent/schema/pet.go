package schema

import (
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
		field.Int("weight").
			Optional(),
		field.Time("birthday").
			Optional(),
		field.Bytes("tag_id").
			Optional(),
		field.Int("height").
			Optional().
			Nillable(),
	}
}

// Edges of the Pet.
func (Pet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("categories", Category.Type).
			Ref("pets").
			Annotations(entoas.Groups("pet:create")),
		edge.From("owner", User.Type).
			Ref("pets").
			Unique().
			Required().
			Annotations(entoas.Groups("pet:create")),
		edge.To("friends", Pet.Type),
	}
}

// Annotations of the Pet.
func (Pet) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entoas.CreateOperation(entoas.OperationGroups("pet:create")),
	}
}
