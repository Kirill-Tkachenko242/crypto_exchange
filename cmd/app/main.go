package main

import (
	"log"

	"github.com/gin-contrib/cors" // Подключаем библиотеку CORS
	"github.com/gin-gonic/gin"

	"crypto_exchange/internal/api"
	"crypto_exchange/internal/config"
	"crypto_exchange/internal/database"
)

func main() {
	// Загрузка конфигураций
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	// Инициализация базы данных
	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Ошибка получения sql.DB: %v", err)
		}
		sqlDB.Close()
	}()

	// Инициализация маршрутизатора
	router := gin.Default()

	// Добавляем CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:80"}, // URL фронтенда
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Передаём маршрутизатор в API
	api.SetupRoutes(router, cfg, db)

	// Запуск сервера
	log.Printf("Сервер запущен на порту %s", cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
