package jwt

import (
	jwtValue "github.com/christian-gama/nutrai-api/internal/auth/domain/value/jwt"
	userValue "github.com/christian-gama/nutrai-api/internal/auth/domain/value/user"
)

// Subject represents the subject claim of a JWT token. It's the primary identity that the token
// refers to and is used to identify the user for whom the token was issued. In this case, the
// Subject contains the user's email, which will be encoded into the JWT token.
type Subject struct {
	Email userValue.Email `json:"email" faker:"email"`
}

// Generator is the interface that wraps the Generate method.
type Generator interface {
	// Generate is a method that generates a new JWT token with claims based on the provided
	// subject. The method should validate the subject, sets the claims, and signs the token before
	// returning it.
	Generate(data *Subject) (jwtValue.Token, error)
}
