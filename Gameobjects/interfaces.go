package Gameobjects

import (
	gf "../Gamefields"
)

type FieldObject interface {
	// A String Identifier for the Object
	AsString() string
	// The gf.Field the object is currently IN
	CurrentField() gf.Field
	// Is the gf.Field Object still Active
	Alive() bool
}
