package config

import (
	"github.com/spf13/viper"
	"log"
)

func init() {
	viper.SetDefault("port", 3000)
	viper.SetDefault("log_level", "info")
}

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err == nil {
		log.Printf("config file used %v", viper.ConfigFileUsed())
	}

	for _, key := range viper.AllKeys() {
		log.Printf("%v: %v", key, viper.Get(key))
	}
}

func GetPort() int {
	return viper.GetInt("port")
}

func GetLogLevel() string {
	return viper.GetString("log_level")
}
