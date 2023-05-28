package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var E数据库连接字符串 = "root@tcp(127.0.0.1:3310)/gotest?charset=utf8&parseTime=true&loc=Local"

func InitDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root@tcp(127.0.0.1:3310)/gotest?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	db.DB().SetConnMaxLifetime(10)
	db.DB().SetMaxIdleConns(10)
	DB = db
	return db
}
