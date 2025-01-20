package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"errors"
	"github.com/google/uuid"
	"time"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("gid"),
		field.String("name").Validate(func(s string) error {
			if len(s) == 0 {
				return errors.New("name should not be empty")
			}
			return nil
		}),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.To("users", User.Type),
	}
}
