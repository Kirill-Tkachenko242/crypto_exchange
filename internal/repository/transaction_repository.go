package repository

import (
	"crypto_exchange/internal/models"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db}
}

func (r *TransactionRepository) Create(transaction *models.Transaction) error {
	// Получаем текущий баланс пользователя
	var user models.User
	if err := r.db.First(&user, transaction.UserID).Error; err != nil {
		return err
	}

	return r.db.Create(transaction).Error
}

func (r *TransactionRepository) GetByUserID(userID uint) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("User").Where("user_id = ?", userID).Order("created_at desc").Find(&transactions).Error
	return transactions, err
}
