package main

import (
	"log"
	"net/http"

	"github.com/alimrndev/go-api/config"
	"github.com/alimrndev/go-api/router"
)

func main() {
	// Load konfigurasi dari file .env
	conf := config.LoadConfig()

	// Inisialisasi koneksi ke database
	config.InitDB(conf)
	defer config.CloseDB()

	// Lakukan hal lain sesuai kebutuhan aplikasi Anda
	// Inisialisasi router
	router := router.NewRouter()

	// Mulai server
	log.Println("Server started on", conf.BaseURL)
	log.Fatal(http.ListenAndServe(":8080", router))
}
