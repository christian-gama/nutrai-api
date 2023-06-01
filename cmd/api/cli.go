package main

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal"
	httpserver "github.com/christian-gama/nutrai-api/internal/core/infra/http/server"
	"github.com/spf13/cobra"
)

var (
	envFile string
	cmd     = &cobra.Command{
		Use: "api",
		Run: run,
	}
)

func init() {
	cmd.PersistentFlags().StringVarP(&envFile, "env-file", "e", "", "environment file")
	cmd.MarkPersistentFlagRequired("env-file")
}

func run(cmd *cobra.Command, args []string) {
	internal.Bootstrap(envFile)
	httpserver.Start(context.Background())
}
