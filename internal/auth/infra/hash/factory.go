package hash

import "github.com/christian-gama/nutrai-api/internal/auth/domain/hasher"

func MakeHasher() hasher.Hasher {
	return New()
}
