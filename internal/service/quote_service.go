package service

import (
	"crypto_exchange/internal/pkg/external_api"
	"errors"
)

type QuoteService struct {
	apiClient *external_api.CryptoAPIClient
}

func NewQuoteService() *QuoteService {
	return &QuoteService{
		apiClient: external_api.NewCryptoAPIClient(),
	}
}

func (s *QuoteService) GetQuotes(symbols []string) ([]map[string]interface{}, error) {
	return s.apiClient.GetQuotes(symbols)
}

func (s *QuoteService) GetPrice(symbol string) (float64, error) {
	quotes, err := s.apiClient.GetQuotes([]string{symbol})
	if err != nil {
		return 0, err
	}

	if len(quotes) == 0 {
		return 0, errors.New("не удалось получить цену")
	}

	price := quotes[0]["price"].(float64)
	return price, nil
}
