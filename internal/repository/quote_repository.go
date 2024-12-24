package repository

import (
	"crypto_exchange/internal/models"

	"gorm.io/gorm"
)

type QuoteRepository struct {
	db *gorm.DB
}

func NewQuoteRepository(db *gorm.DB) *QuoteRepository {
	return &QuoteRepository{db}
}

func (r *QuoteRepository) SaveQuote(quote *models.Quote) error {
	return r.db.Create(quote).Error
}

func (r *QuoteRepository) GetQuote(symbol string) (*models.Quote, error) {
	var quote models.Quote
	err := r.db.Where("symbol = ?", symbol).First(&quote).Error
	return &quote, err
}

func (r *QuoteRepository) GetAllQuotes() ([]models.Quote, error) {
	var quotes []models.Quote
	err := r.db.Find(&quotes).Error
	return quotes, err
}
