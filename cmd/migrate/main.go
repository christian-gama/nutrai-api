package main

import (
	"fmt"

	"github.com/christian-gama/nutrai-api/internal/core/infra/bench"
	"github.com/christian-gama/nutrai-api/internal/core/infra/log"
)

func init() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	duration := bench.Duration(func() {
		cmd.Execute()
	})

	log.MakeWithCaller().Infof("\tFinished in %s", duration)
}
