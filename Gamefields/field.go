package Gamefields

import (
	c "../Constants"
)

func NewField(left, top c.Border) Field {
	return &field{
		borders: [2]c.Border{left, top},
	}
}

type field struct {
	borders [2]c.Border
	cords   [2]int
}

func (this *field) Coordinates() (int, int) {
	return this.cords[0], this.cords[1]
}

func (this *field) LeftBorder() c.Border {
	return this.borders[0]
}

func (this *field) TopBorder() c.Border {
	return this.borders[1]
}

func (this *field) SetCoordinates(row, col int) {
	this.cords[0] = row
	this.cords[1] = col
}
