package Stringview

import (
	"errors"
	"strings"

	c "../Constants"

	gb "../Gameboard"
	gf "../Gamefields"
	fo "../Gameobjects"
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
		verticalBorders: map[c.Border]string{
			c.WALL:      VERTICALWALL,
			c.DOOR:      VERTICALDOOR,
			c.FREE:      VERTICALFREE,
			c.CROSSWALK: VERTICALCROSSWALK,
		},
		horizontalBorders: map[c.Border]string{
			c.WALL:      HORIZONTALWALL,
			c.DOOR:      HORIZONTALDOOR,
			c.FREE:      HORIZONTALFREE,
			c.CROSSWALK: HORIZONTALCROSSWALK,
		},
	}
}

type stringView struct {
	verticalBorders   map[c.Border]string
	horizontalBorders map[c.Border]string
}

func (this *stringView) CompleteBoard(board gb.Board) string {
	boardString := EMPTYSTRING

	row := 0
	for {
		fields, ok := board.FieldsByRow(row)
		if !ok {
			break
		}

		str := NEWLINE + NEWLINE

		for _, field := range fields {
			fieldObjects := board.BoardObjectsByField(field)
			interior := this.fieldInterior(fieldObjects)
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
	bottomBorder := strings.Repeat(PIPE+HORIZONTALWALL, len(row0))
	return boardString + NEWLINE + bottomBorder + PIPE
}

func (this *stringView) OneField(field gf.Field, room string) string {
	top, _ := this.horizontalBorders[field.TopBorder()]

	left, _ := this.verticalBorders[field.LeftBorder()]

	body, _ := matchStringsByNewline(left, room)

	return PIPE + top + NEWLINE + body
}

func (this *stringView) fieldInterior(fieldObjects []fo.FieldObject) string {
	room := EMPTYROOM

	max := strings.Count(room, SPACE)

	i := 1
	for _, obj := range fieldObjects {
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
