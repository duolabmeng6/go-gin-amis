package dal

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var MySQLDSN = "root@tcp(127.0.0.1:3310)/gotest?charset=utf8&parseTime=true&loc=Local"

var DB *gorm.DB
var once sync.Once

func init() {
	fmt.Println("初始化数据库")
	once.Do(func() {
		DB = ConnectDB()
	})
}

func ConnectDB() (conn *gorm.DB) {
	var err error
	conn, err = gorm.Open(mysql.Open(MySQLDSN))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	return conn
}
