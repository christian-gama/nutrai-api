package main

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/bench"
	"github.com/christian-gama/nutrai-api/internal/core/infra/log"
	"github.com/christian-gama/nutrai-api/pkg/screen"
)

func init() {
	screen.Clear()
}

func main() {
	duration := bench.Duration(func() {
		cmd.Execute()
	})

	log.Infof("\tFinished in %s", duration)
}
