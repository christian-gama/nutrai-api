package routing

import (
	"github.com/gin-gonic/gin"
)

// Register registers all the routes for a group.
func Register(group *gin.RouterGroup, routes ...*Routing) (*gin.RouterGroup, []*Routing) {
	for _, route := range routes {
		route.Register(group)
	}

	return group, routes
}
