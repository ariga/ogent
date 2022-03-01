// Code generated by entc, DO NOT EDIT.

package user

import (
	"fmt"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldFavoriteCatBreed holds the string denoting the favorite_cat_breed field in the database.
	FieldFavoriteCatBreed = "favorite_cat_breed"
	// FieldFavoriteDogBreed holds the string denoting the favorite_dog_breed field in the database.
	FieldFavoriteDogBreed = "favorite_dog_breed"
	// EdgePets holds the string denoting the pets edge name in mutations.
	EdgePets = "pets"
	// EdgeBestFriend holds the string denoting the best_friend edge name in mutations.
	EdgeBestFriend = "best_friend"
	// Table holds the table name of the user in the database.
	Table = "users"
	// PetsTable is the table that holds the pets relation/edge.
	PetsTable = "pets"
	// PetsInverseTable is the table name for the Pet entity.
	// It exists in this package in order to avoid circular dependency with the "pet" package.
	PetsInverseTable = "pets"
	// PetsColumn is the table column denoting the pets relation/edge.
	PetsColumn = "user_pets"
	// BestFriendTable is the table that holds the best_friend relation/edge.
	BestFriendTable = "users"
	// BestFriendColumn is the table column denoting the best_friend relation/edge.
	BestFriendColumn = "user_best_friend"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldAge,
	FieldFavoriteCatBreed,
	FieldFavoriteDogBreed,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "users"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_best_friend",
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

// FavoriteCatBreed defines the type for the "favorite_cat_breed" enum field.
type FavoriteCatBreed string

// FavoriteCatBreed values.
const (
	FavoriteCatBreedSiamese FavoriteCatBreed = "siamese"
	FavoriteCatBreedBengal  FavoriteCatBreed = "bengal"
	FavoriteCatBreedLion    FavoriteCatBreed = "lion"
	FavoriteCatBreedTiger   FavoriteCatBreed = "tiger"
	FavoriteCatBreedLeopard FavoriteCatBreed = "leopard"
	FavoriteCatBreedOther   FavoriteCatBreed = "other"
)

func (fcb FavoriteCatBreed) String() string {
	return string(fcb)
}

// FavoriteCatBreedValidator is a validator for the "favorite_cat_breed" field enum values. It is called by the builders before save.
func FavoriteCatBreedValidator(fcb FavoriteCatBreed) error {
	switch fcb {
	case FavoriteCatBreedSiamese, FavoriteCatBreedBengal, FavoriteCatBreedLion, FavoriteCatBreedTiger, FavoriteCatBreedLeopard, FavoriteCatBreedOther:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for favorite_cat_breed field: %q", fcb)
	}
}

// FavoriteDogBreed defines the type for the "favorite_dog_breed" enum field.
type FavoriteDogBreed string

// FavoriteDogBreed values.
const (
	FavoriteDogBreedKuro FavoriteDogBreed = "Kuro"
)

func (fdb FavoriteDogBreed) String() string {
	return string(fdb)
}

// FavoriteDogBreedValidator is a validator for the "favorite_dog_breed" field enum values. It is called by the builders before save.
func FavoriteDogBreedValidator(fdb FavoriteDogBreed) error {
	switch fdb {
	case FavoriteDogBreedKuro:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for favorite_dog_breed field: %q", fdb)
	}
}
