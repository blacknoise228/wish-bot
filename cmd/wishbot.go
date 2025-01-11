/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"wish-bot/internal/api/telegram"
	"wish-bot/internal/config"
	db "wish-bot/internal/db/sqlc"
	"wish-bot/internal/service"

	"github.com/spf13/cobra"
)

// wishbotCmd represents the wishbot command
var wishbotCmd = &cobra.Command{
	Use:   "wishbot",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("wishbot called")

		ctx := context.Background()

		cfg := config.GetConfigs()

		db := db.InitDB(ctx, cfg.Postgres)

		services := service.NewServices(db.Queries, &cfg)

		tg := telegram.NewTelegram(&cfg, services)

		tg.StartBot(ctx)

	},
}

func init() {
	rootCmd.AddCommand(wishbotCmd)
}
