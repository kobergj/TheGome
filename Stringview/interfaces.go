package Stringview

type Border int

const (
	WALL Border = iota
	DOOR
	FREE
	CROSSWALK
)

type View interface {
	// Return the Board as a string
	CompleteBoard(Board) string
}

type Board interface {
	// Return the sections of the board
	FieldsByRow(int) ([]Field, bool)
	// The BoardObject inside the given Field
	BoardObjectsByField(int, int) <-chan FieldObject
}

type Field interface {
	// Borders
	LeftBorder() Border
	TopBorder() Border
	// Position of the Field inside the board
	Coordinates() (int, int)
}

type FieldObject interface {
	// A String Identifier for the Object
	AsString() string
	// Position of the Object inside the board
	Coordinates() (int, int)
}
