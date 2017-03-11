package Stringview

import (
	// c "../Constants"
	gb "../Gameboard"
	// gf "../Gamefields"
	// fo "../Gameobjects"
)

// Main Implemented Interface

type View interface {
	// Return the gb.Board as a string
	CompleteBoard(gb.Board) string
}
