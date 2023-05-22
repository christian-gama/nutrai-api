package store

import (
	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"
	"github.com/gin-gonic/gin"
)

const currentUser = "currentUser"

// SetUser sets the current user in the Gin context.
func SetUser(ctx *gin.Context, user *user.User) {
	ctx.Set(currentUser, user)
}

// GetUser returns the current user from the Gin context.
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
