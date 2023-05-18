package jwt

import (
	"time"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/jwt"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/value"
	"github.com/christian-gama/nutrai-api/internal/core/domain/uuid"
	"github.com/christian-gama/nutrai-api/internal/core/infra/env"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	_jwt "github.com/golang-jwt/jwt"
)

// generatorImpl is the implementation of jwt.Generator.
type generatorImpl struct {
	token     *_jwt.Token
	uuid      uuid.Generator
	tokenType jwt.TokenType
	duration  time.Duration
}

// NewGenerator returns a new instance of Generator.
func NewGenerator(
	uuid uuid.Generator,
	tokenType jwt.TokenType,
	duration time.Duration,
) jwt.Generator {
	return &generatorImpl{
		token:     _jwt.New(_jwt.SigningMethodHS256),
		uuid:      uuid,
		tokenType: tokenType,
		duration:  duration,
	}
}

// Generate implements jwt.Generator.
func (g *generatorImpl) Generate(subject *jwt.Subject) (value.Token, error) {
	if err := g.Validate(subject); err != nil {
		return "", err
	}

	claims := g.token.Claims.(_jwt.MapClaims)
	claims["exp"] = time.Now().Add(g.duration).Unix()
	claims["iat"] = time.Now().Unix()
	claims["sub"] = map[string]any{
		"id":    subject.ID,
		"email": subject.Email,
	}
	claims["iss"] = env.App.Host
	claims["aud"] = env.App.Host
	claims["type"] = g.tokenType
	claims["nbf"] = time.Now().Unix()
	claims["jti"] = g.uuid.Generate()

	signed, err := g.token.SignedString([]byte(env.Jwt.Secret))
	if err != nil {
		return "", err
	}

	return value.Token(signed), nil
}

func (g *generatorImpl) Validate(subject *jwt.Subject) error {
	if subject == nil {
		return errutil.NewErrInternal(errutil.NewErrRequired("subject").Error())
	}

	if err := subject.ID.Validate(); err != nil {
		return errutil.NewErrInternal(err.Error())
	}

	if err := subject.Email.Validate(); err != nil {
		return errutil.NewErrInternal(err.Error())
	}

	return nil
}
