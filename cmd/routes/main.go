package main

import (
	"fmt"

	"github.com/christian-gama/nutrai-api/internal"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/env"
)

func main() {
	env.Load(".env.dev")
	fmt.Print("\033[H\033[2J")
	fmt.Println("Listing all routes:")

	PrintApiRoutes()
}

func PrintApiRoutes() {
	_, apiGroup, apiRoutes := internal.ApiRoutes()
	for _, route := range apiRoutes {
		route.Print(apiGroup)
	}
}
