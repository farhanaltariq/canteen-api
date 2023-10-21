package model

import "time"

type Canteen struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	CanteenName string    `json:"canteen_name" gorm:"not null"`
	Location    string    `json:"location"`
	Contact     string    `json:"contact"`
	ImagePath   string    `json:"image_path"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CanteenData struct {
	CanteenName string `json:"canteen_name"`
	Location    string `json:"location"`
	Contact     string `json:"contact"`
	ImagePath   string `json:"image_path"`
}

type FetchCanteenData struct {
	Businesses []Business `json:"businesses"`
}

type Business struct {
	Alias        string     `json:"alias"`
	Categories   []Category `json:"categories"`
	DisplayPhone string     `json:"display_phone"`
	ImageURL     string     `json:"image_url"`
	Location     Location   `json:"location"`
}

type Category struct {
	Alias string `json:"alias"`
	Title string `json:"title"`
}

type Location struct {
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	City     string `json:"city"`
}
