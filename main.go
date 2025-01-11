/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"wish-bot/cmd"
	shopcfg "wish-bot/core/shop/config"
	wishcfg "wish-bot/core/wishbot/config"
)

func main() {

	wishcfg.LoadConfigs("./config/wishbot.yaml")

	shopcfg.LoadConfigs("./config/shop.yaml")

	cmd.Execute()

}
