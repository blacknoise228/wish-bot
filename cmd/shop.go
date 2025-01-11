/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"wish-bot/core/shop"

	"github.com/spf13/cobra"
)

// shopCmd represents the shop command
var shopCmd = &cobra.Command{
	Use:   "shop",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("shop called")

		shop.ShopBotService()
	},
}

func init() {
	rootCmd.AddCommand(shopCmd)
}
