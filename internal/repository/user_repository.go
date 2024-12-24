package repository

import (
	"crypto_exchange/internal/models"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Если пользователь не найден
		}
		return nil, err // Другие ошибки
	}
	return &user, nil
}

func (r *UserRepository) UpdateBalance(userID uint, amount float64) error {
	return r.db.Model(&models.User{}).Where("id = ?", userID).Update("balance", amount).Error
}

func (r *UserRepository) GetByID(userID uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, userID).Error
	return &user, err
}
