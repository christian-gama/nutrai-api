package jwt

import (
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
)

// Payload is the payload of the JWT token.
type Payload struct {
	Aud  string         `json:"aud"`
	Exp  int64          `json:"exp"`
	Iat  int64          `json:"iat"`
	Iss  string         `json:"iss"`
	Jti  coreValue.UUID `json:"jti"`
	Nbf  int64          `json:"nbf"`
	Sub  Subject        `json:"sub"`
	Type TokenType      `json:"type"`
}
