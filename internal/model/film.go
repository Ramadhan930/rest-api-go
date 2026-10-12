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
	CreateAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdateAt  time.Time `gorm:"autoUpdateTime" json:"update_at"`
}
