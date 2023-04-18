package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DBPort     string `mapstructure:"DB_PORT"`
	DBUsername string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
}

func GetConfig() (cfg *Config) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("error reading config", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal(err)
	}

	return cfg
}
