package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config adalah struct untuk konfigurasi aplikasi
type Config struct {
	BaseURL    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

// LoadConfig membaca variabel lingkungan dari file .env dan mengembalikan struct Config
func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return &Config{
		BaseURL:    os.Getenv("BASE_URL"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}
}
