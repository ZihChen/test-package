package stock

import (
	"test-package-02/app/global/structs"
)

func (b *buss)GetStockList() (list []structs.StockList){
	stocks := b.stock.GetStock()

	//list = []structs.StockList{}

	for _, stock := range stocks {

		tmp := structs.StockList{}

		tmp.Symbol = stock.Symbol
		tmp.Name = stock.Name
		tmp.Type = stock.Type

		list = append(list, tmp)
	}
	return list
}

func (b *buss)GetStockById() {

}

func (b *buss)GetStocksByIds() {

}