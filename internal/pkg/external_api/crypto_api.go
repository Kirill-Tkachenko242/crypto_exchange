package external_api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type CryptoAPIClient struct {
}

func NewCryptoAPIClient() *CryptoAPIClient {
	return &CryptoAPIClient{}
}

func (c *CryptoAPIClient) GetQuotes(symbols []string) ([]map[string]interface{}, error) {
	// Формируем параметры запроса
	ids := strings.ToLower(strings.Join(symbols, ","))
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=usd", ids)

	// Логируем URL
	fmt.Println("Request URL:", url)

	// Выполняем HTTP GET запрос
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Ошибка выполнения запроса: %v\n", err)
		return nil, fmt.Errorf("ошибка выполнения запроса: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Ошибка от CoinGecko API: статус %d\n", resp.StatusCode)
		return nil, fmt.Errorf("ошибка от CoinGecko API: статус %d", resp.StatusCode)
	}

	// Парсим JSON-ответ
	var result map[string]map[string]float64
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Printf("Ошибка парсинга JSON: %v\n", err)
		return nil, fmt.Errorf("ошибка парсинга JSON: %v", err)
	}

	// Логируем результат
	fmt.Printf("API Response: %+v\n", result)

	// Преобразуем результат в список карт
	var quotes []map[string]interface{}
	for symbol, data := range result {
		quotes = append(quotes, map[string]interface{}{
			"symbol": symbol,
			"price":  data["usd"],
		})
	}

	// Логируем конечный результат
	fmt.Printf("Parsed Quotes: %+v\n", quotes)

	return quotes, nil
}
