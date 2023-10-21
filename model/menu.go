package model

import "time"

type Menu struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name" gorm:"not null"`
	Price     int       `json:"price"`
	Stock     int       `json:"stock"`
	ImagePath string    `json:"image_path"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MenuData struct {
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Stock     int    `json:"stock"`
	ImagePath string `json:"image_path"`
}
