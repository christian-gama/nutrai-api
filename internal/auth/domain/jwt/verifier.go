package jwt

import value "github.com/christian-gama/nutrai-api/internal/auth/domain/value/jwt"

// Verifier is the interface that wraps the Verify method.
type Verifier interface {
	// Verify verifies a JWT token.
	Verify(token value.Token) (*Claims, error)
}
