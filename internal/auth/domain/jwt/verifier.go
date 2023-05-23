package jwt

import value "github.com/christian-gama/nutrai-api/internal/auth/domain/value/jwt"

// Verifier is the interface that wraps the Verify method.
type Verifier interface {
	// Verify is a method that accepts a JWT token and validates it. If the token is valid, it
	// returns the claims in it. If the token is invalid, it returns an error.
	Verify(token value.Token) (*Claims, error)
}
