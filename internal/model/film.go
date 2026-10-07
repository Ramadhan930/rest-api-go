package model

import (
	"time"
)

type Film struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Judul     string    `gorm:"not null" json:"judul"`
	Sutradara string    `gorm:"not null" json:"sutradara"`
	Tahun     int       `gorm:"not null" json:"Tahun"`
	Genre     string    `gorm:"not null" json:"genre"`
	Rating    float64   `json:"rating"`
	CreateAt  time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}
