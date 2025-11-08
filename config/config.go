package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseUrl    string
	Port           string
	AllowedOrigins string
	AllowedMethods string
	CloudinaryURL  string
}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func New() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env file can not be loaded")
	}

	return &Config{
		DatabaseUrl:    GetEnv("DATABASE_URL", "postgresql://root:root@localhost:5432/alpaca"),
		Port:           GetEnv("PORT", "8000"),
		AllowedOrigins: GetEnv("ALLOWED_ORIGINS", "*"),
		AllowedMethods: GetEnv("ALLOWED_METHODS", "POST,PUT,DELETE"),
		CloudinaryURL:  GetEnv("CLOUDINARY_URL", "cloudinary://"),
	}
}
