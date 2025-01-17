/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"wish-bot/core/wishbot"

	"github.com/spf13/cobra"
)

// botCmd represents the wishbot command
var botCmd = &cobra.Command{
	Use:   "wishbot",
	Short: "start wish telegram bot for users",
	Long:  "start wish telegram bot for users",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("wishbot called")

		wishbot.WishBotService()
	},
}

func init() {
	rootCmd.AddCommand(botCmd)
}
