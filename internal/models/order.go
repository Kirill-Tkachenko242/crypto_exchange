package models

import (
	"time"
)

type Order struct {
	ID        uint    `gorm:"primaryKey"`
	UserID    uint    `gorm:"not null"`
	Type      string  `gorm:"size:10;not null"` // "buy" or "sell"
	Amount    float64 `gorm:"not null"`
	Price     float64 `gorm:"not null"`
	CreatedAt time.Time
	User      User `gorm:"foreignKey:UserID"`
}
