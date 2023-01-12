package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var d *gorm.DB

func Connect() {
	dsn := "root:ritvik@tcp(127.0.0.1:3306)/hello?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	d = db
}

func GetDB() *gorm.DB {
	return d
}
