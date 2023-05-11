package command

// CheckCredentialsInput is the input for CheckCredentials.
type CheckCredentialsInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
