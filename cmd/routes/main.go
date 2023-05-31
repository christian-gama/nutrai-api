package main

import (
	"fmt"
	"os"

	"github.com/christian-gama/nutrai-api/internal"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/router"
	"github.com/christian-gama/nutrai-api/pkg/screen"
	"github.com/fatih/color"
)

func init() {
	screen.Clear()

	os.Setenv("CONFIG_DEBUG", "false")
	os.Setenv("CONFIG_LOG_LEVEL", "panic")
}

func main() {
	internal.Bootstrap(".env.dev")
	printRoutes()
}

// printRoutes prints all routes in a colorized and formatted way.
func printRoutes() {
	for _, route := range router.Router.Routes() {
		fmt.Printf("%-12s\t%s\n", method(route.Method), path(route.Path))
	}
}

// method returns a colorized string based on the HTTP method.
func method(method string) string {
	methodsMap := map[http.Method]func(a ...any) string{
		http.MethodGet:    color.New(color.FgGreen, color.Bold).SprintFunc(),
		http.MethodPost:   color.New(color.FgYellow, color.Bold).SprintFunc(),
		http.MethodPut:    color.New(color.FgCyan, color.Bold).SprintFunc(),
		http.MethodDelete: color.New(color.FgRed, color.Bold).SprintFunc(),
		http.MethodPatch:  color.New(color.FgMagenta, color.Bold).SprintFunc(),
	}
	return methodsMap[http.Method(method)](method)
}

// path returns a colorized string.
func path(path string) string {
	return color.New(color.FgHiWhite).SprintFunc()(path)
}
