package config

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB adalah variabel untuk koneksi database
var DB *gorm.DB

// InitDB berfungsi untuk inisialisasi koneksi ke database
func InitDB(cfg *Config) {
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPassword)
	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = db
}

// CloseDB berfungsi untuk menutup koneksi database
func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
