package Gameobjects

import (
	gf "../Gamefields"
)

type FakeFieldObject struct {
	IsAlive    bool
	Current    gf.Field
	Identifier string
}

func (this *FakeFieldObject) AsString() string {
	return this.Identifier
}

func (this *FakeFieldObject) Alive() bool {
	return this.IsAlive
}

func (this *FakeFieldObject) CurrentField() gf.Field {
	return this.Current
}
