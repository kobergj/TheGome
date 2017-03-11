package Gamefields

import (
	c "../Constants"
)

// Main Implemented Interface

type Field interface {
	// c.Borders
	LeftBorder() c.Border
	TopBorder() c.Border
	// Set Coordinates
	SetCoordinates(int, int)
}
