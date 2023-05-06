package hash

import (
	"github.com/christian-gama/nutrai-api/internal/user/domain/hasher"
	"golang.org/x/crypto/bcrypt"
)

// hashImpl is the implementation of the hasher.Hasher interface.
type hashImpl struct{}

// New returns a new Hash instance.
func New() hasher.Hasher {
	return &hashImpl{}
}

// Hash implements the hasher.Hasher interface.
func (h *hashImpl) Hash(value string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Compare implements the hasher.Hasher interface.
func (h *hashImpl) Compare(value, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(value))
}
