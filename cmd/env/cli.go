package main

import (
	"fmt"
	"log"
	"os"

	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/pkg/reflection"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	envFile string
	cmd     = &cobra.Command{
		Use: "env",
		Run: run,
	}
)

func init() {
	os.Stdout.Write([]byte("\033[H\033[2J"))

	cmd.PersistentFlags().
		StringVarP(&envFile, "env-file", "e", "", "Path to environment config file")
	cmd.MarkPersistentFlagRequired("env-file")
}

func run(cmd *cobra.Command, args []string) {
	checkEnvFile()
	env.NewLoader(envFile).Load()

	for envName, envStruct := range envsMap {
		fmt.Println(color.New(color.FgHiMagenta).Sprintf("\nENVIRONMENT: %s", envName))
		reflection.IterateStructFields(envStruct, func(opts *reflection.FieldIterationOptions) {
			fmt.Printf(
				"%v: %v\n",
				color.New(color.FgGreen).Sprintf(opts.FieldName),
				color.New(color.FgYellow).Sprintf(fmt.Sprint(opts.Field.Interface())),
			)
		})
	}
}

func checkEnvFile() {
	if envFile == "" {
		log.Fatal("Please specify an environment file with the flag -e")
	}

	if _, err := os.Stat(envFile); os.IsNotExist(err) {
		log.Fatalf("The file %s does not exist", envFile)
	}
}

var envsMap = map[string]any{
	"APP":      env.App,
	"DB":       env.DB,
	"JWT":      env.Jwt,
	"RABBITMQ": env.RabbitMQ,
	"MAILER":   env.Mailer,
	"MAILTRAP": env.Mailtrap,
	"SENDGRID": env.Sendgrid,
	"REDIS":    env.Redis,
}
