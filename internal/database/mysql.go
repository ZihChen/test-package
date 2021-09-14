package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"test-package-02/app/global"
	"test-package-02/app/model"
	"time"
)

var connectPool *gorm.DB

type dbCon struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type DBInterface interface {
	DBConnect() (*gorm.DB, error)
	CheckTableIsExist()
}

func NewConnect() DBInterface {
	return &dbCon{}
}

func (d *dbCon) DBConnect () (*gorm.DB, error){
	var err error

	if connectPool != nil {
		return connectPool, nil
	}

	dsn := composeConfString()
	connectPool, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	sqlPool, err:= connectPool.DB()

	if err != nil {
		panic("failed to connect database")
	}

	// 限制最大開啟的連線數
	sqlPool.SetMaxIdleConns(100)
	// 限制最大閒置連線數
	sqlPool.SetMaxOpenConns(2000)
	// 空閒連線 timeout 時間
	sqlPool.SetConnMaxLifetime(15 * time.Second)

	return connectPool, err
}

func (d *dbCon) CheckTableIsExist () {
	db, errApi := d.DBConnect()

	if errApi != nil {
		panic("failed to connect database")
	}

	err := db.AutoMigrate(&model.Stock{})

	if err != nil {
		panic(err.Error())
	}
}

func composeConfString() string {
	Host := global.Config.Host
	Username := global.Config.Username
	Password := global.Config.Password
	Database := global.Config.Database

	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?timeout=5s&charset=utf8mb4&parseTime=True&loc=Local", Username, Password, Host, Database)
}