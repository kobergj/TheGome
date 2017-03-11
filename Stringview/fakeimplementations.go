package Stringview

type FakeBoard struct {
	fields       map[int][]Field
	boardObjects []FieldObject
}

func (this *FakeBoard) FieldsByRow(i int) ([]Field, bool) {
	fields, ok := this.fields[i]
	return fields, ok
}

func (this *FakeBoard) BoardObjectsByField(row, col int) <-chan FieldObject {
	foChannel := make(chan FieldObject)

	go func() {
		defer close(foChannel)
		for _, bo := range this.boardObjects {
			if y, x := bo.Coordinates(); y == row && x == col {
				foChannel <- bo
			}
		}
	}()

	return foChannel
}

type FakeField struct {
	cords   []int
	borders []Border
}

func (this *FakeField) Coordinates() (int, int) {
	return this.cords[0], this.cords[1]
}

func (this *FakeField) LeftBorder() Border {
	return this.borders[0]
}

func (this *FakeField) TopBorder() Border {
	return this.borders[1]
}

type FakeFieldObject struct {
	identifier string
	cords      []int
}

func (this *FakeFieldObject) AsString() string {
	return this.identifier
}

func (this *FakeFieldObject) Coordinates() (int, int) {
	return this.cords[0], this.cords[1]
}
