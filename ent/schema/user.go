package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("key"),
		field.String("id").NotEmpty().Unique(),
		field.String("email").NotEmpty().Unique(),
		field.Int64("email_verified_at").Optional(),
		field.String("password"),
		field.String("signup_methods").Optional(),
		field.String("given_name").Optional(),
		field.String("family_name").Optional(),
		field.String("middle_name").Optional(),
		field.String("nickname").Optional(),
		field.String("gender").Optional(),
		field.String("birthdate").Optional(),
		field.String("phone_number").Optional(),
		field.Int64("phone_number_verified_at").Optional(),
		field.String("picture").Optional(),
		field.String("roles"),
		field.Int64("revoked_timestamp").Optional(),
		field.Bool("is_multi_factor_auth_enabled").Optional(),
		field.Int64("updated_at"),
		field.Int64("created_at"),
		field.String("app_data").Optional(),
		field.Bool("tenant_owner").Default(false).Nillable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

// Annotations for the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "authorizer_users"},
	}
}
