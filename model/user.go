package model

import "time"

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Username  string    `json:"username" gorm:"type:varchar(255);not null unique"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserCredential struct {
	Username string `json:"username" gorm:"type:varchar(255);not null"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
}
