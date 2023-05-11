package command

// DeleteUserInput represents the input data for the DeleteUser command.
type DeleteUserInput struct {
	ID uint `form:"id" faker:"uint"`
}
