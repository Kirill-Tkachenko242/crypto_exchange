package api

import (
	"crypto_exchange/internal/api/handlers"
	"crypto_exchange/internal/api/middleware"
	"crypto_exchange/internal/config"
	"crypto_exchange/internal/repository"
	"crypto_exchange/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, cfg *config.Config, db *gorm.DB) {
	// Middleware
	router.Use(middleware.Logging())

	// Инициализация обработчиков
	authHandler := handlers.NewAuthHandler(cfg, db)
	tradeHandler := handlers.NewTradeHandler(cfg, db)
	quoteHandler := handlers.NewQuoteHandler(service.NewQuoteService())
	transactionHandler := handlers.NewTransactionHandler(
		service.NewTransactionService(repository.NewTransactionRepository(db)),
	)

	// Маршруты API
	api := router.Group("/api")
	{
		// Регистрация и вход
		api.POST("/register", authHandler.Register)
		api.POST("/login", authHandler.Login)

		// Защищённые маршруты
		secured := api.Group("/")
		secured.Use(middleware.AuthMiddleware(cfg.JWTSecretKey))
		{
			secured.GET("/quotes", quoteHandler.GetQuotes)
			secured.GET("/transactions", transactionHandler.GetTransactions)
			secured.POST("/buy", tradeHandler.Buy)
			secured.POST("/sell", tradeHandler.Sell)
		}
	}
}
