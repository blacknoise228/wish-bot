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
	Use:   "bot",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("wishbot called")

		wishbot.WishBotService()
	},
}

func init() {
	rootCmd.AddCommand(botCmd)
}
