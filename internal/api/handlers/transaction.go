package handlers

import (
	"crypto_exchange/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	transactionService *service.TransactionService
}

func NewTransactionHandler(transactionService *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService}
}

func (h *TransactionHandler) GetTransactions(c *gin.Context) {
	// Получаем ID пользователя из контекста (получено через middleware)
	userID := c.MustGet("userID").(uint)

	// Получаем историю транзакций через сервис
	transactions, err := h.transactionService.GetTransactionsByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить транзакции"})
		return
	}

	c.JSON(http.StatusOK, transactions)
}
