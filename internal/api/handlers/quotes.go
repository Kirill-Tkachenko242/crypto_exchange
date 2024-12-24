package handlers

import (
	"crypto_exchange/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QuoteHandler struct {
	quoteService *service.QuoteService
}

func NewQuoteHandler(quoteService *service.QuoteService) *QuoteHandler {
	return &QuoteHandler{quoteService}
}

func (h *QuoteHandler) GetQuotes(c *gin.Context) {
	// Получаем символы криптовалют из параметров запроса
	symbols := c.QueryArray("symbols") // Например, ?symbols=bitcoin&symbols=ethereum
	if len(symbols) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не указаны символы"})
		return
	}

	// Получаем котировки через сервис
	quotes, err := h.quoteService.GetQuotes(symbols)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить котировки"})
		return
	}

	c.JSON(http.StatusOK, quotes)
}
