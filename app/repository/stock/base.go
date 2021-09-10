package stock

import (
	"sync"
	"test-package-02/app/model"
	"test-package-02/internal/database"
)

type BaseInterface interface {
	GetBySymbol(symbol string) (stock model.Stock, err error)
	GetByIds()
	GetStock() (stockDB []model.Stock)
	Create(stock map[string]interface{})
	Update()
	Delete()
}

type repo struct {
	DB database.DBInterface
}

var singleton *repo
var once sync.Once

func NewRepo() BaseInterface {
	once.Do(func() {
		singleton = &repo{
			DB: database.NewConnect(),
		}
	})
	return singleton
}