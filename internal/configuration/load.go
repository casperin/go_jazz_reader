package configuration

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("port", "8000")
	viper.SetDefault("db_user", "postgres")
	viper.SetDefault("db_pass", "postgres")

	// Load viper
	viper.AddConfigPath("$HOME/.jazz_reader/")
	viper.AutomaticEnv()

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	viper.WatchConfig()
}
