package routing

import "github.com/gin-gonic/gin"

func Register(group *gin.RouterGroup, routes ...*Routing) {
	for _, route := range routes {
		route.Register(group)
	}
}
