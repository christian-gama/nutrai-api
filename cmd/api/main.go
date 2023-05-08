package main

import (
	"fmt"
	"net/http"

	"github.com/christian-gama/nutrai-api/internal/shared/infra/env"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/sql"
)

func main() {
	env.Load(".env.dev")

	sql.MakePostgres()
	fmt.Println("Server running at port 8080")

	http.ListenAndServe(":8080", nil)
}
