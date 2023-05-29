package Controllers

import (
	"errors"
	"go-gin-amis/app/Http/Requests"
	"go-gin-amis/app/dal/model"
	"go-gin-amis/app/serv"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	用户操作 serv.E用户操作
}

func (b *UsersController) Init() {
	b.用户操作 = serv.S用户操作
}

func (b *UsersController) Index(c *gin.Context, req *Requests.UsersIndexRequest) (gin.H, error) {
	articles, total, err := b.用户操作.E查询用户列表(req.Keywords, req.Page, req.PerPage, req.OrderBy, req.OrderDir)
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

func (b *UsersController) Create(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "create",
	})
}
func (b *UsersController) Store(c *gin.Context, req *Requests.UsersStoreRequest) (gin.H, error) {
	// 插入数据库
	articleData := new(model.User)
	articleData.Username = req.Username
	articleData.Password = req.Password

	err := b.用户操作.E创建用户(articleData)
	if err != nil {
		return nil, err
	}

	return gin.H{
		"status": 0,
		"msg":    "",
		"data":   articleData,
	}, nil
}

func (b *UsersController) Show(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "show",
	})
}
func (b *UsersController) Edit(c *gin.Context, req *Requests.UsersIdRequest) (gin.H, error) {
	article, err := b.用户操作.FindOne(req.Id)
	if err != nil {
		return nil, err
	}
	return gin.H{
		"status": 0,
		"msg":    "",
		"data":   article,
	}, nil
}
func (b *UsersController) Update(c *gin.Context, req *Requests.UsersUpdateRequest) (gin.H, error) {

	// 查询文章内容
	article, err := b.用户操作.FindOne(req.Id)
	if err != nil {
		return nil, err
	}

	article.Username = req.Username
	article.Password = req.Password

	// 更新文章内容
	err = b.用户操作.Update(article)
	if err != nil {
		return nil, err
	}

	return gin.H{
		"status": 0,
		"msg":    "更新成功",
		"data":   article,
	}, nil
}
func (b *UsersController) Destroy(c *gin.Context, req *Requests.UsersIdRequest) (gin.H, error) {
	err := b.用户操作.Delete(req.Id)
	if err != nil {
		return nil, err
	}
	return gin.H{
		"status": 0,
		"msg":    "删除成功",
		"data":   "",
	}, nil
}

func (b *UsersController) BulkDelete(c *gin.Context, req *Requests.UsersIdsRequest) (gin.H, error) {
	// 批量删除 ids 参数类似于 1,2,3 需要分割为,然后一个一个删除
	// 分割 ids
	idsArr := strings.Split(req.Ids, ",")
	// 遍历删除
	for _, id := range idsArr {
		// 删除
		idint, _ := strconv.ParseInt(id, 10, 64)
		err := b.用户操作.Delete(idint)
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
