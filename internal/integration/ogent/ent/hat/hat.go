// Code generated by ent, DO NOT EDIT.

package hat

import (
	"fmt"
)

const (
	// Label holds the string label denoting the hat type in the database.
	Label = "hat"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// EdgeWearer holds the string denoting the wearer edge name in mutations.
	EdgeWearer = "wearer"
	// Table holds the table name of the hat in the database.
	Table = "hats"
	// WearerTable is the table that holds the wearer relation/edge.
	WearerTable = "hats"
	// WearerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	WearerInverseTable = "users"
	// WearerColumn is the table column denoting the wearer relation/edge.
	WearerColumn = "user_favorite_hat"
)

// Columns holds all SQL columns for hat fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldType,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "hats"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_favorite_hat",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeDad      Type = "dad"
	TypeTrucker  Type = "trucker"
	TypeSnapback Type = "snapback"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeDad, TypeTrucker, TypeSnapback:
		return nil
	default:
		return fmt.Errorf("hat: invalid enum value for type field: %q", _type)
	}
}
