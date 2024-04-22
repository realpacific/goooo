package initializer

import (
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("port", 3000)
	viper.SetDefault("log_level", "info")
}

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		Logger.Fatalln("invalid %w", err)
	}

	Logger.Debugf("-----------")
	for _, key := range viper.AllKeys() {
		Logger.Debugf("%v: %v", key, viper.Get(key))
	}
	Logger.Debugf("-----------")
}
