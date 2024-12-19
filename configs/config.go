package configs

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"dbname"`
	} `mapstructure:"database"`
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	JWT struct {
		Secret string `mapstructure:"secret"`
	} `mapstructure:"jwt"`
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	var configuration Config
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	if err := viper.Unmarshal(&configuration); err != nil {
		log.Fatalf("Unable to decode config: %v", err)
	}

	return &configuration
}
