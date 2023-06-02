package Controllers

import (
	"encoding/json"
	"fmt"
	"github.com/duolabmeng6/ecli/cmd"
	"github.com/duolabmeng6/goefun/ecore"
	"github.com/duolabmeng6/goefun/edb"
	"github.com/duolabmeng6/goefun/egin"
	"github.com/gin-gonic/gin"
	"go-gin-amis/app/dal"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/text/unicode/norm"
)

type AutoController struct {
	db *edb.Mysql数据库操作类
}

func (b *AutoController) Init() {
	b.db = dal.E数据库操作

}
func (b *AutoController) GetAllTableName(c *gin.Context) {
	info, err := b.db.GetAllTableName()
	fmt.Println(info, err)
	if err != nil {
		c.JSON(200, gin.H{
			"status": 1,
			"msg":    "查询表结构失败",
			"data":   "",
		})
		return
	}

	var options []gin.H
	for _, v := range info {
		options = append(options, gin.H{
			"label": v,
			"value": v,
		})
	}
	c.JSON(200, gin.H{
		"status": 0,
		"msg":    "所有表名称",
		"data": gin.H{
			//"label":   "选项",
			//"type":    "select",
			//"name":    "select",
			"options": options,
		},
	})
	return
}
func (b *AutoController) Get(c *gin.Context) {
	var req struct {
		TableName string `i:"table_name" rule:"required" msg:"table_name 必填"`
	}
	if err := egin.Verify(c, &req); err != nil {
		c.JSON(200, gin.H{
			"status": 1,
			"msg":    err.Error(),
			"data":   "",
		})
		return
	}
	tableName := req.TableName
	info, err := b.db.GetTableInfo(tableName)
	fmt.Println(info, err)
	if err != nil {
		c.JSON(200, gin.H{
			"status": 1,
			"msg":    "查询表结构失败",
			"data":   "",
		})
		return
	}
	tableCol := make([]gin.H, 0)
	for _, k := range info.Keys() {
		v, _ := info.Get(k)
		tableCol = append(tableCol, gin.H{
			"name":      k,
			"translate": cmd.E常见英文翻译函数(k),
			"type":      cmd.E常见字段对应的输入控件(v["dataType"].(string)),
			"dataType":  v["dataType"],
		})
	}
	//RootPath := "/Users/chensuilong/Desktop/goproject/go-gin-amis/"
	RootPath := ecore.E取运行源文件路径()
	RootPath = ecore.E文件取父目录(RootPath)
	RootPath = ecore.E文件取父目录(RootPath)
	RootPath = ecore.E文件取父目录(RootPath)

	c.JSON(200, gin.H{
		"status": 0,
		"msg":    "",
		"data": gin.H{
			"tableName":      req.TableName,
			"tableTranslate": cmd.E常见英文翻译函数(req.TableName),
			"ctlPath":        RootPath + "/app/Http/Controllers/" + tableName + "Controller.go",
			"dataPath":       RootPath + "/app/serv/" + tableName + "Serv.go",
			"curdPath":       RootPath + "/public/pages/" + tableName + "Page.json",
			"tableCol":       tableCol,
		},
	})
}

func (b *AutoController) Store(c *gin.Context) {
	var req struct {
		CtlPath  string `json:"ctlPath"`
		DataPath string `json:"dataPath"`
		CurdPath string `json:"curdPath"`
		TableCol []struct {
			DataType  string `json:"dataType"`
			Name      string `json:"name"`
			Translate string `json:"translate"`
			Type      string `json:"type"`
		} `json:"tableCol"`
		TableName      string `json:"tableName"`
		TableTranslate string `json:"tableTranslate"`
	}
	c.BindJSON(&req)
	//{"ctlPath":"/Users/chensuilong/Desktop/goproject/go-gin-amis/app/Http/Controllers","dataPath":"/Users/chensuilong/Desktop/goproject/go-gin-amis/app/serv","tableCol":[{"dataType":"varchat","name":"username","translate":"用户名","type":"input-text"},{"dataType":"varchat","name":"password","translate":"密码","type":"input-text"}],"tableName":"users","tableTranslate":"用户"}

	columnsTable := make([]cmd.Column, 0)
	columnsLook := make([]cmd.Column, 0)
	columnsEdit := make([]cmd.Column, 0)
	columnsAdd := make([]cmd.Column, 0)
	// 遍历 req.TableCol
	for _, v := range req.TableCol {
		column := cmd.Column{Name: v.Name, Label: v.Translate}

		column.Type = "text"
		columnsTable = append(columnsTable, column)

		column.Type = "static"
		columnsLook = append(columnsLook, column)

		column.Type = v.Type
		columnsEdit = append(columnsEdit, column)

		column.Type = v.Type
		columnsAdd = append(columnsAdd, column)

	}
	marshal, _ := json.MarshalIndent(columnsTable, "", "  ")
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

	data := cmd.H{
		"TableName": cases.Title(language.English).String(norm.NFC.String(req.TableName)),
		"tableName": req.TableName,
		"TableCol":  TableCol,
		"ColLook":   ColLook,
		"ColEdit":   ColEdit,
		"ColAdd":    ColAdd,
	}

	RootPath := ecore.E取运行源文件路径()
	RootPath = ecore.E文件取父目录(RootPath)
	RootPath = ecore.E文件取父目录(RootPath)
	RootPath = ecore.E文件取父目录(RootPath)
	tmplFilePath := RootPath + "/cmd/ectl/tpl/amis-curd.json"
	err := cmd.RenderTemplateToFile(tmplFilePath, data, req.CurdPath)
	if err != nil {
		c.JSON(200, gin.H{
			"status": 1,
			"msg":    "渲染模板并输出文件时发生错误:" + err.Error(),
		})
		return
	}

	tmplFilePath = RootPath + "/cmd/ectl/tpl/amis操作类.go"
	err = cmd.RenderTemplateToFile(tmplFilePath, data, req.DataPath)
	if err != nil {
		c.JSON(200, gin.H{
			"status": 1,
			"msg":    "渲染模板并输出文件时发生错误:" + err.Error(),
		})
		return
	}

	tmplFilePath = RootPath + "/cmd/ectl/tpl/amisController.go"
	err = cmd.RenderTemplateToFile(tmplFilePath, data, req.CtlPath)
	if err != nil {
		c.JSON(200, gin.H{
			"status": 1,
			"msg":    "渲染模板并输出文件时发生错误:" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 0,
		"msg":    "文件生成成功",
		"data":   req,
	})
}
