package stock

import (
	"test-package-02/app/model"
)

func (r *repo)GetStock() (stockDB []model.Stock) {
	db, dbErr := r.DB.DBConnect()

	if dbErr != nil {
		return
	}

	if err := db.Find(&stockDB).Error; err != nil {

		return
	}

	return
}

func (r *repo)GetBySymbol(symbol string) (stock model.Stock, apiErr error){
	db, apiErr := r.DB.DBConnect()
	if apiErr != nil {
		return
	}

	if err := db.Where("symbol = ?", symbol).Find(&stock).Error; err != nil {
		return
	}

	return
}

func (r *repo)GetByIds() {

}