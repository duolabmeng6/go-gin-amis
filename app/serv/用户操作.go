package serv

import (
	"fmt"
	"github.com/duolabmeng6/goefun/edb"
	"go-gin-amis/app/dal"
	"go-gin-amis/app/dal/WrapTime"
	"go-gin-amis/app/dal/model"
	"go-gin-amis/app/dal/query"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var S用户操作 = New用户操作(dal.Edb, dal.DB.Debug())

type E用户操作 struct {
	DB    *gorm.DB
	Q     *query.Query
	edb   *edb.MySQLQueryBuilder
	table string
}

func New用户操作(edb *edb.MySQLQueryBuilder, db *gorm.DB) E用户操作 {
	return E用户操作{
		DB:    db,
		Q:     query.Use(db),
		edb:   edb,
		table: "users",
	}
}
func (c E用户操作) E查询用户(用户名 string) (*model.User, error) {
	var err error
	u := c.Q.User
	user, err := u.Where(u.Username.Eq(用户名)).First()
	if err != nil {
		return nil, err
	}
	return user, err
}
func (c E用户操作) E创建用户(用户 *model.User) error {
	var err error
	//查询用户名是否存在
	u := c.Q.User
	user, err := u.Where(u.Username.Eq(用户.Username)).First()
	// err 不等于  record not found
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if user != nil {
		return fmt.Errorf("%s 用户已存在", 用户.Username)
	}

	err = u.Create(用户)
	if err != nil {
		return fmt.Errorf("用户创建失败,%s", err.Error())
	}
	return err
}
func (c E用户操作) E增加积分(用户 *model.User, 积分 int64, 备注 string) error {
	err := c.Q.Transaction(func(tx *query.Query) error {
		u := tx.User
		ur := tx.UserIntegralRecord
		user, err := u.Clauses(clause.Locking{Strength: "UPDATE"}).Where(u.ID.Eq(用户.ID)).First()
		if err != nil {
			return err
		}

		user.Integral += 积分
		_, err = u.Where(u.ID.Eq(用户.ID)).Update(u.Integral, user.Integral)
		if err != nil {
			return err
		}
		//fmt.Println("update", update.RowsAffected, update.Error)

		urdata := &model.UserIntegralRecord{}
		urdata.UserID = user.ID
		urdata.BeforeIntegral = 用户.Integral
		urdata.AfterIntegral = user.Integral
		urdata.ChangeValue = 积分
		urdata.IntegralType = 0
		urdata.Remarks = 备注
		urdata.OrderType = 0
		urdata.OrderID = 0
		if err := ur.Create(urdata); err != nil {
			return err
		}
		return nil //返回 nil 提交事务
	})
	if err != nil {
		return err
	}

	return err
}
func (c E用户操作) E消费积分(用户 *model.User, 积分 int64, 备注 string) error {
	//开始事务 行级锁
	err := c.Q.Transaction(func(tx *query.Query) error {
		u := tx.User
		ur := tx.UserIntegralRecord
		user, err := u.Clauses(clause.Locking{Strength: "UPDATE"}).Where(u.ID.Eq(用户.ID)).First()
		if err != nil {
			return err
		}
		if user.Integral < 积分 {
			return fmt.Errorf("积分不足")
		}
		user.Integral -= 积分
		_, err = u.Where(u.ID.Eq(用户.ID)).Update(u.Integral, user.Integral)
		if err != nil {
			return err
		}
		//fmt.Println("update", update.RowsAffected, update.Error)

		urdata := &model.UserIntegralRecord{}
		urdata.UserID = user.ID
		urdata.BeforeIntegral = 用户.Integral
		urdata.AfterIntegral = user.Integral
		urdata.ChangeValue = 积分
		urdata.IntegralType = 0
		urdata.Remarks = 备注
		urdata.OrderType = 0
		urdata.OrderID = 0
		if err := ur.Create(urdata); err != nil {
			return err
		}
		return nil //返回 nil 提交事务
	})
	return err
}

type UserIntegralRecordOutput struct {
	ID            int64             `json:"id"`
	UserID        int64             `json:"user_id"`
	ChangeValue   int64             `json:"change_value"`
	AfterIntegral int64             `json:"after_integral"`
	CreatedAt     WrapTime.WrapTime `json:"created_at"`
}

func (c E用户操作) E查询充值记录(用户 *model.User, 当前页数 int64, 每页显示多少条 int64) ([]UserIntegralRecordOutput, error) {
	var outputs []UserIntegralRecordOutput
	ur := c.Q.UserIntegralRecord
	err := ur.Where(ur.UserID.Eq(用户.ID)).
		Offset(int((当前页数-1)*每页显示多少条)).
		Limit(int(每页显示多少条)).
		Select(ur.ID, ur.UserID, ur.ChangeValue, ur.AfterIntegral, ur.CreatedAt).
		Scan(&outputs)
	return outputs, err
}

func (c E用户操作) Index(keywords string, perPage int64, page int64, orderBy string, orderDir string) ([]map[string]interface{}, int64, error) {
	db := c.edb.From(c.table).
		Select("*").
		OrderBy(orderBy, orderDir).
		Paginate(perPage, page)
	if keywords != "" {
		db = db.OrWhere("username", "like", "%"+keywords+"%").
			OrWhere("id", "=", keywords)
	}
	articles, err := db.Get()
	count, err := db.Count()
	return articles, count, err
}

// Insert
func (c E用户操作) Insert(articles map[string]interface{}) (int64, error) {
	id, err := c.edb.From(c.table).Insert(articles)
	return id, err
}

// FindOne
func (c *E用户操作) FindOne(id int64) (map[string]interface{}, error) {
	articles, err := c.edb.From(c.table).Where("id", "=", id).First()
	return articles, err
}

// Update
func (c *E用户操作) Update(articles map[string]interface{}) error {
	_, err := c.edb.From(c.table).
		Where("id", "=", articles["id"]).
		Update(articles)
	return err
}

// Delete
func (c *E用户操作) Delete(id int64) error {
	_, err := c.edb.From(c.table).Where("id", "=", id).Delete()
	return err
}
