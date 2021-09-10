package stock

import "test-package-02/app/model"

func (r *repo)Create(stock map[string]interface{}) {

	db, dbErr := r.DB.DBConnect()
	if dbErr != nil {
		return
	}

	transaction := db.Begin()

	stockDB := model.Stock{}

	if queryErr := transaction.Where("symbol = ?", stock["symbol"]).Assign(stock).FirstOrCreate(&stockDB).Error; queryErr != nil {
		transaction.Rollback()
		panic(queryErr.Error())
		return
	}

	transaction.Commit()
	return
}