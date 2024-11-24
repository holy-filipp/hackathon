// Code generated by ent, DO NOT EDIT.

package ent

import (
	"calendar-backend/internal/storage/postgresql/ent/schema"
	"calendar-backend/internal/storage/postgresql/ent/sports"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	sportsFields := schema.Sports{}.Fields()
	_ = sportsFields
	// sportsDescID is the schema descriptor for id field.
	sportsDescID := sportsFields[0].Descriptor()
	// sports.DefaultID holds the default value on creation for the id field.
	sports.DefaultID = sportsDescID.Default.(func() uuid.UUID)
}
