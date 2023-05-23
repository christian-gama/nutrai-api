package command

// SaveExceptionInput is the command to save a new exception.
type SaveExceptionInput struct {
	Message string `json:"message"`
	Stack   string `json:"stack"`
}
