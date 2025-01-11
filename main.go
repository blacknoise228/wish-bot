/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"wish-bot/cmd"
	"wish-bot/internal/config"
)

func main() {
	cmd.Execute()
	config.LoadConfigs("./config/config.yaml")
}
