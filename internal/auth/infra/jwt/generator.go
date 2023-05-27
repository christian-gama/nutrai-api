package jwt

import (
	"context"
	"time"

	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/jwt"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/token"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/repo"
	value "github.com/christian-gama/nutrai-api/internal/auth/domain/value/jwt"
	"github.com/christian-gama/nutrai-api/internal/core/domain/uuid"
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
	_jwt "github.com/golang-jwt/jwt"
)

// generatorImpl is the implementation of jwt.Generator.
type generatorImpl struct {
	token     *_jwt.Token
	uuid      uuid.Generator
	tokenType jwt.TokenType
	duration  time.Duration
	tokenRepo repo.Token
}

// NewGenerator returns a new instance of Generator.
func NewGenerator(
	uuid uuid.Generator,
	tokenType jwt.TokenType,
	duration time.Duration,
	tokenRepo repo.Token,
) jwt.Generator {
	errutil.MustBeNotEmpty("jwt.TokenType", tokenType)
	errutil.MustBeNotEmpty("jwt.Duration", duration)
	errutil.MustBeNotEmpty("repo.Token", tokenRepo)

	return &generatorImpl{
		token:     _jwt.New(_jwt.SigningMethodHS256),
		uuid:      uuid,
		tokenType: tokenType,
		duration:  duration,
		tokenRepo: tokenRepo,
	}
}

// Generate implements jwt.Generator.
func (g *generatorImpl) Generate(subject *jwt.Subject, persist bool) (value.Token, error) {
	if err := g.validate(subject); err != nil {
		return "", err
	}

	claims := g.setClaims(subject)
	if persist {
		if _, err := g.tokenRepo.Save(context.Background(), repo.SaveTokenInput{
			Token: token.NewToken().
				SetEmail(subject.Email).
				SetJti(claims["jti"].(coreValue.UUID)).
				SetExpiresAt(g.getExpiresAtDuration(claims["exp"].(int64))),
		}); err != nil {
			return "", err
		}
	}

	signed, err := g.signToken()
	if err != nil {
		return "", err
	}

	return value.Token(signed), nil
}

func (g *generatorImpl) getExpiresAtDuration(unix int64) time.Duration {
	expirationTime := time.Unix(unix, 0)
	expiresAt := time.Until(expirationTime)
	return expiresAt
}

// setClaims is a helper method that sets the claims of a JWT token based on the provided subject.
// Claims include various details like audience, expiry time, issuer, subject etc. It validates
// these claims before setting them into the token.
func (g *generatorImpl) setClaims(subject *jwt.Subject) _jwt.MapClaims {
	claims := g.token.Claims.(_jwt.MapClaims)
	claims["aud"] = env.App.Host
	claims["exp"] = time.Now().Add(g.duration).Unix()
	claims["iat"] = time.Now().Unix()
	claims["iss"] = env.App.Host
	claims["jti"] = g.uuid.Generate()
	claims["nbf"] = time.Now().Unix()
	claims["sub"] = map[string]any{"email": subject.Email}
	claims["type"] = g.tokenType

	return claims
}

// signToken is a helper method that signs the JWT token with a secret key, which is used for later
// verification of the token's authenticity.
func (g *generatorImpl) signToken() (string, error) {
	signed, err := g.token.SignedString([]byte(env.Jwt.Secret))
	if err != nil {
		return "", err
	}

	return signed, nil
}

// validate is a helper method that checks the validity of a JWT subject. It ensures the subject is
// not nil and its email is valid. If not, it returns an error.
func (g *generatorImpl) validate(subject *jwt.Subject) error {
	if subject == nil {
		return errors.InternalServerError(errors.Required("subject").Error())
	}

	if subject.Email == "" {
		return errors.InternalServerError(errors.Required("subject.Email").Error())
	}

	return nil
}
