package Stringview

type FakeBoard struct {
	fields map[int][]Field
}

func (this *FakeBoard) FieldsByRow(i int) ([]Field, bool) {
	fields, ok := this.fields[i]
	return fields, ok
}

type FakeField struct {
	survivors []FieldObject
	zombies   []FieldObject
	cords     []int
	borders   []Border
}

func (this *FakeField) Coordinates() (int, int) {
	return this.cords[0], this.cords[1]
}

func (this *FakeField) Survivors() []FieldObject {
	return this.survivors
}

func (this *FakeField) Zombies() []FieldObject {
	return this.zombies
}

func (this *FakeField) LeftBorder() Border {
	return this.borders[0]
}

func (this *FakeField) TopBorder() Border {
	return this.borders[1]
}

type FakeFieldObject struct {
	identifier string
}

func (this *FakeFieldObject) AsString() string {
	return this.identifier
}
