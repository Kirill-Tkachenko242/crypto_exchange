package models

import (
	"time"
)

type User struct {
	ID        uint    `gorm:"primaryKey"`
	Name      string  `gorm:"size:100;not null"`
	Email     string  `gorm:"uniqueIndex;size:100;not null"`
	Password  string  `gorm:"not null"`
	Balance   float64 `gorm:"default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
