package structs

type StockDetail struct {
	Symbol string `json:"symbol" validate:"required"`
	Name   string `json:"name" validate:"required"`
	Type   string `json:"type" validate:"required"`
}