package database

import (
	"fmt"
	"log"
	"rest-api-go/internal/config"
	"rest-api-go/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cfg *config.Config) *gorm.DB {
	//koneksi ke database
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		log.Fatal("gagal koneksi ke database")
	}

	fmt.Println("berhasil koneksi ke database")

	//AutoMigrate: buat tabel otomatis dari struct
	err = db.AutoMigrate(&model.Film{})
	if err != nil {
		log.Fatal("gagal migrasi database")
	}

	fmt.Println("migrasi database berhasil")

	return db
}
