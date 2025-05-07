package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURL   string `json:"db_url"`
	APPPORT string `json:"app_port"`
	APPHOST string `json:"app_host"`
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		DBURL:   os.Getenv("DB_URL"),
		APPPORT: os.Getenv("APP_PORT"),
		APPHOST: os.Getenv("APP_HOST"),
	}
}
