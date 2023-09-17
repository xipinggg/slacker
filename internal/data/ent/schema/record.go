package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Record holds the schema definition for the Record entity.
type Record struct {
	ent.Schema
}

// Fields of the Record.
func (Record) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.String("creator_id").Immutable(),
		field.String("type").Immutable(),
		field.Time("begin_time").Default(time.Now).Immutable(),
		field.Time("end_time").Nillable().Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Nillable().Default(time.Now),
	}
}

// Edges of the Record.
func (Record) Edges() []ent.Edge {
	return nil
}

func (Record) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
	}
}
