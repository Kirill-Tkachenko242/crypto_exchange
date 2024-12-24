package handlers

import (
	"crypto_exchange/internal/config"
	"crypto_exchange/internal/models"
	"crypto_exchange/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(cfg *config.Config, db *gorm.DB) *AuthHandler {
	return &AuthHandler{
		authService: service.NewAuthService(cfg, db),
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var user models.User

	// Привязка JSON к модели пользователя
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	log.Printf("Регистрация пользователя: %s", user.Email)

	// Вызов регистрации в сервисе
	err := h.authService.Register(&user)
	if err != nil {
		if err.Error() == "пользователь с таким email уже существует" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при регистрации пользователя"})
		log.Printf("Ошибка при регистрации пользователя: %v", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Пользователь успешно зарегистрирован"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var credentials models.LoginCredentials
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.authService.Login(credentials.Email, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
