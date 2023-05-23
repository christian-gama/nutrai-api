package store

import (
	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"
	"github.com/gin-gonic/gin"
)

const currentUser = "currentUser"

// SetUser is a function that stores the current authenticated user in the Gin context.
// The Gin context is a per-request context, it carries data across API boundaries and between
// middleware. This allows you to pass the authenticated user's information to further processing
// functions or middleware.
func SetUser(ctx *gin.Context, user *user.User) {
	ctx.Set(currentUser, user)
}

// GetUser is a function that retrieves the current authenticated user's information from the Gin
// context. This is useful when you need to access user information in subsequent processing
// functions or middleware. If the user is not found in the context, or the user information is
// invalid, appropriate errors are returned.
func GetUser(ctx *gin.Context) (*user.User, error) {
	u, ok := ctx.Get(currentUser)
	if !ok {
		return nil, ErrUserNotFound
	}

	user, ok := u.(*user.User)
	if !ok {
		return nil, ErrUserIsInvalid
	}

	return user, nil
}
