package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User schema
type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().StructTag(`json:"id"`),
		field.String("email").Unique().StructTag(`json:"email"`),
		field.String("name").StructTag(`json:"name"`),
	}
}

func (User) Edges() []ent.Edge {
	return nil
}
