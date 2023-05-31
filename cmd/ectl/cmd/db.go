package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/duolabmeng6/goefun/edb"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/text/unicode/norm"
)

var dbCmd = &cobra.Command{
	Use:   "db:create",
	Short: "自动生成数据库对应的CURD模型以及amis的配置文件 使用方法 ecli db:create users",
	Long:  `现在开始吧`,
	Run: func(cmd *cobra.Command, args []string) {
		connectionString, _ := cmd.Flags().GetString("connection")

		if connectionString == "" {
			fmt.Println("请提供数据库连接字符串。")
			return
		}

		tableName, _ := cmd.Flags().GetString("table")

		if tableName == "" {
			fmt.Println("请提供表名称。")
			return
		}

		fmt.Println("数据库连接:", connectionString)
		fmt.Println("正在使用表创建数据库:", tableName)
		db := edb.NewMysql数据库操作类()
		err := db.E连接数据库(connectionString)
		if err != nil {
			fmt.Println("数据库连接失败。")
			return
		}

		info, err := db.GetTableInfo(tableName)
		if err != nil {
			fmt.Println("查询表结构失败。")
			return
		}
		//info=0 提示错误
		if len(info) == 0 {
			fmt.Println("查询表结构失败。")
			return
		}
		//    map[string]map[string]interface {}{
		//        "created_at": {
		//            "comment":  "创建时间",
		//            "dataType": "timestamp",
		//        },

		columns := make([]Column, 0)
		columnsLook := make([]Column, 0)
		columnsEdit := make([]Column, 0)
		columnsAdd := make([]Column, 0)
		for k, v := range info {
			字段名称 := k
			字段类型 := v["dataType"]
			字段备注 := v["comment"]
			fmt.Println(字段名称, 字段类型, 字段备注)
			column := Column{Name: 字段名称, Label: 字段名称, Type: "text"}
			columns = append(columns, column)

			column = Column{Name: 字段名称, Label: 字段名称, Type: "static"}
			columnsLook = append(columnsLook, column)

			column = Column{Name: 字段名称, Label: 字段名称, Type: "input-text"}
			columnsEdit = append(columnsEdit, column)

			column = Column{Name: 字段名称, Label: 字段名称, Type: "input-text"}
			columnsAdd = append(columnsAdd, column)

		}
		marshal, _ := json.MarshalIndent(columns, "", "  ")
		TableCol := string(marshal)
		TableCol = TableCol[1 : len(TableCol)-1]

		marshal, _ = json.MarshalIndent(columnsLook, "", "  ")
		ColLook := string(marshal)
		ColLook = ColLook[1 : len(ColLook)-1]

		marshal, _ = json.MarshalIndent(columnsEdit, "", "  ")
		ColEdit := string(marshal)
		ColEdit = ColEdit[1 : len(ColEdit)-1]

		marshal, _ = json.MarshalIndent(columnsAdd, "", "  ")
		ColAdd := string(marshal)
		ColAdd = ColAdd[1 : len(ColAdd)-1]

		tmplFilePath := "/Use" +
			"rs/chensuilong/Desktop/goproject/go-gin-amis/cmd/ectl/tpl/amis-curd.json"
		outputFilePath := "/Users/chensuilong/Desktop/goproject/go-gin-amis/public/pages/" + tableName + "Page.json"
		data := H{
			"TableName": cases.Title(language.English).String(norm.NFC.String(tableName)),
			"tableName": tableName,
			"TableCol":  TableCol,
			"ColLook":   ColLook,
			"ColEdit":   ColEdit,
			"ColAdd":    ColAdd,
		}
		err = RenderTemplateToFile(tmplFilePath, data, outputFilePath)
		if err != nil {
			fmt.Println("渲染模板并输出文件时发生错误:", err)
			return
		}

		//
		//tmplFilePath := "/Use" +
		//	"rs/chensuilong/Desktop/goproject/go-gin-amis/cmd/ectl/tpl/amis操作类.go"
		//outputFilePath := "/Users/chensuilong/Desktop/goproject/go-gin-amis/cmd/ectl/test/" + tableName + "Serv.go"
		//data := H{
		//	"TableName": cases.Title(language.English).String(norm.NFC.String(tableName)),
		//	"tableName": tableName,
		//}
		//err = RenderTemplateToFile(tmplFilePath, data, outputFilePath)
		//if err != nil {
		//	fmt.Println("渲染模板并输出文件时发生错误:", err)
		//	return
		//}
		//
		//tmplFilePath = "/Users/chensuilong/Desktop/goproject/go-gin-amis/cmd/ectl/tpl/amisController.go"
		//outputFilePath = "/Users/chensuilong/Desktop/goproject/go-gin-amis/cmd/ectl/test/" + tableName + "Controller.go"
		//data = H{
		//	"TableName": cases.Title(language.English).String(norm.NFC.String(tableName)),
		//	"tableName": tableName,
		//}
		//err = RenderTemplateToFile(tmplFilePath, data, outputFilePath)
		//if err != nil {
		//	fmt.Println("渲染模板并输出文件时发生错误:", err)
		//	return
		//}
		//fmt.Println("文件生成成功:", outputFilePath)

	},
}

func init() {
	rootCmd.AddCommand(dbCmd)

	dbCmd.Flags().StringP("connection", "c", "", "数据库连接字符串")
	dbCmd.Flags().StringP("table", "t", "", "表名称")
	dbCmd.MarkFlagRequired("connection")
	dbCmd.MarkFlagRequired("table")
}
