package Controllers

import (
	"github.com/duolabmeng6/goefun/ecore"
	"github.com/duolabmeng6/goefun/egin"
	"github.com/gin-gonic/gin"
)

type AmisPagesController struct {
}

func (b *AmisPagesController) Index(c *gin.Context) {
	//查找目录 ./public/pages/*.json 文件返回文件列表
	var files []string
	ecore.E文件枚举("./public/pages/", ".json", &files, true, false)
	// 输出为 json  name content
	// name 为文件名
	// content 为文件内容
	var jsonList []map[string]string
	for _, value := range files {
		jsonList = append(jsonList, map[string]string{
			"name":   ecore.E文件取文件名(value, false),
			"config": ecore.E读入文本(value),
		})
	}
	c.JSON(200, gin.H{
		"status": 0,
		"msg":    "",
		"data": map[string]interface{}{
			"items": jsonList,
			"total": len(jsonList),
		},
	})
}

func (b *AmisPagesController) Create(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "create",
	})
}
func (b *AmisPagesController) Store(c *gin.Context) {
	name := egin.I(c, "json.name", "")
	config := egin.I(c, "json.config", "")
	if name == "" {
		c.JSON(200, gin.H{
			"status": 1,
			"msg":    "name不能为空",
			"data":   "",
		})
		return
	}
	//保存文件
	ecore.E写到文件("./public/pages/"+name+".json", ecore.E到字节集(config))
	c.JSON(200, gin.H{
		"status": 0,
		"msg":    "",
		"data":   "",
	})
}

func (b *AmisPagesController) Show(c *gin.Context) {
	pageName := egin.I(c, "id", "")
	if pageName == "" {
		c.JSON(200, gin.H{
			"status": 1,
			"msg":    "id不能为空",
			"data":   "",
		})
		return
	}

	c.HTML(200, "amis_show.html", gin.H{
		"pageTitle":     pageName,
		"pageSchemaApi": "GET:/pages/" + pageName + ".json",
		"getConfigAddr": "/pages/" + pageName + ".json",
	})
}
func (b *AmisPagesController) Edit(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "edit",
	})
}
func (b *AmisPagesController) Update(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "update",
	})
}
func (b *AmisPagesController) Destroy(c *gin.Context) {
	pageName := egin.I(c, "id", "")
	if pageName == "" {
		c.JSON(200, gin.H{
			"status": 1,
			"msg":    "name不能为空",
			"data":   "",
		})
		return
	}

	ecore.E删除文件("./public/pages/" + pageName + ".json")

	c.JSON(200, gin.H{
		"status": 0,
		"msg":    "删除成功",
		"data":   "",
	})
}
