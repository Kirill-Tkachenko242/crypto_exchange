package service

import (
	"crypto_exchange/internal/config"
	"crypto_exchange/internal/models"
	"crypto_exchange/internal/repository"
	"errors"

	"gorm.io/gorm"
)

type TradeService struct {
	userRepo        *repository.UserRepository
	transactionRepo *repository.TransactionRepository
	quoteService    *QuoteService
	db              *gorm.DB
}

func NewTradeService(cfg *config.Config, db *gorm.DB) *TradeService {
	return &TradeService{
		userRepo:        repository.NewUserRepository(db),
		transactionRepo: repository.NewTransactionRepository(db),
		quoteService:    NewQuoteService(),
		db:              db,
	}
}

func (s *TradeService) Buy(userID uint, symbol string, amount float64) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// Получаем пользователя
		user, err := s.userRepo.GetByID(userID)
		if err != nil {
			return err
		}

		// Получаем текущую цену
		price, err := s.quoteService.GetPrice(symbol)
		if err != nil {
			return err
		}

		totalCost := amount * price

		// Проверяем баланс
		if user.Balance < totalCost {
			return errors.New("недостаточно средств")
		}

		// Обновляем баланс
		user.Balance -= totalCost
		err = s.userRepo.UpdateBalance(userID, user.Balance)
		if err != nil {
			return err
		}

		// Создаём транзакцию
		transaction := &models.Transaction{
			UserID: userID,
			Type:   "buy",
			Amount: totalCost,
		}
		err = s.transactionRepo.Create(transaction)
		if err != nil {
			return err
		}

		return nil
	})
}

func (s *TradeService) Sell(userID uint, symbol string, amount float64) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// Получаем пользователя
		user, err := s.userRepo.GetByID(userID)
		if err != nil {
			return err
		}

		// Получаем текущую цену
		price, err := s.quoteService.GetPrice(symbol)
		if err != nil {
			return err
		}
		totalGain := amount * price

		// Обновляем баланс
		user.Balance += totalGain
		err = s.userRepo.UpdateBalance(userID, user.Balance)
		if err != nil {
			return err
		}

		// Создаём транзакцию
		transaction := &models.Transaction{
			UserID: userID,
			Type:   "sell",
			Amount: totalGain,
		}
		err = s.transactionRepo.Create(transaction)
		if err != nil {
			return err
		}

		return nil
	})
}
