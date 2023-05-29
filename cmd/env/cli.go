package main

import (
	"fmt"
	"os"

	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/notify/domain/model/mail"
	value "github.com/christian-gama/nutrai-api/internal/notify/domain/value/mail"
	"github.com/christian-gama/nutrai-api/internal/notify/infra/mailer"
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
	cmd.PersistentFlags().StringVarP(&envFile, "env-file", "e", "", "environment file")
}

type AnStruct struct {
	A string
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

	err := mailer.MakeMailer().
		Send(mail.NewMail().
			SetSubject("Test").
			SetTo([]*value.To{{Email: "christiangsilva9@gmail.com", Name: "Christian"}}).
			SetTemplatePath(mailer.Welcome).
			SetContext("Olá!", ""),
		)
	if err != nil {
		fmt.Println(err)
	}
}

func checkEnvFile() {
	if envFile == "" {
		fmt.Println("Please specify an environment file with the flag -e")
		os.Exit(1)
	}

	if _, err := os.Stat(envFile); os.IsNotExist(err) {
		fmt.Printf("The file %s does not exist\n", envFile)
		os.Exit(1)
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
