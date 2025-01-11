/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"wish-bot/internal/config"
	"wish-bot/internal/db/migration"

	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrate called")

		cfg := config.GetConfigs()

		if cfg.Migrations.Migrate {
			migration.RunDBMigrate(&cfg.Postgres)
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
