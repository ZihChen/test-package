package stock

import (
	"sync"
	"test-package-02/app/global/structs"
	"test-package-02/app/repository/stock"
)

type Interface interface {
	GetStockList() (list []structs.StockList)
	GetStockById()
	GetStocksByIds()
	CreateStock(stock *structs.StockDetail) (err error)
	UpdateStock()
	DeleteStock()
}

var singleton *buss
var once sync.Once

type buss struct {
	stock stock.BaseInterface
}

func NewBusiness() Interface{
	once.Do(func() {
		singleton = &buss{
			stock: stock.NewRepo(),
		}
	})
	return singleton
}
