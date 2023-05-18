package jwt

import "github.com/christian-gama/nutrai-api/internal/auth/domain/value"

// Verifier is the interface that wraps the Verify method.
type Verifier interface {
	// Verify verifies a JWT token.
	Verify(token value.Token) (*Payload, error)
}
