package builders

import (
	co "../Constants"
	gb "../Gameboard"
	gf "../Gamefields"
)

func BuildChapter00() gb.Board {
	rows := 6
	cols := 5
	board := gb.NewBoard(rows, cols)

	var field gf.Field

	// Row 1
	field = gf.NewField(co.WALL, co.WALL)
	board.AddField(field)

	field = gf.NewField(co.WALL, co.WALL)
	board.AddField(field)

	field = gf.NewField(co.WALL, co.WALL)
	board.AddField(field)

	field = gf.NewField(co.WALL, co.WALL)
	board.AddField(field)

	field = gf.NewField(co.WALL, co.WALL)
	board.AddField(field)

	// Row 2
	field = gf.NewField(co.WALL, co.CROSSWALK)
	board.AddField(field)

	field = gf.NewField(co.DOOR, co.DOOR)
	board.AddField(field)

	field = gf.NewField(co.WALL, co.CROSSWALK)
	board.AddField(field)

	field = gf.NewField(co.WALL, co.DOOR)
	board.AddField(field)

	field = gf.NewField(co.WALL, co.CROSSWALK)
	board.AddField(field)

	// Row 3
	field = gf.NewField(co.WALL, co.CROSSWALK)
	board.AddField(field)

	field = gf.NewField(co.WALL, co.DOOR)
	board.AddField(field)

	field = gf.NewField(co.WALL, co.CROSSWALK)
	board.AddField(field)

	field = gf.NewField(co.CROSSWALK, co.DOOR)
	board.AddField(field)

	field = gf.NewField(co.CROSSWALK, co.CROSSWALK)
	board.AddField(field)

	// Row 4
	field = gf.NewField(co.WALL, co.FREE)
	board.AddField(field)

	field = gf.NewField(co.WALL, co.DOOR)
	board.AddField(field)

	field = gf.NewField(co.WALL, co.FREE)
	board.AddField(field)

	field = gf.NewField(co.CROSSWALK, co.FREE)
	board.AddField(field)

	field = gf.NewField(co.CROSSWALK, co.FREE)
	board.AddField(field)

	// Row 5
	field = gf.NewField(co.WALL, co.WALL)
	board.AddField(field)

	field = gf.NewField(co.DOOR, co.DOOR)
	board.AddField(field)

	field = gf.NewField(co.DOOR, co.CROSSWALK)
	board.AddField(field)

	field = gf.NewField(co.WALL, co.DOOR)
	board.AddField(field)

	field = gf.NewField(co.FREE, co.WALL)
	board.AddField(field)

	// Row 6
	field = gf.NewField(co.WALL, co.FREE)
	board.AddField(field)

	field = gf.NewField(co.DOOR, co.DOOR)
	board.AddField(field)

	field = gf.NewField(co.DOOR, co.CROSSWALK)
	board.AddField(field)

	field = gf.NewField(co.WALL, co.FREE)
	board.AddField(field)

	field = gf.NewField(co.FREE, co.FREE)
	board.AddField(field)

	return board

}
