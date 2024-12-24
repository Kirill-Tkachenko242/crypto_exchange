package models

import (
	"time"
)

type Transaction struct {
	ID        uint    `gorm:"primaryKey"`
	UserID    uint    `gorm:"not null"`
	Type      string  `gorm:"size:20;not null"` // "deposit", "withdraw", "trade"
	Amount    float64 `gorm:"not null"`
	CreatedAt time.Time
	User      User `gorm:"foreignKey:UserID"`
}
