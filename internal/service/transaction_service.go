package service

import (
	"crypto_exchange/internal/models"
	"crypto_exchange/internal/repository"
)

type TransactionService struct {
	transactionRepo *repository.TransactionRepository
}

func NewTransactionService(transactionRepo *repository.TransactionRepository) *TransactionService {
	return &TransactionService{transactionRepo}
}

func (s *TransactionService) GetTransactionsByUser(userID uint) ([]models.Transaction, error) {
	return s.transactionRepo.GetByUserID(userID)
}
