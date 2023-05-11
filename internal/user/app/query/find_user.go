package query

// FindUserInput is the input data of the user query.
type FindUserOutput struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}
