package hash

import "github.com/christian-gama/nutrai-api/internal/user/domain/hasher"

func MakeHasher() hasher.Hasher {
	return New()
}
