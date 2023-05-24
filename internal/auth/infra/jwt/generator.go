package jwt

import (
	"time"

	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/jwt"
	value "github.com/christian-gama/nutrai-api/internal/auth/domain/value/jwt"
	"github.com/christian-gama/nutrai-api/internal/core/domain/uuid"
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
	if err := g.validate(subject); err != nil {
		return "", err
	}

	g.setClaims(subject)

	signed, err := g.signToken()
	if err != nil {
		return "", err
	}

	return value.Token(signed), nil
}

// setClaims is a helper method that sets the claims of a JWT token based on the provided subject.
// Claims include various details like audience, expiry time, issuer, subject etc. It validates
// these claims before setting them into the token.
func (g *generatorImpl) setClaims(subject *jwt.Subject) {
	claims := g.token.Claims.(_jwt.MapClaims)
	claims["aud"] = env.App.Host
	claims["exp"] = time.Now().Add(g.duration).Unix()
	claims["iat"] = time.Now().Unix()
	claims["iss"] = env.App.Host
	claims["jti"] = g.uuid.Generate()
	claims["nbf"] = time.Now().Unix()
	claims["sub"] = map[string]any{"email": subject.Email}
	claims["type"] = g.tokenType
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
		return errutil.InternalServerError(errutil.Required("subject").Error())
	}

	if err := subject.Email.Validate(); err != nil {
		return errutil.InternalServerError(err.Error())
	}

	return nil
}
