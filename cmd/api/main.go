package main

import (
	"fmt"
	"log"
	"rest-api-go/internal/config"
	"rest-api-go/internal/model"
	"rest-api-go/internal/repository"
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
	filmRepo := repository.NewFilmRepository(db)

	fmt.Printf("server berjalan di port %s\n", cfg.AppPort)

	// 1. create
	filmBaru := &model.Film{
		Judul:     "interstellar",
		Sutradara: "Chirstoper Nolan",
		Tahun:     2014,
		Genre:     "sci-fi",
		Rating:    9.0,
	}

	if err := filmRepo.Create(filmBaru); err != nil {
		log.Fatal("gagal create: ", err)
	}
	fmt.Println("film ditambahkan, ID: ", filmBaru.ID)

	// 2. GetAll
	films, err := filmRepo.GetAll()
	if err != nil {
		log.Fatal("gagal getAll:", err)
	}
	fmt.Println("daftar film:")
	for _, f := range films {
		fmt.Printf(" [%d] %s (%d) -%s | rating: %.1f\n", f.ID, f.Judul, f.Tahun, f.Sutradara, f.Rating)
	}

	// 3. GetByID
	film, err := filmRepo.GetByID(filmBaru.ID)
	if err != nil {
		log.Fatal("gagal getByID: ", err)
	}
	fmt.Println("ditemukan: ", film.Judul)

	// 4. Update
	film.Rating = 9.5
	if err := filmRepo.Update(film); err != nil {
		log.Fatal("gagal update:", err)
	}
	fmt.Println("Rating diupdate jadi:", film.Rating)

}
