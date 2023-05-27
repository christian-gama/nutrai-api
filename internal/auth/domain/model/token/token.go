package token

import (
	"time"

	userValue "github.com/christian-gama/nutrai-api/internal/auth/domain/value/user"
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
)

type Token struct {
	Email     userValue.Email
	Jti       coreValue.UUID
	ExpiresAt time.Duration
}

func NewToken() *Token {
	return &Token{}
}

func (r *Token) Validate() (*Token, error) {
	var errs *errutil.Error

	if r.Email == "" {
		errs = errutil.Append(errors.Required("email is required"))
	}

	if r.Jti == "" {
		errs = errutil.Append(errors.Required("jti is required"))
	}

	if errs != nil {
		return nil, errs
	}

	return r, nil
}

func (r *Token) SetEmail(email userValue.Email) *Token {
	r.Email = email
	return r
}

func (r *Token) SetJti(jti coreValue.UUID) *Token {
	r.Jti = jti
	return r
}

func (r *Token) SetExpiresAt(expiresAt time.Duration) *Token {
	r.ExpiresAt = expiresAt
	return r
}
