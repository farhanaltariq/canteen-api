package model

import "time"

type CanteenMenu struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	CanteenID int       `gorm:"primaryKey" json:"canteen_id"`
	MenuID    int       `gorm:"primaryKey" json:"menu_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CanteenMenuData struct {
	CanteenID int `json:"canteen_id"`
	MenuID    int `json:"menu_id"`
}

type CanteenMenuDetail struct {
	CanteenName string `json:"canteen_name"`
	Menu        []Menu `json:"menu"`
}
