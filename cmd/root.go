package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use:   "goooo",
	Short: "Demo application",
	Long:  `Demo application`,
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		log.Fatalln("error executing command %w", err)
	}
}

func init() {
	cobra.OnInitialize(readConfigFile)
}

func readConfigFile() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("invalid %w", err)
	}

	log.Info("-----------")
	for _, key := range viper.AllKeys() {
		log.Infof("%v: %v", key, viper.Get(key))
	}
	log.Info("-----------")
}
