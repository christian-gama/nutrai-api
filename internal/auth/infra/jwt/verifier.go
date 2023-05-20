package jwt

import (
	"fmt"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/jwt"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/value"
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	"github.com/christian-gama/nutrai-api/internal/core/infra/env"
	_jwt "github.com/golang-jwt/jwt"
)

// verifierImpl is the implementation of the Verifier interface.
type verifierImpl struct{}

// NewVerifier creates a new JWT verifier.
func NewVerifier() jwt.Verifier {
	return &verifierImpl{}
}

// Verify implements the jwt.Verifier interface.
func (s *verifierImpl) Verify(t value.Token) (*jwt.Payload, error) {
	token, err := _jwt.Parse(t.String(), keyFunc)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(_jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}

	return s.getPayload(claims), nil
}

// getPayload converts the claims to a Payload.
func (s *verifierImpl) getPayload(claims _jwt.MapClaims) *jwt.Payload {
	sub := claims["sub"].(map[string]any)
	data := jwt.Subject{
		Email: value.Email(sub["email"].(string)),
	}

	return &jwt.Payload{
		Aud:  claims["aud"].(string),
		Exp:  int64(claims["exp"].(float64)),
		Iat:  int64(claims["iat"].(float64)),
		Iss:  claims["iss"].(string),
		Jti:  coreValue.UUID(claims["jti"].(string)),
		Nbf:  int64(claims["nbf"].(float64)),
		Sub:  data,
		Type: jwt.TokenType(claims["type"].(string)),
	}
}

// keyFunc implements the jwt.Keyfunc func type.
func keyFunc(token *_jwt.Token) (any, error) {
	if _, ok := token.Method.(*_jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(env.Jwt.Secret), nil
}
