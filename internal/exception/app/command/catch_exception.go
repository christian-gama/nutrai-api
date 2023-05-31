package command

// CatchExceptionInput is the command to save a new exception.
type CatchExceptionInput struct {
	Message string `json:"message"`
	Stack   string `json:"stack"`
}
