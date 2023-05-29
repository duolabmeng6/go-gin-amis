package serv

import (
	"github.com/duolabmeng6/goefun/edb"
	"go-gin-amis/app/dal"
)

var S文章操作 = New文章操作(dal.Edb)

type E文章操作 struct {
	edb   *edb.MySQLQueryBuilder
	table string
}

func New文章操作(edb *edb.MySQLQueryBuilder) E文章操作 {
	return E文章操作{
		edb:   edb,
		table: "articles",
	}
}

func (c E文章操作) Index(keywords string, perPage int64, page int64, orderBy string, orderDir string) ([]map[string]interface{}, int64, error) {
	db := c.edb.From(c.table).
		Select("*").
		OrderBy(orderBy, orderDir).
		Paginate(perPage, page)
	if keywords != "" {
		db = db.OrWhere("title", "like", "%"+keywords+"%").
			OrWhere("id", "=", keywords)
	}
	articles, err := db.Get()
	count, err := db.Count()
	return articles, count, err
}

// Insert
func (c E文章操作) Insert(articles map[string]interface{}) (int64, error) {
	id, err := c.edb.From(c.table).Insert(articles)
	return id, err
}

// FindOne
func (c *E文章操作) FindOne(id int64) (map[string]interface{}, error) {
	articles, err := c.edb.From(c.table).Where("id", "=", id).First()
	return articles, err
}

// Update
func (c *E文章操作) Update(articles map[string]interface{}) error {
	_, err := c.edb.From(c.table).
		Where("id", "=", articles["id"]).
		Update(articles)
	return err
}

// Delete
func (c *E文章操作) Delete(id int64) error {
	_, err := c.edb.From(c.table).Where("id", "=", id).Delete()
	return err
}
