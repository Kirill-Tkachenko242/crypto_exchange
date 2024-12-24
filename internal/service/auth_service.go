package service

import (
	"crypto_exchange/internal/config"
	"crypto_exchange/internal/models"
	"crypto_exchange/internal/repository"
	"crypto_exchange/internal/utils"
	"errors"
	"log"

	"gorm.io/gorm"
)

type AuthService struct {
	userRepo  *repository.UserRepository
	jwtSecret string
}

func NewAuthService(cfg *config.Config, db *gorm.DB) *AuthService {
	return &AuthService{
		userRepo:  repository.NewUserRepository(db),
		jwtSecret: cfg.JWTSecretKey,
	}
}

func (s *AuthService) Register(user *models.User) error {
	log.Printf("Проверка пользователя с email: %s", user.Email)

	// Проверка на существующего пользователя
	existingUser, err := s.userRepo.GetByEmail(user.Email)
	if err != nil {
		log.Printf("Ошибка при проверке email: %v", err)
		return err
	}
	if existingUser != nil {
		log.Printf("Пользователь с email %s уже существует", user.Email)
		return errors.New("пользователь с таким email уже существует")
	}

	// Хэширование пароля
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Printf("Ошибка хэширования пароля: %v", err)
		return err
	}
	user.Password = hashedPassword

	// Сохранение пользователя в базе данных
	err = s.userRepo.Create(user)
	if err != nil {
		log.Printf("Ошибка при создании пользователя: %v", err)
		return err
	}

	log.Printf("Пользователь с email %s успешно создан", user.Email)
	return nil
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return "", errors.New("неверный email или пароль")
	}

	// Проверка пароля
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("неверный email или пароль")
	}

	// Генерация JWT
	token, err := utils.GenerateJWT(user.ID, s.jwtSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}
