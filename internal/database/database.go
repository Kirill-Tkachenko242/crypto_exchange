package database

import (
	"fmt"

	"crypto_exchange/internal/config"
	"crypto_exchange/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *config.Config) (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Автоматическая миграция моделей
	err = db.AutoMigrate(&models.User{}, &models.Order{}, &models.Transaction{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
