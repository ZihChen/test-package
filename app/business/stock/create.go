package stock

import (
	"encoding/json"
	"test-package-02/app/global/structs"
)

func (b *buss)CreateStock(stock *structs.StockDetail) (err error) {
	stockDB, err := b.stock.GetBySymbol(stock.Symbol)

	if err != nil {
		return
	}

	if stockDB.Symbol != "" {
		return
	}

	tmp := structs.StockDetail{}
	tmp.Symbol = stock.Symbol
	tmp.Name = stock.Name
	tmp.Type = stock.Type

	myMap := make(map[string]interface{})

	byteData, err := json.Marshal(tmp)

	json.Unmarshal(byteData, &myMap)

	b.stock.Create(myMap)

	return
}