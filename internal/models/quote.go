package models

type Quote struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
}
