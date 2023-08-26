package config

import "github.com/spf13/viper"

func Init() {
	config := viper.NewWithOptions()
	config.SetConfigType("yaml")
}
