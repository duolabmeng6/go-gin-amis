package Controllers

import (
	"errors"
	"github.com/duolabmeng6/goefun/egin"
	"go-gin-amis/app/serv"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type {{ .TableName }}Controller struct {
	{{ .TableName }} serv.E{{ .TableName }}Serv
}

func (b *{{ .TableName }}Controller) Init() {
	b.{{ .TableName }} = serv.S{{ .TableName }}Serv
}

func (b *{{ .TableName }}Controller) Index(c *gin.Context, req *struct {
	Keywords string `i:"keywords"`
	PerPage  int64  `i:"perPage" rule:"required" msg:"PerPage 必填"`
	Page     int64  `i:"page" rule:"required" msg:"Page 必填"`
	OrderBy  string `i:"orderBy" default:"id"`
	OrderDir string `i:"orderDir" default:"desc"`
}) (gin.H, error) {

	articles, total, err := b.{{ .TableName }}.Index(req.Keywords, req.PerPage, req.Page, req.OrderBy, req.OrderDir)
	if err != nil {
		return nil, err
	}
	if total == 0 {
		return nil, errors.New("没有数据")
	}
	return gin.H{
		"items": articles,
		"total": total,
	}, nil
}

func (b *{{ .TableName }}Controller) Create(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "create",
	})
}
func (b *{{ .TableName }}Controller) Store(c *gin.Context) (gin.H, error) {
	articleData := egin.IAll(c)
	id, err := b.{{ .TableName }}.Insert(articleData)
	if err != nil {
		return nil, err
	}
	article, err := b.{{ .TableName }}.FindOne(id)
	if err != nil {
		return nil, err
	}
	return gin.H{
		"status": 0,
		"msg":    "",
		"data":   article,
	}, nil
}

func (b *{{ .TableName }}Controller) Show(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "show",
	})
}
func (b *{{ .TableName }}Controller) Edit(c *gin.Context, req *struct {
	Id int64 `i:"id" rule:"required" msg:"id 必填"`
}) (gin.H, error) {
	article, err := b.{{ .TableName }}.FindOne(req.Id)
	if err != nil {
		return nil, err
	}
	return gin.H{
		"status": 0,
		"msg":    "",
		"data":   article,
	}, nil
}
func (b *{{ .TableName }}Controller) Update(c *gin.Context) (gin.H, error) {
	articleData := egin.IAll(c)
	err := b.{{ .TableName }}.Update(articleData)
	if err != nil {
		return nil, err
	}
	return gin.H{
		"status": 0,
		"msg":    "更新成功",
		//"data":   article,
	}, nil
}
func (b *{{ .TableName }}Controller) Destroy(c *gin.Context, req *struct {
	Id int64 `i:"id" rule:"required" msg:"id 必填"`
}) (gin.H, error) {

	err := b.{{ .TableName }}.Delete(req.Id)
	if err != nil {
		return nil, err
	}
	return gin.H{
		"status": 0,
		"msg":    "删除成功",
		"data":   "",
	}, nil
}

func (b *{{ .TableName }}Controller) BulkDelete(c *gin.Context, req *struct {
	Ids string `i:"ids" rule:"required" msg:"ids 必填"`
}) (gin.H, error) {

	// 批量删除 ids 参数类似于 1,2,3 需要分割为,然后一个一个删除
	// 分割 ids
	idsArr := strings.Split(req.Ids, ",")
	// 遍历删除
	for _, id := range idsArr {
		// 删除
		idint, _ := strconv.ParseInt(id, 10, 64)
		err := b.{{ .TableName }}.Delete(idint)
		if err != nil {
			return nil, err
		}
	}
	return gin.H{
		"status": 0,
		"msg":    "删除成功",
		"data":   "",
	}, nil
}
