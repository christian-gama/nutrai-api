package token

import (
	"time"

	userValue "github.com/christian-gama/nutrai-api/internal/auth/domain/value/user"
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
)

// Token represents a Token model, containing essential credentials and personal identification
// information for an individual user. This includes a unique identifier (jti), email address, and
// expiration time.

type Token struct {
	Email     userValue.Email `faker:"email"`
	Jti       coreValue.UUID  `faker:"uuid_hyphenated"`
	ExpiresAt time.Duration   `faker:"unix_time"`
}

// NewToken returns a new token instance.
func NewToken() *Token {
	return &Token{}
}

func (r Token) String() string {
	return "token"
}

// Validate returns an error if the token is invalid.
func (r *Token) Validate() (*Token, error) {
	var errs *errutil.Error

	if r.Email == "" {
		errs = errutil.Append(errors.Required("Email"))
	}

	if r.Jti == "" {
		errs = errutil.Append(errors.Required("Jti"))
	}

	if r.ExpiresAt == 0 {
		errs = errutil.Append(errors.Required("ExpiresAt"))
	}

	if errs != nil {
		return nil, errs
	}

	return r, nil
}

// SetEmail sets the token email.
func (r *Token) SetEmail(email userValue.Email) *Token {
	r.Email = email
	return r
}

// SetJti sets the token jti.
func (r *Token) SetJti(jti coreValue.UUID) *Token {
	r.Jti = jti
	return r
}

// SetExpiresAt sets the token expiration time.
func (r *Token) SetExpiresAt(expiresAt time.Duration) *Token {
	r.ExpiresAt = expiresAt
	return r
}
