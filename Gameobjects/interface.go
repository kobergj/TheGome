package Gameobjects

import (
	gf "../Gamefields"
)

type FieldObject interface {
	// A String Identifier for the Object
	AsString() string
	// The gf.Field the object is currently IN
	CurrentField() gf.Field
	// Is the FieldObject still Active
	Alive() bool
	// Move to the given field
	Move(gf.Field)
}
