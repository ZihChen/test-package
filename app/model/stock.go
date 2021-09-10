package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Stock struct {
	gorm.Model
	Symbol string `json:"symbol" gorm:"column:symbol;type:varchar(250)"`
	Name string `json:"name" gorm:"column:name;type:varchar(250)"`
	Type string `json:"type" gorm:"column:type;type:varchar(250)"`
	LatestRefresh time.Time `json:"latest_refresh" gorm:"column:latest_refresh;type:TIMESTAMP;default:CURRENT_TIMESTAMP"`
}