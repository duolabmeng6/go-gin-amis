package serv

import (
	"github.com/duolabmeng6/goefun/edb"
	"go-gin-amis/app/dal"
)

var SUsersServ = NewUsersServ(dal.Edb)

type EUsersServ struct {
	edb   *edb.MySQLQueryBuilder
	table string
}

func NewUsersServ(edb *edb.MySQLQueryBuilder) EUsersServ {
	return EUsersServ{
		edb:   edb,
		table: "users",
	}
}

func (c EUsersServ) Index(keywords string, perPage int64, page int64, orderBy string, orderDir string) ([]map[string]interface{}, int64, error) {
	db := c.edb.From(c.table).
		Select("*").
		OrderBy(orderBy, orderDir).
		Paginate(perPage, page)
	if keywords != "" {
		db = db.OrWhere("id", "=", keywords)
	}
	info, err := db.Get()
	count, err := db.Count()
	return info, count, err
}

// Insert
func (c EUsersServ) Insert(info map[string]interface{}) (int64, error) {
	id, err := c.edb.From(c.table).Insert(info)
	return id, err
}

// FindOne
func (c *EUsersServ) FindOne(id int64) (map[string]interface{}, error) {
	info, err := c.edb.From(c.table).Where("id", "=", id).First()
	return info, err
}

// Update
func (c *EUsersServ) Update(info map[string]interface{}) error {
	_, err := c.edb.From(c.table).
		Where("id", "=", info["id"]).
		Update(info)
	return err
}

// Delete
func (c *EUsersServ) Delete(id int64) error {
	_, err := c.edb.From(c.table).Where("id", "=", id).Delete()
	return err
}
