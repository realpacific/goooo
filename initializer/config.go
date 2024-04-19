package initializer

import "github.com/spf13/viper"

func init() {
	viper.SetDefault("port", 3000)
	viper.SetDefault("log_level", "info")
}
