package main

import (
	build "../TheGome/Builders"
	view "../TheGome/Stringview"

	"fmt"
)

func main() {
	view := view.NewStringView()
	board := build.BuildChapter00()

	fmt.Println(view.CompleteBoard(board))
	return
}
