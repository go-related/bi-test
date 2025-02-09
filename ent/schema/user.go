package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"errors"
	"github.com/google/uuid"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("fid"),
		field.Int("age").Positive(),
		field.String("name").Validate(func(s string) error {
			if len(s) == 0 {
				return errors.New("name should not be empty")
			}
			return nil
		}),
		field.String("nickname").
			Unique(),
		field.String("email").
			Optional(),
		field.String("phone").
			Unique(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("group", Group.Type),
		edge.To("friends", User.Type),
	}
}
