package app

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitConfig() *viper.Viper {
	// get config
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath(".")
	err := config.ReadInConfig() // Find and read the config file
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return config
}
