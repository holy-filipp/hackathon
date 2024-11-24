package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Events holds the schema definition for the Events entity.
type Events struct {
	ent.Schema
}

// Fields of the Events.
func (Events) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("name"),
		field.String("sex_age"),
		field.String("discipline"),
		field.Time("time_start"),
		field.Time("time_end"),
		field.String("venue"),
		field.Int("participants"),
		field.String("sport_id"),
	}
}

// Edges of the Events.
func (Events) Edges() []ent.Edge {
	return nil
}
