/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"wish-bot/core/wishbot/config"
	"wish-bot/migration"

	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "run migrations",
	Long:  "run migrations",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrate called")

		cfg := config.GetConfigs()

		if cfg.Migrations.Migrate {
			fmt.Println("Migration: start")
			migration.RunDBMigrate(&cfg.Postgres)
		}

		fmt.Println("Done")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
