package main

import (
	"context"
	"log"

	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/core/infra/bench"
	"github.com/christian-gama/nutrai-api/internal/notify/domain/model/mail"
	value "github.com/christian-gama/nutrai-api/internal/notify/domain/value/mail"
	"github.com/christian-gama/nutrai-api/internal/notify/infra/mailer"
	"github.com/spf13/cobra"
)

var (
	envFile string
	to      string
	subject string
	name    string
	body    string
	cmd     = &cobra.Command{
		Use: "mail",
		Run: run,
	}
)

func init() {
	cmd.PersistentFlags().
		StringVarP(&envFile, "env-file", "e", "", "Path to environment config file")
	cmd.MarkPersistentFlagRequired("env-file")

	cmd.PersistentFlags().StringVarP(&to, "to", "t", "", "Recipient address")
	cmd.MarkPersistentFlagRequired("to")

	cmd.PersistentFlags().StringVarP(&subject, "subject", "s", "", "Email subject line")
	cmd.MarkPersistentFlagRequired("subject")

	cmd.PersistentFlags().StringVarP(&body, "body", "b", "", "Content for email body")
	cmd.MarkPersistentFlagRequired("body")

	cmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Recipient user's name")
	cmd.MarkPersistentFlagRequired("name")
}

func run(cmd *cobra.Command, args []string) {
	env.NewLoader(envFile).Load()

	mail, err := mail.NewMail().
		SetSubject(subject).
		SetTo(&value.To{Email: to, Name: name}).
		SetTemplatePath(mail.WelcomeTemplate).
		SetContext(value.Context{"Name": name}).
		Validate()
	if err != nil {
		log.Fatal(err)
	}

	duration := bench.Duration(func() {
		err := mailer.MakeMailer().Send(context.Background(), mail)
		if err != nil {
			log.Fatal(err)
		}
	})

	log.Printf("[%s] - Sent email to '%s' in %s", env.Mailer.Provider, to, duration)
}
