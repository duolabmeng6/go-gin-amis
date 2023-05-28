package Controllers

import (
	"go-gin-amis/app/Http/Requests"
	"go-gin-amis/app/Models"
	"go-gin-amis/config"
	"strconv"
	"strings"

	"github.com/duolabmeng6/goefun/edb"
	"github.com/duolabmeng6/goefun/egin"
	"github.com/gin-gonic/gin"
)

type ArticlesController struct {
	ArticlesService *Models.ArticlesService
	edb             edb.Mysql数据库操作类
}

func (b *ArticlesController) Init() {
	b.ArticlesService = Models.NewArticlesService(config.DB)
	b.edb = *edb.NewMysql数据库操作类()
	b.edb.E连接数据库(config.E数据库连接字符串)

}

func (b *ArticlesController) Index(c *gin.Context) {
	var req Requests.ArticlesIndexRequest
	if err := egin.Verify(c, &req); err != nil {
		c.JSON(500, gin.H{
			"messagex": err.Error(),
		})
		return
	}
	//query := "SELECT id,title,content,created_at,updated_at FROM articles" + edb.OrderBY(req.OrderBy, req.OrderDir, "title", "desc") + edb.LIMIT(req.Page, req.PerPage)

	dialect := edb.NewMySQLQueryBuilder()
	qb := dialect.From("articles").
		Select("id", "title", "content", "created_at", "updated_at").
		//Where("created_at", ">", "2020-01-01 00:00:00").
		OrderBy(req.OrderBy, req.OrderDir).
		Limit(int(req.Page), int(req.PerPage))
	if req.Keywords != "" {
		qb.Where("title", "like", "%"+req.Keywords+"%")
	}
	query, param := qb.ToSQL()
	println(query, param)

	result, err := b.edb.QueryRaw(query, param)

	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	query2, param2 := qb.Count().ToSQL()
	total, _ := b.edb.CountRaw(query2, param2)

	c.JSON(200, gin.H{
		"query": query,
		"items": result,
		"total": total,
	})
	// articles, total, err := b.ArticlesService.Index(req.Keywords, req.PerPage, req.Page, req.OrderBy, req.OrderDir)
	// if err != nil {
	// 	c.JSON(500, gin.H{
	// 		"messagex": err.Error(),
	// 	})
	// 	return
	// }
	// c.JSON(200, gin.H{
	// 	"items": articles,
	// 	"total": total,
	// })
}

func (b *ArticlesController) Create(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "create",
	})
}
func (b *ArticlesController) Store(c *gin.Context, req *Requests.ArticlesStoreRequest) (gin.H, error) {
	// 插入数据库
	articleData := new(Models.Articles)
	articleData.Title = req.Title
	articleData.Content = req.Content

	id, err := b.ArticlesService.Insert(articleData)
	if err != nil {
		return nil, err
	}
	article, err := b.ArticlesService.FindOne(id)
	if err != nil {
		return nil, err
	}
	return gin.H{
		"id":         article.Id,
		"title":      article.Title,
		"content":    article.Content,
		"created_at": article.CreatedAt.Format("2006-01-02 15:04:05"),
		"updated_at": article.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (b *ArticlesController) Show(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "show",
	})
}
func (b *ArticlesController) Edit(c *gin.Context, req *Requests.ArticlesIdRequest) (gin.H, error) {
	article, err := b.ArticlesService.FindOne(req.Id)
	if err != nil {
		return nil, err
	}
	return gin.H{
		"status": 0,
		"msg":    "",
		"data": gin.H{
			"id":         article.Id,
			"title":      article.Title,
			"content":    article.Content,
			"created_at": article.CreatedAt.Format("2006-01-02 15:04:05"),
			"updated_at": article.CreatedAt.Format("2006-01-02 15:04:05"),
		},
	}, nil
}
func (b *ArticlesController) Update(c *gin.Context, req *Requests.ArticlesUpdateRequest) (gin.H, error) {

	// 查询文章内容
	article, err := b.ArticlesService.FindOne(req.Id)
	if err != nil {
		return nil, err
	}

	article.Title = req.Title
	article.Content = req.Content

	// 更新文章内容
	err = b.ArticlesService.Update(article)
	if err != nil {
		return nil, err
	}

	return gin.H{
		"status": 0,
		"msg":    "更新成功",
		"data": gin.H{
			"id":         article.Id,
			"title":      article.Title,
			"content":    article.Content,
			"created_at": article.CreatedAt.Format("2006-01-02 15:04:05"),
			"updated_at": article.CreatedAt.Format("2006-01-02 15:04:05"),
		},
	}, nil
}
func (b *ArticlesController) Destroy(c *gin.Context, req *Requests.ArticlesIdRequest) (gin.H, error) {
	err := b.ArticlesService.Delete(req.Id)
	if err != nil {
		return nil, err
	}
	return gin.H{
		"status": 0,
		"msg":    "删除成功",
		"data":   "",
	}, nil
}

func (b *ArticlesController) BulkDelete(c *gin.Context, req *Requests.ArticlesIdsRequest) (gin.H, error) {
	// 批量删除 ids 参数类似于 1,2,3 需要分割为,然后一个一个删除
	// 分割 ids
	idsArr := strings.Split(req.Ids, ",")
	// 遍历删除
	for _, id := range idsArr {
		// 删除
		idint, _ := strconv.ParseInt(id, 10, 64)
		err := b.ArticlesService.Delete(idint)
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
