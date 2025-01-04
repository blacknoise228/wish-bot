package config

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var configs = Config{}

func LoadConfigs(configPath string) {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix("FITAVATAR")
	viper.AutomaticEnv()

	if configPath != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configPath)
		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())

		} else {
			log.Fatal(err)
		}
	}

	if err := viper.Unmarshal(&configs); err != nil {
		log.Fatal(err)
	}

	buf := bytes.NewBuffer(nil)
	_ = yaml.NewEncoder(buf).Encode(configs)
	if configs.App.Environment == "dev" {
		fmt.Println("Effective configuration:")
		fmt.Println(buf.String())
	}

}

func GetConfigs() Config {
	return configs
}
