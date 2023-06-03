package command

// RecoveryInput is the command to save a new exception.
type RecoveryInput struct {
	Message string `json:"message"`
	Stack   string `json:"stack"`
}
