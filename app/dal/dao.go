package dal

import (
	"fmt"
	"github.com/duolabmeng6/goefun/edb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var MySQLDSN = "root@tcp(127.0.0.1:3310)/gotest?charset=utf8&parseTime=true&loc=Local"

var DB *gorm.DB
var once sync.Once
var Edb *edb.MySQLQueryBuilder

func init() {
	fmt.Println("初始化数据库")
	once.Do(func() {
		DB = ConnectDB()
		数据库操作 := edb.NewMysql数据库操作类()
		数据库操作.E连接数据库(MySQLDSN)
		Edb = edb.NewMySQL查询构建器(数据库操作)
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
