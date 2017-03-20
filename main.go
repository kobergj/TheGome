package main

import (
	build "../TheGome/Builders"
	cntr "../TheGome/Controller"
	view "../TheGome/Stringview"
)

func main() {
	controller := cntr.NewController()
	view := view.NewStringView()
	board := build.BuildChapter00()

	controller.VisualizeString(view.CompleteBoard(board))

	err := controller.ExecuteInput()
	if err != nil {
		panic(err)
	}
	return
}
