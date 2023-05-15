package main

import (
	"context"
	"os"

	"github.com/christian-gama/nutrai-api/internal"
	l "github.com/christian-gama/nutrai-api/internal/core/infra/log"
	"github.com/spf13/cobra"
)

var (
	envFile string
	log     = l.MakeLogWithCaller(1)
	cmd     = &cobra.Command{
		Use: "api",
		Run: run,
	}
)

func init() {
	os.Stdout.Write([]byte("\033[H\033[2J"))
	cmd.PersistentFlags().StringVarP(&envFile, "env-file", "e", "", "environment file")
}

func run(cmd *cobra.Command, args []string) {
	checkEnvFile()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	internal.Bootstrap(ctx, log, envFile)
}

func checkEnvFile() {
	if envFile == "" {
		log.Fatalf("Please specify an environment file with the flag -e")
	}

	if _, err := os.Stat(envFile); os.IsNotExist(err) {
		log.Fatalf("The file %s does not exist", envFile)
	}
}
