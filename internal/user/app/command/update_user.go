package command

// UpdateUserInput represents the input data for the UpdateUser command.
type UpdateUserInput struct {
	Email string `json:"email" faker:"email"`
	Name  string `json:"name" faker:"name"`
}
