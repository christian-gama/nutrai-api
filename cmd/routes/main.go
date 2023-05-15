package main

import (
	"fmt"

	"github.com/christian-gama/nutrai-api/internal"
	"github.com/christian-gama/nutrai-api/internal/core/infra/env"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func init() {
	env.Load(".env.dev")
	env.Config.Debug = false
}

func main() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Listing all routes:")

	engine := internal.LoadEngine()
	routes := engine.Routes()
	printRoutes(routes)
}

// printRoutes prints all routes in a colorized and formatted way.
func printRoutes(routes gin.RoutesInfo) {
	for _, route := range routes {
		fmt.Printf("%-2s\t%s\n", method(route.Method), path(route.Path))
	}
}

// method returns a colorized string based on the HTTP method.
func method(method string) string {
	methodsMap := map[http.Method]func(a ...any) string{
		http.MethodGet:    color.New(color.FgGreen, color.Bold).SprintFunc(),
		http.MethodPost:   color.New(color.FgMagenta, color.Bold).SprintFunc(),
		http.MethodPut:    color.New(color.FgCyan, color.Bold).SprintFunc(),
		http.MethodDelete: color.New(color.FgRed, color.Bold).SprintFunc(),
	}
	return methodsMap[http.Method(method)](method)
}

// path returns a colorized string.
func path(path string) string {
	return color.New(color.FgHiYellow).SprintFunc()(path)
}
