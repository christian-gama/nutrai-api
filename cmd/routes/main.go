package main

import (
	"fmt"

	"github.com/christian-gama/nutrai-api/internal"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/env"
	"github.com/gin-gonic/gin"
)

func main() {
	env.Load(".env.dev")
	fmt.Print("\033[H\033[2J")
	fmt.Println("Listing all routes:")

	engine := internal.Routes()
	routes := engine.Routes()

	PrintRoutes(routes)
}

func PrintRoutes(routes gin.RoutesInfo) {
	for _, route := range routes {
		fmt.Printf("%-5s %s\n", route.Method, route.Path)
	}
}
