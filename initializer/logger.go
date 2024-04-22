package initializer

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

var Logger = &logrus.Logger{
	Out:       os.Stderr,
	Formatter: new(logrus.TextFormatter),
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.DebugLevel,
}

func init() {
	Logger.SetOutput(os.Stdout)
	logLevel := getLogLevel()
	Logger.SetLevel(logLevel)
}

func getLogLevel() logrus.Level {
	config := viper.GetString("log_level")
	var logLevel logrus.Level
	switch config {
	case "warn":
		logLevel = logrus.WarnLevel
	case "error":
		logLevel = logrus.ErrorLevel
	case "debug":
		logLevel = logrus.DebugLevel
	case "info":
		logLevel = logrus.InfoLevel
	default:
		logLevel = logrus.InfoLevel
	}
	return logLevel
}
