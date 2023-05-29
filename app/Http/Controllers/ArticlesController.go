package Controllers

import (
	"errors"
	"github.com/duolabmeng6/goefun/egin"
	"go-gin-amis/app/serv"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ArticlesController struct {
	文章操作 serv.E文章操作
}

func (b *ArticlesController) Init() {
	b.文章操作 = serv.S文章操作
}

func (b *ArticlesController) Index(c *gin.Context, req *struct {
	Keywords string `i:"keywords"`
	PerPage  int64  `i:"perPage" rule:"required" msg:"PerPage 必填"`
	Page     int64  `i:"page" rule:"required" msg:"Page 必填"`
	OrderBy  string `i:"orderBy" default:"id"`
	OrderDir string `i:"orderDir" default:"desc"`
}) (gin.H, error) {

	articles, total, err := b.文章操作.Index(req.Keywords, req.PerPage, req.Page, req.OrderBy, req.OrderDir)
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

func (b *ArticlesController) Create(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "create",
	})
}
func (b *ArticlesController) Store(c *gin.Context) (gin.H, error) {
	articleData := egin.IAll(c)
	id, err := b.文章操作.Insert(articleData)
	if err != nil {
		return nil, err
	}
	article, err := b.文章操作.FindOne(id)
	if err != nil {
		return nil, err
	}
	return gin.H{
		"status": 0,
		"msg":    "",
		"data":   article,
	}, nil
}

func (b *ArticlesController) Show(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "show",
	})
}
func (b *ArticlesController) Edit(c *gin.Context, req *struct {
	Id int64 `i:"id" rule:"required" msg:"id 必填"`
}) (gin.H, error) {
	article, err := b.文章操作.FindOne(req.Id)
	if err != nil {
		return nil, err
	}
	return gin.H{
		"status": 0,
		"msg":    "",
		"data":   article,
	}, nil
}
func (b *ArticlesController) Update(c *gin.Context) (gin.H, error) {
	articleData := egin.IAll(c)
	err := b.文章操作.Update(articleData)
	if err != nil {
		return nil, err
	}
	return gin.H{
		"status": 0,
		"msg":    "更新成功",
		//"data":   article,
	}, nil
}
func (b *ArticlesController) Destroy(c *gin.Context, req *struct {
	Id int64 `i:"id" rule:"required" msg:"id 必填"`
}) (gin.H, error) {

	err := b.文章操作.Delete(req.Id)
	if err != nil {
		return nil, err
	}
	return gin.H{
		"status": 0,
		"msg":    "删除成功",
		"data":   "",
	}, nil
}

func (b *ArticlesController) BulkDelete(c *gin.Context, req *struct {
	Ids string `i:"ids" rule:"required" msg:"ids 必填"`
}) (gin.H, error) {

	// 批量删除 ids 参数类似于 1,2,3 需要分割为,然后一个一个删除
	// 分割 ids
	idsArr := strings.Split(req.Ids, ",")
	// 遍历删除
	for _, id := range idsArr {
		// 删除
		idint, _ := strconv.ParseInt(id, 10, 64)
		err := b.文章操作.Delete(idint)
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
