package jwt

import (
	"fmt"

	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/jwt"
	jwtValue "github.com/christian-gama/nutrai-api/internal/auth/domain/value/jwt"
	userValue "github.com/christian-gama/nutrai-api/internal/auth/domain/value/user"
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	_jwt "github.com/golang-jwt/jwt"
)

// verifierImpl is the implementation of the Verifier interface.
type verifierImpl struct {
	tokenType jwt.TokenType
}

// NewVerifier creates a new JWT verifier.
func NewVerifier(tokenType jwt.TokenType) jwt.Verifier {
	errutil.MustBeNotEmpty("jwt.TokenType", tokenType)

	return &verifierImpl{
		tokenType: tokenType,
	}
}

// Verify implements the jwt.Verifier interface.
func (s *verifierImpl) Verify(t jwtValue.Token) (*jwt.Claims, error) {
	token, err := _jwt.Parse(t.String(), keyFunc)
	if err != nil {
		return nil, err
	}

	mapClaims, ok := token.Claims.(_jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}

	claims := s.getClaims(mapClaims)
	if err := s.validate(claims); err != nil {
		return nil, err
	}

	return claims, nil
}

// getClaims is a helper method that converts the claims from a JWT token into a structured
// jwt.Claims object. It extracts and formats the relevant fields from the raw claims.
func (s *verifierImpl) getClaims(mapClaims _jwt.MapClaims) *jwt.Claims {
	sub := mapClaims["sub"].(map[string]any)
	data := jwt.Subject{
		Email: userValue.Email(sub["email"].(string)),
	}

	return &jwt.Claims{
		Aud:  mapClaims["aud"].(string),
		Exp:  int64(mapClaims["exp"].(float64)),
		Iat:  int64(mapClaims["iat"].(float64)),
		Iss:  mapClaims["iss"].(string),
		Jti:  coreValue.UUID(mapClaims["jti"].(string)),
		Nbf:  int64(mapClaims["nbf"].(float64)),
		Sub:  data,
		Type: jwt.TokenType(mapClaims["type"].(string)),
	}
}

// keyFunc is a helper function that validates the signing method of a JWT token and returns the
// secret key used for signing the token. If the signing method is not HMAC, it returns an error.
func keyFunc(token *_jwt.Token) (any, error) {
	if _, ok := token.Method.(*_jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(env.Jwt.Secret), nil
}

// validate is a helper method that validates the claims of a JWT token. It checks if the token is
// of the expected type and if the audience is the same as the application host.
func (s *verifierImpl) validate(claims *jwt.Claims) error {
	if claims.Type != s.tokenType {
		return ErrInvalidToken
	}

	if claims.Aud != env.App.Host {
		return ErrInvalidToken
	}

	return nil
}
