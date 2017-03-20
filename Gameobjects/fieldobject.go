package Gameobjects

import (
	gf "../Gamefields"
)

func NewFieldObject(identifier string) FieldObject {
	return &fieldObject{
		alive:      true,
		identifier: identifier,
	}
}

type fieldObject struct {
	alive        bool
	identifier   string
	currentField gf.Field
}

func (this *fieldObject) Alive() bool {
	return this.alive
}

func (this *fieldObject) AsString() string {
	return this.identifier
}

func (this *fieldObject) CurrentField() gf.Field {
	return this.currentField
}

func (this *fieldObject) Move(field gf.Field) {
	this.currentField = field
	return
}
