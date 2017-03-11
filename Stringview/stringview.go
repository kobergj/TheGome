package Stringview

import (
	"errors"
	"strings"
)

const (
	// The size of one field in characters
	FIELDSIZE = 3
	// String Constants
	NEWLINE     = "\n"
	EMPTYSTRING = ""
	SPACE       = " "
	PIPE        = "|"
	PLUS        = "+"
	// String Identifiers
	HORIZONTALDOOR      = "- -"
	HORIZONTALWALL      = "---"
	HORIZONTALFREE      = "   "
	HORIZONTALCROSSWALK = " = "
	VERTICALDOOR        = "|\n "
	VERTICALWALL        = "|\n|"
	VERTICALFREE        = " \n "
	VERTICALCROSSWALK   = "=\n="
	EMPTYROOM           = "   \n   "
	// Error Messages
	ERR_STRINGSDONTMATCH = "Tried to concatenate multiline string, but failed"
)

func NewStringView() View {
	return &stringView{
		verticalBorders: map[Border]string{
			WALL:      VERTICALWALL,
			DOOR:      VERTICALDOOR,
			FREE:      VERTICALFREE,
			CROSSWALK: VERTICALCROSSWALK,
		},
		horizontalBorders: map[Border]string{
			WALL:      HORIZONTALWALL,
			DOOR:      HORIZONTALDOOR,
			FREE:      HORIZONTALFREE,
			CROSSWALK: HORIZONTALCROSSWALK,
		},
	}
}

type stringView struct {
	verticalBorders   map[Border]string
	horizontalBorders map[Border]string
}

func (this *stringView) CompleteBoard(board Board) string {
	boardString := EMPTYSTRING

	row := 0
	for {
		fields, ok := board.FieldsByRow(row)
		if !ok {
			break
		}

		str := NEWLINE + NEWLINE

		for col, field := range fields {
			interior := this.fieldInterior(board.BoardObjectsByField(row, col))
			str, _ = matchStringsByNewline(str, this.OneField(field, interior))
		}

		rowfinish := PIPE + NEWLINE + VERTICALWALL
		str, _ = matchStringsByNewline(str, rowfinish)

		if row > 0 {
			boardString += NEWLINE
		}

		boardString += str

		row++
	}

	row0, _ := board.FieldsByRow(0)
	bottomborder := strings.Repeat(PIPE+HORIZONTALWALL, len(row0))
	return boardString + NEWLINE + bottomborder + PIPE
}

func (this *stringView) OneField(field Field, room string) string {
	top, _ := this.horizontalBorders[field.TopBorder()]

	left, _ := this.verticalBorders[field.LeftBorder()]

	body, _ := matchStringsByNewline(left, room)

	return PIPE + top + NEWLINE + body
}

func (this *stringView) fieldInterior(objectFeed <-chan FieldObject) string {
	room := EMPTYROOM

	max := strings.Count(room, SPACE)

	i := 1
	for obj := range objectFeed {
		id := obj.AsString()
		if i == max {
			id = PLUS
		}
		room = strings.Replace(room, SPACE, id, 1)
		i++
	}

	return room
}

func matchStringsByNewline(a, b string) (string, error) {
	aSplitted := strings.Split(a, NEWLINE)
	bSplitted := strings.Split(b, NEWLINE)

	for len(aSplitted) != len(bSplitted) {
		println(ERR_STRINGSDONTMATCH)
		println("A:")
		println(a)
		println("B:")
		println(b)
		return EMPTYSTRING, errors.New(ERR_STRINGSDONTMATCH)
	}

	str := ""
	for i := 0; i < len(aSplitted); i++ {
		if i > 0 {
			str += NEWLINE
		}
		str += aSplitted[i]
		str += bSplitted[i]

	}

	return str, nil
}
