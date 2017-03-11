package Gameboard

import (
	gf "../Gamefields"
	fo "../Gameobjects"
)

// Main Implemented Interface

type Board interface {
	// Return the sections of the board
	FieldsByRow(int) ([]gf.Field, bool)
	// The BoardObject inside the given gf.Field
	BoardObjectsByField(gf.Field) []fo.FieldObject
	// Add a gf.Field with the specified Borders to the specified row on the Board
	AddField(gf.Field)
	// Add a gf.Field Object
	AddFieldObject(fo.FieldObject)
}
