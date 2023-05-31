package main

import (
	"strconv"

	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/core/infra/migrate"
	"github.com/spf13/cobra"
)

var envFile string

func init() {
	cmd.PersistentFlags().
		StringVarP(&envFile, "env-file", "e", "", "Path to environment config file")
	cmd.MarkPersistentFlagRequired("env-file")

	cmd.AddCommand(upCmd)
	cmd.AddCommand(dropCmd)
	cmd.AddCommand(forceCmd)
	cmd.AddCommand(downCmd)
	cmd.AddCommand(versionCmd)
	cmd.AddCommand(stepsCmd)
	cmd.AddCommand(resetCmd)
}

var cmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Long:  "Run database migrations using the up and down commands",
}

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Run migrations up",
	Long:  "Run all available migrations to bring the database schema to the latest version",
	Run: func(cmd *cobra.Command, args []string) {
		env.NewLoader(envFile).Load()
		migrate.MakeMigrate().Up()
	},
}

var dropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Drop all tables",
	Long:  "Drop all tables in the database, use with caution",
	Run: func(cmd *cobra.Command, args []string) {
		env.NewLoader(envFile).Load()
		migrate.MakeMigrate().Drop()
	},
}

var forceCmd = &cobra.Command{
	Use:        "force",
	Short:      "Force a specific migration version",
	Long:       "Force a specific migration version",
	Args:       cobra.ExactArgs(1),
	ArgAliases: []string{"version"},
	Run: func(cmd *cobra.Command, args []string) {
		env.NewLoader(envFile).Load()

		versionStr := args[0]
		version, err := strconv.Atoi(versionStr)
		if err != nil {
			panic(err)
		}

		migrate.MakeMigrate().Force(version)
	},
}

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Run migrations down",
	Long:  "Run all available migrations to bring the database schema to the previous version",
	Run: func(cmd *cobra.Command, args []string) {
		env.NewLoader(envFile).Load()
		migrate.MakeMigrate().Down()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the current migration version",
	Long:  "Version returns the currently active migration version. Return an error if no version is set.",
	Run: func(cmd *cobra.Command, args []string) {
		env.NewLoader(envFile).Load()
		migrate.MakeMigrate().Version()
	},
}

var stepsCmd = &cobra.Command{
	Use:   "steps",
	Short: "Migrate up or down based on the step.",
	Long:  "Looks at the currently active migration version. Migrates up if n > 0, and down if n < 0.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		env.NewLoader(envFile).Load()

		stepsStr := args[0]
		steps, err := strconv.Atoi(stepsStr)
		if err != nil {
			panic(err)
		}

		migrate.MakeMigrate().Steps(steps)
	},
}

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset the database",
	Long:  "Reset the database by downing all migrations and then running them all again",
	Run: func(cmd *cobra.Command, args []string) {
		env.NewLoader(envFile).Load()
		migrate.MakeMigrate().Reset()
	},
}
