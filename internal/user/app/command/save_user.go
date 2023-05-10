package command

// SaveUserInput represents the input data for the SaveUser command.
type SaveUserInput struct {
	Email    string `json:"email" faker:"email"`
	Password string `json:"password" faker:"len=8"`
	Name     string `json:"name" faker:"name"`
}
