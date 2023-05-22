package http

import "github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"

// Meta is the metadata of a controller.
type Meta struct {
	Method      func() Method
	Path        func() Path
	IsPublic    func() bool
	Params      func() Params
	CurrentUser func() *user.User
}
