package jwt

import (
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
)

// Claims is the claim of the JWT token. Please refer to
// https://datatracker.ietf.org/doc/html/rfc7519#section-4.1 for more information.
type Claims struct {
	// Aud (Audience) identifies the recipients that the JWT is intended for. It can be a single
	// recipient or multiple recipients.
	Aud string `json:"aud"`

	// Exp (Expiration Time) is a timestamp that identifies the time on or after which the JWT must
	// not be accepted for processing. The value is typically a NumericDate: a JSON numeric value
	// representing the number of seconds from 1970-01-01T00:00:00Z UTC until the specified UTC
	// date/time.
	Exp int64 `json:"exp"`

	// Iat (Issued At) is a timestamp that identifies the time at which the JWT was issued. Like
	// 'exp', this is usually a NumericDate.
	Iat int64 `json:"iat"`

	// Iss (Issuer) identifies the principal that issued the JWT. This could be a server, a service,
	// an identity provider, or some other kind of principal.
	Iss string `json:"iss"`

	// Jti (JWT ID) provides a unique identifier for the JWT. It can be used to prevent the JWT from
	// being replayed by maintaining a record of used JTI values.
	Jti coreValue.UUID `json:"jti"`

	// Nbf (Not Before) is a timestamp that identifies the time before which the JWT must not be
	// accepted for processing. Like 'exp' and 'iat', this is usually a NumericDate.
	Nbf int64 `json:"nbf"`

	// Sub (Subject) identifies the principal that is the subject of the JWT. In other words, it's
	// the identity that the claim is about. The Subject value could be a user's id, email or
	// username depending upon the system.
	Sub Subject `json:"sub"`

	// Type (Token Type) is a custom field, not defined in the JWT standard. It's likely used to
	// differentiate different types of tokens within the same application (like access token,
	// refresh token, etc.).
	Type TokenType `json:"type"`
}
