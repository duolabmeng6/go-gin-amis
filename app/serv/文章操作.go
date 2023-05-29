package serv

import (
	"github.com/duolabmeng6/goefun/ecore"
	"go-gin-amis/app/Http/Requests"
	"go-gin-amis/app/dal"
	"go-gin-amis/app/dal/model"
	"go-gin-amis/app/dal/query"
	"gorm.io/gorm"
)

var S文章操作 = New文章操作(dal.DB.Debug())

type E文章操作 struct {
	db *gorm.DB
	q  *query.Query
}

func New文章操作(db *gorm.DB) E文章操作 {
	return E文章操作{
		db: db,
		q:  query.Use(db),
	}
}

func (c E文章操作) Index(keywords string, perPage int64, page int64, OrderBy string, OrderDir string) (*[]Requests.ArticleItems, int64, error) {
	var articles []Requests.ArticleItems
	a := c.q.Article.
		Offset(int((page - 1) * perPage)).
		Limit(int(perPage))

	if keywords != "" {
		a = a.Or(c.q.Article.Title.Like("%" + keywords + "%"))
		a = a.Or(c.q.Article.ID.Eq(ecore.E到整数(keywords)))
	}
	if OrderBy == "id" && OrderDir == "desc" {
		a.Order(c.q.Article.ID.Desc())
	} else {
		a.Order(c.q.Article.ID)
	}

	if err := a.Scan(&articles); err != nil {
		return nil, 0, err
	}
	count, err := a.Count()

	return &articles, count, err
}
func (c E文章操作) Insert(articles *model.Article) (int64, error) {
	err := c.q.Article.Create(articles)
	return articles.ID, err
}

// FindOne
func (c *E文章操作) FindOne(id int64) (*model.Article, error) {

	articles, err := c.q.Article.Where(c.q.Article.ID.Eq(id)).First()
	return articles, err
}

// Update
func (c *E文章操作) Update(articles *model.Article) error {
	return c.q.Article.Save(articles)
}

// Delete
func (c *E文章操作) Delete(id int64) error {
	result, _ := c.q.Article.Where(c.q.Article.ID.Eq(id)).Delete()
	return result.Error
}
