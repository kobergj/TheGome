package Gameboard

import (
	gf "../Gamefields"
	fo "../Gameobjects"
)

type FakeBoard struct {
	Fields map[int][]gf.Field

	FieldObjects [][]fo.FieldObject
	Index        int
}

func (this *FakeBoard) FieldsByRow(i int) ([]gf.Field, bool) {
	fields, ok := this.Fields[i]
	return fields, ok
}

func (this *FakeBoard) BoardObjectsByField(field gf.Field) []fo.FieldObject {
	this.Index += 1
	if this.Index >= len(this.FieldObjects) {
		return nil
	}
	return this.FieldObjects[this.Index]
}

func (this *FakeBoard) AddField(gf.Field) {
	return
}

func (this *FakeBoard) AddFieldObject(fo.FieldObject) {
	return
}
