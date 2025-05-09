package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() {
	// Load from .env first
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig() // Try to load .env
	if err != nil {
		log.Println("No .env file found, continuing...")
	}

	// Load from config.yaml
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// Read from config.yaml (optional)
	if err := viper.MergeInConfig(); err != nil {
		log.Println("config.yaml not found, continuing...")
	}

	viper.AutomaticEnv() // Allow env vars to override config
}
