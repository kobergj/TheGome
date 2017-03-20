package Stringview

import (
	gb "../Gameboard"
)

// Main Implemented Interface

type View interface {
	// Return the gb.Board as a string
	CompleteBoard(gb.Board) string
}
