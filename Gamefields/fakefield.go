package Gamefields

import (
	co "../Constants"
)

type FakeField struct {
	Cords   [2]int
	Borders []co.Border
}

func (this *FakeField) SetCoordinates(row int, col int) {
	this.Cords[0] = row
	this.Cords[1] = col
	return
}

func (this *FakeField) LeftBorder() co.Border {
	return this.Borders[0]
}

func (this *FakeField) TopBorder() co.Border {
	return this.Borders[1]
}
