package Controller

import (
	"errors"
	"fmt"
)

func NewController() Controller {
	return &controller{}
}

type controller struct {
	actions []func() error
}

func (this *controller) VisualizeString(str string) {
	fmt.Println(str)
	return
}

func (this *controller) SetActions(actions []func() error) {
	this.actions = actions
	return
}

func (this *controller) ExecuteInput() error {
	var i int

	_, err := fmt.Scanln(&i)
	if err != nil {
		return err
	}

	if i >= len(this.actions) {
		return errors.New(fmt.Sprintf("Wrong input: '%s'", i))
	}

	err = this.actions[i]()
	if err != nil {
		return err
	}

	return nil
}
