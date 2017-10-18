package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func setLogLevel(level string) {
	switch level {
	case "WARN":
		log.SetLevel(log.WarnLevel)
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "FATAL":
		log.SetLevel(log.FatalLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	}
}

func Load(configFile string) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)
	viper.ReadInConfig()
	setLogLevel(viper.GetString("app.logLevel"))
}
