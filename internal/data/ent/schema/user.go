package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the Users.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("mysql_id").Unique().Immutable().SchemaType(map[string]string{
			"mysql": "INT AUTO_INCREMENT",
		}),
		field.String("id").Unique().Immutable().MinLen(2).MaxLen(64).DefaultFunc(uuid.NewString),
		field.String("name").Unique().MinLen(2).MaxLen(64),
		field.Time("created_at").Immutable().Optional().Default(time.Now),
		field.Time("updated_at").Nillable().Optional().Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("mysql_id", "id").Unique(),
	}
}
