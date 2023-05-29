package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var E数据库连接字符串 = "root@tcp(127.0.0.1:3310)/gotest?charset=utf8&parseTime=true&loc=Local"

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "../../app/dal/query",
		ModelPkgPath: "../../app/dal/model",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	gormdb, _ := gorm.Open(mysql.Open(E数据库连接字符串))
	g.UseDB(gormdb) // reuse your gorm db

	//生成所有表的模型
	//g.GenerateAllTable()

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(g.GenerateAllTable()...)

	//g.ApplyBasic(
	//	g.GenerateModel("user_integral_record",
	//		gen.FieldType("created_at", `formatTime.WrapTime`),
	//	),
	//)

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	//g.ApplyInterface(func(Querier) {}, model.User{}, model.Company{})

	// Generate the code
	g.Execute()
}
