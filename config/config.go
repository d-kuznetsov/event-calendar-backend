package config

import (
	"log"

	"github.com/spf13/viper"
)

type config struct {
	ServerUri  string `mapstructure:"SERVER_URI"`
	ClientUri  string `mapstructure:"CLIENT_URI"`
	DbUri      string `mapstructure:"DB_URI"`
	DbName     string `mapstructure:"DB_NAME"`
	SigningKey string `mapstructure:"SIGNING_KEY"`
}

var cfg config

func GetConfig() config {
	return cfg
}

func init() {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal(err)
	}
}
