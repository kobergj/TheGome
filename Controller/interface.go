package Controller

type Controller interface {
	// Visualize the given string aka print
	VisualizeString(string)
	// Register the actions. Should be called before ExecuteInput
	SetActions([]func() error)
	// Should return error if input is invalid. TODO: define 'invalid'
	ExecuteInput() error
}
