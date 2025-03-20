package config

import (
	"log"

	"github.com/spf13/viper"
)

var Instance *Config

type Config struct {
	AppName       string
	Port          int
	RedisConfig   *RedisConfig
	PostgreConfig *PostgreConfig
}

type RedisConfig struct {
	Host     string
	Password string
	Port     int
	DB       int
}

type PostgreConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Schema   string
	Port     int
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Account service load default config")
		log.Println(err.Error())
		return loadDefaultConfig(), nil
	}
	var config *Config
	err = viper.Unmarshal(&config)
	return config, err
}

func loadDefaultConfig() *Config {
	return &Config{}
}
