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
	Short: "start shop telegram bot for managers",
	Long:  "start shop telegram bot for managers",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("shop called")

		shop.ShopBotService()
	},
}

func init() {
	rootCmd.AddCommand(shopCmd)
}
