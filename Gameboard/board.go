package Gameboard

import (
	"reflect"

	gf "../Gamefields"
	fo "../Gameobjects"
)

func NewBoard(rows, cols int) Board {
	fields := make([][]gf.Field, rows)
	for i := range fields {
		fields[i] = make([]gf.Field, 0, cols)
	}
	return &board{
		fields: fields,
	}
}

type board struct {
	fields       [][]gf.Field
	fieldObjects []fo.FieldObject
}

func (this *board) FieldsByRow(row int) ([]gf.Field, bool) {
	if row >= len(this.fields) {
		return nil, false
	}

	return this.fields[row], true
}

func (this *board) BoardObjectsByField(field gf.Field) []fo.FieldObject {
	fos := make([]fo.FieldObject, 0)

	for _, obj := range this.fieldObjects {
		cur := obj.CurrentField()

		if reflect.DeepEqual(cur, field) && obj.Alive() {
			fos = append(fos, obj)
		}
	}

	return fos
}

func (this *board) AddField(field gf.Field) {
	for x, row := range this.fields {
		y := len(row)
		if y < cap(row) {
			field.SetCoordinates(x, y)
			row = append(row, field)
			this.fields[x] = row
			return
		}
	}
	panic("Too Many fields added")
}

func (this *board) AddFieldObject(obj fo.FieldObject) {
	this.fieldObjects = append(this.fieldObjects, obj)
	return
}
