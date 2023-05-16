package main

import (
	"context"
	"log"
	"os"

	"github.com/christian-gama/nutrai-api/internal"
	"github.com/christian-gama/nutrai-api/internal/core/infra/server"
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
	os.Stdout.Write([]byte("\033[H\033[2J"))
	cmd.PersistentFlags().StringVarP(&envFile, "env-file", "e", "", "environment file")
}

func run(cmd *cobra.Command, args []string) {
	checkEnvFile()
	internal.Bootstrap(envFile)
	server.Start(context.Background())
}

func checkEnvFile() {
	if envFile == "" {
		log.Fatalf("Please specify an environment file with the flag -e")
	}

	if _, err := os.Stat(envFile); os.IsNotExist(err) {
		log.Fatalf("The file %s does not exist", envFile)
	}
}
