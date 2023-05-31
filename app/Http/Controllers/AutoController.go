package Controllers

import (
	"encoding/json"
	"fmt"
	"github.com/duolabmeng6/ecli/cmd"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/text/unicode/norm"
)

type AutoController struct {
}

func (b *AutoController) Init() {
}

func (b *AutoController) Get(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 0,
		"msg":    "",
		"data": gin.H{
			"tableName":      "users",
			"tableTranslate": "用户",
			"ctlPath":        "/Users/chensuilong/Desktop/goproject/go-gin-amis/app/Http/Controllers/usersController.go",
			"dataPath":       "/Users/chensuilong/Desktop/goproject/go-gin-amis/app/serv/usersServ.go",
			"curdPath":       "/Users/chensuilong/Desktop/goproject/go-gin-amis/public/pages/usersPage.json",
			"tableCol": []gin.H{
				gin.H{
					"name":      "username",
					"translate": "用户名",
					"type":      "input-text",
					"dataType":  "varchat",
				},
				gin.H{
					"name":      "password",
					"translate": "密码",
					"type":      "input-text",
					"dataType":  "varchat",
				},
			},
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
	tmplFilePath := "/Users/chensuilong/Desktop/goproject/go-gin-amis/cmd/ectl/tpl/amis-curd.json"
	err := cmd.RenderTemplateToFile(tmplFilePath, data, req.CurdPath)
	if err != nil {
		c.JSON(200, gin.H{
			"status": 1,
			"msg":    fmt.Println("渲染模板并输出文件时发生错误:", err),
		})
		return
	}

	tmplFilePath = "/Users/chensuilong/Desktop/goproject/go-gin-amis/cmd/ectl/tpl/amis操作类.go"
	err = cmd.RenderTemplateToFile(tmplFilePath, data, req.DataPath)
	if err != nil {
		c.JSON(200, gin.H{
			"status": 1,
			"msg":    fmt.Println("渲染模板并输出文件时发生错误:", err),
		})
		return
	}

	tmplFilePath = "/Users/chensuilong/Desktop/goproject/go-gin-amis/cmd/ectl/tpl/amisController.go"
	err = cmd.RenderTemplateToFile(tmplFilePath, data, req.CtlPath)
	if err != nil {
		c.JSON(200, gin.H{
			"status": 1,
			"msg":    fmt.Println("渲染模板并输出文件时发生错误:", err),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 0,
		"msg":    "文件生成成功",
		"data":   req,
	})
}
