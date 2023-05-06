package hash

import (
	"github.com/christian-gama/nutrai-api/internal/user/domain/hasher"
	value "github.com/christian-gama/nutrai-api/internal/user/domain/value/user"
	"golang.org/x/crypto/bcrypt"
)

// hashImpl is the implementation of the hasher.Hasher interface.
type hashImpl struct{}

// New returns a new Hash instance.
func New() hasher.Hasher {
	return &hashImpl{}
}

// Hash implements the hasher.Hasher interface.
func (h *hashImpl) Hash(password value.Password) (value.Password, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return value.Password(bytes), nil
}

// Compare implements the hasher.Hasher interface.
func (h *hashImpl) Compare(password value.Password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
