package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Configuration struct {
	Port        string `mapstructure:"PORT"`
	DB_USER     string `mapstructure:"DB_USER"`
	DB_PASSWORD string `mapstructure:"DB_PASSWORD"`
	DB_NAME     string `mapstructure:"DB_NAME"`
	DB_HOST     string `mapstructure:"DB_HOST"`
}

var Default Configuration

func Init(confPath string) {
	c := Configuration{}
	viper.SetConfigFile(confPath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Warning(err, "Config file not found")
	}
	if err := viper.Unmarshal(&c); err != nil {
		log.Panic(err, "Error Unmarshal Viper Config")
	}

	Default = c
}
