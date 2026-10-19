package service

import (
	"errors"
	"strings"
	"time"

	"rest-api-go/internal/model"
	"rest-api-go/internal/repository"

	"gorm.io/gorm"
)

// Interface: kontrak
type FilmService interface {
	CreateFilm(film *model.Film) error
	GetAllFilms() ([]model.Film, error)
	GetFilmByID(id uint) (*model.Film, error)
	UpdateFilm(film *model.Film) error
	DeleteFilm(id uint) error
}

// Struct: implementasi
type filmService struct {
	repo repository.FilmRepository
}

// Constructor
func NewFilmService(repo repository.FilmRepository) FilmService {
	return &filmService{repo: repo}
}

// validasi: function private (huruf kecil)
func validateFilm(film *model.Film) error {
	// 1. Cek judul
	if strings.TrimSpace(film.Judul) == "" {
		return errors.New("judul tidak boleh kosong")
	}
	// 2. Cek sutradara
	if strings.TrimSpace(film.Sutradara) == "" {
		return errors.New("sutradara tidak boleh kosong")
	}
	// 3. Cek tahun
	tahunSekarang := time.Now().Year()
	if film.Tahun < 1888 || film.Tahun > tahunSekarang {
		return errors.New("tahun tidak valid")
	}
	// 4. Cek rating
	if film.Rating < 0 || film.Rating > 10 {
		return errors.New("rating harus antara 0 dan 10")
	}
	return nil
}

// CreateFilm: validasi dulu, baru simpan
func (s *filmService) CreateFilm(film *model.Film) error {
	if err := validateFilm(film); err != nil {
		return err
	}

	// cek apakah film sudah ada
	existingFilm, err := s.repo.GetByJudulAndTahun(film.Judul, film.Tahun)
	if err == nil && existingFilm != nil {
		return errors.New("film dengan judul dan tahun tersebut sudah ada")
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	return s.repo.Create(film)
}

// GetAllFilms: langsung ke repository
func (s *filmService) GetAllFilms() ([]model.Film, error) {
	return s.repo.GetAll()
}

// GetFilmByID: langsung ke repository
func (s *filmService) GetFilmByID(id uint) (*model.Film, error) {
	return s.repo.GetByID(id)
}

// UpdateFilm: validasi dulu, baru update
func (s *filmService) UpdateFilm(film *model.Film) error {
	if err := validateFilm(film); err != nil {
		return err
	}
	return s.repo.Update(film)
}

// DeleteFilm: langsung ke repository
func (s *filmService) DeleteFilm(id uint) error {
	return s.repo.Delete(id)
}
