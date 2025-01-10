package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTPPort    string
	DatabaseURI string
	RedisURI    string
	JWTSecret   string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		HTTPPort:    os.Getenv("HTTP_PORT"),
		DatabaseURI: os.Getenv("DB_URI"),
		RedisURI:    os.Getenv("REDIS_URI"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
	}
}
