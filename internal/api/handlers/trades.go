package handlers

import (
	"crypto_exchange/internal/config"
	"crypto_exchange/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TradeHandler struct {
	tradeService *service.TradeService
}

func NewTradeHandler(cfg *config.Config, db *gorm.DB) *TradeHandler {
	return &TradeHandler{
		tradeService: service.NewTradeService(cfg, db),
	}
}

func (h *TradeHandler) Buy(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	var request struct {
		Amount float64 `json:"amount"`
		Symbol string  `json:"symbol"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.tradeService.Buy(userID, request.Symbol, request.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Покупка успешно выполнена"})
}

func (h *TradeHandler) Sell(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	var request struct {
		Amount float64 `json:"amount"`
		Symbol string  `json:"symbol"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.tradeService.Sell(userID, request.Symbol, request.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Продажа успешно выполнена"})
}
