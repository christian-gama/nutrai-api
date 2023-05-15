package routing

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/router"
)

// Register registers all the routes for a group.
func Register(groupName string, routes ...*Routing) {
	group := router.Engine.Group(groupName)

	for _, route := range routes {
		route.Register(group)
	}
}
