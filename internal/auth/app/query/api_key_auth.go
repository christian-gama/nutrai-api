package query

// ApiKeyAuthInput is the query to check if the JWT token is valid and find the user associated with
// it.
type ApiKeyAuthInput struct {
	Key string `json:"key" validate:"required"`
}

// ApiKeyAuthOutput is the output of the AuthInput query.
type ApiKeyAuthOutput struct {
	Key string `json:"key"`
}
