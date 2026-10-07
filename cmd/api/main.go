package main

import (
	"fmt"
	"log"
	"rest-api-go/internal/config"
	"rest-api-go/pkg/database"
)

func main() {
	//load konfigurasi
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Gagal load config: ", err)
	}

	//koneksi ke database
	db := database.ConnectDB(cfg)
	_ = db // sementara, nanti akan digunakan

	fmt.Printf("server berjalan di port %s\n", cfg.AppPort)

}
