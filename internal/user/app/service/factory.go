package service

import (
	"github.com/christian-gama/nutrai-api/internal/user/infra/hash"
)

func MakeHashPasswordHandler() HashPasswordHandler {
	return NewHashPasswordHandler(hash.MakeHasher())
}
