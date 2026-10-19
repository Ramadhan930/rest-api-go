package main

import (
	"log"

	"rest-api-go/internal/config"
	"rest-api-go/internal/handler"
	"rest-api-go/internal/repository"
	"rest-api-go/internal/router"
	"rest-api-go/internal/service"
	"rest-api-go/pkg/database"
)

func main() {
	// 1. Load Config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Gagal load config:", err)
	}

	// 2. Koneksi Database
	db := database.ConnectDB(cfg)

	// 3. Inisialisasi Layer (Repository -> Service -> Handler)
	filmRepo := repository.NewFilmRepository(db)
	filmService := service.NewFilmService(filmRepo)
	filmHandler := handler.NewFilmHandler(filmService)

	// 4. Setup Router
	r := router.SetupRouter(filmHandler)

	// 5. Jalankan Server
	log.Printf("Server berjalan di http://localhost:%s", cfg.AppPort)
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatal("Gagal menjalankan server:", err)
	}
}
