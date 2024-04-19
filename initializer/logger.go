package initializer

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
	logLevel := getLogLevel()
	log.SetLevel(logLevel)
}

func getLogLevel() log.Level {
	config := viper.GetString("log_level")
	var logLevel log.Level
	switch config {
	case "warn":
		logLevel = log.WarnLevel
	case "error":
		logLevel = log.ErrorLevel
	case "debug":
		logLevel = log.DebugLevel
	case "info":
		logLevel = log.InfoLevel
	default:
		logLevel = log.InfoLevel
	}
	return logLevel
}
