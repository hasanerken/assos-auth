// Code generated by ent, DO NOT EDIT.

package ent

import (
	"assos/ent/user"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Key holds the value of the "key" field.
	Key string `json:"key,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// EmailVerifiedAt holds the value of the "email_verified_at" field.
	EmailVerifiedAt int64 `json:"email_verified_at,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// SignupMethods holds the value of the "signup_methods" field.
	SignupMethods string `json:"signup_methods,omitempty"`
	// GivenName holds the value of the "given_name" field.
	GivenName string `json:"given_name,omitempty"`
	// FamilyName holds the value of the "family_name" field.
	FamilyName string `json:"family_name,omitempty"`
	// MiddleName holds the value of the "middle_name" field.
	MiddleName string `json:"middle_name,omitempty"`
	// Nickname holds the value of the "nickname" field.
	Nickname string `json:"nickname,omitempty"`
	// Gender holds the value of the "gender" field.
	Gender string `json:"gender,omitempty"`
	// Birthdate holds the value of the "birthdate" field.
	Birthdate string `json:"birthdate,omitempty"`
	// PhoneNumber holds the value of the "phone_number" field.
	PhoneNumber string `json:"phone_number,omitempty"`
	// PhoneNumberVerifiedAt holds the value of the "phone_number_verified_at" field.
	PhoneNumberVerifiedAt int64 `json:"phone_number_verified_at,omitempty"`
	// Picture holds the value of the "picture" field.
	Picture string `json:"picture,omitempty"`
	// Roles holds the value of the "roles" field.
	Roles string `json:"roles,omitempty"`
	// RevokedTimestamp holds the value of the "revoked_timestamp" field.
	RevokedTimestamp int64 `json:"revoked_timestamp,omitempty"`
	// IsMultiFactorAuthEnabled holds the value of the "is_multi_factor_auth_enabled" field.
	IsMultiFactorAuthEnabled bool `json:"is_multi_factor_auth_enabled,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt int64 `json:"created_at,omitempty"`
	// AppData holds the value of the "app_data" field.
	AppData string `json:"app_data,omitempty"`
	// TenantOwner holds the value of the "tenant_owner" field.
	TenantOwner  *bool `json:"tenant_owner,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldIsMultiFactorAuthEnabled, user.FieldTenantOwner:
			values[i] = new(sql.NullBool)
		case user.FieldEmailVerifiedAt, user.FieldPhoneNumberVerifiedAt, user.FieldRevokedTimestamp, user.FieldUpdatedAt, user.FieldCreatedAt:
			values[i] = new(sql.NullInt64)
		case user.FieldID, user.FieldKey, user.FieldEmail, user.FieldPassword, user.FieldSignupMethods, user.FieldGivenName, user.FieldFamilyName, user.FieldMiddleName, user.FieldNickname, user.FieldGender, user.FieldBirthdate, user.FieldPhoneNumber, user.FieldPicture, user.FieldRoles, user.FieldAppData:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				u.ID = value.String
			}
		case user.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				u.Key = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldEmailVerifiedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field email_verified_at", values[i])
			} else if value.Valid {
				u.EmailVerifiedAt = value.Int64
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldSignupMethods:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field signup_methods", values[i])
			} else if value.Valid {
				u.SignupMethods = value.String
			}
		case user.FieldGivenName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field given_name", values[i])
			} else if value.Valid {
				u.GivenName = value.String
			}
		case user.FieldFamilyName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field family_name", values[i])
			} else if value.Valid {
				u.FamilyName = value.String
			}
		case user.FieldMiddleName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field middle_name", values[i])
			} else if value.Valid {
				u.MiddleName = value.String
			}
		case user.FieldNickname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nickname", values[i])
			} else if value.Valid {
				u.Nickname = value.String
			}
		case user.FieldGender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value.Valid {
				u.Gender = value.String
			}
		case user.FieldBirthdate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field birthdate", values[i])
			} else if value.Valid {
				u.Birthdate = value.String
			}
		case user.FieldPhoneNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone_number", values[i])
			} else if value.Valid {
				u.PhoneNumber = value.String
			}
		case user.FieldPhoneNumberVerifiedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field phone_number_verified_at", values[i])
			} else if value.Valid {
				u.PhoneNumberVerifiedAt = value.Int64
			}
		case user.FieldPicture:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field picture", values[i])
			} else if value.Valid {
				u.Picture = value.String
			}
		case user.FieldRoles:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field roles", values[i])
			} else if value.Valid {
				u.Roles = value.String
			}
		case user.FieldRevokedTimestamp:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field revoked_timestamp", values[i])
			} else if value.Valid {
				u.RevokedTimestamp = value.Int64
			}
		case user.FieldIsMultiFactorAuthEnabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_multi_factor_auth_enabled", values[i])
			} else if value.Valid {
				u.IsMultiFactorAuthEnabled = value.Bool
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Int64
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Int64
			}
		case user.FieldAppData:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field app_data", values[i])
			} else if value.Valid {
				u.AppData = value.String
			}
		case user.FieldTenantOwner:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field tenant_owner", values[i])
			} else if value.Valid {
				u.TenantOwner = new(bool)
				*u.TenantOwner = value.Bool
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("key=")
	builder.WriteString(u.Key)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("email_verified_at=")
	builder.WriteString(fmt.Sprintf("%v", u.EmailVerifiedAt))
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(u.Password)
	builder.WriteString(", ")
	builder.WriteString("signup_methods=")
	builder.WriteString(u.SignupMethods)
	builder.WriteString(", ")
	builder.WriteString("given_name=")
	builder.WriteString(u.GivenName)
	builder.WriteString(", ")
	builder.WriteString("family_name=")
	builder.WriteString(u.FamilyName)
	builder.WriteString(", ")
	builder.WriteString("middle_name=")
	builder.WriteString(u.MiddleName)
	builder.WriteString(", ")
	builder.WriteString("nickname=")
	builder.WriteString(u.Nickname)
	builder.WriteString(", ")
	builder.WriteString("gender=")
	builder.WriteString(u.Gender)
	builder.WriteString(", ")
	builder.WriteString("birthdate=")
	builder.WriteString(u.Birthdate)
	builder.WriteString(", ")
	builder.WriteString("phone_number=")
	builder.WriteString(u.PhoneNumber)
	builder.WriteString(", ")
	builder.WriteString("phone_number_verified_at=")
	builder.WriteString(fmt.Sprintf("%v", u.PhoneNumberVerifiedAt))
	builder.WriteString(", ")
	builder.WriteString("picture=")
	builder.WriteString(u.Picture)
	builder.WriteString(", ")
	builder.WriteString("roles=")
	builder.WriteString(u.Roles)
	builder.WriteString(", ")
	builder.WriteString("revoked_timestamp=")
	builder.WriteString(fmt.Sprintf("%v", u.RevokedTimestamp))
	builder.WriteString(", ")
	builder.WriteString("is_multi_factor_auth_enabled=")
	builder.WriteString(fmt.Sprintf("%v", u.IsMultiFactorAuthEnabled))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", u.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", u.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("app_data=")
	builder.WriteString(u.AppData)
	builder.WriteString(", ")
	if v := u.TenantOwner; v != nil {
		builder.WriteString("tenant_owner=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
