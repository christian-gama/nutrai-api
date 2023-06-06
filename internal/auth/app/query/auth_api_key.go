package query

// AuthApiKeyInput is the query to check if the JWT token is valid and find the user associated with
// it.
type AuthApiKeyInput struct {
	Key string `json:"key" validate:"required"`
}

// AuthApiKeyOutput is the output of the AuthInput query.
type AuthApiKeyOutput struct {
	Key string `json:"key"`
}
