package repository

import (
	"rest-api-go/internal/model"

	"gorm.io/gorm"
)

// interface: kontak
type FilmRepository interface {
	Create(film *model.Film) error
	GetAll() ([]model.Film, error)
	GetByID(id uint) (*model.Film, error)
	Update(film *model.Film) error
	Delete(id uint) error
}

// struct: implementasi nyata
type filmRepository struct {
	db *gorm.DB
}

// constructor: cara membuat instance

func NewFilmRepository(db *gorm.DB) FilmRepository {
	return &filmRepository{db: db}
}

// implementasi 5 method
func (r *filmRepository) Create(film *model.Film) error {
	return r.db.Create(film).Error
}

func (r *filmRepository) GetAll() ([]model.Film, error) {
	var films []model.Film
	err := r.db.Find(&films).Error
	return films, err
}

func (r *filmRepository) GetByID(id uint) (*model.Film, error) {
	var film model.Film
	err := r.db.First(&film, id).Error
	if err != nil {
		return nil, err
	}
	return &film, nil
}

func (r *filmRepository) Update(film *model.Film) error {
	return r.db.Save(film).Error
}

func (r *filmRepository) Delete(id uint) error {
	return r.db.Delete(&model.Film{}, id).Error
}
