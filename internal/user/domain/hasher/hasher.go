package hasher

import value "github.com/christian-gama/nutrai-api/internal/user/domain/value/user"

type Hasher interface {
	// Hash returns the hashed password.
	Hash(password value.Password) (value.Password, error)

	// Compare returns an error if the password does not match the hash.
	Compare(password value.Password, hash value.Password) error
}
