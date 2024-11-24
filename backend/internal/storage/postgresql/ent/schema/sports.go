package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Sports holds the schema definition for the Sports entity.
type Sports struct {
	ent.Schema
}

// Fields of the Sports.
func (Sports) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("name"),
	}
}

// Edges of the Sports.
func (Sports) Edges() []ent.Edge {
	return nil
}
