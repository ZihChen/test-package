package stock

import "test-package-02/app/business/stock"

type Handler struct {
	buss stock.Interface
}

func NewHandler() *Handler {
	return &Handler{
		buss: stock.NewBusiness(),
	}
}